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
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 应用余量查询 (appKey:masterSecret)
//  - 功能说明：查询应用余量，应用余量指分配给某个应用单独使用的短信余量。
//  - 调用地址：GET `/v1/accounts/app`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#%E5%BA%94%E7%94%A8%E4%BD%99%E9%87%8F%E6%9F%A5%E8%AF%A2-api
func (s *apiv1) GetAppBalance(ctx context.Context) (*AppBalanceGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/accounts/app",
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AppBalanceGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AppBalanceGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	AppBalance    *uint64        `json:"app_balance,omitempty"`  // 全类型短信余量
	AppVoice      *uint64        `json:"app_voice,omitempty"`    // 语音短信余量
	AppIndustry   *uint64        `json:"app_industry,omitempty"` // 行业短信余量
	AppMarket     *uint64        `json:"app_market,omitempty"`   // 营销短信余量
}

func (rs *AppBalanceGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
