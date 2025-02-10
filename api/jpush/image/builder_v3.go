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

package image

import (
	"encoding/base64"
	"errors"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 用于构建和配置 Image API v3 访问客户端的构建器。
type APIv3Builder struct {
	client       api.Client
	host         string
	appKey       string
	masterSecret string
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

// 【必填】设置 API 的应用标识。
func (b *APIv3Builder) SetAppKey(appKey string) *APIv3Builder {
	if appKey == "" {
		b.err = errors.New("`appKey` cannot be empty")
	}
	b.appKey = appKey
	return b
}

// 【必填】设置 API 的应用主密钥。
func (b *APIv3Builder) SetMasterSecret(masterSecret string) *APIv3Builder {
	if masterSecret == "" {
		b.err = errors.New("`masterSecret` cannot be empty")
	}
	b.masterSecret = masterSecret
	return b
}

// 【可选】设置 API 的日志记录器，默认为 api.DefaultJPushLogger。
func (b *APIv3Builder) SetLogger(logger jiguang.Logger) *APIv3Builder {
	b.logger = logger
	return b
}

// 【可选】设置 API 的 HTTP 日志记录级别，用于指定记录 API 的 HTTP 请求和响应的日志信息的详细程度，默认为 api.HttpLogLevelBasic。
//  - 可用的级别：api.HttpLogLevelNone、api.HttpLogLevelBasic、api.HttpLogLevelHeaders、api.HttpLogLevelFull；
//  - 若要禁用 HTTP 日志记录，可使用 DisableHttpLogging。
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
	if b.appKey == "" || b.masterSecret == "" {
		return (*apiv3)(nil), errors.New("both `appKey` and `masterSecret` cannot be empty")
	}

	client := api.NewHttpClient(b.client, b.logger, b.httpLogLevel)
	proto := client.DetectProto(b.host)
	creds := base64.StdEncoding.EncodeToString([]byte(b.appKey + ":" + b.masterSecret))

	return &apiv3{
		client: client,
		proto:  proto,
		host:   b.host,
		auth:   "Basic " + creds,
	}, nil
}

// apiv3 内部实现了 APIv3，是 Image API v3 的默认访问客户端。
type apiv3 struct {
	client api.HttpClient
	proto  string
	host   string
	auth   string
}
