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
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 验证码验证
//   - 功能说明：验证验证码是否有效。
//   - 调用地址：POST `/v1/codes/{msgID}/valid`，其中 `msgID` 为发送验证码短信时返回的消息 ID，`code` 为用户输入的验证码。
//   - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E9%AA%8C%E8%AF%81%E7%A0%81%E9%AA%8C%E8%AF%81-api
func (s *apiv1) VerifyCode(ctx context.Context, msgID, code string) (*CodeVerifyResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if msgID == "" {
		return nil, errors.New("`msgID` cannot be empty")
	}
	if code == "" {
		return nil, errors.New("`code` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/codes/" + msgID + "/valid",
		Auth:   s.auth,
		Body:   &codeVerifyParam{Code: code},
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &CodeVerifyResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 验证码验证参数
type codeVerifyParam struct {
	Code string `json:"code"` // 【必填】验证码
}

// 验证码验证结果
type CodeVerifyResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	IsValid       bool           `json:"is_valid"` // 验证码是否有效
}

func (rs *CodeVerifyResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
