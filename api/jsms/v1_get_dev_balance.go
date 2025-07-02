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

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 账号余量查询
//
// 使用 devKey:devSecret 认证。
//   - 功能说明：查询账号余量，账号余量指未分配给某个应用，属于账号共享的短信余量。
//   - 调用地址：GET `/v1/accounts/dev`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#%E8%B4%A6%E5%8F%B7%E4%BD%99%E9%87%8F%E6%9F%A5%E8%AF%A2-api
func (s *apiv1) GetDevBalance(ctx context.Context) (*DevBalanceGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}
	if s.devAuth == "" {
		return nil, errors.New("please set the `devKey` and `devSecret` required for this API")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/accounts/dev",
		Auth:   s.devAuth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &DevBalanceGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type DevBalanceGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	DevBalance    *uint64        `json:"dev_balance,omitempty"`  // 全类型短信余量
	DevVoice      *uint64        `json:"dev_voice,omitempty"`    // 语音短信余量
	DevIndustry   *uint64        `json:"dev_industry,omitempty"` // 行业短信余量
	DevMarket     *uint64        `json:"dev_market,omitempty"`   // 营销短信余量
}

func (rs *DevBalanceGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
