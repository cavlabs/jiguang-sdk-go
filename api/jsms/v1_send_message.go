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

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 发送单条模板短信
//   - 功能说明：发送单条模板短信。
//   - 调用地址：POST `/v1/messages`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E5%8D%95%E6%9D%A1%E6%A8%A1%E6%9D%BF%E7%9F%AD%E4%BF%A1-api
func (s *apiv1) SendMessage(ctx context.Context, param *MessageSendParam) (*MessageSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/messages",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &MessageSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 单条模板短信发送参数
type MessageSendParam struct {
	Mobile     string                 `json:"mobile"`              // 【必填】手机号码
	SignID     int                    `json:"sign_id,omitempty"`   // 【可选】签名 ID，该字段为空则使用应用默认签名
	TempID     int64                  `json:"temp_id"`             // 【必填】模板 ID
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 【可选】模板参数，需要替换的参数名和参数值的键值对
}

// 单条模板短信发送结果
type MessageSendResult = sendResult
