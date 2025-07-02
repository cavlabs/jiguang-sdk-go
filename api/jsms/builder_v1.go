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

package jsms

import (
	"encoding/base64"
	"errors"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jsms/callback"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// 用于构建和配置 JSMS API v1 访问客户端的构建器。
type APIv1Builder struct {
	client                api.Client
	host                  string
	appKey                string
	masterSecret          string
	devKey                string
	devSecret             string
	logger                jiguang.Logger
	httpLogLevel          api.HttpLogLevel
	callbackEnabled       bool
	callbackConfigOptions []callback.ConfigOption
	err                   error
}

func NewAPIv1Builder() *APIv1Builder {
	return &APIv1Builder{
		client:       api.DefaultClient,
		host:         api.HostJSmsV1,
		logger:       api.DefaultJSmsLogger,
		httpLogLevel: api.HttpLogLevelBasic,
	}
}

// 【可选】设置 API 的客户端，用于发送 HTTP 请求，默认为 api.DefaultClient。
func (b *APIv1Builder) SetClient(client api.Client) *APIv1Builder {
	b.client = client
	return b
}

// 【可选】设置 API 的 Host 基础 URL，默认为 api.HostJSmsV1。
func (b *APIv1Builder) SetHost(host string) *APIv1Builder {
	if host == "" {
		b.err = errors.New("`host` cannot be empty")
	}
	b.host = host
	return b
}

// 【必填】设置 API 的应用标识。
func (b *APIv1Builder) SetAppKey(appKey string) *APIv1Builder {
	if appKey == "" {
		b.err = errors.New("`appKey` cannot be empty")
	}
	b.appKey = appKey
	return b
}

// 【必填】设置 API 的应用主密钥。
func (b *APIv1Builder) SetMasterSecret(masterSecret string) *APIv1Builder {
	if masterSecret == "" {
		b.err = errors.New("`masterSecret` cannot be empty")
	}
	b.masterSecret = masterSecret
	return b
}

// 【可选】设置 API 的开发者标识。
//   - 当需要同时使用 “账号余量查询” API 接口时，请务必同时设置 `devKey`；
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#http-%E9%AA%8C%E8%AF%81
func (b *APIv1Builder) SetDevKey(devKey string) *APIv1Builder {
	if devKey == "" {
		b.err = errors.New("`devKey` cannot be empty")
	}
	b.devKey = devKey
	return b
}

// 【可选】设置 API 的开发者密钥。
//   - 当需要同时使用 “账号余量查询” API 接口时，请务必同时设置 `devSecret`；
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#http-%E9%AA%8C%E8%AF%81
func (b *APIv1Builder) SetDevSecret(devSecret string) *APIv1Builder {
	if devSecret == "" {
		b.err = errors.New("`devSecret` cannot be empty")
	}
	b.devSecret = devSecret
	return b
}

// 【可选】设置 API 的日志记录器，默认为 api.DefaultJSmsLogger。
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

// 【可选】启用回调接口服务。
func (b *APIv1Builder) EnableCallback(callbackConfigOptions ...CallbackConfigOption) *APIv1Builder {
	b.callbackEnabled = true
	b.callbackConfigOptions = callbackConfigOptions
	return b
}

func (b *APIv1Builder) Build() (APIv1, error) {
	if b.err != nil {
		return (*apiv1)(nil), b.err
	}
	if b.appKey == "" || b.masterSecret == "" {
		return (*apiv1)(nil), errors.New("both `appKey` and `masterSecret` cannot be empty")
	}

	client := api.NewHttpClient(b.client, b.logger, b.httpLogLevel)
	proto := client.DetectProto(b.host)
	creds := base64.StdEncoding.EncodeToString([]byte(b.appKey + ":" + b.masterSecret))

	var devCreds string
	if b.devKey != "" && b.devSecret != "" {
		devCreds = base64.StdEncoding.EncodeToString([]byte(b.devKey + ":" + b.devSecret))
	}

	var (
		srv *Callback
		err error
	)
	if b.callbackEnabled {
		opts := []callback.ConfigOption{callback.WithLogger(b.logger)}
		if len(b.callbackConfigOptions) > 0 {
			opts = append(opts, b.callbackConfigOptions...)
		}
		srv, err = callback.NewServer(b.appKey, b.masterSecret, opts...)
	}

	v1 := &apiv1{
		client:   client,
		proto:    proto,
		host:     b.host,
		auth:     "Basic " + creds,
		callback: srv,
	}
	if devCreds != "" {
		v1.devAuth = "Basic " + devCreds
	}
	return v1, err
}

// apiv1 内部实现了 APIv1，是 JSMS API v1 的默认访问客户端。
type apiv1 struct {
	client   api.HttpClient
	proto    string
	host     string
	auth     string
	devAuth  string
	callback *Callback
}
