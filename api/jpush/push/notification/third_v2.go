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

package notification

// # 自定义消息转厂商通知内容（v2 版本）
//
// Push API 发起自定义消息类型的推送请求时，针对 Android 设备，如果 APP 长连接不在线，则消息没法及时的下发，针对这种情况，极光推出了 “自定义消息转厂商通知” 的功能。
//
// 也就是说，针对用户一些重要的自定义消息，可以申请开通极光 VIP 厂商通道功能，开通后，通过 APP 长连接不在线时没法及时下发的消息，可以通过厂商通道下发以厂商通知形式展示，及时提醒到用户。
// 极光内部会有去重处理，您不用担心消息重复下发问题。
//
// Android、IOS、HMOS 三者必须有其一，可以三者并存！
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
type ThirdV2 struct {
	Android *Android `json:"android,omitempty"` // 【可选】Android 平台，自定义消息转 Android 通知内容体。
	IOS     *IOS     `json:"ios,omitempty"`     // 【可选】iOS 平台，自定义消息转 iOS 通知内容体。
	HMOS    *HMOS    `json:"hmos,omitempty"`    // 【可选】HarmonyOS 平台，自定义消息转 HarmonyOS 平台 通知内容体。
}
