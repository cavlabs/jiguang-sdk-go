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
	"log"

	"github.com/cavlabs/jiguang-sdk-go/api/jsms"
)

type callbackReplyDataProcessor struct{}

func (p callbackReplyDataProcessor) Process(data jsms.CallbackReplyData) {
	// 请在这里处理「用户回复消息」SMS_REPLY 回执数据。
	log.Printf("| JSMS | SMS_REPLY     >>> %s", data.RawJSON)
}

type callbackReportDataProcessor struct{}

func (p callbackReportDataProcessor) Process(data jsms.CallbackReportData) {
	// 请在这里处理「短信送达状态」SMS_REPORT 回执数据。
	log.Printf("| JSMS | SMS_REPORT    >>> %s", data.RawJSON)
}

type callbackTemplateDataProcessor struct{}

func (p callbackTemplateDataProcessor) Process(data jsms.CallbackTemplateData) {
	// 请在这里处理「模板审核结果」SMS_TEMPLATE 回执数据。
	log.Printf("| JSMS | SMS_TEMPLATE  >>> %s", data.RawJSON)
}

type callbackSignDataProcessor struct{}

func (p callbackSignDataProcessor) Process(data jsms.CallbackSignData) {
	// 请在这里处理「签名审核结果」SMS_SIGN 回执数据。
	log.Printf("| JSMS | SMS_SIGN      >>> %s", data.RawJSON)
}
