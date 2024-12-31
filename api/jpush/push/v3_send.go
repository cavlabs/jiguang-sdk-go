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

// 普通推送
//  - 功能说明：向某单个设备或者某设备列表推送一条通知或者消息。推送的内容只能是 JSON 表示的一个推送对象。
//	- 调用地址：POST `/v3/push`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push
func (p *apiv3) Send(ctx context.Context, param *SendParam) (*SendResult, error) {
	return p.CustomSend(ctx, param)
}

// 自定义推送：如果遇到 Send 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
func (p *apiv3) CustomSend(ctx context.Context, param interface{}) (*SendResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push",
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

type SendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	MsgID         string         `json:"msg_id,omitempty"` // 推送消息 ID
	SendNo        string         `json:"sendno,omitempty"` // 推送序号
}

func (rs *SendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
