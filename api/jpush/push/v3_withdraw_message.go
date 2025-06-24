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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 推送撤销
//   - 功能说明：撤销操作首先会从服务端尝试撤销（Android 消息，排队中/发送中状态可以服务端撤销；
//     iOS 消息，排队中状态可以服务端撤销）；其次，针对 Push SDK（JPush Android SDK v3.5.0 及以上和 JPush iOS SDK v3.2.8 及以上），
//     会尝试从设备端撤销已展示但未被点击的消息。
//   - 调用地址：DELETE `/v3/push/{msgID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%8E%A8%E9%80%81%E6%92%A4%E9%94%80-api
func (p *apiv3) WithdrawMessage(ctx context.Context, msgID string) (*WithdrawMessageResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if msgID == "" {
		return nil, errors.New("`msgID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodDelete,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/" + msgID,
		Auth:   p.auth,
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &WithdrawMessageResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type WithdrawMessageResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *WithdrawMessageResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
