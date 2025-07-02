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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 创建推送计划（VIP）
//   - 功能说明：API 层面支持按计划推送，具体推送方式参考 [推送计划使用指南]，此接口创建一个推送计划。最多允许创建 1000 个计划。
//   - 调用地址：POST `/v3/push_plan/create`
//   - 接口文档：[docs.jiguang.cn]
//
// [推送计划使用指南]: https://docs.jiguang.cn/jpush/console/config_manage/push_plan#%E9%80%9A%E8%BF%87-api-%E6%8E%A8%E9%80%81
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_plan#%E5%88%9B%E5%BB%BA%E6%8E%A8%E9%80%81%E8%AE%A1%E5%88%92
func (p *apiv3) CreatePlan(ctx context.Context, param *PlanCreateParam) (*PlanCreateResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push_plan/create",
		Auth:   p.auth,
		Body:   param,
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &PlanCreateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type PlanCreateParam struct {
	Name        string `json:"name"`                  // 计划标识
	Description string `json:"description,omitempty"` // 计划描述
}

type PlanCreateResult struct {
	*api.Response `json:"-"`
	*api.CodeError
}

func (rs *PlanCreateResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
