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

package push

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// # 查询推送计划列表（VIP）
//   - 功能说明：查询推送计划列表。
//   - 调用地址：GET `/v3/push_plan/list?page={page}&page_size={pageSize}&info={info}&send_source={sendSource}`；
//     `page` 为查询页码，`pageSize` 为每页记录条数，`info` 表示推送计划或者推送标识，只要其中之一匹配到即可（模糊查询），`sendSource` 为创建来源，0 表示 API，1 表示 web（控制台创建），不传或者非 0 和 1 值表示不区分；
//     `page` 和 `pageSize` 不传（为 0）则取默认值（分别为 1 和 50）。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_plan#%E6%9F%A5%E8%AF%A2%E6%8E%A8%E9%80%81%E8%AE%A1%E5%88%92%E5%88%97%E8%A1%A8
func (p *apiv3) ListPlans(ctx context.Context, page, pageSize int, info string, sendSource int) (*PlansListResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		page = 50
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  p.proto,
		URL:    fmt.Sprintf("%s/v3/push_plan/list?page=%d&page_size=%d&info=%s&send_source=%d", p.host, page, pageSize, info, sendSource),
		Auth:   p.auth,
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &PlansListResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type PlansListResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	Data *PlansData `json:"data,omitempty"` // 推送计划列表数据
}

// 推送计划列表数据
type PlansData struct {
	Total    int          `json:"total,omitempty"`     // 总记录数
	Page     int          `json:"page,omitempty"`      // 当前页码
	PageSize int          `json:"page_size,omitempty"` // 每页记录条数
	Detail   []PlanDetail `json:"detail,omitempty"`    // 推送计划详情列表
}

// 推送计划详情
type PlanDetail struct {
	Name         string                `json:"name,omitempty"`           // 计划标识
	Description  string                `json:"description,omitempty"`    // 计划描述
	PushCount    int                   `json:"push_count,omitempty"`     // 推送次数
	CreateTime   jiguang.LocalDateTime `json:"create_time,omitempty"`    // 创建时间
	LastUsedTime jiguang.LocalDateTime `json:"last_used_time,omitempty"` // 最后使用时间
	SendSource   int                   `json:"send_source,omitempty"`    // 创建来源，0 表示 API，1 表示 web（控制台创建）
}

func (rs *PlansListResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
