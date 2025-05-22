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

package push

import (
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/audience"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/callback"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/cid"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/liveactivity"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/message"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/hmos"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/options"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/send"
)

// ↓↓↓ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↓↓↓

type (
	// # 推送参数
	SendParam = send.Param
	// # 模板参数
	//
	// 创建模板时，开发者设置的变量参数。
	TemplateParam = send.TemplateParam

	// # 获取推送唯一标识 (CID) 结果
	CidGetResult = cid.GetResult

	// # 推送设备对象
	//
	// 表示一条推送可以被推送到哪些设备列表。
	//
	// 确认推送设备对象，JPush 提供了多种方式，比如：别名、标签、注册 ID、分群、广播等，详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#audience%EF%BC%9A%E6%8E%A8%E9%80%81%E7%9B%AE%E6%A0%87
	Audience = audience.Audience
	// # 文件推送对象
	//  - 可用于包括 SendByFile（文件立即推送）和 ScheduleSend（文件定时推送）等相关接口。
	//  - 详见 [文件推送 API] 文档说明。
	// [文件推送 API]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%96%87%E4%BB%B6%E6%8E%A8%E9%80%81-api
	FileAudience = audience.File

	// # 推送可选项
	//
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#options%EF%BC%9A%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0
	Options = options.Options
	// # 推送请求下发通道
	ThirdPartyChannel = options.ThirdPartyChannel
	// # 推送请求下发通道的策略和属性参数
	ThirdPartyChannelOptions = options.ThirdPartyChannelOptions
	// # 极光 WebPortal 的附加属性
	PortalExtraOptions = options.PortalExtraOptions

	// # 通知内容
	//
	// “通知” 对象，是一条推送的实体内容对象之一（另一个是 “消息”），是会作为 “通知” 推送到客户端的，详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification%EF%BC%9A%E9%80%9A%E7%9F%A5
	Notification = notification.Notification
	// # Android 平台上的通知
	AndroidNotification = notification.Android
	// # iOS 平台上的 APNs 通知
	IosNotification = notification.IOS
	// # 鸿蒙（HarmonyOS）平台上的通知
	HmosNotification = notification.HMOS
	// # 快应用平台上的通知
	QuickAppNotification = notification.QuickApp
	// # 自定义消息转厂商通知内容（v1 版本）
	//
	// Push API 发起自定义消息类型的推送请求时，针对 Android 设备，如果 APP 长连接不在线，则消息没法及时的下发，针对这种情况，极光推出了 “自定义消息转厂商通知” 的功能。
	//
	// 也就是说，针对用户一些重要的自定义消息，可以申请开通极光 VIP 厂商通道功能，开通后，通过 APP 长连接不在线时没法及时下发的消息，可以通过厂商通道下发以厂商通知形式展示，及时提醒到用户。
	// 极光内部会有去重处理，您不用担心消息重复下发问题。
	//
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// Deprecated: 已过时，推荐使用 ThirdNotificationV2。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
	ThirdNotification = notification.Third
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
	ThirdNotificationV2 = notification.ThirdV2
	// # 指定通知点击跳转页面
	NotificationIntent = notification.Intent

	// # 自定义消息内容
	//
	// 自定义消息，又称作：应用内消息，透传消息。
	//  - 此部分内容不会展示到通知栏上，JPush SDK 收到消息内容后透传给 APP，需要 APP 自行处理；
	//  - iOS 在推送应用内消息通道（非 APNs）获取此部分内容，需 APP 处于前台；
	//  - 鸿蒙平台（HarmonyOS）从 2024.08.13 开始支持自定义消息，对应 JPush HarmonyOS SDK v1.1.0 版本。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#message%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
	CustomMessage = message.Custom

	// # 实时活动内容
	//  - 实时活动消息要求使用 iOS P8 证书，对应 [极光 WebPortal 集成设置中 iOS 鉴权方式需要选择「Token Authentication 配置」] 方式；
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [极光 WebPortal 集成设置中 iOS 鉴权方式需要选择「Token Authentication 配置」]: https://docs.jiguang.cn/jpush/console/push_setting/integration_set#ios
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#live_activity%EF%BC%9A%E5%AE%9E%E6%97%B6%E6%B4%BB%E5%8A%A8%E6%B6%88%E6%81%AF
	LiveActivityMessage = liveactivity.Message
	// # iOS 的实时活动消息
	LiveActivityIosMessage = liveactivity.IosMessage
	// # iOS 实时活动通知内容
	LiveActivityIosAlertMessage = liveactivity.IosAlertMessage
	// # 实时活动事件类型
	LiveActivityEvent = liveactivity.Event

	// # 应用内增强提醒
	//  - 此功能生效需 Android push SDK ≥ v3.9.0、iOS push SDK ≥ v3.4.0，若低于此版本按照原流程执行；
	//  - 面向于通知栏消息类型，需搭配 Notification 参数一起使用，对于通知权限关闭的用户可设置启用此功能。此功能启用后，当用户前台运行 APP 时，会通过应用内消息的方式展示通知栏消息内容。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#inapp_message%EF%BC%9A%E5%BA%94%E7%94%A8%E5%86%85%E5%A2%9E%E5%BC%BA%E6%8F%90%E9%86%92
	InAppMessage = message.InApp

	// # 短信渠道补充送达内容
	//  - 需要先把用户的手机号码与设备标识 Registration ID 匹配；
	//  - 短信补发：在指定时间之内，判断推送是否成功，若没有达到成功标准则补发短信，请设置 DelayTime 为非 0 值；
	//  - 短信并发：极光推送支持同时下发推送和短信，请设置 DelayTime 为 0。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1
	SmsMessage = message.SMS

	// # 回调参数
	//  - 通过指定 Callback 参数，方便用户临时变更回调 URL 或者回调带上其自定义参数，满足其日常业务需求；
	//  - 此功能仅针对极光 VIP 用户提供，主要提供消息送达、点击回执数据。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#callback%EF%BC%9A%E5%9B%9E%E8%B0%83%E5%8F%82%E6%95%B0
	Callback = callback.Callback
	// # 回调数据类型
	CallbackType = callback.Type
)

