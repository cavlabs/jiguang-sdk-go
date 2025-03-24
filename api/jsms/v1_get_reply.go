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
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jsms/inquire"
)

// # 获取上行短信消息
//
// SMS_REPLY - 用户回复消息。
//   - 功能说明：获取 appKey 下的上行短信消息；API 调用频率为 5 秒/次，每次调用最多可获取 500 条数据；每条上行短信消息数据最多保留七天，即极光收到运营商返回上行消息后的七天。
//   - 调用地址：GET `/v1/reply`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire#%E8%8E%B7%E5%8F%96%E4%B8%8A%E8%A1%8C%E7%9F%AD%E4%BF%A1%E6%B6%88%E6%81%AF
func (s *apiv1) GetReply(ctx context.Context) (*ReplyGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/reply",
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ReplyGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ReplyGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError      `json:"error,omitempty"`
	Data          []inquire.ReplyData `json:"-"` // 用户回复消息数据
}

func (rs *ReplyGetResult) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' && data[len(data)-1] == ']' {
		var replies []inquire.ReplyData
		if err := json.Unmarshal(data, &replies); err != nil {
			return err
		}
		rs.Data = replies
		return nil
	}
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ReplyGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ReplyGetResult
	aux := (*Alias)(rs)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (rs ReplyGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else {
		return json.Marshal(rs.Data)
	}
}

func (rs *ReplyGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
