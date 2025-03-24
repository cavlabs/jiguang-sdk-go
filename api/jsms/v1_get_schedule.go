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

// # 定时短信查询
//   - 功能说明：查询模板短信定时发送任务。
//   - 调用地址：GET `/v1/schedule/{scheduleID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%9F%A5%E8%AF%A2api
func (s *apiv1) GetSchedule(ctx context.Context, scheduleID string) (*ScheduleGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if scheduleID == "" {
		return nil, errors.New("`scheduleID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/schedule/" + scheduleID,
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ScheduleGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ScheduleGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError            `json:"error,omitempty"`
	ScheduleID    string                    `json:"schedule_id,omitempty"` // 任务 ID
	SendTime      *jiguang.LocalDateTime    `json:"send_time,omitempty"`   // 发送时间
	TempID        int64                     `json:"temp_id,omitempty"`     // 模板 ID
	Recipients    []ScheduleRecipientResult `json:"recipients,omitempty"`  // 接收者列表
}

// 定时短信查询接收者结果
type ScheduleRecipientResult struct {
	MsgID      string                 `json:"msg_id,omitempty"`    // 消息 ID
	Mobile     string                 `json:"mobile"`              // 手机号码
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 模板参数，需要替换的参数名和参数值的键值对
}

func (rs *ScheduleGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
