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
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 获取通道 Token
//   - 功能说明：获取通道 Token。
//   - 调用地址：GET `/v1/token?type={channelType}`
//   - 接口文档：[docs.jiguang.cn]
//
// channelType 为所需要获取的通道的类型，通过渠道鉴权，可以取到本渠道授权的通道的 Token 值，它的取值说明如下：
//   - wechatoa：微信公众号；
//   - wechatmp：微信小程序；
//   - dingtalkcc：钉钉；
//   - wechatwk：企业微信。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_token
func (u *apiv1) GetChannelToken(ctx context.Context, channelType string) (*ChannelTokenGetResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if channelType == "" {
		return nil, errors.New("`channelType` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  u.proto,
		URL:    u.host + "/v1/token?type=" + channelType,
		Auth:   u.auth,
	}
	resp, err := u.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ChannelTokenGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ChannelTokenGetResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	Token     string `json:"token,omitempty"`      // 该通道的 Token 值。
	ExpiresIn int64  `json:"expires_in,omitempty"` // 有效时间，单位：秒。
}

func (rs *ChannelTokenGetResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ChannelTokenGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ChannelTokenGetResult
	aux := struct {
		ExpiresIn json.RawMessage `json:"expires_in,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(rs),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var strExpiresIn string
	if err := json.Unmarshal(aux.ExpiresIn, &strExpiresIn); err == nil {
		if strExpiresIn != "" {
			rs.ExpiresIn, _ = strconv.ParseInt(strExpiresIn, 10, 64)
		}
		return nil
	} else {
		var intExpiresIn int64
		if err = json.Unmarshal(aux.ExpiresIn, &intExpiresIn); err == nil {
			rs.ExpiresIn = intExpiresIn
		} /* else {
			return err
		}*/
	}

	return nil
}

func (rs ChannelTokenGetResult) MarshalJSON() ([]byte, error) {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ChannelTokenGetResult` 的方法，包括 `MarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ChannelTokenGetResult
	aux := struct {
		ExpiresIn string `json:"expires_in"`
		Alias
	}{
		ExpiresIn: strconv.FormatInt(rs.ExpiresIn, 10),
		Alias:     (Alias)(rs),
	}
	return json.Marshal(&aux)
}

func (rs *ChannelTokenGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
