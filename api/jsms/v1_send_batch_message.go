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

package jsms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 发送批量模板短信
//   - 功能说明：发送批量模板短信。
//   - 调用地址：POST `/v1/messages/batch`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E6%89%B9%E9%87%8F%E6%A8%A1%E6%9D%BF%E7%9F%AD%E4%BF%A1-api
func (s *apiv1) SendBatchMessage(ctx context.Context, param *BatchMessageSendParam) (*BatchMessageSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/messages/batch",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &BatchMessageSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 批量模板短信发送参数
type BatchMessageSendParam struct {
	SignID     int                       `json:"sign_id,omitempty"` // 【可选】签名 ID，该字段为空则使用应用默认签名
	TempID     int64                     `json:"temp_id"`           // 【必填】模板 ID
	Tag        string                    `json:"tag,omitempty"`     // 【可选】标签
	Recipients []BatchSendRecipientParam `json:"recipients"`        // 【必填】接收者列表
}

// 批量模板短信接收者参数
type BatchSendRecipientParam struct {
	Mobile     string                 `json:"mobile"`              // 【必填】手机号码
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 【可选】模板参数，需要替换的参数名和参数值的键值对
}

// 批量模板短信发送结果
type BatchMessageSendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError             `json:"error,omitempty"`
	SuccessCount  *uint64                    `json:"success_count,omitempty"` // 成功发送的数量
	FailureCount  *uint64                    `json:"failure_count,omitempty"` // 发送失败的数量
	Recipients    []BatchSendRecipientResult `json:"recipients,omitempty"`    // 发送失败接收者结果列表
}

func (rs *BatchMessageSendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 批量模板短信发送接收者结果
type BatchSendRecipientResult struct {
	ErrorCode    string                 `json:"error_code,omitempty"`    // 错误码
	ErrorMessage string                 `json:"error_message,omitempty"` // 错误信息
	MsgID        string                 `json:"msg_id,omitempty"`        // 消息 ID
	Mobile       string                 `json:"mobile"`                  // 手机号码
	TempParams   map[string]interface{} `json:"temp_para,omitempty"`     // 模板参数，需要替换的参数名和参数值的键值对
}

func (rs *BatchSendRecipientResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `BatchSendRecipientResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias BatchSendRecipientResult
	aux := struct {
		ErrorCode json.RawMessage `json:"error_code"`
		MsgID     json.RawMessage `json:"msg_id"`
		*Alias
	}{
		Alias: (*Alias)(rs),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var strErrorCode string // 字符串类型的 `error_code`
	if err := json.Unmarshal(aux.ErrorCode, &strErrorCode); err == nil {
		if strErrorCode != "" { // omitempty
			rs.ErrorCode = strErrorCode
		}
	} else {
		var intErrorCode int // 兼容数字类型的 `error_code`
		if err = json.Unmarshal(aux.ErrorCode, &intErrorCode); err == nil {
			if intErrorCode != 0 { // omitempty
				rs.ErrorCode = strconv.Itoa(intErrorCode)
			}
		} else {
			// return err
			rs.ErrorCode = string(aux.ErrorCode)
		}
	}

	var strMsgID string // 字符串类型的 `msg_id`
	if err := json.Unmarshal(aux.MsgID, &strMsgID); err == nil {
		if strMsgID != "" { // omitempty
			rs.MsgID = strMsgID
		}
	} else {
		var intMsgID int64 // 兼容数字类型的 `msg_id`
		if err = json.Unmarshal(aux.MsgID, &intMsgID); err == nil {
			if intMsgID != 0 { // omitempty
				rs.MsgID = strconv.FormatInt(intMsgID, 10)
			}
		} else {
			// return err
			rs.MsgID = string(aux.MsgID)
		}
	}

	return nil
}

func (rs *BatchSendRecipientResult) IsSuccess() bool {
	return rs == nil || rs.ErrorCode == "50000"
}
