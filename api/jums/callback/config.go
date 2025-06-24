/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

const (
	defaultAddr = ":8089"     // 默认监听地址
	defaultPath = "/callback" // 默认回调路径
)

// 标志位：是否已经设置了自定义的回调数据处理器
const (
	flagTargetValid int16 = 1 << iota
	flagTargetInvalid
	flagSentSucc
	flagSentFail
	flagReceivedSucc
	flagReceivedFail
	flagClick
	flagRetractedSucc
	flagRetractedFail
)

// ---------------------------------------------------------------------------------------------------------------------

// 回调接口服务配置。
type config struct {
	addr          string            // 监听地址 (如 ":8089")，默认为 ":8089"
	path          string            // 回调路径 (如 "/callback")，默认为 "/callback"
	logger        jiguang.Logger    // 日志打印器，用于记录回调接口服务的日志，默认为 api.DefaultJUmsLogger
	checkAuth     bool              // 是否开启安全校验，默认开启
	handler       http.Handler      // HTTP Handler，可自定义处理回调请求，默认为使用 net/http 实现的一个简单的 Handler
	flag          int16             // 标志位，用于标记是否已经设置了自定义的回调数据处理器，从低位到高位分别表示：TargetValid、TargetInvalid、SentSucc、SentFail、ReceivedSucc、ReceivedFail、Click、RetractedSucc、RetractedFail
	targetValid   DataProcessor     // 目标有效 (0) 回调数据处理器，为 nil 时不处理
	targetInvalid DataProcessor     // 目标无效 (1) 回调数据处理器，为 nil 时不处理
	sentSucc      DataProcessor     // 提交成功 (2) 回调数据处理器，为 nil 时不处理
	sentFail      DataProcessor     // 提交失败 (3) 回调数据处理器，为 nil 时不处理
	receivedSucc  DataProcessor     // 送达成功 (4) 回调数据处理器，为 nil 时不处理
	receivedFail  DataProcessor     // 送达失败 (5) 回调数据处理器，为 nil 时不处理
	click         DataProcessor     // 点击 (6) 回调数据处理器，为 nil 时不处理
	retractedSucc DataProcessor     // 撤回成功 (7) 回调数据处理器，为 nil 时不处理
	retractedFail DataProcessor     // 撤回失败 (8) 回调数据处理器，为 nil 时不处理
	unified       DataListProcessor // 统一的回调数据列表处理器，为 nil 时不处理
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

// 自定义配置回调接口服务监听地址，默认为 ":8089"。
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

// 自定义配置回调接口服务的日志打印器，默认为 api.DefaultJUmsLogger。
func WithLogger(logger jiguang.Logger) ConfigOption {
	return loggerOption{logger}
}

// ---------------------------------------------------------------------------------------------------------------------

// 是否开启安全校验配置选项。
type checkAuthOption bool

func (o checkAuthOption) apply(c *config) error {
	c.checkAuth = bool(o)
	return nil
}

// 自定义配置回调接口服务是否开启安全校验，默认开启。
func WithCheckAuth(checkAuth bool) ConfigOption {
	return checkAuthOption(checkAuth)
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

// 目标有效 (0) 回调数据处理器配置选项。
type targetValidDataProcessorOption struct {
	processor DataProcessor
}

func (o targetValidDataProcessorOption) apply(c *config) error {
	c.targetValid = o.processor
	c.flag |= flagTargetValid
	return nil
}

// 自定义配置 目标有效 (0) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithTargetValidDataProcessor(processor DataProcessor) ConfigOption {
	return targetValidDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 目标无效 (1) 回调数据处理器配置选项。
type targetInvalidDataProcessorOption struct {
	processor DataProcessor
}

func (o targetInvalidDataProcessorOption) apply(c *config) error {
	c.targetInvalid = o.processor
	c.flag |= flagTargetInvalid
	return nil
}

// 自定义配置 目标无效 (1) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithTargetInvalidDataProcessor(processor DataProcessor) ConfigOption {
	return targetInvalidDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 提交成功 (2) 回调数据处理器配置选项。
type sentSuccDataProcessorOption struct {
	processor DataProcessor
}

func (o sentSuccDataProcessorOption) apply(c *config) error {
	c.sentSucc = o.processor
	c.flag |= flagSentSucc
	return nil
}

// 自定义配置 提交成功 (2) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithSentSuccDataProcessor(processor DataProcessor) ConfigOption {
	return sentSuccDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 提交失败 (3) 回调数据处理器配置选项。
type sentFailDataProcessorOption struct {
	processor DataProcessor
}

func (o sentFailDataProcessorOption) apply(c *config) error {
	c.sentFail = o.processor
	c.flag |= flagSentFail
	return nil
}

// 自定义配置 提交失败 (3) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithSentFailDataProcessor(processor DataProcessor) ConfigOption {
	return sentFailDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 送达成功 (4) 回调数据处理器配置选项。
type receivedSuccDataProcessorOption struct {
	processor DataProcessor
}

func (o receivedSuccDataProcessorOption) apply(c *config) error {
	c.receivedSucc = o.processor
	c.flag |= flagReceivedSucc
	return nil
}

// 自定义配置 送达成功 (4) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithReceivedSuccDataProcessor(processor DataProcessor) ConfigOption {
	return receivedSuccDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 送达失败 (5) 回调数据处理器配置选项。
type receivedFailDataProcessorOption struct {
	processor DataProcessor
}

func (o receivedFailDataProcessorOption) apply(c *config) error {
	c.receivedFail = o.processor
	c.flag |= flagReceivedFail
	return nil
}

// 自定义配置 送达失败 (5) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithReceivedFailDataProcessor(processor DataProcessor) ConfigOption {
	return receivedFailDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 点击 (6) 回调数据处理器配置选项。
type clickDataProcessorOption struct {
	processor DataProcessor
}

func (o clickDataProcessorOption) apply(c *config) error {
	c.click = o.processor
	c.flag |= flagClick
	return nil
}

// 自定义配置 点击 (6) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithClickDataProcessor(processor DataProcessor) ConfigOption {
	return clickDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 撤回成功 (7) 回调数据处理器配置选项。
type retractedSuccDataProcessorOption struct {
	processor DataProcessor
}

func (o retractedSuccDataProcessorOption) apply(c *config) error {
	c.retractedSucc = o.processor
	c.flag |= flagRetractedSucc
	return nil
}

// 自定义配置 撤回成功 (7) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithRetractedSuccDataProcessor(processor DataProcessor) ConfigOption {
	return retractedSuccDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 撤回失败 (8) 回调数据处理器配置选项。
type retractedFailDataProcessorOption struct {
	processor DataProcessor
}

func (o retractedFailDataProcessorOption) apply(c *config) error {
	c.retractedFail = o.processor
	c.flag |= flagRetractedFail
	return nil
}

// 自定义配置 撤回失败 (8) 回调数据处理器。注：你的自定义处理器需要实现 DataProcessor 接口。
func WithRetractedFailDataProcessor(processor DataProcessor) ConfigOption {
	return retractedFailDataProcessorOption{processor}
}

// ---------------------------------------------------------------------------------------------------------------------

// 统一的回调数据列表处理器配置选项（所有类型的回调数据汇总到此处理器中进行统一处理）。
type dataListProcessorOption struct {
	processor DataListProcessor
}

func (o dataListProcessorOption) apply(c *config) error {
	c.unified = o.processor
	c.flag = 0
	return nil
}

// 统一的回调数据列表处理器（所有类型的回调数据汇总到此处理器中进行统一处理）。注：你的自定义处理器需要实现 DataListProcessor 接口。
//   - 如果你不希望根据不同的回调数据类型使用不同的处理器，可以仅配置使用此选项。
func WithDataListProcessor(processor DataListProcessor) ConfigOption {
	return dataListProcessorOption{processor}
}
