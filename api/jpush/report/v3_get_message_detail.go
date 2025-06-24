/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package report

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 消息统计详情（VIP-新）
//   - 功能说明：获取指定 msgIDs（最多支持 100 个）的消息统计数据。与 “送达统计详情” 不同的是，该 API 提供了更多的针对单个 msgID 的统计数据。
//   - 调用地址：GET `/v3/messages/detail`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E6%B6%88%E6%81%AF%E7%BB%9F%E8%AE%A1%E8%AF%A6%E6%83%85%EF%BC%88vip-%E6%96%B0%EF%BC%89
func (r *apiv3) GetMessageDetail(ctx context.Context, msgIDs []string) (*MessageDetailGetResult, error) {
	if r == nil {
		return nil, api.ErrNilJPushReportAPIv3
	}

	l := len(msgIDs)
	if l == 0 {
		return nil, errors.New("`msgIDs` cannot be empty")
	}
	if l > 100 {
		return nil, errors.New("`msgIDs` cannot be more than 100")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  r.proto,
		URL:    r.host + "/v3/messages/detail?msg_ids=" + strings.Join(msgIDs, ","),
		Auth:   r.auth,
	}
	resp, err := r.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &MessageDetailGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type MessageDetailGetResult struct {
	*api.Response  `json:"-"`
	Error          *api.CodeError  `json:"error,omitempty"`
	MessageDetails []MessageDetail `json:"-"`
}

// # 消息统计详情
type MessageDetail struct {
	MsgID   string   `json:"msg_id"`            // 推送消息 ID
	Details *Details `json:"details,omitempty"` // 消息统计详情指标，同一条消息可能包含通知栏消息、自定义消息和应用提醒消息等不同类型

	// 以下为 2021.09.01 前旧体系指标 ↓↓↓

	JPush         *LegacyJPush         `json:"jpush,omitempty"`          // 极光通道统计数据，走极光通道下发的普通 Android 用户通知/自定义消息以及 iOS 用户自定义消息总体情况
	AndroidPns    *LegacyAndroidPns    `json:"android_pns,omitempty"`    // Android 厂商通道统计数据，走厂商通道下发统计数据
	IOS           *LegacyIos           `json:"ios,omitempty"`            // iOS 平台统计数据
	QuickAppJPush *LegacyQuickAppJPush `json:"quickapp_jpush,omitempty"` // 快应用平台极光通道统计数据
	QuickAppPns   *LegacyQuickAppPns   `json:"quickapp_pns,omitempty"`   // 快应用平台 Android 厂商通道统计数据
}

// # 消息统计详情指标定义
//
// 2021.09.01 新体系指标。
type Details struct {
	Notification  *MessageStats      `json:"notification,omitempty"`  // 通知栏消息类型的数据汇总统计
	CustomMessage *MessageStats      `json:"message,omitempty"`       // 自定义消息的数据汇总统计
	InApp         *MessageStats      `json:"inapp,omitempty"`         // 应用内提醒消息的数据汇总统计
	LiveActivity  *LiveActivityStats `json:"live_activity,omitempty"` // 实时活动消息的数据汇总统计
}

// # 消息生命周期状态定义
//
// 2021.09.01 新体系指标。
type MessageStats struct {
	Target   *uint64 `json:"target,omitempty"`   // 有效目标，将推送任务所选定的目标人群，经过有效性筛选后的目标设备数量
	Sent     *uint64 `json:"sent,omitempty"`     // 发送数量，有效目标设备中，极光服务器实际成功创建了发送任务的设备数量
	Received *uint64 `json:"received,omitempty"` // 送达数量，通知消息发送后，实际送达至设备终端的数量，5 天之后的送达数量不被计算在内，华为、荣耀、魅族和 iOS 需要配置回调送达数据才更加精准
	Display  *uint64 `json:"display,omitempty"`  // 展示数量，通知消息送达后，实际在设备终端成功展示的数量，5 天之后的展示数量不被计算在内
	Click    *uint64 `json:"click,omitempty"`    // 点击数量，通知消息成功展示后，实际被用户点击的数量，5 天之后的点击数量不被计算在内

	// sub_*: 对不同类型的消息，按平台按发送通道汇总统计指标 ↓↓↓

	SubAndroid  *AndroidStats  `json:"sub_android,omitempty"`  // Android 平台的数据汇总统计
	SubIos      *IosStats      `json:"sub_ios,omitempty"`      // iOS 平台的数据汇总统计
	SubQuickApp *QuickAppStats `json:"sub_quickapp,omitempty"` // 快应用平台的数据汇总统计
	SubHmos     *HmosStats     `json:"sub_hmos,omitempty"`     // 鸿蒙平台的数据汇总统计，2024.08.13 新增
}

// # 实时活动消息生命周期状态定义
//
// 2022.12.22 实时活动指标。
type LiveActivityStats struct {
	Target   *uint64 `json:"target,omitempty"`   // 有效目标汇总，将推送任务所选定的目标人群，经过有效性筛选后的目标设备数量
	Sent     *uint64 `json:"sent,omitempty"`     // 发送数量汇总，有效目标设备中，极光服务器实际成功创建了发送任务的设备数量
	Received *uint64 `json:"received,omitempty"` // 送达数量汇总，通知消息发送后，实际送达至设备终端的数量，5 天之后的送达数量不被计算在内，华为、荣耀、魅族和 iOS 需要配置回调送达数据才更加精准
	Display  *uint64 `json:"display,omitempty"`  // 展示数量汇总，通知消息送达后，实际在设备终端成功展示的数量，5 天之后的展示数量不被计算在内
	Click    *uint64 `json:"click,omitempty"`    // 点击数量汇总，通知消息成功展示后，实际被用户点击的数量，5 天之后的点击数量不被计算在内

	// sub_*: 按发送通道汇总统计指标 ↓↓↓

	SubIos *ChannelStats `json:"sub_ios,omitempty"` // iOS 平台统计指标
}

// # Android 平台的数据汇总统计
type AndroidStats struct {
	Jiguang *ChannelStats `json:"jg_android,omitempty"` // 极光通道统计指标
	Xiaomi  *ChannelStats `json:"xiaomi,omitempty"`     // 小米厂商通道统计指标
	Huawei  *ChannelStats `json:"huawei,omitempty"`     // 华为厂商通道统计指标
	Honor   *ChannelStats `json:"honor,omitempty"`      // 荣耀厂商通道统计指标
	Meizu   *ChannelStats `json:"meizu,omitempty"`      // 魅族厂商通道统计指标
	OPPO    *ChannelStats `json:"oppo,omitempty"`       // OPPO 厂商通道统计指标
	Vivo    *ChannelStats `json:"vivo,omitempty"`       // vivo 厂商通道统计指标
	ASUS    *ChannelStats `json:"asus,omitempty"`       // 华硕厂商通道统计指标
	FCM     *ChannelStats `json:"fcm,omitempty"`        // FCM 通道统计指标
	Tuibida *ChannelStats `json:"tuibida,omitempty"`    // 推必达通道统计指标
	NIO     *ChannelStats `json:"nio,omitempty"`        // 蔚来通道统计指标，2025.01.13 新增
}

// # iOS 平台的数据汇总统计
type IosStats struct {
	Jiguang *ChannelStats `json:"jg_ios,omitempty"` // 极光通道统计指标
	VoIP    *ChannelStats `json:"voip,omitempty"`   // VoIP 通道统计指标
	APNs    *ChannelStats `json:"apns,omitempty"`   // APNs 通道统计指标
}

// # 快应用平台的数据汇总统计
type QuickAppStats struct {
	Jiguang *ChannelStats `json:"quick_jg,omitempty"`     // 极光通道统计指标
	Xiaomi  *ChannelStats `json:"quick_xiaomi,omitempty"` // 小米厂商通道统计指标
	Huawei  *ChannelStats `json:"quick_huawei,omitempty"` // 华为厂商通道统计指标
	OPPO    *ChannelStats `json:"quick_oppo,omitempty"`   // OPPO 厂商通道统计指标
}

// # 鸿蒙平台的数据汇总统计
type HmosStats struct {
	Hmpns *ChannelStats `json:"hmpns,omitempty"` // HMPNs 发送通道统计指标
}

// # 不同平台发送通道的数据汇总统计
type ChannelStats struct {
	Target   *uint64 `json:"target,omitempty"`   // 有效目标，将推送任务所选定的目标人群，经过有效性筛选后的目标设备数量
	Sent     *uint64 `json:"sent,omitempty"`     // 发送数量，有效目标设备中，极光服务器实际成功创建了发送任务的设备数量
	Received *uint64 `json:"received,omitempty"` // 送达数量，通知消息发送后，实际送达至设备终端的数量，5 天之后的送达数量不被计算在内，华为、荣耀、魅族和 iOS 需要配置回调送达数据才更加精准
	Display  *uint64 `json:"display,omitempty"`  // 展示数量，通知消息送达后，实际在设备终端成功展示的数量，5 天之后的展示数量不被计算在内
	Click    *uint64 `json:"click,omitempty"`    // 点击数量，通知消息成功展示后，实际被用户点击的数量，5 天之后的点击数量不被计算在内
}

// # 极光通道统计数据
//
// 2021.09.01 前旧体系指标。
type LegacyJPush struct {
	Target     *uint64 `json:"target,omitempty"`      // 有效目标
	OnlinePush *uint64 `json:"online_push,omitempty"` // 在线推送数，该字段 2021.09.01 失效
	Sent       *uint64 `json:"sent,omitempty"`        // 发送数量
	Received   *uint64 `json:"received,omitempty"`    // 送达数量
	Display    *uint64 `json:"display,omitempty"`     // 展示数量，2021.08.17 新增指标
	Click      *uint64 `json:"click,omitempty"`       // 通知栏消息点击数量
	MsgClick   *uint64 `json:"msg_click,omitempty"`   // 自定义消息点击数量
}

// # Android 厂商通道统计数据
//
// 2021.09.01 前旧体系指标。
type LegacyAndroidPns struct {
	PnsTarget   *uint64 `json:"pns_target,omitempty"`   // 通过厂商通道推送目标数
	PnsSent     *uint64 `json:"pns_sent,omitempty"`     // 推送到厂商通道成功数
	PnsReceived *uint64 `json:"pns_received,omitempty"` // 厂商推送送达设备数，2020.03.24 新增指标
	PnsDisplay  *uint64 `json:"pns_display,omitempty"`  // 厂商推送展示数，2021.08.17 新增指标

	// *_detail: 按发送通道汇总统计详情 ↓↓↓

	XiaomiDetail *LegacyChannelStats `json:"xm_detail,omitempty"`    // 推送到小米通道详情
	HuaweiDetail *LegacyChannelStats `json:"hw_detail,omitempty"`    // 推送到华为通道详情
	HonorDetail  *LegacyChannelStats `json:"honor_detail,omitempty"` // 推送到荣耀通道详情，2022.06.20 新增
	MeizuDetail  *LegacyChannelStats `json:"mz_detail,omitempty"`    // 推送到魅族通道详情
	OppoDetail   *LegacyChannelStats `json:"oppo_detail,omitempty"`  // 推送到 OPPO 通道详情
	VivoDetail   *LegacyChannelStats `json:"vivo_detail,omitempty"`  // 推送到 vivo 通道详情
	AsusDetail   *LegacyChannelStats `json:"asus_detail,omitempty"`  // 推送到华硕通道详情
	FcmDetail    *LegacyChannelStats `json:"fcm_detail,omitempty"`   // 推送到 FCM 通道详情
	NioDetail    *LegacyChannelStats `json:"nio_detail,omitempty"`   // 推送到蔚来通道详情，2025.01.13 新增
}

// # 推送到厂商通道详情
//
// 2021.09.01 前旧体系指标。
type LegacyChannelStats struct {
	Target   *uint64 `json:"target,omitempty"`   // 推送目标数
	Sent     *uint64 `json:"sent,omitempty"`     // 推送成功数
	Received *uint64 `json:"received,omitempty"` // 送达设备数，2020.03.24 新增指标
	Display  *uint64 `json:"display,omitempty"`  // 推送展示数，2021.08.17 新增指标
}

// # iOS 平台统计数据
//
// 2021.09.01 前旧体系指标。
type LegacyIos struct {
	ApnsTarget   *uint64 `json:"apns_target,omitempty"`   // APNs 通知推送目标数
	ApnsSent     *uint64 `json:"apns_sent,omitempty"`     // APNs 通知推送成功数，发送到 APNs 服务器成功
	ApnsReceived *uint64 `json:"apns_received,omitempty"` // APNs 通知送达展示数，APNs 服务器下发到设备并成功展示，统计该项请参考【集成指南高级功能-通知展示统计】：https://docs.jiguang.cn/jpush/client/iOS/ios_guide_new#%E9%80%9A%E7%9F%A5%E5%B1%95%E7%A4%BA%E7%BB%9F%E8%AE%A1
	ApnsClick    *uint64 `json:"apns_click,omitempty"`    // APNs 通知用户点击数
	ApnsDisplay  *uint64 `json:"apns_display,omitempty"`  // APNs 通知展示数，2021.08.17 新增指标
	MsgTarget    *uint64 `json:"msg_target,omitempty"`    // 自定义消息推送目标数
	MsgReceived  *uint64 `json:"msg_received,omitempty"`  // 自定义消息推送送达数
	MsgClick     *uint64 `json:"msg_click,omitempty"`     // 自定义消息用户点击数，iOS 由于其特殊性，自定义消息无点击指标统计，但会有此字段返回
	MsgDisplay   *uint64 `json:"msg_display,omitempty"`   // 自定义消息展示数，2021.08.17 新增指标
}

// # 快应用平台极光通道统计数据
//
// 2021.09.01 前旧体系指标。
type LegacyQuickAppJPush struct {
	Target     *uint64 `json:"target,omitempty"`      // 推送目标数
	OnlinePush *uint64 `json:"online_push,omitempty"` // 在线推送数，该字段 2021.09.01 失效
	Received   *uint64 `json:"received,omitempty"`    // 推送送达数
	Click      *uint64 `json:"click,omitempty"`       // 用户点击数
	MsgClick   *uint64 `json:"msg_click,omitempty"`   // 自定义消息点击数
}

// # 快应用平台 Android 厂商通道统计数据
//
// 2021.09.01 前旧体系指标。
type LegacyQuickAppPns struct {
	PnsTarget *uint64 `json:"pns_target,omitempty"` // 推送目标数
	PnsSent   *uint64 `json:"pns_sent,omitempty"`   // 推送成功数
}

func (rs *MessageDetailGetResult) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' && data[len(data)-1] == ']' {
		var details []MessageDetail
		if err := json.Unmarshal(data, &details); err != nil {
			return err
		}
		rs.MessageDetails = details
		return nil
	}
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `MessageDetailGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias MessageDetailGetResult
	aux := (*Alias)(rs)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (rs MessageDetailGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else {
		return json.Marshal(rs.MessageDetails)
	}
}

func (rs *MessageDetailGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
