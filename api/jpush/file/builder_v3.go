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

package file

import (
	"encoding/base64"
	"errors"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// 用于构建和配置 File API v3 访问客户端的构建器。
type APIv3Builder struct {
	client       api.Client
	host         string
	authKey      string
	authSecret   string
	logger       jiguang.Logger
	httpLogLevel api.HttpLogLevel
	err          error
}

func NewAPIv3Builder() *APIv3Builder {
	return &APIv3Builder{
		client:       api.DefaultClient,
		host:         api.HostJPushPushV3,
		logger:       api.DefaultJPushLogger,
		httpLogLevel: api.HttpLogLevelBasic,
	}
}

// 【可选】设置 API 的客户端，用于发送 HTTP 请求，默认为 api.DefaultClient。
func (b *APIv3Builder) SetClient(client api.Client) *APIv3Builder {
	b.client = client
	return b
}

// 【可选】设置 API 的 Host 基础 URL，默认为 api.HostJPushPushV3。
func (b *APIv3Builder) SetHost(host string) *APIv3Builder {
	if host == "" {
		b.err = errors.New("`host` cannot be empty")
	}
	b.host = host
	return b
}

// 【必填】设置 API 的应用标识 `appKey` 或开发者标识 `devKey`。
func (b *APIv3Builder) SetAuthKey(authKey string) *APIv3Builder {
	if authKey == "" {
		b.err = errors.New("`authKey` (`appKey` or `devKey`) cannot be empty")
	}
	b.authKey = authKey
	return b
}

// 【必填】设置 API 的应用主密钥 `masterSecret` 或开发者密钥 `devSecret`。
func (b *APIv3Builder) SetAuthSecret(authSecret string) *APIv3Builder {
	if authSecret == "" {
		b.err = errors.New("`authSecret` (`masterSecret` or `devSecret`) cannot be empty")
	}
	b.authSecret = authSecret
	return b
}

// 【可选】设置 API 的日志记录器，默认为 api.DefaultJPushLogger。
func (b *APIv3Builder) SetLogger(logger jiguang.Logger) *APIv3Builder {
	b.logger = logger
	return b
}

// 【可选】设置 API 的 HTTP 日志记录级别，用于指定记录 API 的 HTTP 请求和响应的日志信息的详细程度，默认为 api.HttpLogLevelBasic。
//   - 可用的级别：api.HttpLogLevelNone、api.HttpLogLevelBasic、api.HttpLogLevelHeaders、api.HttpLogLevelFull；
//   - 若要禁用 HTTP 日志记录，可使用 DisableHttpLogging。
func (b *APIv3Builder) SetHttpLogLevel(httpLogLevel api.HttpLogLevel) *APIv3Builder {
	b.httpLogLevel = httpLogLevel
	return b
}

// 【可选】禁用 API 的 HTTP 日志记录，即不记录 API 的 HTTP 请求和响应的日志信息。
func (b *APIv3Builder) DisableHttpLogging() *APIv3Builder {
	b.httpLogLevel = api.HttpLogLevelNone
	return b
}

func (b *APIv3Builder) Build() (APIv3, error) {
	if b.err != nil {
		return (*apiv3)(nil), b.err
	}
	if b.authKey == "" || b.authSecret == "" {
		return (*apiv3)(nil), errors.New("both `authKey` (`appKey`/`devKey`) and `authSecret` (`masterSecret`/`devSecret`) cannot be empty")
	}

	client := api.NewHttpClient(b.client, b.logger, b.httpLogLevel)
	proto := client.DetectProto(b.host)
	creds := base64.StdEncoding.EncodeToString([]byte(b.authKey + ":" + b.authSecret))

	return &apiv3{
		client: client,
		proto:  proto,
		host:   b.host,
		auth:   "Basic " + creds,
	}, nil
}

// apiv3 内部实现了 APIv3，是 File API v3 的默认访问客户端。
type apiv3 struct {
	client api.HttpClient
	proto  string
	host   string
	auth   string
}
