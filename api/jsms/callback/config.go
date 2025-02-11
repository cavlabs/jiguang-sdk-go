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

package callback

import (
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

const (
	defaultAddr = ":8088"     // 默认监听地址
	defaultPath = "/callback" // 默认回调路径
)

// 标志位：是否已经设置了自定义的回执数据回调处理器
const (
	flagReply = 1 << iota
	flagReport
	flagTemplate
	flagSign
)

// ---------------------------------------------------------------------------------------------------------------------

// 回调接口服务配置。
type config struct {
	addr     string                // 监听地址 (如 ":8088")，默认为 ":8088"
	path     string                // 回调路径 (如 "/callback")，默认为 "/callback"
	logger   jiguang.Logger        // 日志打印器，用于记录回调接口服务的日志，默认为 api.DefaultJSmsLogger
	handler  http.Handler          // HTTP Handler，可自定义处理回调请求，默认为使用 net/http 实现的一个简单的 Handler
	flag     int8                  // 标志位，用于标记是否已经设置了自定义的回执数据回调处理器，从低位到高位分别表示：SMS_REPLY、SMS_REPORT、SMS_TEMPLATE、SMS_SIGN
	reply    ReplyDataProcessor    // 「用户回复消息」SMS_REPLY 回执数据回调处理器，为 nil 时不处理
	report   ReportDataProcessor   // 「短信送达状态」SMS_REPORT 回执数据回调处理器，为 nil 时不处理
	template TemplateDataProcessor // 「模板审核结果」SMS_TEMPLATE 回执数据回调处理器，为 nil 时不处理
	sign     SignDataProcessor     // 「签名审核结果」SMS_SIGN 回执数据回调处理器，为 nil 时不处理
}

// ---------------------------------------------------------------------------------------------------------------------

// 回调接口服务配置选项。
type ConfigOption interface {
	apply(*config) error
}

// ---------------------------------------------------------------------------------------------------------------------

// 监听地址配置选项。
type addrOption string

func (o addrOption) apply(c *config) error {
	addr := string(o)
	if addr == "" {
		return errors.New("`addr` cannot be empty")
	}
	c.addr = addr
	return nil
}

// 自定义配置回调接口服务监听地址，默认为 ":8088"。
func WithAddr(addr string) ConfigOption {
	return addrOption(addr)
}

// ---------------------------------------------------------------------------------------------------------------------

// 回调路径配置选项。
type pathOption string

func (o pathOption) apply(c *config) error {
	path := string(o)
	if path == "" {
		return errors.New("`path` cannot be empty")
	}
	c.path = path
	return nil
}

// 自定义配置回调接口服务回调路径，默认为 "/callback"。
func WithPath(path string) ConfigOption {
	return pathOption(path)
}

// ---------------------------------------------------------------------------------------------------------------------

// 日志打印器配置选项。
type loggerOption struct {
	logger jiguang.Logger
}

func (o loggerOption) apply(c *config) error {
	if o.logger == nil {
		return errors.New("`logger` cannot be nil")
	}
	c.logger = o.logger
	return nil
}

// 自定义配置回调接口服务的日志打印器，默认为 api.DefaultJSmsLogger。
func WithLogger(logger jiguang.Logger) ConfigOption {
	return loggerOption{logger}
}

// ---------------------------------------------------------------------------------------------------------------------

// HTTP Handler 配置选项。
type httpHandlerOption struct {
	handler http.Handler
}

func (o httpHandlerOption) apply(c *config) error {
	if o.handler == nil {
		return errors.New("HTTP `handler` cannot be nil")
	}
	c.handler = o.handler
	return nil
}

// 自定义配置回调接口服务的 HTTP Handler，默认为使用 net/http 实现的一个简单的 Handler。
func WithHttpHandler(handler http.Handler) ConfigOption {
	return httpHandlerOption{handler}
}

// ---------------------------------------------------------------------------------------------------------------------

// 「用户回复消息」SMS_REPLY 回执数据回调处理器配置选项。
type replyDataProcessorOption struct {
	reply ReplyDataProcessor
}

func (o replyDataProcessorOption) apply(c *config) error {
	c.reply = o.reply
	c.flag |= flagReply
	return nil
}

// 自定义配置「用户回复消息」SMS_REPLY 回执数据回调处理器。注：你的自定义处理器需要实现 ReplyDataProcessor 接口。
func WithReplyDataProcessor(reply ReplyDataProcessor) ConfigOption {
	return replyDataProcessorOption{reply}
}

// ---------------------------------------------------------------------------------------------------------------------

// 「短信送达状态」SMS_REPORT 回执数据回调处理器配置选项。
type reportDataProcessorOption struct {
	report ReportDataProcessor
}

func (o reportDataProcessorOption) apply(c *config) error {
	c.report = o.report
	c.flag |= flagReport
	return nil
}

// 自定义配置「短信送达状态」SMS_REPORT 回执数据回调处理器。注：你的自定义处理器需要实现 ReportDataProcessor 接口。
func WithReportDataProcessor(report ReportDataProcessor) ConfigOption {
	return reportDataProcessorOption{report}
}

// ---------------------------------------------------------------------------------------------------------------------

// 「模板审核结果」SMS_TEMPLATE 回执数据回调处理器配置选项。
type templateDataProcessorOption struct {
	template TemplateDataProcessor
}

func (o templateDataProcessorOption) apply(c *config) error {
	c.template = o.template
	c.flag |= flagTemplate
	return nil
}

// 自定义配置「模板审核结果」SMS_TEMPLATE 回执数据回调处理器。注：你的自定义处理器需要实现 TemplateDataProcessor 接口。
func WithTemplateDataProcessor(template TemplateDataProcessor) ConfigOption {
	return templateDataProcessorOption{template}
}

// ---------------------------------------------------------------------------------------------------------------------

// 「签名审核结果」SMS_SIGN 回执数据回调处理器配置选项。
type signDataProcessorOption struct {
	sign SignDataProcessor
}

func (o signDataProcessorOption) apply(c *config) error {
	c.sign = o.sign
	c.flag |= flagSign
	return nil
}

// 自定义配置「签名审核结果」SMS_SIGN 回执数据回调处理器。注：你的自定义处理器需要实现 SignDataProcessor 接口。
func WithSignDataProcessor(sign SignDataProcessor) ConfigOption {
	return signDataProcessorOption{sign}
}
