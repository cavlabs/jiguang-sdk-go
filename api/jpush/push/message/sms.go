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

package message

// # 短信渠道补充送达内容
//   - 需要先把用户的手机号码与设备标识 Registration ID 匹配；
//   - 短信补发：在指定时间之内，判断推送是否成功，若没有达到成功标准则补发短信，请设置 DelayTime 为非 0 值；
//   - 短信并发：极光推送支持同时下发推送和短信，请设置 DelayTime 为 0。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1
type SMS struct {
	// 【必填】短信补充的内容模板 ID，没有填写该字段即表示不使用短信补充功能。
	TempID int64 `json:"temp_id"`
	// 【可选】签名 ID，该字段为空则使用应用默认签名。
	SignID int `json:"signid,omitempty"`
	// 【必填】短信发送的延迟时间，若在设定的时间内没有推送成功，则下发短信。
	//  - 设置为 0，表示立即发送短信，即通知和短信并发；
	//  - 设置为非 0，表示若在设定的时间内没有推送成功，则进行短信补发；
	//  - 单位为秒，不能超过 24 小时；
	//  - 该参数仅对 Android 和 iOS 平台有效。
	DelayTime int `json:"delay_time"`
	// 【可选】是否对补发短信的用户进行活跃过滤。
	ActiveFilter *bool `json:"active_filter,omitempty"`
	// 【可选】短信模板中的参数。
	TempParams map[string]interface{} `json:"temp_para,omitempty"`
}
