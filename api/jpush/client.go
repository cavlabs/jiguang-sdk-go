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

package jpush

import (
	"errors"

	"github.com/cavlabs/jiguang-sdk-go/v2/api"
	"github.com/cavlabs/jiguang-sdk-go/v2/jiguang"
)

// Client 聚合 JPush 各 REST 能力。
type Client struct {
	Push     Push
	Device   Device
	File     File
	Image    Image
	Schedule Schedule
	Report   Report
	GPush    GPush
	GReport  GReport
	Admin    Admin
}

type options struct {
	httpClient   api.Client
	logger       jiguang.Logger
	httpLogLevel api.HttpLogLevel
	httpLogSet   bool

	appKey       string
	masterSecret string

	groupKey          string
	groupMasterSecret string

	devKey    string
	devSecret string
}

// Option 配置 Client。
type Option func(*options)

// WithHTTPClient 设置底层 HTTP 客户端，默认 api.DefaultClient。
func WithHTTPClient(c api.Client) Option {
	return func(o *options) {
		if c != nil {
			o.httpClient = c
		}
	}
}

// WithLogger 设置日志记录器。
func WithLogger(logger jiguang.Logger) Option {
	return func(o *options) { o.logger = logger }
}

// WithHttpLogLevel 设置 HTTP 日志级别，默认 api.HttpLogLevelBasic。
func WithHttpLogLevel(level api.HttpLogLevel) Option {
	return func(o *options) {
		o.httpLogLevel = level
		o.httpLogSet = true
	}
}

// DisableHttpLogging 关闭 HTTP 日志。
func DisableHttpLogging() Option {
	return WithHttpLogLevel(api.HttpLogLevelNone)
}

// WithApp 配置应用级鉴权（Push / Device / File / Image / Schedule / Report 等）。
func WithApp(appKey, masterSecret string) Option {
	return func(o *options) {
		o.appKey = appKey
		o.masterSecret = masterSecret
	}
}

// WithGroup 配置分组鉴权（GPush / GReport）。
func WithGroup(groupKey, groupMasterSecret string) Option {
	return func(o *options) {
		o.groupKey = groupKey
		o.groupMasterSecret = groupMasterSecret
	}
}

// WithDev 配置开发者鉴权（Admin；GPush 文件能力可选）。
func WithDev(devKey, devSecret string) Option {
	return func(o *options) {
		o.devKey = devKey
		o.devSecret = devSecret
	}
}

// New 创建 JPush 产品客户端。缺少某类凭证不会导致 New 失败。
func New(opts ...Option) (*Client, error) {
	o := &options{
		httpClient:   api.DefaultClient,
		httpLogLevel: api.HttpLogLevelBasic,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	if o.httpClient == nil {
		return nil, errors.New("jpush: http client cannot be nil")
	}
	return &Client{
		Push:     newPush(o),
		Device:   newDevice(o),
		File:     newFile(o),
		Image:    newImage(o),
		Schedule: newSchedule(o),
		Report:   newReport(o),
		GPush:    newGPush(o),
		GReport:  newGReport(o),
		Admin:    newAdmin(o),
	}, nil
}

func (o *options) hasApp() bool {
	return o.appKey != "" && o.masterSecret != ""
}

func (o *options) hasGroup() bool {
	return o.groupKey != "" && o.groupMasterSecret != ""
}

func (o *options) hasDev() bool {
	return o.devKey != "" && o.devSecret != ""
}
