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

package api

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"

	sdk "github.com/calvinit/jiguang-sdk-go"
)

// 极光 REST API 的 HTTP 请求原始请求封装。
type Request struct {
	Method string      // 请求方法
	Proto  string      // 协议版本，如 "HTTP/1.0"、"HTTP/1.1"、"HTTP/2.0" 等。
	URL    string      // 请求完整 URL
	Auth   string      // 请求授权信息
	Header http.Header // 自定义请求头
	Body   interface{} // 请求正文负载
}

// defaultUserAgent 是默认的用户代理字符串，用于携带的请求头 `User-Agent` 标识。
const defaultUserAgent = "calvinit/jiguang-sdk-go@" + sdk.Version

// newApplicationJSONRequest 创建一个带有 JSON 正文负载的 HTTP 请求。
func newApplicationJSONRequest(ctx context.Context, req *Request) (*http.Request, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if req.Proto != "" {
		httpReq.Proto = req.Proto
	}

	for k, v := range req.Header {
		httpReq.Header[k] = v
	}

	httpReq.Header.Set("Authorization", req.Auth)
	httpReq.Header.Set("Content-Type", "application/json;charset=UTF-8")
	httpReq.Header.Set("User-Agent", defaultUserAgent)

	return httpReq, nil
}

// newMultipartFormDataRequest 创建一个带有多部分表单数据正文负载的 HTTP 请求。
func newMultipartFormDataRequest(ctx context.Context, req *Request) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := req.Body.(MultipartFormDataBody).Prepare(writer); err != nil {
		return nil, err
	}

	_ = writer.Close()

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.URL, body)
	if err != nil {
		return nil, err
	}

	if req.Proto != "" {
		httpReq.Proto = req.Proto
	}

	for k, v := range req.Header {
		httpReq.Header[k] = v
	}

	httpReq.Header.Set("Authorization", req.Auth)
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("User-Agent", defaultUserAgent)

	return httpReq, nil
}
