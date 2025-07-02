// Copyright 2025 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"net/http"
	"strings"
)

// 极光 REST API 的 HTTP 请求原始响应封装。
type Response struct {
	StatusCode int         `json:"-"` // 状态码
	Header     http.Header `json:"-"` // 响应头部信息
	RawBody    []byte      `json:"-"` // 原始响应正文
	Rate       Rate        `json:"-"` // API 频率控制信息
}

// 判断响应是否无正文内容。
func (resp *Response) IsNoContent() bool {
	return resp.StatusCode == http.StatusNoContent || len(resp.RawBody) == 0
}

func (resp *Response) String() string {
	contentType := resp.Header.Get("Content-Type")
	var body string
	if resp.IsNoContent() {
		body = "<no content>"
	} else {
		if strings.Contains(contentType, "json") || strings.Contains(contentType, "text") {
			body = string(resp.RawBody)
		} else {
			body = fmt.Sprintf("<binary> len %d", len(resp.RawBody))
		}
	}
	return fmt.Sprintf("StatusCode: %d, Header: %v, Content-Type: %s, Body: %v",
		resp.StatusCode, resp.Header, contentType, body)
}

// 当前在一个时间窗口内可调用次数。如果未设置或无法解析，则为 0。
func (resp *Response) RateLimit() int {
	return resp.Rate.Limit
}

// 当前的时间窗口剩余的可用次数。如果未设置或无法解析，则为 0。
func (resp *Response) RateRemaining() int {
	return resp.Rate.Remaining
}

// 当前距离时间窗口重置剩余的秒数。如果未设置或无法解析，则为 0。
func (resp *Response) RateReset() int64 {
	return resp.Rate.Reset
}
