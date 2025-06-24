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

package message

// # 应用内增强提醒
//   - 此功能生效需 Android push SDK ≥ v3.9.0、iOS push SDK ≥ v3.4.0，若低于此版本按照原流程执行；
//   - 面向于通知栏消息类型，需搭配 Notification 参数一起使用，对于通知权限关闭的用户可设置启用此功能。此功能启用后，当用户前台运行 APP 时，会通过应用内消息的方式展示通知栏消息内容。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#inapp_message%EF%BC%9A%E5%BA%94%E7%94%A8%E5%86%85%E5%A2%9E%E5%BC%BA%E6%8F%90%E9%86%92
type InApp struct {
	// 【必填】面向通知栏消息，是否启用应用内增强提醒功能。
	Enabled bool `json:"inapp_message"`
}
