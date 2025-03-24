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

// # 撤回消息
//   - 功能说明：撤回消息发送，[撤回效果说明]。
//   - 调用地址：POST `/v1/retract/{msgID}`，`msgID` 为请求发送消息成功后返回的 ID 值，一次仅允许传 1 个 `msgID`。
//   - 接口文档：[docs.jiguang.cn]
//
// [撤回效果说明]: https://docs.jiguang.cn/jums/server/rest_api_jums_retract_message#%E6%92%A4%E5%9B%9E%E6%95%88%E6%9E%9C%E8%AF%B4%E6%98%8E
// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_template_message
func (u *apiv1) Retract(ctx context.Context, msgID string) (*RetractResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if msgID == "" {
		return nil, errors.New("`msgID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  u.proto,
		URL:    u.host + "/v1/retract/" + msgID,
		Auth:   u.auth,
	}
	resp, err := u.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &RetractResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RetractResult struct {
	*api.Response `json:"-"`
	*api.CodeError
}

func (rs *RetractResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
