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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 推送校验
//  - 功能说明：该 API 只用于验证推送调用是否能够成功，与推送 API 的区别在于：不向用户发送任何消息。
//	- 调用地址：POST `/v3/push/validate`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%8E%A8%E9%80%81%E6%A0%A1%E9%AA%8C-api
func (p *apiv3) ValidateSend(ctx context.Context, param *SendParam) (*SendResult, error) {
	return p.ValidateCustomSend(ctx, param)
}

// 自定义推送校验：如果遇到 ValidateSend 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
func (p *apiv3) ValidateCustomSend(ctx context.Context, param interface{}) (*SendResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/validate",
		Auth:   p.auth,
		Body:   param,
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
