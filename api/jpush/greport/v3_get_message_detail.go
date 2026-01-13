// Copyright 2025 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package greport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/report"
)

// # 消息统计详情（VIP）
//   - 功能说明：针对分组应用，获取指定 groupMsgIDs（最多支持 10 个）的消息统计数据。只支持查询 30 天以内的推送信息。
//   - 调用地址：GET `/v3/group/messages/detail`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E5%88%86%E7%BB%84%E7%BB%9F%E8%AE%A1-%E6%B6%88%E6%81%AF%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
func (gr *apiv3) GetMessageDetail(ctx context.Context, groupMsgIDs []string) (*MessageDetailGetResult, error) {
	if gr == nil {
		return nil, api.ErrNilJPushGroupReportAPIv3
	}

	l := len(groupMsgIDs)
	if l == 0 {
		return nil, errors.New("`groupMsgIDs` cannot be empty")
	}
	if l > 10 {
		return nil, errors.New("`groupMsgIDs` cannot be more than 10")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  gr.proto,
		URL:    gr.host + "/v3/group/messages/detail?group_msgids=" + strings.Join(groupMsgIDs, ","),
		Auth:   gr.auth,
	}
	resp, err := gr.client.Request(ctx, req)
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

// 分组消息统计详情
type MessageDetail struct {
	GroupMsgID string           `json:"group_msgid"`           // 分组推送消息 ID
	JPush      *JPushStats      `json:"jpush,omitempty"`       // 极光通道统计数据，走极光通道下发的普通 Android 用户通知/自定义消息以及 iOS 用户自定义消息总体情况
	AndroidPns *AndroidPnsStats `json:"android_pns,omitempty"` // Android 厂商通道统计数据，走厂商通道下发统计数据
	IOS        *IosStats        `json:"ios,omitempty"`         // iOS 平台统计数据
	HMOS       *HmosStats       `json:"hmos,omitempty"`        // 鸿蒙平台统计数据
}

// ↓↓↓ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↓↓↓

type (
	JPushStats      = report.LegacyJPush
	AndroidPnsStats = report.LegacyAndroidPns
	IosStats        = report.LegacyIos
)

// ↑↑↑ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↑↑↑

// HmosStats 鸿蒙平台统计数据（2024.08.13 新增指标）
type HmosStats struct {
	HmpnsTarget   *uint64 `json:"hmpns_target,omitempty"`   // 鸿蒙通知推送目标数
	HmpnsSent     *uint64 `json:"hmpns_sent,omitempty"`     // 鸿蒙通知成功推送数，发送到 HMPNs 服务器成功
	HmpnsReceived *uint64 `json:"hmpns_received,omitempty"` // 鸿蒙通知送达展示数，HMPNs 服务器下发到设备并成功展示
	HmpnsClick    *uint64 `json:"hmpns_click,omitempty"`    // 鸿蒙通知用户点击数
	MsgTarget     *uint64 `json:"msg_target,omitempty"`     // 鸿蒙自定义消息推送目标数
	MsgReceived   *uint64 `json:"msg_received,omitempty"`   // 鸿蒙自定义消息推送送达数
	MsgClick      *uint64 `json:"msg_click,omitempty"`      // 鸿蒙自定义消息用户点击数
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
	}
	return json.Marshal(rs.MessageDetails)
}

func (rs *MessageDetailGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
