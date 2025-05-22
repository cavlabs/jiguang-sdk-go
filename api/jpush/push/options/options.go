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

package options

import "github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/style"

// # 推送可选项
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#options%EF%BC%9A%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0
type Options struct {
	// 【可选】推送序号。
	//  - 纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回；
	//  - 值为 0 表示该消息无 SendNo，所以字段取值范围为非 0 的整数。
	SendNo int64 `json:"sendno,omitempty"`
	// 【可选】离线消息保留时长（单位：秒）。
	//  - 推送当前用户不在线时，为该用户保留多长时间的离线消息，以便其上线时再次推送；
	//  - 默认 86400（1 天），普通用户最长 3 天，VIP 用户最长 10 天，设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到；
	//  - 该字段对 iOS 的 Notification 消息无效。
	TimeToLive *int64 `json:"time_to_live,omitempty"`
	// 【可选】要覆盖的消息 ID。
	//
	// 如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 MsgID 就会产生覆盖效果，即：
	//  - 该 MsgID 离线收到的消息是覆盖后的内容，即使该 MsgID Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知；
	//  - 覆盖功能起作用的时限是：1 天，如果在覆盖指定时限内该 MsgID 不存在，则返回 1003 错误，提示不是一次有效的消息覆盖操作，当前的消息不会被推送；
	//  - 该字段仅对 Android 有效，且仅支持极光通道、小米通道、魅族通道、OPPO 通道、FCM 通道、荣耀通道和华为通道（EMUI 10 及以上的设备）。
	OverrideMsgID int64 `json:"override_msg_id,omitempty"`
	// 【可选】APNs 是否生产环境。
	//
	// 该字段仅对 iOS 的 Notification 有效，如果不指定则为推送生产环境。
	ApnsProduction *bool `json:"apns_production,omitempty"`
	// 【可选】更新 iOS 通知的标识符。
	//  - APNs 新通知如果匹配到当前通知中心有相同 ApnsCollapseID 字段的通知，则会用新通知内容来更新它，并使其置于通知中心首位；
	//  - ApnsCollapseID 长度不可超过 64 字节。
	ApnsCollapseID string `json:"apns_collapse_id,omitempty"`
	// 【可选】定速推送时长（单位：分钟）。
	//  - 又名缓慢推送，把原本尽可能快的推送速度，降低下来，给定的 n 分钟内，均匀地向这次推送的目标用户推送，最大值为 1400；
	//  - 最多能同时存在 20 条定速推送；
	//  - 未设置则不是定速推送。
	BigPushDuration int `json:"big_push_duration,omitempty"`
	// 【可选】推送请求下发通道。
	//  - 目前只支持 xiaomi、huawei、honor、meizu、oppo、vivo、fcm、nio 类型用户，可以一个或者多个同时存在，未传递的通道类型其对应的厂商下发走「默认下发逻辑」：
	//
	//  1. 免费用户：Distribution 默认为 secondary_push，DistributionFcm 默认为 secondary_fcm_push；
	//
	//  2. VIP 用户：Distribution 默认为 first_ospush，DistributionFcm 默认为 fcm。
	//  - 仅针对配置了厂商用户使用有效，详情参考 [third_party_channel 说明]。
	//
	// [third_party_channel 说明]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#third_party_channel-%E8%AF%B4%E6%98%8E
	ThirdPartyChannel *ThirdPartyChannel `json:"third_party_channel,omitempty"`
	// 【可选】消息类型分类。
	//
	// 极光不对指定的消息类型进行判断或校准，会以开发者自行指定的消息类型适配 Android 厂商通道。不填默认为 0。
	//  - 0：代表运营消息；
	//  - 1：代表系统消息。
	// 此字段优先级最高，会覆盖 ThirdPartyChannel 的 vivo 厂商的 Classification 设置的值。
	Classification *int `json:"classification,omitempty"`
	// 【可选】目标转化事件。
	//
	// 目标转化事件支持传递「自定义事件」和「极光预置事件」，目前支持 Android 和 iOS 平台（要求 JPush SDK ≥ v5.0.0 ，且 JCore ≥ v4.2.0），支持通知消息和应用内消息两种消息类型。
	//  - 自定义事件：需集成极光分析 SDK，开发者在极光分析产品中自行创建的业务事件（如：加入购物车、浏览商品等），详情参考 [如何创建自定义事件] 和 [SDK 如何上报自定义事件]；
	//  - 极光预置事件：极光推送 SDK 默认支持，无需开发者创建，也无需集成极光分析 SDK，系统已预置；目标支持的预置事件有 2 个：jg_app_show（应用切换到前台）、jg_app_hide（应用切换到后台）。
	// 代码示例：`{"options": {"target_event": ["jg_app_show"]}}`。
	//
	// [如何创建自定义事件]: https://docs.jiguang.cn/public_service/dataCenter/metadata/metaEvent
	// [SDK 如何上报自定义事件]: https://docs.jiguang.cn/public_service/client/Android/sdk_api#%E4%B8%8A%E6%8A%A5%E8%87%AA%E5%AE%9A%E4%B9%89%E4%BA%8B%E4%BB%B6
	TargetEvent []string `json:"target_event,omitempty"`
	// 【可选】测试消息标识。
	//  - 指定鸿蒙平台通知和自定义消息推送配置，优先级大于 HMOS 通知体内的 TestMessage 字段（同样适配鸿蒙自定义消息，如果推送鸿蒙自定义消息，请传递此字段）；
	//  - 请注意区别于 TestMode 功能字段，TestMessage 仅用于适配厂商的测试消息功能，并非表示处于测试模式下推送。
	TestMessage *bool `json:"test_message,omitempty"`
	// 【可选】华为回执 ID。
	//  - 指定鸿蒙平台通知和自定义消息推送配置，优先级大于 HMOS 通知体内的 ReceiptID 字段。
	ReceiptID string `json:"receipt_id,omitempty"`
	// 【可选】是否使用亮屏推送。
	//  - true：使用亮屏推送，false：不使用亮屏推送，默认值 false；
	//  - 此功能为增值付费服务，需要额外申请权限；
	//  - 当使用亮屏推送时，建议同时设置 NeedBackup 为 true；
	//  - 此功能仅支持单纯通知消息，不支持自定义消息或者通知+自定义消息推送，否则请求会返回 code 码 1035；
	//  - 此功能不支持定速推送，否则请求会返回 code 码 1035；
	//  - 亮屏推送支持的时间范围是每天 7:00 - 22:00；
	//  - 亮屏推送对于 Android 厂商用户的下发策略固定为在线走极光，离线走厂商。
	ActivePush *bool `json:"active_push,omitempty"`
	// 【可选】是否使用亮屏推送兜底策略。
	//  - true：使用亮屏兜底策略，false：不使用亮屏兜底策略，默认值 false；
	//  - 若此字段指定为 true，则 ActivePush 字段值必须为 true；
	//  - 是否使用兜底策略主要是确认离线消息到期后的处理逻辑；
	//  - 当使用兜底策略下发时：如果是厂商用户（离线消息到期后 0～5 分钟之内通过厂商通道下发），如果是非厂商用户（离线消息到期后，如果用户是在线状态则直接下发；如果用户离线则丢弃）。
	//  例如上午 8 点推送此条消息，设置了离线时间 2 小时。在 8:00 - 10:00 内，设备亮屏则会触发消息下发。剩余未发送的用户，在到达 10:00 后，0～5 分钟之内剩余消息走厂商通道下发。
	//  - 当不使用兜底策略下发时：离线消息到期后未下发的直接丢弃，不区分是否厂商用户。
	NeedBackup *bool `json:"need_backup,omitempty"`
	// 【可选】推送计划标识。
	//  - 需先创建计划标识值，创建步骤参考 [推送计划文档]。
	// [推送计划文档]: https://docs.jiguang.cn/jpush/console/config_manage/push_plan
	BusinessOperationCode string `json:"business_operation_code,omitempty"`
	// 【可选】是否测试模式推送。
	//  - false：正式模式推送消息（默认值），true：测试模式推送消息；
	//  - 测试模式推送消息仅推送给到测试设备；
	//  - 功能逻辑可参考文档 [测试模式]；
	//  - 请注意区分区别 TestMessage 字段：TestMessage 仅用于适配厂商的测试消息功能，并非表示处于测试模式下推送；TestMode 则表示请求在极光平台下发消息时就已经控制，消息是否仅下发给到测试设备；
	//  - 此功能为增值付费服务，需要额外申请权限。
	// [测试模式]: https://docs.jiguang.cn/jpush/console/push_manage/testmode
	TestMode *bool `json:"test_mode,omitempty"`
	// 【可选】是否设置个性化文案。
	AlternateSet *bool `json:"alternate_set,omitempty"`
	// 【可选】地理围栏配置参数。
	Geofence map[string]interface{} `json:"geofence,omitempty"`
	// 【可选】极光 WebPortal 的附加属性。
	PortalExtra *PortalExtraOptions `json:"portal_extra,omitempty"`
	// 【可选】自定义消息转厂商通知功能版本。
	//  - 可选值：v1、v2，为空则默认使用 v1 版本，如果使用 v2 版本则必须指定此字段值；
	//  - 推荐使用 v2 版本，支持 Android、iOS、HarmonyOS 三个平台；
	//  - v1 版本仅支持 Android 平台，且后续将不再拓展支持新功能字段，仅维持现状。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
	Notification3rdVer string `json:"notification_3rd_ver,omitempty"`
}

