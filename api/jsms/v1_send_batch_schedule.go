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
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 批量定时短信提交
//   - 功能说明：提交批量模板短信定时发送任务。
//   - 调用地址：POST `/v1/schedule/batch`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E6%89%B9%E9%87%8F%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%8F%90%E4%BA%A4-api
func (s *apiv1) SendBatchSchedule(ctx context.Context, param *BatchScheduleSendParam) (*BatchScheduleSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/schedule/batch",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &BatchScheduleSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type BatchScheduleSendParam struct {
	SendTime   jiguang.LocalDateTime         `json:"send_time"`         // 【必填】发送时间。
	SignID     int                           `json:"sign_id,omitempty"` // 【可选】签名 ID，该参数为空则使用应用默认签名；注意：签名审核通过后，可在控制台设为默认签名。
	TempID     int64                         `json:"temp_id"`           // 【必填】模板 ID。
	Tag        string                        `json:"tag,omitempty"`     // 【可选】标签，仅用作标示该短信的别名，不在短信中展示，最多不超过 10 个字。
	Recipients []BatchScheduleRecipientParam `json:"recipients"`        // 【必填】接收者列表，上限 1000 个。
}

// 批量定时短信接收者参数
type BatchScheduleRecipientParam struct {
	Mobile     string                 `json:"mobile"`              // 【必填】手机号码。
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 【可选】模板参数，需要替换的参数名和参数值的键值对。
}

type BatchScheduleSendResult struct {
	*api.Response     `json:"-"`
	Error             *api.CodeError                        `json:"error,omitempty"`
	SuccessCount      *uint64                               `json:"success_count,omitempty"`      // 成功发送的数量
	FailureCount      *uint64                               `json:"failure_count,omitempty"`      // 发送失败的数量
	FailureRecipients []BatchScheduleFailureRecipientResult `json:"failure_recipients,omitempty"` // 发送失败接收者结果列表
}

func (rs *BatchScheduleSendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 批量定时短信发送失败接收者结果
type BatchScheduleFailureRecipientResult struct {
	ErrorCode    string                 `json:"error_code,omitempty"`    // 错误码
	ErrorMessage string                 `json:"error_message,omitempty"` // 错误信息
	Mobile       string                 `json:"mobile"`                  // 手机号码
	TempParams   map[string]interface{} `json:"temp_para,omitempty"`     // 模板参数，需要替换的参数名和参数值的键值对
}

func (rs *BatchScheduleFailureRecipientResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `BatchScheduleFailureRecipientResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias BatchScheduleFailureRecipientResult
	aux := struct {
		ErrorCode json.RawMessage `json:"error_code"`
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

	return nil
}

func (rs *BatchScheduleFailureRecipientResult) IsSuccess() bool {
	return rs == nil || rs.ErrorCode == "50000"
}