const (
	BroadcastAuds = audience.All // 广播推送，表示推送给所有设备 (all)。

	HmosPushTypeAlert      = hmos.PushTypeAlert      // 华为场景化消息类型：通知消息 (0)
	HmosPushTypeSubscribe  = hmos.PushTypeSubscribe  // 华为场景化消息类型：授权订阅消息 (0)
	HmosPushTypeFormUpdate = hmos.PushTypeFormUpdate // 华为场景化消息类型：卡片刷新消息 (1)
	HmosPushTypeExtension  = hmos.PushTypeExtension  // 华为场景化消息类型：通知扩展消息 (2)
	HmosPushTypeBackground = hmos.PushTypeBackground // 华为场景化消息类型：后台消息 (6)
	HmosPushTypeLiveView   = hmos.PushTypeLiveView   // 华为场景化消息类型：实况窗消息 (7)
	HmosPushTypeVoIPCall   = hmos.PushTypeVoIPCall   // 华为场景化消息类型：应用内通话消息 (10)

	LiveActivityEventStart  = liveactivity.EventStart  // 实时活动事件类型：start 创建
	LiveActivityEventUpdate = liveactivity.EventUpdate // 实时活动事件类型：update 更新
	LiveActivityEventEnd    = liveactivity.EventEnd    // 实时活动事件类型：end 结束

	CallbackTypeReceived = callback.Received // 回调数据类型：送达回执 (1)
	CallbackTypeClicked  = callback.Clicked  // 回调数据类型：点击回执 (2)
	CallbackTypePush     = callback.Push     // 回调数据类型：推送成功回执 (8)
)

// ↑↑↑ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↑↑↑
