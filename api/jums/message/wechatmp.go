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

package message

// [微信小程序] 模板消息。
type Wechatmp struct {
	// 【必填】模板 ID。
	TemplateID string `json:"template_id"`
	// 【可选】点击模板卡片后的跳转页面，仅限本小程序内的页面。
	//  - 支持带参数，示例：index?foo=bar
	//  - 该字段不填则模板无跳转。
	Page string `json:"page,omitempty"`
	// 【必填】模板内容。
	//  - 格式形如：{"key1":{"value": any},"key2":{"value":any}}
	Data map[string]interface{} `json:"data"`
	// 【可选】跳转小程序类型，developer 为开发版、trial 为体验版、formal 为正式版，默认为正式版。
	MiniProgramState string `json:"miniprogram_state,omitempty"`
	// 【可选】进入小程序查看的语言类型，支持 zh_CN (简体中文)、en_US (英文)、zh_HK (繁體中文-香港)、zh_TW (正體中文-臺灣)，默认为 zh_CN。
	Lang string `json:"lang,omitempty"`
}
