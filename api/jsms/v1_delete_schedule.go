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
)

// 定时短信删除
//  - 功能说明：删除模板短信定时发送任务。
//  - 调用地址：DELETE `/v1/schedule/{scheduleID}`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E5%88%A0%E9%99%A4-api
func (s *apiv1) DeleteSchedule(ctx context.Context, scheduleID string) (*ScheduleDeleteResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if scheduleID == "" {
		return nil, errors.New("`scheduleID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodDelete,
		Proto:  s.proto,
		URL:    s.host + "/v1/schedule/" + scheduleID,
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ScheduleDeleteResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type ScheduleDeleteResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *ScheduleDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
