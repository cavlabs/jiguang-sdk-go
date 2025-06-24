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

package jsms

import (
	"errors"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jsms/callback"
)

// # 获取配置好的回调接口服务
func (s *apiv1) Callback() (*Callback, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}
	if s.callback == nil {
		return nil, errors.New("JSMS callback server is disabled")
	}
	return s.callback, nil
}

// ↓↓↓ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↓↓↓

type (
	Callback                      = callback.Server                // 回调接口服务核心结构。
	CallbackConfigOption          = callback.ConfigOption          // 回调接口服务配置选项。
	CallbackReplyData             = callback.ReplyData             // 「用户回复消息」SMS_REPLY 回执数据。
	CallbackReportData            = callback.ReportData            // 「短信送达状态」SMS_REPORT 回执数据。
	CallbackTemplateData          = callback.TemplateData          // 「模板审核结果」SMS_TEMPLATE 回执数据。
	CallbackSignData              = callback.SignData              // 「签名审核结果」SMS_SIGN 回执数据。
	CallbackReplyDataProcessor    = callback.ReplyDataProcessor    // 「用户回复消息」SMS_REPLY 回执数据回调处理器。
	CallbackReportDataProcessor   = callback.ReportDataProcessor   // 「短信送达状态」SMS_REPORT 回执数据回调处理器。
	CallbackTemplateDataProcessor = callback.TemplateDataProcessor // 「模板审核结果」SMS_TEMPLATE 回执数据回调处理器。
	CallbackSignDataProcessor     = callback.SignDataProcessor     // 「签名审核结果」SMS_SIGN 回执数据回调处理器。
)

var (
	// 自定义配置回调接口服务监听地址，默认为 ":8088"。
	WithCallbackAddr = callback.WithAddr
	// 自定义配置回调接口服务回调路径，默认为 "/callback"。
	WithCallbackPath = callback.WithPath
	// 自定义配置回调接口服务的日志打印器，默认与 JSMS API v1 使用相同的日志打印器。
	WithCallbackLogger = callback.WithLogger
	// 自定义配置回调接口服务的 HTTP Handler，默认为使用 net/http 实现的一个简单的 Handler。
	WithCallbackHttpHandler = callback.WithHttpHandler
	// 自定义配置「用户回复消息」SMS_REPLY 回执数据回调处理器。注：你的自定义处理器需要实现 CallbackReplyDataProcessor 接口。
	WithCallbackReplyDataProcessor = callback.WithReplyDataProcessor
	// 自定义配置「短信送达状态」SMS_REPORT 回执数据回调处理器。注：你的自定义处理器需要实现 CallbackReportDataProcessor 接口。
	WithCallbackReportDataProcessor = callback.WithReportDataProcessor
	// 自定义配置「模板审核结果」SMS_TEMPLATE 回执数据回调处理器。注：你的自定义处理器需要实现 CallbackTemplateDataProcessor 接口。
	WithCallbackTemplateDataProcessor = callback.WithTemplateDataProcessor
	// 自定义配置「签名审核结果」SMS_SIGN 回执数据回调处理器。注：你的自定义处理器需要实现 CallbackSignDataProcessor 接口。
	WithCallbackSignDataProcessor = callback.WithSignDataProcessor
)

// ↑↑↑ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↑↑↑