// # 推送请求下发通道
type ThirdPartyChannel struct {
	Xiaomi *ThirdPartyChannelOptions `json:"xiaomi,omitempty"` // 小米通道策略和属性参数。
	Huawei *ThirdPartyChannelOptions `json:"huawei,omitempty"` // 华为通道策略和属性参数。
	Honor  *ThirdPartyChannelOptions `json:"honor,omitempty"`  // 荣耀通道策略和属性参数。
	Meizu  *ThirdPartyChannelOptions `json:"meizu,omitempty"`  // 魅族通道策略和属性参数。
	OPPO   *ThirdPartyChannelOptions `json:"oppo,omitempty"`   // OPPO 通道策略和属性参数。
	Vivo   *ThirdPartyChannelOptions `json:"vivo,omitempty"`   // vivo 通道策略和属性参数。
	FCM    *ThirdPartyChannelOptions `json:"fcm,omitempty"`    // FCM 通道策略和属性参数。
	NIO    *ThirdPartyChannelOptions `json:"nio,omitempty"`    // 蔚来通道策略和属性参数。
}

// # 推送请求下发通道的策略和属性参数
type ThirdPartyChannelOptions struct {
	// 【可选】通知栏消息下发逻辑。
	//  - first_ospush（VIP）：成功注册厂商通道的设备走厂商通道，仅注册极光通道的设备走极光通道；
	//  - ospush（VIP）：表示推送强制走厂商通道下发。需要特别注意，只要指定此值的厂商对应配额不够时，推送请求会失败，返回 1012 错误码：
	//
	//  举例：假设指定一个小米用户的 Registration ID 推送，请求时针对小米、OPPO 等厂商通道都指定了 ospush，且 OPPO 厂商通道都配额已经用完，则推送同样会返回 1012 错误，提示厂商配额不足。
	//  - jpush：表示推送强制走极光通道下发；
	//  - secondary_push：表示推送优先走极光，极光不在线再走厂商，厂商作为辅助（建议此种方式）。
	Distribution string `json:"distribution,omitempty"`
	// 【可选】通知栏消息 FCM + 国内厂商组合类型下发逻辑。
	//  - jpush：表示推送强制走极光通道下发；
	//  - fcm（VIP）：表示推送强制走 FCM 通道下发；
	//  - pns（VIP）：表示推送强制走 小米 / 华为 / 荣耀 / 魅族 / OPPO / vivo 通道下发；
	//  - secondary_fcm_push：表示针对 FCM + 国内厂商组合类型用户，推送优先走极光，极光不在线再走 FCM 通道，FCM 作为辅助；
	//  - secondary_pns_push：表示针对 FCM + 国内厂商组合类型用户，推送优先走极光，极光不在线再走厂商通道，厂商作为辅助。
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	// 【可选】自定义消息国内厂商类型下发逻辑。
	//  定义国内厂商类型用户下发自定义消息的逻辑，此功能仅支持 huawei、honor 通道，需 Android push SDK ≥ v3.9.0。
	// 注意：小米推送于 2022.09.12 0 点起停止提供透传消息下发的服务，届时您将无法通过小米通道发送透传消息，请注意调整下发策略。
	//  - jpush：表示推送强制走极光通道下发；
	//  - first_ospush（VIP）：成功注册厂商通道的设备走厂商通道，仅注册极光通道的设备走极光通道；
	//  - secondary_push：表示推送优先走极光，极光不在线再走厂商，厂商作为辅助。
	DistributionCustomize string `json:"distribution_customize,omitempty"`

	// 【可选】通知栏消息分类。
	//  - 为了适配 小米、华为、OPPO、蔚来 手机厂商通知栏消息分类，由开发者自行向手机厂商申请，具体申请规则参考 [厂商消息分类使用指南]；
	//  - 注意华为数据处理位置为中国区的应用不支持该字段，详情参见 [华为自定义通知渠道]；
	//  - Android 下也有 ChannelID 字段，若本字段有填充，则优先使用，若无填充则以 Android 的 ChannelID 的定义为准；
	//  - 特别注意：由于 OPPO 厂商 2024.11.20 实施 [OPPO 消息分类新规]，建议您同时填写此字段和 Category, NotifyLevel 字段。
	// [厂商消息分类使用指南]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id
	// [华为自定义通知渠道]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/android-custom-chan-0000001050040122
	// [OPPO 消息分类新规]: https://open.oppomobile.com/new/developmentDoc/info?id=13189
	ChannelID string `json:"channel_id,omitempty"`
	// 【可选】是否跳过配额判断及扣除，目前仅对小米和 OPPO 有效，默认为 false。
	//  - true：表示跳过判断及跳过扣除极光侧的配额；
	//  - false：表示不跳过判断及跳过扣除极光侧的配额。
	SkipQuota *bool `json:"skip_quota,omitempty"`
	// 【可选】通知栏消息分类。
	//
	// vivo 手机厂商通知栏消息分类，不填默认为 0。
	//
	// 此字段优先级较低，会被 Options 的 Classification 设置的值覆盖，请您务必设置 Options 的 Classification 值。
	//  - 0：代表运营消息；
	//  - 1：代表系统消息。
	// 目前 vivo 对系统消息分类较为严格，参考 [具体规则]。
	//
	// 关于 Classification 和 SkipQuota 字段说明：
	//  - 不传递 Classification 字段，但传递 SkipQuota 时，应用是否扣除配额以客户传递的 SkipQuota 为准，需开发者自己管理配额；
	//  - 传递 Classification 时，会忽略 SkipQuota 值，极光会按照 [厂商系统消息、运营消息分类] 规则自动判断是否扣除配额，帮助开发者管理配额；
	//  - 蔚来厂商根据该字段确定发送给厂商的 Category 字段，填 0 表示 mobile_marketing（运营消息），填 1 表示 mobile_service（系统消息）。
	// [具体规则]: https://dev.vivo.com.cn/documentCenter/doc/359
	// [厂商系统消息、运营消息分类]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id
	Classification *int `json:"classification,omitempty"`
	// 【可选】通知栏消息类型。
	//
	// 对应 vivo 的 pushMode 字段，不填默认为 0。详情参考 [dev.vivo.com.cn] 文档说明。
	//  - 0：表示正式推送；
	//  - 1：表示测试推送。
	// [dev.vivo.com.cn]: https://dev.vivo.com.cn/documentCenter/doc/362#w2-98542835
	PushMode *int `json:"push_mode,omitempty"`
	// 【可选】华为、荣耀通知栏消息智能分类。
	//
	// 为了适配华为、荣耀手机厂商的通知栏消息智能分类，对应 华为 / 荣耀 的「云端通知」importance 字段，不填充则不下发。
	//  - LOW：一般消息。
	//  - NORMAL：重要消息。
	//  - HIGH：非常重要消息（仅华为支持）。
	// 说明：华为从 2023.09.15 开始基于《华为消息分类标准》对其云端通知和本地通知进行共同管控推送，开发者通过极光服务发起推送时，请注意此字段传值要符合华为官方「云端通知 importance」取值要求，要和 [华为云端通知 category 取值] 要求对应。
	//
	// 参考文档：
	//  - [华为通知消息智能分类]；
	//  - [荣耀通知消息分类标准]。
	//
	// [华为云端通知 category 取值]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section153801515616
	// [华为通知消息智能分类]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835
	// [荣耀通知消息分类标准]: https://developer.hihonor.com/cn/kitdoc?category=%E5%9F%BA%E7%A1%80%E6%9C%8D%E5%8A%A1&kitId=11002&navigation=guides&docId=notification-class.md&token=
	Importance string `json:"importance,omitempty"`
	// 【可选】华为、vivo、OPPO 厂商消息场景标识。
	//
	// 为了适配华为、vivo、OPPO 手机厂商消息，用于标识「云端通知」消息类型，确定消息提醒方式，对特定类型消息加快发送。
	//
	// 对应值及其说明参考：[华为]、[vivo]、[OPPO]。
	//
	// 注意事项：
	//  - 华为需完成 [自分类权益申请]；
	//  - 华为从 2023.09.15 开始基于《华为消息分类标准》对其云端通知和本地通知进行共同管控推送，开发者通过极光服务发起推送时，请注意此字段传值要符合华为官方 [华为云端通知 category 取值] 要求；
	//  - vivo 具体规则参考 [vivo 官方说明]；
	//  - OPPO 于 2024.11.20 实施消息分类新规，具体规则参考 [OPPO 官方说明]。
	//
	// [华为]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-References/https-send-api-0000001050986197#ZH-CN_TOPIC_0000001134031085__p5203378238
	// [vivo]: https://dev.vivo.com.cn/documentCenter/doc/359#w2-67805227
	// [OPPO]: https://open.oppomobile.com/new/developmentDoc/info?id=13189
	// [自分类权益申请]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id#%E5%8D%8E%E4%B8%BA%E6%B6%88%E6%81%AF%E5%88%86%E7%B1%BB%E4%BD%BF%E7%94%A8%E6%8C%87%E5%8D%97
	// [华为云端通知 category 取值]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-classification-0000001149358835#section153801515616
	// [vivo 官方说明]: https://dev.vivo.com.cn/documentCenter/doc/359#w1-36109489
	// [OPPO 官方说明]: https://open.oppomobile.com/new/developmentDoc/info?id=13189
	Category string `json:"category,omitempty"`
	// 【可选】OPPO 通知栏消息提醒等级。
	//  - 官方取值定义：1-通知栏、2-通知栏+锁屏、16-通知栏+锁屏+横幅+震动+铃声，请开发者按照官网定义传递，极光仅做透传处理；
	//  - 根据官方说明 NotifyLevel 字段，仅对「服务与通讯类」消息生效；
	//  - 使用 NotifyLevel 参数时，Category 参数必传。
	NotifyLevel int `json:"notify_level,omitempty"`
	// 【可选】华为自定义铃声。
	//  - 铃声文件必须存放在应用的 /res/raw 路径下，例如 /res/raw/shake.mp3，对应 Sound 值参数为 /raw/shake，无需后缀，支持的格式包括 MP3、WAV、MPEG 等；
	//  - 仅首次给应用推送 [服务与通讯消息] 时设置有效，需要配合 DefaultSound 一起使用，详情参考 [如何实现自定义铃声] 文档说明。
	// [服务与通讯消息]: https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/message-classification-0000001149358835#section5101818813
	// [如何实现自定义铃声]: https://docs.jiguang.cn/jpush/practice/custom_ringtone#%E5%8D%8E%E4%B8%BA%E9%80%9A%E9%81%93%E9%80%9A%E7%9F%A5%E5%AE%9E%E7%8E%B0
	Sound string `json:"sound,omitempty"`
	// 【可选】华为默认铃声控制开关。
	//
	// 华为官方说明，首次给应用推送 [服务与通讯消息] 时携带 Sound 字段且 DefaultSound 值设置为 false。
	//
	// 注意：由于铃声是通知渠道的属性，因此铃声仅在首次创建渠道（设置 Sound）有效，后续无法修改。
	//  - true：使用系统默认铃声；
	//  - false：使用 Sound 自定义铃声。
	//
	// [服务与通讯消息]: https://developer.huawei.com/consumer/cn/doc/HMSCore-Guides/message-classification-0000001149358835#section5101818813
	DefaultSound *bool `json:"default_sound,omitempty"`
	// 【可选】华为厂商自定义消息优先级。
	//
	// 为了适配华为手机厂商自定义消息的优先级：
	//  - HIGH：非常重要消息，HIGH 级别消息到达用户手机时可强制拉起应用进程。
	//  - NORMAL：重要消息。
	// 设置为 HIGH 需要向华为申请特殊权限，详见 [developer.huawei.com] 文档说明。
	//
	// [developer.huawei.com]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/faq-0000001050042183#section037425218509
	Urgency string `json:"urgency,omitempty"`
	// 【可选】华为回执 ID。
	//  - 仅华为通道有效。
	//  - 输入一个唯一的回执 ID 指定本次下行消息的回执地址及配置，该回执 ID 可以在 [华为回执参数配置] 中查看。
	// [华为回执参数配置]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/msg-receipt-guide-0000001050040176#ZH-CN_TOPIC_0000001087208860__li15263162510251
	ReceiptID string `json:"receipt_id,omitempty"`
	// 【可选】vivo 回执 ID。
	//  - 仅 vivo 通道有效。
	//  - 输入一个唯一的回执 ID 指定本次下行消息的回执地址及配置，该回执 ID 可以在 [vivo 回执参数配置]。
	// [vivo 回执参数配置]: https://dev.vivo.com.cn/documentCenter/doc/681#w2-33657032
	CallbackID string `json:"callback_id,omitempty"`
	// 【可选】华为消息类型，仅华为通道有效。
	//  - 0：普通消息（默认值）；
	//  - 1：测试消息。
	// 每个应用每日可发送测试消息 500 条且不受 [每日单设备推送数量上限要求] 限制。
	//
	// [每日单设备推送数量上限要求]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-restriction-description-0000001361648361#section104849311415
	TargetUserType *int `json:"target_user_type,omitempty"`
	// 【可选】厂商消息大图标样式。
	//  - 支持 华为 / 荣耀 / OPPO 厂商，使用详情参见 [设置图标文档]；
	//  - 优先使用厂商字段，厂商字段没有填充，则使用 [Android 里面定义 LargeIcon 字段] (large_icon)；
	//  - 小米从 2023.08 开始不再支持推送时动态设置小图标、右侧图标、大图片功能；
	//  - 华为、荣耀支持极光的 MediaID 及网络 https 路径；
	//  - OPPO 支持极光的 MediaID 及 OPPO 厂商的大图标 ID；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	// [设置图标文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
	// [Android 里面定义 LargeIcon 字段]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#android
	LargeIcon string `json:"large_icon,omitempty"`
	// 【可选】厂商消息小图标样式。
	//  - 目前支持 华为 / 荣耀 厂商，使用详情参见 [设置图标文档]；
	//  - 优先使用厂商字段，厂商字段没有填充，则使用 [Android 里面定义 SmallIcon 字段] (small_icon_uri)。
	//  - 华为、荣耀支持极光的 MediaID 及厂商本地路径。(小米从 2023.08 开始不再支持推送时动态设置小图标、右侧图标、大图片功能，建议开发者不要继续使用小米相关特性功能)。
	// [设置图标文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
	// [Android 里面定义 SmallIcon 字段]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#android
	SmallIcon string `json:"small_icon_uri,omitempty"`
	// 【可选】小米厂商小图标样式颜色。
	//  - 为了适配小米厂商的消息小图标样式颜色，不填充默认是灰色 (小米官方后续不再支持自定义小图标，建议开发者不要继续使用小米小图标相关特性功能)。
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	// 注意：小米从 2023.08 开始不再支持推送时动态设置小图标、右侧图标、大图片功能，开发者可不再设置此字段值。
	SmallIconColor string `json:"small_icon_color,omitempty"`
	// 【可选】厂商消息 大文本 / Inbox / 大图片 样式。
	//
	// 用来指定厂商的通知栏样式类型，JPush Android SDK v3.9.0 版本以上才支持该字段，默认为 0，其他枚举值：
	//  - style.BigText：大文本通知栏样式 (1)；
	//  - style.Inbox：文本条目通知栏样式 (2)；
	//  - style.BigPicture：大图片通知栏样式 (3)。
	Style style.Style `json:"style,omitempty"`
	// 【可选】厂商消息大文本样式。
	//  - 为了适配厂商的消息大文本样式, 目前支持 小米 / 华为 / 荣耀 / OPPO 厂商。
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 BigText 字段；
	//  - 其中小米最多支持 128 个字符 (一个英文或一个中文算一个字符)，配合小米 Style 使用，OPPO 最多也是支持 128 个字符，配合 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	BigText string `json:"big_text,omitempty"`
	// 【可选】厂商消息 Inbox 样式。
	//  - 为了适配厂商的消息 Inbox 样式, 目前支持华为厂商；
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 Inbox 字段，配合华为 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	Inbox map[string]interface{} `json:"inbox,omitempty"`
	// 【可选】厂商消息大图片样式。
	//  - 为了适配厂商的消息大图片样式，目前支持 OPPO 厂商:
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 BigPicture 字段，配合各自厂商的 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	BigPicture string `json:"big_pic_path,omitempty"`
	// 【可选】是否使用自身通道设置样式。
	//
	// 是否只使用自身通道设置的样式，不使用 Android 里面设置的样式，默认为 false，JPush Android SDK v3.9.0 版本以上才支持该字段。
	//  - true：只使用自身通道设置的样式；
	//  - false：可使用 Android 里面设置的样式。
	OnlyUseVendorStyle *bool `json:"only_use_vendor_style,omitempty"`
	// 【可选】基于第三方审核结果。
	//
	// 目前支持 [华为] / [OPPO] / [vivo] 厂商。
	//
	// 此处直接使用第三方审核结果的返回值原数据填充即可，开发者无需关心各个厂商原始协议，对应推必安信息审核 API 响应内容，详见 [tuibian.mobileservice.cn] 文档说明。
	//
	// [华为]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/android-3rd-party-review-0000001050166008
	// [OPPO]: https://open.oppomobile.com/new/developmentDoc/info?id=11344
	// [vivo]: https://dev.vivo.com.cn/documentCenter/doc/585
	// [tuibian.mobileservice.cn]: https://tuibian.mobileservice.cn/
	AuditResponse map[string]interface{} `json:"auditResponse,omitempty"`
}
