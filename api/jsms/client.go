// Copyright 2026 cavlabs/jiguang-sdk-go authors.
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
	"errors"

	"github.com/cavlabs/jiguang-sdk-go/v2/api"
	"github.com/cavlabs/jiguang-sdk-go/v2/jiguang"
)

// Client 是 JSMS 产品入口；默认指向 API v1。
type Client struct {
	APIv1
	V1 APIv1
}

type clientOptions struct {
	httpClient   api.Client
	logger       jiguang.Logger
	httpLogLevel api.HttpLogLevel
	httpLogSet   bool

	appKey       string
	masterSecret string
	devKey       string
	devSecret    string

	callbackEnabled bool
	callbackOpts    []CallbackConfigOption
}

// ClientOption 配置 JSMS Client。
type ClientOption func(*clientOptions)

// WithHTTPClient 设置底层 HTTP 客户端。
func WithHTTPClient(c api.Client) ClientOption {
	return func(o *clientOptions) {
		if c != nil {
			o.httpClient = c
		}
	}
}

// WithLogger 设置日志记录器。
func WithLogger(logger jiguang.Logger) ClientOption {
	return func(o *clientOptions) { o.logger = logger }
}

// WithHttpLogLevel 设置 HTTP 日志级别。
func WithHttpLogLevel(level api.HttpLogLevel) ClientOption {
	return func(o *clientOptions) {
		o.httpLogLevel = level
		o.httpLogSet = true
	}
}

// DisableHttpLogging 关闭 HTTP 日志。
func DisableHttpLogging() ClientOption {
	return WithHttpLogLevel(api.HttpLogLevelNone)
}

// WithApp 配置应用级鉴权。
func WithApp(appKey, masterSecret string) ClientOption {
	return func(o *clientOptions) {
		o.appKey = appKey
		o.masterSecret = masterSecret
	}
}

// WithDev 配置开发者鉴权（账号余量等接口可选）。
func WithDev(devKey, devSecret string) ClientOption {
	return func(o *clientOptions) {
		o.devKey = devKey
		o.devSecret = devSecret
	}
}

// WithCallback 启用回调服务。
func WithCallback(opts ...CallbackConfigOption) ClientOption {
	return func(o *clientOptions) {
		o.callbackEnabled = true
		o.callbackOpts = opts
	}
}

// New 创建 JSMS 产品客户端。未配置 App 凭证时仍可 New，调用 API 会返回 api.ErrNilJSmsAPIv1。
func New(opts ...ClientOption) (*Client, error) {
	o := &clientOptions{
		httpClient:   api.DefaultClient,
		httpLogLevel: api.HttpLogLevelBasic,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	if o.httpClient == nil {
		return nil, errors.New("jsms: http client cannot be nil")
	}

	b := NewAPIv1Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.appKey != "" && o.masterSecret != "" {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	if o.devKey != "" && o.devSecret != "" {
		b.SetDevKey(o.devKey).SetDevSecret(o.devSecret)
	}
	if o.callbackEnabled {
		b.EnableCallback(o.callbackOpts...)
	}
	a, _ := b.Build()
	return &Client{APIv1: a, V1: a}, nil
}
