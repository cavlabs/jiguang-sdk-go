/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
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

package admin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 创建极光 APP
//   - 功能说明：在开发者账号下创建一个 APP。
//   - 调用地址：POST `/v1/app`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#%E5%88%9B%E5%BB%BA%E6%9E%81%E5%85%89-app
func (a *apiv1) CreateApp(ctx context.Context, param *AppCreateParam) (*AppCreateResult, error) {
	if a == nil {
		return nil, api.ErrNilJPushAdminAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  a.proto,
		URL:    a.host + "/v1/app",
		Auth:   a.auth,
		Body:   param,
	}
	resp, err := a.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AppCreateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AppCreateParam struct {
	AppName     string `json:"app_name"`             // 应用名称
	PackageName string `json:"android_package"`      // 应用包名（Android）
	GroupName   string `json:"group_name,omitempty"` // 应用分组名称
}

type AppCreateResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	AppKey        string         `json:"app_key,omitempty"`         // 应用标识
	PackageName   string         `json:"android_package,omitempty"` // 应用包名（Android）
	IsNewCreated  *bool          `json:"is_new_created,omitempty"`  // 是否新创建的应用
	MasterSecret  string         `json:"master_secret,omitempty"`   // 应用主密钥，用于服务器端 API 调用时与 AppKey 配合使用达到鉴权的目的
}

func (rs *AppCreateResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
