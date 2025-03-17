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
	"encoding/base64"
	"errors"

	"github.com/calvinit/jiguang-sdk-go/api"

	"github.com/calvinit/jiguang-sdk-go/api/jums/callback"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 用于构建和配置 JUMS API v1 访问客户端的构建器。
type APIv1Builder struct {
	client                api.Client
	host                  string
	channelKey            string
	masterSecret          string
	accessKey             string
	accessMasterSecret    string
	logger                jiguang.Logger
	httpLogLevel          api.HttpLogLevel
	callbackEnabled       bool
	callbackConfigOptions []callback.ConfigOption
	err                   error
}

func NewAPIv1Builder() *APIv1Builder {
	return &APIv1Builder{
		client:       api.DefaultClient,
		host:         api.HostJUmsV1,
		logger:       api.DefaultJUmsLogger,
		httpLogLevel: api.HttpLogLevelBasic,
	}
}

// 【可选】设置 API 的客户端，用于发送 HTTP 请求，默认为 api.DefaultClient。
func (b *APIv1Builder) SetClient(client api.Client) *APIv1Builder {
	b.client = client
	return b
}

// 【可选】设置 API 的 Host 基础 URL，默认为 api.HostJUmsV1。
func (b *APIv1Builder) SetHost(host string) *APIv1Builder {
	if host == "" {
		b.err = errors.New("`host` cannot be empty")
	}
	b.host = host
	return b
}

// 【必填】设置 API 的渠道标识。
func (b *APIv1Builder) SetChannelKey(channelKey string) *APIv1Builder {
	if channelKey == "" {
		b.err = errors.New("`channelKey` cannot be empty")
	}
	b.channelKey = channelKey
	return b
}

// 【必填】设置 API 的渠道主密钥。
func (b *APIv1Builder) SetMasterSecret(masterSecret string) *APIv1Builder {
	if masterSecret == "" {
		b.err = errors.New("`masterSecret` cannot be empty")
	}
	b.masterSecret = masterSecret
	return b
}

// 【可选】设置 API 的全局访问标识。
//  - 当需要同时使用 “用户管理” API 接口时，请务必同时设置 `accessKey`；
//  - 参考文档：https://docs.jiguang.cn/jums/server/ums_server_overview#%E9%89%B4%E6%9D%83%E6%96%B9%E5%BC%8F
func (b *APIv1Builder) SetAccessKey(accessKey string) *APIv1Builder {
	if accessKey == "" {
		b.err = errors.New("`accessKey` cannot be empty")
	}
	b.accessKey = accessKey
	return b
}

// 【可选】设置 API 的全局访问主密钥。
//  - 当需要同时使用 “用户管理” API 接口时，请务必同时设置 `accessMasterSecret`；
//  - 参考文档：https://docs.jiguang.cn/jums/server/ums_server_overview#%E9%89%B4%E6%9D%83%E6%96%B9%E5%BC%8F
func (b *APIv1Builder) SetAccessMasterSecret(accessMasterSecret string) *APIv1Builder {
	if accessMasterSecret == "" {
		b.err = errors.New("`accessMasterSecret` cannot be empty")
	}
	b.accessMasterSecret = accessMasterSecret
	return b
}

// 【可选】设置 API 的日志记录器，默认为 api.DefaultJUmsLogger。
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
	if b.channelKey == "" || b.masterSecret == "" {
		return (*apiv1)(nil), errors.New("both `channelKey` and `masterSecret` cannot be empty")
	}

	client := api.NewHttpClient(b.client, b.logger, b.httpLogLevel)
	proto := client.DetectProto(b.host)
	creds := base64.StdEncoding.EncodeToString([]byte(b.channelKey + ":" + b.masterSecret))

	var accessCreds string
	if b.accessKey != "" && b.accessMasterSecret != "" {
		accessCreds = base64.StdEncoding.EncodeToString([]byte(b.accessKey + ":" + b.accessMasterSecret))
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
		srv, err = callback.NewServer(b.channelKey, b.masterSecret, opts...)
	}

	v1 := &apiv1{
		client:   client,
		proto:    proto,
		host:     b.host,
		auth:     "Basic " + creds,
		callback: srv,
	}
	if accessCreds != "" {
		v1.accessAuth = "Basic " + accessCreds
	}
	return v1, err
}

// apiv1 内部实现了 APIv1，是 JUMS API v1 的默认访问客户端。
type apiv1 struct {
	client     api.HttpClient
	proto      string
	host       string
	auth       string
	accessAuth string
	callback   *Callback
}
