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

package admin

import (
	"encoding/base64"
	"errors"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 用于构建和配置 Admin API v1 访问客户端的构建器。
type APIv1Builder struct {
	client       api.Client
	host         string
	devKey       string
	devSecret    string
	logger       jiguang.Logger
	httpLogLevel api.HttpLogLevel
	err          error
}

func NewAPIv1Builder() *APIv1Builder {
	return &APIv1Builder{
		client:       api.DefaultClient,
		host:         api.HostJPushAdminV1,
		logger:       api.DefaultJPushLogger,
		httpLogLevel: api.HttpLogLevelBasic,
	}
}

// 【可选】设置 API 的客户端，用于发送 HTTP 请求，默认为 api.DefaultClient。
func (b *APIv1Builder) SetClient(client api.Client) *APIv1Builder {
	b.client = client
	return b
}

// 【可选】设置 API 的 Host 基础 URL，默认为 api.HostJPushAdminV1。
func (b *APIv1Builder) SetHost(host string) *APIv1Builder {
	if host == "" {
		b.err = errors.New("`host` cannot be empty")
	}
	b.host = host
	return b
}

// 【必填】设置 API 的开发者标识。
func (b *APIv1Builder) SetDevKey(devKey string) *APIv1Builder {
	if devKey == "" {
		b.err = errors.New("`devKey` cannot be empty")
	}
	b.devKey = devKey
	return b
}

// 【必填】设置 API 的开发者密钥。
func (b *APIv1Builder) SetDevSecret(devSecret string) *APIv1Builder {
	if devSecret == "" {
		b.err = errors.New("`devSecret` cannot be empty")
	}
	b.devSecret = devSecret
	return b
}

// 【可选】设置 API 的日志记录器，默认为 api.DefaultJPushLogger。
func (b *APIv1Builder) SetLogger(logger jiguang.Logger) *APIv1Builder {
	b.logger = logger
	return b
}

// 【可选】设置 API 的 HTTP 日志记录级别，用于指定记录 API 的 HTTP 请求和响应的日志信息的详细程度，默认为 api.HttpLogLevelBasic。
//   - 可用的级别：api.HttpLogLevelNone、api.HttpLogLevelBasic、api.HttpLogLevelHeaders、api.HttpLogLevelFull；
//   - 若要禁用 HTTP 日志记录，可使用 DisableHttpLogging。
func (b *APIv1Builder) SetHttpLogLevel(httpLogLevel api.HttpLogLevel) *APIv1Builder {
	b.httpLogLevel = httpLogLevel
	return b
}

// 【可选】禁用 API 的 HTTP 日志记录，即不记录 API 的 HTTP 请求和响应的日志信息。
func (b *APIv1Builder) DisableHttpLogging() *APIv1Builder {
	b.httpLogLevel = api.HttpLogLevelNone
	return b
}

func (b *APIv1Builder) Build() (APIv1, error) {
	if b.err != nil {
		return (*apiv1)(nil), b.err
	}
	if b.devKey == "" || b.devSecret == "" {
		return (*apiv1)(nil), errors.New("both `devKey` and `devSecret` cannot be empty")
	}

	client := api.NewHttpClient(b.client, b.logger, b.httpLogLevel)
	proto := client.DetectProto(b.host)
	creds := base64.StdEncoding.EncodeToString([]byte(b.devKey + ":" + b.devSecret))

	return &apiv1{
		client: client,
		proto:  proto,
		host:   b.host,
		auth:   "Basic " + creds,
	}, nil
}

// apiv1 内部实现了 APIv1，是 Admin API v1 的默认访问客户端。
type apiv1 struct {
	client api.HttpClient
	proto  string
	host   string
	auth   string
}
