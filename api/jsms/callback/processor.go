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

package callback

import (
	"context"

	"github.com/cavlabs/jiguang-sdk-go/api/jsms/inquire"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// ↓↓↓ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↓↓↓

type (
	ReplyData    = inquire.ReplyData    // 「用户回复消息」SMS_REPLY 回执数据。
	ReportData   = inquire.ReportData   // 「短信送达状态」SMS_REPORT 回执数据。
	TemplateData = inquire.TemplateData // 「模板审核结果」SMS_TEMPLATE 回执数据。
	SignData     = inquire.SignData     // 「签名审核结果」SMS_SIGN 回执数据。
)

// ↑↑↑ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↑↑↑

// ---------------------------------------------------------------------------------------------------------------------

// 「用户回复消息」SMS_REPLY 回执数据回调处理器。
type ReplyDataProcessor interface {
	Process(data ReplyData)
}

// 「短信送达状态」SMS_REPORT 回执数据回调处理器。
type ReportDataProcessor interface {
	Process(data ReportData)
}

// 「模板审核结果」SMS_TEMPLATE 回执数据回调处理器。
type TemplateDataProcessor interface {
	Process(data TemplateData)
}

// 「签名审核结果」SMS_SIGN 回执数据回调处理器。
type SignDataProcessor interface {
	Process(data SignData)
}

// ---------------------------------------------------------------------------------------------------------------------

type loggingDataProcessor struct {
	logger jiguang.Logger
}

type loggingReplyDataProcessor loggingDataProcessor

func (p loggingReplyDataProcessor) Process(data ReplyData) {
	p.logger.Debugf(context.TODO(), "SMS_REPLY    : %s", data.RawJSON)
}

type loggingReportDataProcessor loggingDataProcessor

func (p loggingReportDataProcessor) Process(data ReportData) {
	p.logger.Debugf(context.TODO(), "SMS_REPORT   : %s", data.RawJSON)
}

type loggingTemplateDataProcessor loggingDataProcessor

func (p loggingTemplateDataProcessor) Process(data TemplateData) {
	p.logger.Debugf(context.TODO(), "SMS_TEMPLATE : %s", data.RawJSON)
}

type loggingSignDataProcessor loggingDataProcessor

func (p loggingSignDataProcessor) Process(data SignData) {
	p.logger.Debugf(context.TODO(), "SMS_SIGN     : %s", data.RawJSON)
}
