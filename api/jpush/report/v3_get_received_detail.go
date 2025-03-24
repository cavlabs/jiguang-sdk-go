/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
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

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 送达统计详情
//   - 功能说明：获取指定 msgIDs（最多支持 100 个）的送达统计数据。如果一次 API 调用推送有很多对象（比如广播推送），则此 API 返回的统计数据会因为持续有客户端送达而持续增加。
//     此接口会根据消息是通过极光自有通道下发、Android 厂商通道下发进行数据统计区分。
//     每条推送消息的送达统计数据最多保留一个月。即发起推送请求后从最后一个推送送达记录时间点开始保留一个月，如果保留期间有新的送达，将在这个新送达的时间点起再往后保留一个月。
//   - 调用地址：GET `/v3/received/detail`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E9%80%81%E8%BE%BE%E7%BB%9F%E8%AE%A1%E8%AF%A6%E6%83%85
func (r *apiv3) GetReceivedDetail(ctx context.Context, msgIDs []string) (*ReceivedDetailGetResult, error) {
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
		URL:    r.host + "/v3/received/detail?msg_ids=" + strings.Join(msgIDs, ","),
		Auth:   r.auth,
	}
	resp, err := r.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ReceivedDetailGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ReceivedDetailGetResult struct {
	*api.Response   `json:"-"`
	Error           *api.CodeError   `json:"error,omitempty"`
	ReceivedDetails []ReceivedDetail `json:"-"`
}

type ReceivedDetail struct {
	MsgID                 string  `json:"msg_id"`                  // 推送消息 ID
	JPushReceived         *uint64 `json:"jpush_received"`          // 极光通道用户送达数；包含普通 Android 用户的通知 + 自定义消息送达，iOS 用户自定义消息送达；如果无此项数据则为 null。
	AndroidPnsSent        *uint64 `json:"android_pns_sent"`        // Android 厂商用户推送到厂商服务器成功数，计算方式同 Android 厂商成功数；如果无此项数据则为 null。
	AndroidPnsReceived    *uint64 `json:"android_pns_received"`    // Android 厂商用户推送达到设备数，计算方式以厂商回调数据为准；如果无此项数据则为 null。2020.03.24 新增指标。
	IOSApnsSent           *uint64 `json:"ios_apns_sent"`           // iOS 通知推送到 APNs 成功数量。如果无此项数据则为 null。
	IOSApnsReceived       *uint64 `json:"ios_apns_received"`       // iOS 通知送达到设备并成功展示。如果无项数据则为 null。统计该项请参考【集成指南高级功能-通知展示统计】：https://docs.jiguang.cn/jpush/client/iOS/ios_guide_new#%E9%80%9A%E7%9F%A5%E5%B1%95%E7%A4%BA%E7%BB%9F%E8%AE%A1。
	IOSMsgReceived        *uint64 `json:"ios_msg_received"`        // iOS 自定义消息送达数。如果无此项数据则为 null。
	LiveActivitySent      *uint64 `json:"live_acivity_sent"`       // 实时活动消息推送到 APNs 成功的用户数量。
	LiveActivityReceived  *uint64 `json:"live_acivity_received"`   // 实时活动消息送达成功的用户数量。
	WpMpnsSent            *uint64 `json:"wp_mpns_sent"`            // Windows Phone 通知推送到 MPNS 成功数量。如果无此项数据则为 null。
	QuickAppPnsSent       *uint64 `json:"quickapp_pns_sent"`       // 快应用推送走厂商通道请求成功的用户数量。
	QuickAppJPushReceived *uint64 `json:"quickapp_jpush_received"` // 快应用推送走极光通道送达设备成功的用户数量。
	HmosHmpnsSent         *uint64 `json:"hmos_hmpns_sent"`         // 鸿蒙通知推送到厂商服务器成功数。2024.08.13 新增指标。
	HmosHmpnsReceived     *uint64 `json:"hmos_hmpns_received"`     // 鸿蒙通知送达到设备数，计算方式以厂商回调数据为准。2024.08.13 新增指标。
	HmosMsgReceived       *uint64 `json:"hmos_msg_received"`       // 鸿蒙自定义消息送达到设备数，计算方式以厂商回调数据为准。2024.08.13 新增指标。
}

func (rs *ReceivedDetailGetResult) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' && data[len(data)-1] == ']' {
		var details []ReceivedDetail
		if err := json.Unmarshal(data, &details); err != nil {
			return err
		}
		rs.ReceivedDetails = details
		return nil
	}
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ReceivedDetailGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ReceivedDetailGetResult
	aux := (*Alias)(rs)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (rs ReceivedDetailGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else {
		return json.Marshal(rs.ReceivedDetails)
	}
}

func (rs *ReceivedDetailGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
