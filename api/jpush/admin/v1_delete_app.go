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

// 删除极光 APP
//  - 功能说明：删除开发者账号下的指定 APP。
//  - 调用地址：POST `/v1/app/{appKey}/delete`，appKey 为 APP 的唯一标识。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#app-%E5%88%A0%E9%99%A4
func (a *apiv1) DeleteApp(ctx context.Context, appKey string) (*AppDeleteResult, error) {
	if a == nil {
		return nil, api.ErrNilJPushAdminAPIv1
	}

	if appKey == "" {
		return nil, errors.New("`appKey` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  a.proto,
		URL:    a.host + "/v1/app/" + appKey + "/delete",
		Auth:   a.auth,
	}
	resp, err := a.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AppDeleteResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AppDeleteResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Success       string         `json:"success,omitempty"`
}

func (rs *AppDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
