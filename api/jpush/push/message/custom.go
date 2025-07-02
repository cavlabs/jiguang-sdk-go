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

package message

// # 自定义消息内容
//
// 自定义消息，又称作：应用内消息，透传消息。
//   - 此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 APP，需要 APP 自行处理；
//   - iOS 在推送应用内消息通道（非 APNs）获取此部分内容，需 APP 处于前台；
//   - 鸿蒙平台（HarmonyOS）从 2024.08.13 开始支持自定义消息，对应 JPush HarmonyOS SDK v1.1.0 版本。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#message%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
type Custom struct {
	// 【可选】消息标题。
	Title string `json:"title,omitempty"`
	// 【必填】消息内容本身。
	Content string `json:"msg_content"`
	// 【可选】消息内容类型，开发者可根据自身业务定义具体类型。
	ContentType string `json:"content_type,omitempty"`
	// 【可选】个性化文案 - 备用标题
	AlternateTitle string `json:"alternate_title,omitempty"`
	// 【可选】个性化文案 - 备用内容
	AlternateContent string `json:"alternate_content,omitempty"`
	// 【可选】可选参数。
	Extras map[string]interface{} `json:"extras,omitempty"`
}
