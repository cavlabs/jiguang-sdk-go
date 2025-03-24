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

// # 批量定时短信修改
//   - 功能说明：修改批量模板短信定时发送任务。
//   - 调用地址：PUT `/v1/schedule/batch/{scheduleID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E6%89%B9%E9%87%8F%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E4%BF%AE%E6%94%B9-api
func (s *apiv1) UpdateBatchSchedule(ctx context.Context, scheduleID string, param *BatchScheduleUpdateParam) (*BatchScheduleUpdateResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if scheduleID == "" {
		return nil, errors.New("`scheduleID` cannot be empty")
	}
	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPut,
		Proto:  s.proto,
		URL:    s.host + "/v1/schedule/batch/" + scheduleID,
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &BatchScheduleUpdateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type (
	BatchScheduleUpdateParam  = BatchScheduleSendParam
	BatchScheduleUpdateResult = BatchScheduleSendResult
)
