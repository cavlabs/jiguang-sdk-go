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

package jums

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 批量删除用户信息
//  - 功能说明：本 API 将删除用户的唯一 ID 及其所绑定的所有信息，请谨慎操作。该操作必须使用全局 accessKey 进行鉴权。
//  - 调用地址：POST `/v1/user/delete`，userIDs 为要批量删除的用户的唯一标识列表。
//  - 接口文档：https://docs.jiguang.cn/jums/server/rest_api_jums_user#%E6%89%B9%E9%87%8F%E5%88%A0%E9%99%A4-api
func (u *apiv1) BatchDeleteUsers(ctx context.Context, userIDs []string) (*UsersBatchDeleteResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if u.accessAuth == "" {
		return nil, errors.New("please set the `accessKey` and `accessMasterSecret` required for this API")
	}

	if len(userIDs) == 0 {
		return nil, errors.New("`userIDs` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  u.proto,
		URL:    u.host + "/v1/user/delete",
		Auth:   u.accessAuth,
		Body:   userIDs,
	}
	resp, err := u.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &UsersBatchDeleteResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UsersBatchDeleteResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	Data interface{} `json:"data,omitempty"` // [string / UsersBatchDeleteData] 删除成功/失败的详细数据，当请求失败时，数据为空。
}

// 当删除成功或部分成功时，将返回成功/失败的详细数据
type UsersBatchDeleteData struct {
	Success []string                   `json:"success"` // 删除成功的数据
	Fail    []UsersBatchDeleteFailData `json:"fail"`    // 删除失败的数据
}

// 删除失败的详细数据
type UsersBatchDeleteFailData struct {
	UserID  string `json:"userid,omitempty"`  // 用户的唯一标识。
	ErrCode string `json:"errcode,omitempty"` // 当有失败数据时，返回具体的失败错误码。
	ErrMsg  string `json:"errmsg,omitempty"`  // 当有失败数据时，返回具体的失败错误原因。
}

func (rs *UsersBatchDeleteResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `UsersBatchDeleteResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias UsersBatchDeleteResult
	aux := &struct {
		Data json.RawMessage `json:"data,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(rs),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	rs.CodeError = aux.CodeError

	if aux.Data == nil {
		return nil
	}

	var strData string
	if err := json.Unmarshal(aux.Data, &strData); err == nil {
		rs.Data = strData
		return nil
	} else {
		objData := &UsersBatchDeleteData{}
		if err = json.Unmarshal(aux.Data, &objData); err == nil {
			if objData != nil { // omitempty
				rs.Data = objData
			}
		} else {
			// return err
			rs.Data = string(aux.Data)
		}
	}

	return nil
}

func (rs *UsersBatchDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
