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
	"encoding/json"
	"log"

	"github.com/calvinit/jiguang-sdk-go/api/jums"
)

type callbackTargetValidDataProcessor struct{}

func (p callbackTargetValidDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「目标有效」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「目标有效」回调      >>> %s", s)
}

type callbackTargetInvalidDataProcessor struct{}

func (p callbackTargetInvalidDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「目标无效」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「目标无效」回调      >>> %s", s)
}

type callbackSentSuccDataProcessor struct{}

func (p callbackSentSuccDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「提交成功」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「提交成功」回调      >>> %s", s)
}

type callbackSentFailDataProcessor struct{}

func (p callbackSentFailDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「提交失败」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「提交失败」回调      >>> %s", s)
}

type callbackReceivedSuccDataProcessor struct{}

func (p callbackReceivedSuccDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「送达成功」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「送达成功」回调      >>> %s", s)
}

type callbackReceivedFailDataProcessor struct{}

func (p callbackReceivedFailDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「送达失败」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「送达失败」回调      >>> %s", s)
}

type callbackClickDataProcessor struct{}

func (p callbackClickDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「点击」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「　点击　」回调      >>> %s", s)
}

type callbackRetractedSuccDataProcessor struct{}

func (p callbackRetractedSuccDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「撤回成功」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「撤回成功」回调      >>> %s", s)
}

type callbackRetractedFailDataProcessor struct{}

func (p callbackRetractedFailDataProcessor) Process(data jums.CallbackData) {
	// 请在这里处理「撤回失败」回调数据。
	s, _ := json.Marshal(data)
	log.Printf("| JUMS |「撤回失败」回调      >>> %s", s)
}

// ---------------------------------------------------------------------------------------------------------------------

type callbackDataListProcessor struct{}

func (p callbackDataListProcessor) Process(list jums.CallbackDataList) {
	log.Printf("| JUMS | 收到回调消息        --> %s", list.RawJSON)
	for _, data := range list.Items {
		// 请在这里统一处理回调数据。
		s, _ := json.Marshal(data)
		if data.Type == jums.CallbackTypeClick {
			log.Printf("| JUMS |「　点击　」回调      >>> %s", s) // Tip: 这里仅仅是为了打印对齐好看而已！
		} else {
			log.Printf("| JUMS |「%s」回调      >>> %s", data.Type, s)
		}
	}
}
