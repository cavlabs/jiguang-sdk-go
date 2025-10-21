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

import (
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/notification/alert"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/notification/style"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # Android 平台上的通知
type Android struct {
	// 【必填】通知内容。
	//  - 这里指定后会覆盖上级统一指定的 Alert 信息；
	//  - 内容可以为空字符串，表示不展示到通知栏；
	//  - 各推送通道对此字段的限制详见 [推送限制] 文档说明。
	// [推送限制]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#%E7%9B%B8%E5%85%B3%E5%8F%82%E8%80%83
	Alert string `json:"alert"`
	// 【可选】通知标题。
	//  - 如果指定了，则通知里原来展示 APP 名称的地方，将展示 Title；
	//  - 各推送通道对此字段的限制详见 [推送限制] 文档说明。
	// [推送限制]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#%E7%9B%B8%E5%85%B3%E5%8F%82%E8%80%83
	Title string `json:"title,omitempty"`
	// 【可选】通知栏样式 ID。
	//  - Android SDK 可 [设置通知栏样式]；
	//  - 根据样式 ID 来指定通知样式；
	//  - Android 8.0 开始建议采用 [NotificationChannel 配置]。
	// [设置通知栏样式]: https://docs.jiguang.cn/jpush/client/Android/android_api#%E9%80%9A%E7%9F%A5%E6%A0%8F%E6%A0%B7%E5%BC%8F%E5%AE%9A%E5%88%B6-api
	// [NotificationChannel 配置]: https://docs.jiguang.cn/jpush/client/Android/android_api#notificationchannel-%E9%85%8D%E7%BD%AE
	BuilderID int `json:"builder_id,omitempty"`
	// 【可选】Android 通知 ChannelID。
	//  - 根据 ChannelID 来指定通知栏展示效果，不超过 1000 字节；
	//  - Android 8.0 开始可以进行 [NotificationChannel 配置]；
	//  - Options.ThirdPartyChannel 下的蔚来小米、OPPO 和华为厂商参数也有 ChannelID 字段，若有填充，则优先使用，若无填充则以本字段定义为准。
	// [NotificationChannel 配置]: https://docs.jiguang.cn/jpush/client/Android/android_api#notificationchannel-%E9%85%8D%E7%BD%AE
	ChannelID string `json:"channel_id,omitempty"`
	// 【可选】通知栏消息分类条目。
	//  - 完全依赖 ROM 厂商对 Category 的处理策略；
	//  - 华为从 2023.09.15 开始基于《[华为消息分类标准]》对其本地通知进行管控推送，参考《[华为本地通知频次及分类管控通知]》，
	//  此字段值对应华为「本地通知」category 取值，开发者通过极光服务发起推送时如果传递了此字段值，请务必按照华为官方要求传递，
	//  极光会自动适配华为本地通知 importance 取值，无需开发者额外处理；
	//  - 考虑到一次推送包含多个厂商用户的情况，建议此处传递的字段值要和您 APP 开发代码中创建的 channel 效果对应（Category 值一致），最好创建新的 ChannelID，避免曾经已经创建了无法修改；
	//  - 官方 Category 分类取值规则也可参考《[华为消息分类对应表]》。
	// [华为消息分类标准]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section153801515616
	// [华为本地通知频次及分类管控通知]: https://developer.huawei.com/consumer/cn/doc/development/hmscore-common-Guides/push_notice_local-0000001615143510
	// [华为消息分类对应表]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id#%E5%8D%8E%E4%B8%BA%E6%B6%88%E6%81%AF%E5%88%86%E7%B1%BB%E8%AF%B4%E6%98%8E
	Category string `json:"category,omitempty"`
	// 【可选】通知栏展示优先级。
	//  - 默认为 0，范围为 -2～2；
	//  - 华为从 2023.09.15 开始基于《[华为消息分类标准]》对其本地通知进行管控推送，参考《[华为本地通知频次及分类管控通知]》，
	//  开发者通过极光服务发起推送时，如果有传递此字段值，请注意此字段要和 Category 同时使用；
	//  反之，如果传了 Category，没传递此值时极光会自动帮您适配处理优先级；
	//  - Priority = -2 时，对应华为本地通知 importance 级别为 IMPORTANCE_MIN；Priority = 0 时，对应华为本地通知 importance 级别为 IMPORTANCE_DEFAULT；
	//  - 官方消息优先级取值规则也可参考《[华为消息分类对应表]》；
	//  - 极光取值 -2～-1 对应 FCM 取值 normal，极光取值 0～2 对应 FCM 取值 high。
	// [华为消息分类标准]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section153801515616
	// [华为本地通知频次及分类管控通知]: https://developer.huawei.com/consumer/cn/doc/development/hmscore-common-Guides/push_notice_local-0000001615143510
	// [华为消息分类对应表]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id#%E5%8D%8E%E4%B8%BA%E6%B6%88%E6%81%AF%E5%88%86%E7%B1%BB%E8%AF%B4%E6%98%8E
	Priority *int `json:"priority,omitempty"`
	// 【可选】通知栏样式类型，默认为 0，其他枚举值：
	//  - style.BigText：大文本通知栏样式，1；
	//  - style.Inbox：文本条目通知栏样式，2；
	//  - style.BigPicture：大图片通知栏样式，3。
	Style style.Style `json:"style,omitempty"`
	// 【可选】通知提醒方式。
	//  - 可选范围为 -1～7，默认按照 -1 处理 (alert.DefaultAll)。即 0111 二进制，左数第二位代表 lights，第三位代表 vibrate，第四位代表 sound。
	//  - 0：不生效，1：生效。如: alert.DefaultAll = -1，alert.DefaultSound = 1, alert.DefaultVibrate = 2, alert.DefaultLights = 4 的任意 or 组合。
	AlertType alert.Type `json:"alert_type,omitempty"`
	// 【可选】大文本通知栏样式。
	//  - 当 Style = style.BigText 时可用，内容会被通知栏以大文本的形式展示出来；
	//  - 若没有填充 [厂商 BigText]，则也默认使用该字段展示；
	//  - 支持 API 16 以上的 ROM。
	// [厂商 BigText]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	BigText string `json:"big_text,omitempty"`
	// 【可选】文本条目通知栏样式。
	//  - 当 Style = style.Inbox 时可用，JSON 的每个 key 对应的 value 会被当作文本条目逐条展示；
	//  - 若没有填充 [厂商 Inbox]，则默认使用该 Inbox 字段展示；
	//  - 支持 API 16 以上的 ROM。
	// [厂商 Inbox]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	Inbox map[string]interface{} `json:"inbox,omitempty"`
	// 【可选】大图片通知栏样式。
	//  - 当 Style = style.BigPicture 时可用，目前支持 .jpg 和 .png 格式的图片，使用详情参见 [设置大图片文档]；
	//  - 支持网络图片 URL、本地图片的 Path、[极光 MediaID]，如果是 http/https 的 URL，会自动下载；如果要指定开发者准备的本地图片就填 SD 卡的相对路径；
	//  - 若没有填充 [厂商 BigPicture]，则默认使用该字段展示；
	//  - 支持 API 16 以上的 ROM。
	// [设置大图片文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
	// [极光 MediaID]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image
	// [厂商 BigPicture]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	BigPicture string `json:"big_pic_path,omitempty"`
	// 【可选】扩展字段。
	//  - 这里自定义 JSON 格式的 key/value 信息，以供业务使用。
	//  - 针对部分厂商跳转地址异常，可通过 third_url_encode 兼容处理，详情参考 [厂商通道无法跳转问题分析]；
	//  - 当通知内容超过厂商的限制时，厂商通道会推送失败，可以在 Extras 中配置 xx_content_forshort 参数传入对应厂商的通知内容，详情说明如下！
	// xx_content_forshort 参数：
	//  - mipns_content_forshort：【可选】小米通知内容。由于小米官方的通知内容长度限制为 128 个字符以内（中英文都算一个），当通知内容（极光的 Alert 字段的值）长度超过 128 时，小米通道会推送失败。
	//  此时调用极光 API 推送通知时，可使用此字段传入不超过 128 字符的通知内容作为小米通道通知内容；
	//  - oppns_content_forshort：【可选】OPPO 通知内容。由于 OPPO 官方的通知内容长度限制为 200 个字符以内（中英文都算一个），当通知内容（极光的 Alert 字段的值）长度超过 200 时，OPPO 通道会推送失败。
	//  此时调用极光 API 推送通知时，可使用此字段传入不超过 200 字符的通知内容作为 OPPO 通道通知内容；
	//  - vpns_content_forshort：【可选】vivo 通知内容。由于 vivo 官方的通知内容长度限制为 100 个字符以内（1 个汉字等于 2 个英文字符），当通知内容（极光的 Alert 字段的值）长度超过 100 时，vivo 通道会推送失败。
	//  此时调用极光 API 推送通知时，可使用此字段传入不超过 100 字符的通知内容作为 vivo 通道通知内容；
	//  - mzpns_content_forshort：【可选】魅族通知内容。由于魅族官方的通知内容长度限制为 100 个字符以内（中英文都算一个），当通知内容（极光的 Alert 字段的值）长度超过 100 时，魅族通道会推送失败。
	//  此时调用极光 API 推送通知时，可使用此字段传入不超过 100 字符的通知内容作为魅族通道通知内容。
	// [厂商通道无法跳转问题分析]: https://docs.jiguang.cn/jpush/faq/tech_faq#%E5%8E%82%E5%95%86%E9%80%9A%E9%81%93%E6%97%A0%E6%B3%95%E8%B7%B3%E8%BD%AC%EF%BC%9F
	Extras map[string]interface{} `json:"extras,omitempty"`
	// 【可选】通知栏大图标。
	//  - 图标大小不超过 30k（注：从 JPush Android SDK v4.0.0 版本开始，图片大小限制提升至 300k），使用详情参见 [设置图标文档]；
	//  - 支持网络图片 URL、本地图片的 Path、[极光 MediaID]，
	//  如果是 http/https 的 URL，会自动下载；如果要指定开发者准备的本地图片就填 SD 卡的相对路径；
	//  - 此字段值，若是 MediaID, 则对其它厂商通道生效，若非 MediaID，则对走华硕通道和极光通道下发的消息生效，不影响请求走其它厂商通道；
	//  - 若没有填充 [厂商 LargeIcon]，则默认使用该字段展示。
	// [设置图标文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
	// [极光 MediaID]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image
	// [厂商 LargeIcon]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	LargeIcon string `json:"large_icon,omitempty"`
	// 【可选】通知栏小图标。
	//  - 图标大小不超过 30k（注：从 JPush Android SDK v4.0.0 版本开始，图片大小限制提升至 300k），使用详情参见 [设置图标文档]；
	//  - 支持以 http/https 开头的网络图片和通过极光图片上传接口得到的 [MediaID] 值；
	//  - 此字段值，若是 MediaID, 则对其它厂商通道生效，若非 MediaID，则对走华硕通道和极光通道下发的消息生效，不影响请求走其它厂商通道；
	//  - 若没有填充 [厂商 SmallIcon]，则默认使用该字段展示。
	// [设置图标文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
	// [MediaID]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image
	// [厂商 SmallIcon]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	SmallIcon string `json:"small_icon_uri,omitempty"`
	// 【可选】设置通知小图标背景色。
	//  - 该字段仅对消息走极光通道下发生效；
	//  - 该字段能辅助解决部分设备小图标显示灰白情况，但最终还是依赖系统本身支持情况，建议开发者在设计 UI 图标时就做好适配工作；
	//  - 需要搭配 Android JPush SDK v5.5.0 及其以上版本使用。
	IconBgColor string `json:"icon_bg_color,omitempty"`
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
	//  - JPush SDK ≥ v4.2.2，可不再填写本字段，仅设置 Intent 字段即可，但若需兼容旧版 SDK 必须填写该字段。
	UriAction string `json:"uri_action,omitempty"`
	// 【可选】设置角标数字累加值，在原角标的基础上进行累加。
	//  - 此属性目前仅针对华为 EMUI 8.0 及以上、小米 MIUI 6 及以上、vivo、荣耀设备生效；
	//  - 此字段如果不填，表示不改变角标数字（小米设备由于系统控制，不论推送走极光通道下发还是厂商通道下发，即使不传递依旧是默认 +1 的效果）；
	//  - 取值范围为：1～99，若设置了取值范围内的数字，下一条通知栏消息配置的 BadgeAddNum 数据会和原角标数量进行相加，建议取值为 1。
	//  举例：BadgeAddNum 取值为 1，原角标数为 2，发送此角标消息后，应用角标数显示为 3。
	//  - 针对华为和荣耀通道，若 BadgeSetNum 与 BadgeAddNum 同时存在，则以 BadgeSetNum 为准；
	//  若 BadgeAddNum 和 BadgeSetNum 都设置为空，则应用角标数字默认加 1。
	BadgeAddNum *int `json:"badge_add_num,omitempty"`
	// 【可选】设置角标数字固定值。
	//  - 此属性目前仅针对华为 EMUI 8.0 及以上、荣耀设备走厂商通道时生效，若 BadgeSetNum 与 BadgeAddNum 同时存在，则以 BadgeSetNum 为准；
	//  若 BadgeAddNum 和 BadgeSetNum 都设置为空，则应用角标数字默认加 1。
	//  - 取值范围为：0～99，若设置了取值范围内的数字，对应下一条通知栏消息配置的 BadgeSetNum 数字则为角标数值，
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
	// 【可选】定时展示开始时间。
	//  - 此属性不填写，SDK 默认立即展示；此属性填写，则以填写时间点为准才开始展示；
	//  - JPush Android SDK v3.5.0 版本开始支持；
	//  - 目前适配的通道有：极光、OPPO、vivo、魅族。
	ShowBeginTime *jiguang.LocalDateTime `json:"show_begin_time,omitempty"`
	// 【可选】定时展示结束时间。
	//  - 此属性不填写，SDK 会一直展示；此属性填写，则以填写时间点为准，到达时间点后取消展示；
	//  - JPush Android SDK v3.5.0 版本开始支持；
	//  - 目前适配的通道有：极光、OPPO、vivo、魅族。
	ShowEndTime *jiguang.LocalDateTime `json:"show_end_time,omitempty"`
	// 【可选】APP 在前台，通知是否展示。
	//  - 值为 "1" 时，APP 在前台会弹出/展示通知栏消息；
	//  - 值为 "0" 时，APP 在前台不会弹出/展示通知栏消息；
	//  - 默认情况下 APP 在前台会弹出/展示通知栏消息，JPush Android SDK v3.5.8 版本开始支持；
	//  - 目前适配的通道有：极光、华为、小米、vivo。
	DisplayForeground string `json:"display_foreground,omitempty"`
	// 【可选】远程自定义通知按钮。
	//
	// 适配极光通道与华为厂商通道。
	//
	// 注意：此功能从 JPush Android SDK v5.9.0 版本开始支持。
	CustomButtons []struct {
		// 【必填】按钮名称。
		//  - 非空字符串，长度不大于 20 字符。
		Name string `json:"name"`
		// 【必填】按钮动作类型。
		//  - 1：intent 跳转；
		//  - 2：deeplink 跳转；
		//  - 3：进入应用首页；
		//  - 4：清除通知。
		ActionType int `json:"action_type"`
		// 【可选】自定义动作。
		//  - 当 ActionType 为 1 或 2 时必填，且需要是对应跳转格式；
		//  - 长度不大于 300 字符。
		Action string `json:"action"`
		// 【可选】用于在点击按钮后给应用的透传字段。
		//  - 长度不大于 100 字符。
		Data string `json:"data"`
	} `json:"custom_buttons,omitempty"`
}
