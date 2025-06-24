/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 查询签名
//   - 功能说明：查询短信签名。
//   - 调用地址：GET `/v1/sign/{signID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E6%9F%A5%E8%AF%A2%E7%AD%BE%E5%90%8Dapi
func (s *apiv1) GetSign(ctx context.Context, signID int) (*SignGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if signID < 1 {
		return nil, errors.New("invalid `signID`, required: >0")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/sign/" + strconv.Itoa(signID),
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SignGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SignGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	SignID        int            `json:"sign_id,omitempty"`    // 签名 ID
	Sign          string         `json:"sign,omitempty"`       // 签名内容
	Status        int            `json:"status,omitempty"`     // 签名审核状态：1 - 审核通过；2 - 审核不通过
	IsDefault     bool           `json:"is_default,omitempty"` // 签名使用状态：是否使用默认签名
	UseStatus     bool           `json:"use_status,omitempty"` // 签名使用状态：使用中/未使用
}

func (rs *SignGetResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `SignGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias SignGetResult
	aux := struct {
		IsDefault *int `json:"is_default,omitempty"`
		UseStatus *int `json:"use_status,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(rs),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.IsDefault != nil && *aux.IsDefault == 1 {
		rs.IsDefault = true
	}
	if aux.UseStatus != nil && *aux.UseStatus == 1 {
		rs.UseStatus = true
	}

	return nil
}

func (rs SignGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	}
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `SignGetResult` 的方法，包括 `MarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias SignGetResult
	aux := struct {
		IsDefault int `json:"is_default"`
		UseStatus int `json:"use_status"`
		Alias
	}{
		Alias: (Alias)(rs),
	}

	if rs.IsDefault {
		aux.IsDefault = 1
	}
	if rs.UseStatus {
		aux.UseStatus = 1
	}

	return json.Marshal(&aux)
}

func (rs *SignGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
