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

package notification

// # 自定义消息转厂商通知内容（v1 版本）
//
// Push API 发起自定义消息类型的推送请求时，针对 Android 设备，如果 APP 长连接不在线，则消息没法及时的下发，针对这种情况，极光推出了 “自定义消息转厂商通知” 的功能。
//
// 也就是说，针对用户一些重要的自定义消息，可以申请开通极光 VIP 厂商通道功能，开通后，通过 APP 长连接不在线时没法及时下发的消息，可以通过厂商通道下发以厂商通知形式展示，及时提醒到用户。
// 极光内部会有去重处理，您不用担心消息重复下发问题。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// Deprecated: 已过时，推荐使用 ThirdV2。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
type Third struct {
	// 【可选】补发通知标题，如果为空则默认为应用名称。
	Title string `json:"title,omitempty"`
	// 【必填】补发通知的内容，不能为空或空字符串。
	Content string `json:"content"`
	// 【可选】指定跳转页面（使用 Intent 里的 URL 指定点击通知栏后跳转的目标页面）。
	//  - SDK ＜ 422 的版本此字段值仅对走华硕通道和极光自有通道下发生效，不影响请求走其它厂商通道；
	//  - SDK ≥ 422 的版本，API 推送时建议填写 Intent 字段，否则点击通知可能无跳转动作。
	// 支持以下三种类型：
	//  1. 跳转到目标页: intent:#Intent;action=action 路径;component= 包名 /Activity 全名;end
	//
	// 	注：OPPO 和 FCM 通道必须传 "action 路径"，其他厂商必须传 "Activity 全名", 否则将出现对应厂商无法跳转问题！
	//  2. 跳转到 deeplink 地址：scheme://test?key1=val1&key2=val2
	//  3. 应用首页: intent:#Intent;action=android.intent.action.MAIN;end（固定为此地址）
	Intent *Intent `json:"intent,omitempty"`
	// 【可选】指定跳转页面。
	//  - 用于指定开发者想要打开的 activity，值为 activity 节点的 "android:name" 属性值；
	//  - 适配华为、小米、vivo 厂商通道跳转；
	//  - JPush SDK ≥ v4.2.2，可不再填写本字段，仅设置 Intent 字段即可。
	UriActivity string `json:"uri_activity,omitempty"`
	// 【可选】指定跳转页面。
	//  - 用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值；
	//  - 适配 OPPO、FCM 跳转；
	//  - JPush SDK ≥ v4.2.2，可不再填写本字段，仅设置 Intent 字段即可，但若需兼容旧版 SDK 必须填写该字段
	UriAction string `json:"uri_action,omitempty"`
	// 【可选】设置角标数字累加值，在原角标的基础上进行累加，取值范围为：1～99。
	//  - 此属性目前仅针对华为 EMUI 8.0 及以上、小米 MIUI 6 及以上、vivo、荣耀设备生效；
	//  - 此字段如果不填，表示不改变角标数字（小米设备由于系统控制，不论推送走极光通道下发还是厂商通道下发，即使不传递依旧是默认 +1 的效果）；
	//  - 若设置了取值范围内的数字，下一条通知栏消息配置的 BadgeAddNum 数据会和原角标数量进行相加，建议取值为 1。
	//  举例：BadgeAddNum 取值为 1，原角标数为 2，发送此角标消息后，应用角标数显示为 3。
	//  - 针对华为和荣耀通道，若 BadgeSetNum 与 BadgeAddNum 同时存在，则以 BadgeSetNum 为准。
	BadgeAddNum *int `json:"badge_add_num,omitempty"`
	// 【可选】设置角标数字固定值，取值范围为：0～99。
	//  - 此属性目前仅针对华为 EMUI 8.0 及以上、荣耀设备走厂商通道时生效，若 BadgeSetNum 与 BadgeAddNum 同时存在，则以 BadgeSetNum 为准；
	//  - 若设置了取值范围内的数字，对应下一条通知栏消息配置的 BadgeSetNum 数字则为角标数值，
	//  举例：BadgeSetNum 取值为 1，无论应用之前角标数为多少，发送此角标消息后，应用角标数均显示为 1。
	BadgeSetNum *int `json:"badge_set_num,omitempty"`
	// 【可选】桌面图标对应的应用入口 Activity 类，比如 "com.test.badge.MainActivity"。
	//  - 仅华为和荣耀通道推送时生效，此值如果填写非主 Activity 类，以厂商限制逻辑为准；
	//  - 若需要实现角标累加功能，需配合 BadgeAddNum 使用，二者需要共存，缺少其一不可；
	//  - 若需要实现角标固定值功能，需配合 BadgeSetNum 使用，二者需要共存，缺少其一不可。
	BadgeClass string `json:"badge_class,omitempty"`
	// 【可选】铃声。
	//  - 填写 Android 工程中 /res/raw/ 路径下铃声文件名称，无需文件名后缀；
	//  - 注意：针对 Android 8.0 以上，当传递了 ChannelID 时，此属性不生效。
	Sound string `json:"sound,omitempty"`
	// 【可选】Android 通知 ChannelID。
	//  - 根据 ChannelID 来指定通知栏展示效果，不超过 1000 字节；
	//  - Android 8.0 开始可以进行 [NotificationChannel 配置]。
	// [NotificationChannel 配置]: https://docs.jiguang.cn/jpush/client/Android/android_api#notificationchannel-%E9%85%8D%E7%BD%AE
	ChannelID string `json:"channel_id,omitempty"`
	// 【可选】扩展字段。
	//  - 这里自定义 JSON 格式的 key/value 信息，以供业务使用。
	Extras map[string]interface{} `json:"extras,omitempty"`
}
