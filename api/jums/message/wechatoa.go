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

// # 【微信公众号】消息
//
// 当前支持以下 2 种消息：
//   - 模板消息 (0): WechatoaTemplate；
//   - 订阅通知 (1): WechatoaSubscription。
type Wechatoa interface {
	setType() Wechatoa
	getType() int
}

// ---------------------------------------------------------------------------------------------------------------------

// # 【微信公众号】模板消息
type WechatoaTemplate struct {
	// 【必填】消息类型，模板消息类型取值为 0。
	Type int `json:"type"`
	// 【必填】模板 ID。
	TemplateID string `json:"template_id"`
	// 【可选】模板跳转链接。
	URL string `json:"url,omitempty"`
	// 【可选】跳转小程序所需数据，不需跳转小程序可不用传该数据。
	//  - URL 和 MiniProgram 同时不填，无跳转；
	//  - URL 和 MiniProgram 同时填写，优先跳转小程序。
	MiniProgram map[string]interface{} `json:"miniprogram,omitempty"`
	// 【可选】所需跳转到的小程序 AppID。
	//  - 该小程序 AppID 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏。
	AppID string `json:"appid,omitempty"`
	// 【可选】所需跳转到的小程序的具体页面路径。
	//  - 支持带参数，示例：index?foo=bar
	//  - 要求该小程序已发布，暂不支持小游戏。
	PagePath string `json:"pagepath,omitempty"`
	// 【必填】模板内容。
	//  - 格式形如：{"key1":{"value": any,"color":"#173177"},"key2":{"value":any}}
	Data map[string]interface{} `json:"data"`
	// 【可选】模板内容字体颜色。
	//  - 不填默认为黑色。
	Color string `json:"color,omitempty"`
}

var _ Wechatoa = &WechatoaTemplate{Type: 0}

func (m *WechatoaTemplate) setType() Wechatoa { m.Type = 0; return m }
func (m *WechatoaTemplate) getType() int      { return 0 }

// ---------------------------------------------------------------------------------------------------------------------

// # 【微信公众号】订阅通知
type WechatoaSubscription struct {
	// 【必填】消息类型，订阅通知类型取值为 1。
	Type int `json:"type"`
	// 【必填】模板 ID。
	TemplateID string `json:"template_id"`
	// 【可选】跳转网页。
	Page string `json:"page,omitempty"`
	// 【可选】跳转小程序所需数据，不需跳转小程序可不用传该数据。
	//  - Page 和 MiniProgram 同时不填，无跳转；
	//  - Page 和 MiniProgram 同时填写，优先跳转小程序。
	MiniProgram map[string]interface{} `json:"miniprogram,omitempty"`
	// 【可选】所需跳转到的小程序 AppID。
	//  - 该小程序 AppID 必须与发订阅通知的公众号是绑定关联关系，暂不支持小游戏。
	AppID string `json:"appid,omitempty"`
	// 【可选】所需跳转到的小程序的具体页面路径。
	//  - 支持带参数，示例：index?foo=bar
	//  - 要求该小程序已发布，暂不支持小游戏。
	PagePath string `json:"pagepath,omitempty"`
	// 【必填】模板内容。
	//  - 格式形如：{"key1":{"value": any,"color":"#173177"},"key2":{"value":any}}
	Data map[string]interface{} `json:"data"`
}

var _ Wechatoa = &WechatoaSubscription{Type: 1}

func (m *WechatoaSubscription) setType() Wechatoa { m.Type = 1; return m }
func (m *WechatoaSubscription) getType() int      { return 1 }
