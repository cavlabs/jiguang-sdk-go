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

package schedule

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 获取有效的定时任务列表
//   - 功能说明：获取当前有效（EndTime 未过期）的定时任务列表。
//   - 调用地址：GET `/v3/schedules?page={page}`
//   - 接口文档：[docs.jiguang.cn]
//
// 注意事项：
//   - 返回当前 page 请求页的详细的定时任务列表，如 page < 0 则自动重置为 1；
//   - 结果列表按 “创建时间” 排序；
//   - 如果请求页数大于总页数，则 page 为请求值，结果列表为空；
//   - 每页最多返回 50 个定时任务，如请求页实际的个数小于 50，则按实际数量返回。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E8%8E%B7%E5%8F%96%E6%9C%89%E6%95%88%E7%9A%84%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E5%88%97%E8%A1%A8
func (s *apiv3) GetSchedules(ctx context.Context, page int) (*SchedulesGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJPushScheduleAPIv3
	}

	if page < 1 {
		page = 1
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v3/schedules?page=" + strconv.Itoa(page),
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SchedulesGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SchedulesGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	TotalCount    int            `json:"total_count,omitempty"` // 当前有效的定时任务总数
	TotalPages    int            `json:"total_pages,omitempty"` // 总页数
	Page          int            `json:"page,omitempty"`        // 当前页数
	Schedules     []Schedule     `json:"schedules,omitempty"`   // 定时任务详情列表
}

func (rs *SchedulesGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
