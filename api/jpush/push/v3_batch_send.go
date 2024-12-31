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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/callback"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/message"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/options"
)

// 批量单推（Registration ID 方式）
//  - 功能说明：如果您在给每个用户的推送内容都不同的情况下，可以使用此接口。使用此接口前，您需要配合使用 GetCidForPush 接口提前获取到 CID 池。
//	- 调用地址：POST `/v3/push/batch/regid/single`，`pushList` 的 key 为 CID 值，最多支持填写 1000 个。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_single
func (p *apiv3) BatchSendByRegistrationID(ctx context.Context, pushList map[string]BatchPushParam) (*BatchSendResult, error) {
	return p.batchSend(ctx, "regid", pushList)
}

// 批量单推（Alias 方式）
//  - 功能说明：如果您在给每个用户的推送内容都不同的情况下，可以使用此接口。使用此接口前，您需要配合使用 GetCidForPush 接口提前获取到 CID 池。
//	- 调用地址：POST `/v3/push/batch/alias/single`，`pushList` 的 key 为 CID 值，最多支持填写 1000 个。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_single
func (p *apiv3) BatchSendByAlias(ctx context.Context, pushList map[string]BatchPushParam) (*BatchSendResult, error) {
	return p.batchSend(ctx, "alias", pushList)
}

func (p *apiv3) batchSend(ctx context.Context, byType string, pushList map[string]BatchPushParam) (*BatchSendResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	l := len(pushList)
	if l == 0 {
		return nil, errors.New("`pushList` cannot be empty")
	}
	if l > 1000 {
		return nil, errors.New("`pushList` cannot be more than 1000")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/batch/" + byType + "/single",
		Auth:   p.auth,
		Body:   &batchSendParam{PushList: pushList},
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &BatchSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type batchSendParam struct {
	PushList map[string]BatchPushParam `json:"pushlist"` // 【必填】key 为 CID 值，最多支持填写 1000 个。
}

// 批量推送参数。
type BatchPushParam struct {
	// 【必填】推送平台，支持 2 种类型值：
	//  - platform.All：推送到所有平台；
	//  - platform.Android、platform.IOS、platform.HMOS、platform.QuickApp 的组合列表：指定特定推送平台。
	Platform interface{} `json:"platform"`
	// 【必填】推送目标，此处填写的是 Registration ID 值或者 Alias 值。
	Target string `json:"target"`
	// 【可选】推送可选项。
	//  - 详见文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#options%EF%BC%9A%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0。
	Options *options.Options `json:"options,omitempty"`
	// 【可选】通知内容，是被推送到客户端的内容；与 CustomMessage 一起二者必须有其一，可以二者并存。
	//  - 详见文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification%EF%BC%9A%E9%80%9A%E7%9F%A5。
	Notification *notification.Notification `json:"notification,omitempty"`
	// 【可选】自定义消息内容，是被推送到客户端的内容；与 Notification 一起二者必须有其一，可以二者并存。
	//  - 详见文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#message%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF。
	CustomMessage *message.Custom `json:"message,omitempty"`
	// 【可选】短信渠道补充送达内容。
	//  - 详见文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1。
	SmsMessage *message.SMS `json:"sms_message,omitempty"`
	// 【可选】回调参数。
	//  - 详见文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#callback%EF%BC%9A%E5%9B%9E%E8%B0%83%E5%8F%82%E6%95%B0。
	Callback *callback.Callback `json:"callback,omitempty"`
}

type BatchSendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError             `json:"error,omitempty"`
	SendResult    map[string]BatchPushResult `json:"-"` // key 为 CID 值。
}

type BatchPushResult struct {
	MsgID string          `json:"msg_id,omitempty"` // 推送消息 ID
	Error *BatchSendError `json:"error,omitempty"`  // 错误信息
}

func (rs *BatchSendResult) UnmarshalJSON(data []byte) error {
	rs.SendResult = make(map[string]BatchPushResult)

	var aux map[string]json.RawMessage
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
	}

	if len(aux) > 0 {
		for k, v := range aux {
			var result BatchPushResult
			if err := json.Unmarshal(v, &result); err != nil {
				return err
			}
			rs.SendResult[k] = result
		}
	}

	return nil
}

func (rs BatchSendResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else if len(rs.SendResult) > 0 {
		data := make(map[string]BatchPushResult, len(rs.SendResult))
		for k, v := range rs.SendResult {
			data[k] = v
		}
		return json.Marshal(data)
	}
	return []byte("null"), nil
}

func (rs *BatchSendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 批量单推失败错误
type BatchSendError struct {
	api.CodeError
	IllegalRIDs []string `json:"illegal_rids,omitempty"` // 非法 Registration ID 集合
}

func (e *BatchSendError) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}

func (e *BatchSendError) String() string {
	if e == nil {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}
