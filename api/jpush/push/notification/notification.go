/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
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

package notification

// # 通知内容
//
// “通知” 对象，是一条推送的实体内容对象之一（另一个是 “消息”），是会作为 “通知” 推送到客户端的，详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification%EF%BC%9A%E9%80%9A%E7%9F%A5
type Notification struct {
	// 【可选】各个平台的通知内容，详见 [alert] 文档说明
	//  - 通知的内容在各个平台上，都可能只有这一个最基本的属性 Alert。
	//  - 这个位置的 Alert 属性（直接在 Notification 对象下），是一个快捷定义，各平台的 Alert 信息如果都一样，则以此定义为准；如果各平台有定义，则覆盖这里的定义。
	// [alert]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#alert
	Alert string `json:"alert,omitempty"`
	// 【可选】个性化文案 - 备用 Alert 信息。
	AlternateAlert string `json:"alternate_alert,omitempty"`
	// 【可选】Android 平台上的通知，JPush SDK 按照一定的通知栏样式展示，详见 [android] 文档说明
	//
	// [android]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#android
	Android *Android `json:"android,omitempty"`
	// 【可选】iOS 平台上 APNs 通知结构，详见 [ios] 文档说明
	//  - 该通知内容会由 JPush 代理发往 Apple APNs 服务器，并在 iOS 设备上在系统通知的方式呈现；
	//  - 该通知内容满足 APNs 的规范。
	// [ios]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#ios
	IOS *IOS `json:"ios,omitempty"`
	// 【可选】鸿蒙（HarmonyOS）平台上通知结构，JPush SDK 按照一定的通知栏样式展示，详见 [hmos] 文档说明
	//
	// [hmos]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#hmos
	HMOS *HMOS `json:"hmos,omitempty"`
	// 【可选】快应用平台上通知结构，详见 [quickapp] 文档说明
	//  - 该通知内容满足快应用平台的规范。
	// [quickapp]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#quickapp
	QuickApp *QuickApp `json:"quickapp,omitempty"`
	// 【可选】iOS VoIP 功能。该类型推送支持和 iOS 的 Notification 通知并存，详见 [voip] 文档说明
	//  - 任意自定义 key/value 对，会透传给 APP。
	// [voip]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#voip
	VoIP map[string]interface{} `json:"voip,omitempty"`
}
