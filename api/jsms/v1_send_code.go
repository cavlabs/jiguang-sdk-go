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

// # 发送文本验证码短信
//   - 功能说明：发送文本验证码短信。
//   - 调用地址：POST `/v1/codes`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E6%96%87%E6%9C%AC%E9%AA%8C%E8%AF%81%E7%A0%81%E7%9F%AD%E4%BF%A1-api
func (s *apiv1) SendCode(ctx context.Context, param *CodeSendParam) (*CodeSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/codes",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &CodeSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 文本验证码短信发送参数
type CodeSendParam struct {
	Mobile string `json:"mobile"`            // 【必填】手机号码
	SignID int    `json:"sign_id,omitempty"` // 【可选】签名 ID，该字段为空则使用应用默认签名
	TempID int64  `json:"temp_id"`           // 【必填】模板 ID
}

// 文本验证码短信发送结果
type CodeSendResult = sendResult
