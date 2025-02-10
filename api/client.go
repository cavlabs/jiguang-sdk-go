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
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// Client 定义了一个用于发送 HTTP 请求的客户端接口。它被设计为与标准库的 http.Client 相兼容。
type Client interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// ---------------------------------------------------------------------------------------------------------------------

// HttpClient 定义了一个发送带有 JSON 或多部分表单数据正文负载的 HTTP 请求的接口。
type HttpClient interface {
	// 探测给定 URL 对应服务器支持的 HTTP 协议版本，如 "HTTP/1.0"、"HTTP/1.1"、"HTTP/2.0" 等。
	DetectProto(url string) string

	// 使用 JSON 正文 `Content-Type: application/json;charset=UTF-8` 发送 HTTP 请求。
	Request(ctx context.Context, req *Request) (resp *Response, err error)

	// 使用多部分表单数据正文 `Content-Type: multipart/form-data; boundary=...` 发送 HTTP 请求。
	FormRequest(ctx context.Context, req *Request) (resp *Response, err error)
}

// ---------------------------------------------------------------------------------------------------------------------

// HttpClient 接口的内部默认实现，可记录下 HTTP 请求和响应的日志信息。
type loggingHttpClient struct {
	client     Client
	httpLogger *httpLogger
}

// 在 API 没有提供自定义 Client 时使用 DefaultClient。
var DefaultClient Client = &http.Client{Timeout: 30 * time.Second}

func NewHttpClient(client Client, logger jiguang.Logger, level HttpLogLevel) HttpClient {
	lc := loggingHttpClient{}
	if client == nil {
		lc.client = DefaultClient
	} else {
		lc.client = client
	}
	if logger == nil {
		logger = jiguang.NewStdLogger()
	}
	if !level.IsValid() {
		level = HttpLogLevelNone
	}
	lc.httpLogger = newHttpLogger(logger, level)
	return &lc
}

// 探测给定 URL 对应服务器支持的 HTTP 协议版本，如 "HTTP/1.0"、"HTTP/1.1"、"HTTP/2.0" 等。
func (lc *loggingHttpClient) DetectProto(url string) string {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return "HTTP/1.1"
	}
	resp, err := lc.client.Do(req)
	if err != nil {
		return "HTTP/1.1"
	}
	return resp.Proto
}

// 使用 JSON 正文 `Content-Type: application/json;charset=UTF-8` 发送 HTTP 请求。
func (lc *loggingHttpClient) Request(ctx context.Context, req *Request) (resp *Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	applicationJSONRequest, err := newApplicationJSONRequest(ctx, req)
	if err != nil {
		return
	}

	return lc.doRequest(ctx, applicationJSONRequest)
}

// 使用多部分表单数据正文 `Content-Type: multipart/form-data; boundary=...` 发送 HTTP 请求。
func (lc *loggingHttpClient) FormRequest(ctx context.Context, req *Request) (resp *Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	multipartFormDataRequest, err := newMultipartFormDataRequest(ctx, req)
	if err != nil {
		return
	}

	return lc.doRequest(ctx, multipartFormDataRequest)
}

// 执行 HTTP 请求，并根据适当的日志记录级别记录下请求和响应的日志信息。
func (lc *loggingHttpClient) doRequest(ctx context.Context, httpReq *http.Request) (resp *Response, err error) {
	startTime := time.Now()
	lc.logRequest(ctx, httpReq)

	httpResp, err := lc.client.Do(httpReq)
	if err != nil {
		// 如果遇到了错误，并且上下文已经被取消，那么上下文的错误可能更有用。
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}
		return
	}

	lc.logResponse(ctx, httpResp, startTime)

	defer func() { err = httpResp.Body.Close() }()

	rawBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	resp = &Response{
		StatusCode: httpResp.StatusCode,
		Header:     httpResp.Header,
		RawBody:    rawBody,
		Rate:       ParseRate(httpResp),
	}
	return
}

// 根据适当的日志记录级别输出相应的 HTTP 请求日志。
func (lc *loggingHttpClient) logRequest(ctx context.Context, req *http.Request) {
	if lc.httpLogger.Level <= HttpLogLevelNone {
		return
	}

	if lc.httpLogger.Level >= HttpLogLevelBasic {
		lc.httpLogger.Debugf(ctx, "---> %s %s %s", req.Method, req.URL, req.Proto)
	}

	if lc.httpLogger.Level >= HttpLogLevelHeaders {
		for name, values := range req.Header {
			if name == "Authorization" {
				lc.httpLogger.Debugf(ctx, "%s: %s", name, "[REDACTED]")
				continue
			}
			for _, value := range values {
				lc.httpLogger.Debugf(ctx, "%s: %s", name, value)
			}
		}
	}

	if lc.httpLogger.Level == HttpLogLevelFull {
		if req.Body == nil {
			lc.httpLogger.Debug(ctx, "<no content>")
		} else {
			body, _ := io.ReadAll(req.Body)
			req.Body = io.NopCloser(bytes.NewBuffer(body)) // 重置读取后的 body，以便后续读取

			// "" or null
			if len(body) == 0 || (len(body) == 4 && body[0] == 'n' && body[1] == 'u' && body[2] == 'l' && body[3] == 'l') {
				lc.httpLogger.Debug(ctx, "<no content>")
			} else {
				contentType := req.Header.Get("Content-Type")
				if strings.Contains(contentType, "json") || strings.Contains(contentType, "text") {
					lc.httpLogger.Debugf(ctx, "%s", body)
				} else {
					lc.httpLogger.Debugf(ctx, "<binary> len %d", len(body))
				}
			}
		}
	}
}

// 根据适当的日志记录级别输出相应的 HTTP 响应日志。
func (lc *loggingHttpClient) logResponse(ctx context.Context, resp *http.Response, startTime time.Time) {
	if lc.httpLogger.Level <= HttpLogLevelNone {
		return
	}

	statusCode := resp.StatusCode
	if lc.httpLogger.Level >= HttpLogLevelBasic {
		lc.httpLogger.Debugf(ctx, "<--- %s %d %s (%v)", resp.Proto, statusCode, http.StatusText(statusCode), time.Since(startTime))
	}

	if lc.httpLogger.Level >= HttpLogLevelHeaders {
		for name, values := range resp.Header {
			for _, value := range values {
				lc.httpLogger.Debugf(ctx, "%s: %s", name, value)
			}
		}
	}

	if lc.httpLogger.Level == HttpLogLevelFull {
		if resp.Body == nil {
			lc.httpLogger.Debug(ctx, "<no content>")
		} else {
			body, _ := io.ReadAll(resp.Body)
			resp.Body = io.NopCloser(bytes.NewBuffer(body)) // 重置读取后的 body，以便后续读取

			if resp.StatusCode == http.StatusNoContent || len(body) == 0 {
				lc.httpLogger.Debug(ctx, "<no content>")
			} else {
				contentType := resp.Header.Get("Content-Type")
				if strings.Contains(contentType, "json") || strings.Contains(contentType, "text") {
					lc.httpLogger.Debugf(ctx, "%s", body)
				} else {
					lc.httpLogger.Debugf(ctx, "<binary> len %d", len(body))
				}
			}
		}
	}
}
