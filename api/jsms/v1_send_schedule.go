/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
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

package jsms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 单条定时短信提交
//  - 功能说明：提交单条模板短信定时发送任务。
//  - 调用地址：POST `/v1/schedule`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%8D%95%E6%9D%A1%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%8F%90%E4%BA%A4-api
func (s *apiv1) SendSchedule(ctx context.Context, param *ScheduleSendParam) (*ScheduleSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/schedule",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ScheduleSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ScheduleSendParam struct {
	SendTime   jiguang.LocalDateTime  `json:"send_time"`           // 【必填】发送时间。
	Mobile     string                 `json:"mobile"`              // 【必填】手机号码。
	SignID     int                    `json:"sign_id,omitempty"`   // 【可选】签名 ID，该参数为空则使用应用默认签名；注意：签名审核通过后，可在控制台设为默认签名。
	TempID     int64                  `json:"temp_id"`             // 【必填】模板 ID。
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 【可选】模板参数，需要替换的参数名和参数值的键值对。
}

type ScheduleSendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	ScheduleID    string         `json:"schedule_id,omitempty"` // 任务 ID。
}

func (rs *ScheduleSendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
