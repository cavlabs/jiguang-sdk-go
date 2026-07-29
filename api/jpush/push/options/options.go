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

package options

import "github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push/notification/style"

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
	//  - 默认 86400（1 天），普通用户最长 3 天，VIP 用户最长 10 天；设置为 0 表示不保留离线消息，只有推送当前在线的用户可以收到；
	//  - 该字段对 iOS、HarmonyOS、Android 厂商通道，会根据系统推送能力本身支持情况自动适配。
	TimeToLive *int64 `json:"time_to_live,omitempty"`
	// 【可选】要覆盖的消息 ID。
	//
	// 如果当前的推送要覆盖之前的一条推送，这里填写前一条推送的 MsgID 就会产生覆盖效果，即：
	//  - 该 MsgID 离线收到的消息是覆盖后的内容，即使该 MsgID Android 端用户已经收到，如果通知栏还未清除，则新的消息内容会覆盖之前这条通知；
	//  - 覆盖功能起作用的时限是：1 天，如果在覆盖指定时限内该 MsgID 不存在，则返回 1003 错误，提示不是一次有效的消息覆盖操作，当前的消息不会被推送；
	//  - 该字段对 Android 有效，仅支持极光通道、小米通道、OPPO 通道、vivo 通道、FCM 通道、荣耀通道、华为通道（EMUI 10 及以上的设备）和鸿蒙通道。
	OverrideMsgID int64 `json:"override_msg_id,omitempty"`
	// 【可选】APNs 是否生产环境。
	//
	// 该字段仅对 iOS 的 Notification 有效，如果不指定则为推送生产环境。
	//
	// 注意：JPush 服务端 SDK 默认设置为推送「开发环境」。
	//  - true：表示推送生产环境；
	//  - false：表示推送开发环境。
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
	//  - 目前只支持 xiaomi、huawei、honor、meizu、oppo、vivo、fcm、nio、asus、hmos 类型用户，可以一个或者多个同时存在，未传递的通道类型其对应的厂商下发走「默认下发逻辑」：
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
	//
	// 此字段优先级最高，会覆盖 ThirdPartyChannel 的 Vivo.Classification、Honor.Importance 等设置的值。
	// 极光平台默认根据此字段值来判断消息类型，并进行厂商配额扣除处理和厂商消息分类适配逻辑，请务必确保传值准确。
	//
	// 厂商适配使用方式可参考：[厂商消息分类使用指南]
	//
	// [厂商消息分类使用指南]: https://docs.jiguang.cn/jpush/client/Android/android_channel_id
	Classification *int `json:"classification,omitempty"`
	// 【可选】目标转化事件。
	//
	// 目标转化事件支持传递「自定义事件」和「极光预置事件」，目前支持 Android 和 iOS 平台（要求 JPush SDK ≥ v5.0.0 ，且 JCore ≥ v4.2.0），支持通知消息和应用内消息两种消息类型。
	//  - 自定义事件：需集成极光分析 SDK，开发者在极光分析产品中自行创建的业务事件（如：加入购物车、浏览商品等），详情参考 [如何创建自定义事件] 和 [SDK 如何上报自定义事件]；
	//  - 极光预置事件：极光推送 SDK 默认支持，无需开发者创建，也无需集成极光分析 SDK，系统已预置；目标支持的预置事件有 2 个：jg_app_show（应用切换到前台）、jg_app_hide（应用切换到后台）。
	// 代码示例：`{"options": {"target_event": ["jg_app_show"]}}`。
	//
	// 自 2025.10.30 开始，此功能升级为 VIP 用户功能。
	//
	// [如何创建自定义事件]: https://docs.jiguang.cn/public_service/dataCenter/metadata/metaEvent
	// [SDK 如何上报自定义事件]: https://docs.jiguang.cn/public_service/client/Android/sdk_api#%E4%B8%8A%E6%8A%A5%E8%87%AA%E5%AE%9A%E4%B9%89%E4%BA%8B%E4%BB%B6
	TargetEvent []string `json:"target_event,omitempty"`
	// 【可选】是否测试消息标识。
	//  - 优先级大于 HMOS 通知体内的 TestMessage 字段（同样适配鸿蒙自定义消息，如果推送鸿蒙自定义消息，请传递此字段）；
	//  - 请注意区别于 TestModel 功能字段，TestMessage 仅用于适配厂商的测试消息功能，并非表示处于测试模式下推送。
	TestMessage *bool `json:"test_message,omitempty"`
	// 【可选】回执 ID。
	//  - 优先级大于 HMOS 通知体内的 ReceiptID 字段（同样适配鸿蒙自定义消息，如果推送鸿蒙自定义消息，请传递此字段）。
	ReceiptID string `json:"receipt_id,omitempty"`
	// 【可选】是否使用亮屏推送。
	//  - true：使用亮屏推送，false：不使用亮屏推送，默认值 false；
	//  - 适用于实时性要求不高的营销类消息，实时性要求较高的消息不推荐使用；
	//  - 此功能为增值付费服务，需要额外申请权限；
	//  - 当使用亮屏推送时，建议同时设置 NeedBackup 为 true；
	//  - 此功能仅支持单纯通知消息，不支持自定义消息或者通知+自定义消息推送，否则请求会返回 code 码 1035；
	//  - 此功能不支持定速推送，否则请求会返回 code 码 1035；
	//  - 亮屏推送支持的时间范围是每天 7:00 - 22:00；
	//  - 亮屏推送对于 Android 厂商用户的下发策略固定为在线走极光，离线走厂商（如有自行指定厂商下发策略，使用亮屏推送时可不指定策略，或者必须指定为 secondary_push）。
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
	//  - 测试模式推送消息仅推送给到测试设备，请务必先添加测试设备；
	//  - 功能逻辑可参考文档 [测试模式]；
	//  - 请注意区分区别 TestMessage 字段：TestMessage 仅用于适配厂商的测试消息功能，并非表示处于测试模式下推送；TestModel 则表示请求在极光平台下发消息时就已经控制，消息是否仅下发给到测试设备；
	//  - 此功能为增值付费服务，需要额外申请权限。
	// [测试模式]: https://docs.jiguang.cn/jpush/console/push_manage/testmode
	TestModel *bool `json:"test_model,omitempty"`
	// 【可选】是否设置个性化文案。
	AlternateSet *bool `json:"alternate_set,omitempty"`
	// 【可选】地理围栏配置参数。
	Geofence map[string]any `json:"geofence,omitempty"`
	// 【可选】极光 WebPortal 的附加属性。
	PortalExtra *PortalExtraOptions `json:"portal_extra,omitempty"`
	// 【可选】自定义消息转厂商通知功能版本。
	//  - 取值："v1"、"v2"；
	//  - v1：仅支持 Android 平台；
	//  - v2：支持 Android、iOS、HarmonyOS 三个平台，推荐使用 v2 版本。
	// 功能字段和使用逻辑可参考文档 [notification_3rd：自定义消息转厂商通知]。
	//
	// [notification_3rd：自定义消息转厂商通知]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
	Notification3rdVer string `json:"notification_3rd_ver,omitempty"`
	// 【可选】厂商通道消息超长是否自动截断。
	//  - 默认为 true，如果传递的消息体内容发送到厂商通道时发现超长，会自动截断；不期望截断，可以传递 false 关闭；
	//  - 如果在通知消息体下传递了 *pns_content_forshort 字段，会优先使用 *pns_content_forshort 字段作为消息体内容，同时 AutoTruncation 也能生效。
	AutoTruncation *bool `json:"auto_truncation,omitempty"`
	// 【可选】是否启用情景商业 Push。
	//  - 默认为 false：表示普通消息推送；
	//  - true：表示极光情景商业 Push 推送，可「联系商务咨询」。
	// 情景商业 Push 说明：
	//  1. 背景：国内各个厂商，都对消息进行了严格分类管理，并实施了差异化配额管控，比如营销类消息限额大部分情况限额 2 条/设备/应用/天，直接影响 APP 业务关键消息推送，影响用户行为链路转化。
	//  2. 基于上述背景，极光目前和部分厂商有达成深度商业合作，享有付费额度提升特权，突破系统默认运营消息推送条数限制，实现关键营销节点无上限触达，保障关键信息触达目标用户。
	//  3. 此功能为增值付费服务，需要额外申请权限。
	//  4. 使用前需在极光控制台启用「厂商运营超限走商业 Push 通道下发」功能，再通过本字段指定是否使用情景商业 Push。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#%E6%83%85%E6%99%AF%E5%95%86%E4%B8%9Apush-%E8%AF%B4%E6%98%8E
	MktEnable *bool `json:"mkt_enable,omitempty"`
	// 【可选】是否过滤 APP 通知开关关闭的用户。
	//  - 默认为 false：表示不过滤；
	//  - true：表示过滤，极光 VIP 客户专属功能，可「联系商务咨询」。
	NotificationSwitchFilter *bool `json:"notification_switch_filter,omitempty"`
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
	ASUS   *ThirdPartyChannelOptions `json:"asus,omitempty"`   // 华硕通道策略和属性参数。
	HMOS   *ThirdPartyChannelOptions `json:"hmos,omitempty"`   // 鸿蒙通道策略和属性参数。
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
	// 【可选】华为 / 荣耀消息类型，仅华为、荣耀通道有效。
	//  - 0：普通消息（默认值）；
	//  - 1：测试消息。
	// 华为每个应用每日可发送测试消息 500 条且不受 [华为每日单设备推送数量上限要求] 限制；
	// 荣耀每个应用每日可发送测试消息 1000 条且不受 [荣耀每日单设备推送数量上限要求] 限制。
	//
	// [华为每日单设备推送数量上限要求]: https://developer.huawei.com/consumer/cn/doc/development/HMSCore-Guides/message-restriction-description-0000001361648361#section104849311415
	// [荣耀每日单设备推送数量上限要求]: https://developer.honor.com/cn/docs/11002/guides/notification-push-standards#%E4%B8%8D%E5%90%8C%E5%BA%94%E7%94%A8%E7%B1%BB%E5%88%AB%E7%9A%84%E6%8E%A8%E9%80%81%E6%95%B0%E9%87%8F%E4%B8%8A%E9%99%90%E8%A6%81%E6%B1%82
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
	//  - 优先使用厂商字段，厂商字段没有填充，则使用 [Android 里面定义 SmallIcon 字段] (small_icon_uri)；
	//  - 华为、荣耀支持极光的 MediaID 及厂商本地路径。(小米从 2023.08 开始不再支持推送时动态设置小图标、右侧图标、大图片功能，建议开发者不要继续使用小米相关特性功能)；
	//  - 请注意：如果推送的是运营消息（营销消息），荣耀通道不支持携带图标，否则推送会直接返回失败。
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
	//
	// 特别说明：实际展示效果以终端设备为准，由设备系统决定。
	Style style.Style `json:"style,omitempty"`
	// 【可选】厂商消息大文本样式。
	//  - 为了适配厂商的消息大文本样式, 目前支持 小米 / 华为 / 荣耀 / OPPO 厂商。
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 BigText 字段；
	//  - 其中小米最多支持 128 个字符 (一个英文或一个中文算一个字符)，配合小米 Style 使用，OPPO 最多也是支持 128 个字符，配合 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	//
	// 特别说明：实际展示效果以终端设备为准，由设备系统决定。
	BigText string `json:"big_text,omitempty"`
	// 【可选】厂商消息 Inbox 样式。
	//  - 为了适配厂商的消息 Inbox 样式, 目前支持华为厂商；
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 Inbox 字段，配合华为 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	//
	// 特别说明：实际展示效果以终端设备为准，由设备系统决定。
	Inbox map[string]any `json:"inbox,omitempty"`
	// 【可选】厂商消息大图片样式。
	//  - 为了适配厂商的消息大图片样式，目前支持 OPPO 厂商，使用详情参见 [设置大图片文档]；
	//  - 优先使用厂商字段，如果厂商字段没有填充，则使用 Android 里面定义 BigPicture 字段，配合各自厂商的 Style 使用；
	//  - JPush Android SDK v3.9.0 版本以上才支持该字段。
	//
	// 特别说明：实际展示效果以终端设备为准，由设备系统决定。
	//
	// [设置大图片文档]: https://docs.jiguang.cn/jpush/practice/set_icon#android%E3%80%82
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
	AuditResponse map[string]any `json:"auditResponse,omitempty"`
	// 【可选】私信模板 ID。2025.07.14 新增。
	//
	// 仅支持 OPPO 厂商。
	//
	// 下发对应私信模板时必须携带，不支持自拟。
	//
	// 详见：[OPUSH 私信模版校验能力接入说明]
	//
	// [OPUSH 私信模版校验能力接入说明]: https://open.oppomobile.com/documentation/page/info?id=12391
	PrivateMsgTemplateID string `json:"private_msg_template_id,omitempty"`
	// 【可选】标题模板填充参数。2025.07.14 新增。
	//
	// 仅支持 OPPO 厂商。
	//
	// 例：私信模板 ID 标题模板为：`欢迎来到 ${city}$，${city}$ 欢迎您`，此参数内容为：`{"city": "北京"}`。
	//
	// 详见：[OPUSH 私信模版校验能力接入说明]
	//
	// [OPUSH 私信模版校验能力接入说明]: https://open.oppomobile.com/documentation/page/info?id=12391
	PrivateTitleParameters map[string]any `json:"private_title_parameters,omitempty"`
	// 【可选】内容模板填充参数。2025.07.14 新增。
	//
	// 仅支持 OPPO 厂商。
	//
	// 例：私信模板 ID 对应的内容模板为：`欢迎 ${userName}$ 来到 ${city}$`，此参数内容为：`{"userName": "汤姆", "city": "深圳市"}`。
	//
	// 详见：[OPUSH 私信模版校验能力接入说明]
	//
	// [OPUSH 私信模版校验能力接入说明]: https://open.oppomobile.com/documentation/page/info?id=12391
	PrivateContentParameters map[string]any `json:"private_content_parameters,omitempty"`
	// 【可选】消息模板 ID。
	//
	// 仅支持小米厂商。所有私信消息下发时需同时携带 ChannelID 及本字段。
	//
	// 详见：[小米关于消息模板推送新规的更新通知]。
	//
	// [小米关于消息模板推送新规的更新通知]: https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2315
	MiTemplateID string `json:"mi_template_id,omitempty"`
	// 【可选】消息模板参数。
	//
	// 仅支持小米厂商。值为消息模板参数的 JSON 字符串。
	//
	// 例：`{"app_name":"小米商城","order_no":"XM202601130001","order_amount":"11.00","delivery_time":"2026-01-15 18:00"}`。
	//
	// 详见：[小米消息模板接入指南]。
	//
	// [小米消息模板接入指南]: https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2314
	MiTemplateParam string `json:"mi_template_param,omitempty"`
	// 【可选】数字角标操作类型。
	//
	// 目前支持 OPPO 厂商。取值说明：
	//  - 0：覆盖；
	//  - 1：增加。
	//
	// 系统要求：ColorOS 3.1 及以上支持「覆盖」；Android 11 及以上版本支持角标「增加」。
	BadgeOperationType *int `json:"badge_operation_type,omitempty"`
	// 【可选】设置数字角标数量。
	//
	// 目前支持 OPPO 厂商。
	//  - 0：清除数字角标（需要设置 BadgeOperationType = 0（覆盖））；
	//  - 1–99：显示对应的数字角标；
	//  - 大于 99：显示 99+。
	//
	// 当 BadgeOperationType 设置为 1（增加）时，不需要传递此字段，默认 +1。
	BadgeMessageCount *int `json:"badge_message_count,omitempty"`
	// 【可选】是否展示角标。
	//
	// 目前支持 vivo 厂商。该字段透传给 vivo 厂商通道，需要向厂商申请权限，设置后角标自动加 1。
	//
	// 详见：[桌面角标设置方法 - Vpush 接入方法]。
	//
	// 注意：此功能从 JPush Android SDK v5.9.0 版本开始支持。
	//
	// [桌面角标设置方法 - Vpush 接入方法]: https://dev.vivo.com.cn/documentCenter/doc/788
	AddBadge *bool `json:"add_badge,omitempty"`
	// 【可选】华为推送类型。
	//
	// 目前支持华为厂商，极光透传该字段值给华为厂商。
	//
	// 若需使用实况窗通知，请设置本字段为 7，并同时满足：
	//  1. 提供有效的 HwLivePayload；
	//  2. 设置 Classification = 1；
	//  3. 发送策略设为 ospush（或不指定，默认为 ospush）。
	//
	// 推送华为实况窗时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	HwPushType *int `json:"hw_push_type,omitempty"`
	// 【可选】华为实况窗消息体。
	//
	// 用于配置华为实况窗通知，需提前向华为申请相应权限。仅在 HwPushType = 7 时生效。
	//
	// 前置配置：实况窗推送需要独立的校验文件，请在推送前登录极光控制台（Android 厂商集成设置）上传对应的服务密钥文件（与普通厂商 Push 鉴权不同），密钥文件获取参考 [服务账号密钥]。
	//
	// 极光将该字段透传给华为厂商，对应华为 LiveNotificationPayload 字段。具体参数说明请参见 [华为实况窗通知刷新文档]。
	//
	// [服务账号密钥]: https://developer.huawei.com/consumer/cn/doc/start/api-0000001062522591#section3554194116341
	// [华为实况窗通知刷新文档]: https://developer.huawei.com/consumer/cn/doc/HMSCore-References/rest-live-0000001562939968#ZH-CN_TOPIC_0000001700850537__p195121620102511
	HwLivePayload map[string]any `json:"hw_live_payload,omitempty"`
	// 【可选】OPPO 推送类型。
	//
	// 目前支持 OPPO 厂商。枚举值：3（VoIP 消息）、7（实况窗通知）。
	//
	//  - 实况窗通知：设置本字段为 7，并按 OpIntelligentIntent 或 OpDeleteIntentData 说明传参。
	//
	//  - VoIP 通知：参考 OPPO 厂商 [VoIP 消息推送指南] 申请权益。通过极光平台推送时，设置本字段为 3，并同时满足：
	//
	//   1. 提供有效的 VoipExtraData；
	//
	//   2. 设置 Classification = 1；
	//
	//   3. 发送策略设为 ospush（或不指定，默认为 ospush）。
	//
	// 推送 OPPO VoIP 时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// [VoIP 消息推送指南]: https://open.oppomobile.com/documentation/page/info?id=13765
	OpPushType *int `json:"op_push_type,omitempty"`
	// 【可选】意图共享数据结构（OPPO 流体云消息）。
	//
	// 用于配置 OPPO 流体云消息，需提前向 OPPO 申请相应权限。
	//
	// 生效条件：
	//  1. 设置 Classification = 1；
	//  2. 发送策略设为 ospush（或不指定，默认为 ospush）。
	//
	// 前置配置：需在极光控制台（Android 厂商集成设置）完成流体云推送特有鉴权配置（与普通厂商 Push 鉴权不同），鉴权密钥获取参考 [接口认证]。
	//
	// 极光将该字段透传给 OPPO 厂商，对应厂商 IntelligentIntent 字段。具体参数说明参见 OPPO [意图共享数据结构] 文档。校验到本字段不为空时，会认为此次推送是 OPPO 厂商流体云消息。
	//
	// 推送 OPPO 流体云消息时可不必携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// 与 OpDeleteIntentData 二者不可同时存在，二选一。
	//
	// [接口认证]: https://open.oppomobile.com/documentation/page/info?id=13559
	// [意图共享数据结构]: https://open.oppomobile.com/documentation/page/info?id=13565
	OpIntelligentIntent map[string]any `json:"op_intelligent_intent,omitempty"`
	// 【可选】删除意图共享（撤销 OPPO 流体云消息）。
	//
	// 用于撤销 OPPO 流体云消息，需提前向 OPPO 申请相应权限。
	//
	// 生效条件：
	//  1. 设置 Classification = 1；
	//  2. 发送策略设为 ospush（或不指定，默认为 ospush）。
	//
	// 前置配置：需在极光控制台（Android 厂商集成设置）完成流体云推送特有鉴权配置（与普通厂商 Push 鉴权不同），鉴权密钥获取参考 [接口认证]。
	//
	// 极光将该字段透传给 OPPO 厂商，对应厂商销卡协议 data 字段。具体参数说明参见 OPPO [销卡] 文档。校验到本字段不为空时，会认为此次推送是 OPPO 流体云消息。
	//
	// 删除 OPPO 流体云消息时可不必携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// 与 OpIntelligentIntent 二者不可同时存在，二选一。
	//
	// [接口认证]: https://open.oppomobile.com/documentation/page/info?id=13559
	// [销卡]: https://open.oppomobile.com/documentation/page/info?id=13578
	OpDeleteIntentData map[string]any `json:"op_delete_intent_data,omitempty"`
	// 【可选】vivo 推送类型。
	//
	// 目前支持 vivo 厂商。枚举值：2（透传消息）、3（VoIP 消息）、7（原子通知）。
	//
	//  - 透传：设置本字段为 2，并提供有效的 VivoInAppMsg；
	//  - VoIP：设置本字段为 3，并提供有效的 VoipExtraData；
	//  - 原子通知：设置本字段为 7，并提供有效的 VivoLiveMessage。
	//
	// 以上枚举场景均需同时满足：
	//  - 设置 Classification = 1；
	//  - 发送策略为 ospush（或不指定）；
	//  - 可不必携带 Notification、Message 等消息内容顶级字段；
	//  - 如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	VivoPushType *int `json:"vivo_push_type,omitempty"`
	// 【可选】vivo 透传消息体。
	//
	// 用于配置 vivo 透传消息，需提前向 vivo 申请相应权限。传递此字段即视为推送 vivo 透传消息，需同时设置 Classification = 1，发送策略为 ospush（或不指定）。
	//
	// 极光会将字段内容透传至 vivo 厂商，对应厂商字段为 inAppMsg（具体协议需向 vivo 申请获取）。
	//
	// 推送 vivo 透传消息时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	VivoInAppMsg map[string]any `json:"vivo_inapp_msg,omitempty"`
	// 【可选】vivo 原子通知消息体。
	//
	// 用于配置 vivo 原子通知，需提前向 vivo 申请相应权限。传递此字段即视为推送 vivo 原子通知，需同时设置 Classification = 1，发送策略为 ospush（或不指定）。
	//
	// 极光会将字段内容透传至 vivo 厂商，对应厂商字段为 [liveMessage]。
	//
	// 推送 vivo 原子通知时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// [liveMessage]: https://dev.vivo.com.cn/documentCenter/doc/896#s-fdagzbd4
	VivoLiveMessage map[string]any `json:"vivo_liveMessage,omitempty"`
	// 【可选】荣耀推送类型。
	//
	// 目前支持荣耀厂商。枚举值：3（VoIP 消息）。请参考荣耀厂商 [VoIP 能力使用规范] 申请权益。
	//
	// 若需使用 VoIP 消息，请设置本字段为 3，并同时满足：
	//  1. 提供有效的 VoipExtraData；
	//  2. 设置 Classification = 1；
	//  3. 发送策略设为 ospush（或不指定，默认为 ospush）；
	//
	// 推送荣耀 VoIP 消息时可不携带 notification，message 等消息内容顶级字段；如未传递alert字段，极光服务端将自动为 notification.android.alert 赋空值，以确保兼容。
	//
	// [VoIP 能力使用规范]: https://developer.honor.com/cn/docs/11002/guides/notification-voip-standards#2%E3%80%81%E7%94%B3%E8%AF%B7%E6%B5%81%E7%A8%8B
	HonorPushType *int `json:"honor_push_type,omitempty"`
	// 【可选】小米推送类型。
	//
	// 目前支持小米厂商。枚举值：3（VoIP 消息）、7（超级岛通知）。请参考小米厂商 [音视频通话功能限制] 申请权益。
	//
	//  - VoIP：设置本字段为 3，并提供有效的 VoipExtraData，且必须设置 ChannelID；
	//  - 超级岛通知：设置本字段为 7，并提供有效的 MiLivePayload。
	//
	// 以上枚举场景均需同时满足：
	//  - 设置 Classification = 1；
	//  - 发送策略为 ospush（或不指定）；
	//  - 可不必携带 Notification、Message 等消息内容顶级字段；
	//  - 如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// [音视频通话功能限制]: https://dev.mi.com/xiaomihyperos/documentation/detail?pId=1656
	MiPushType *int `json:"mi_push_type,omitempty"`
	// 【可选】小米超级岛通知体。
	//
	// 用于配置小米超级岛通知，需提前向小米申请相应权限。传递此字段即视为推送小米超级岛通知，需同时设置 Classification = 1，发送策略为 ospush（或不指定）。
	//
	// 极光会将字段内容透传至小米厂商，对应厂商字段为 [miui.focus.param 和 miui.focus.pic_XXX]；传参时务必定义小米对应字段 key。
	//
	// 推送小米超级岛通知时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	//
	// [miui.focus.param 和 miui.focus.pic_XXX]: https://dev.mi.com/xiaomihyperos/documentation/detail?pId=2131
	MiLivePayload map[string]any `json:"mi_live_payload,omitempty"`
	// 【可选】VoIP 消息体。
	//
	// 目前支持厂商：荣耀、小米、OPPO、vivo。
	//
	// 若需使用 VoIP 消息，请设置对应厂商的 *PushType = 3，并同时满足：
	//  1. 提供有效的本字段；
	//  2. 设置 Classification = 1；
	//  3. 发送策略设为 ospush（或不指定，默认为 ospush）。
	//
	// 推送 VoIP 消息时可不携带 Notification、Message 等消息内容顶级字段；如未传递 Alert 字段，极光服务端将自动为 Notification.Android.Alert 赋空值，以确保兼容。
	VoipExtraData string `json:"voip_extraData,omitempty"`
	// 【可选】vivo VoIP 消息特有字段。
	//
	// 该字段仅用于 vivo VoIP 消息特性，极光透传厂商，具体特性以厂商为准，参考厂商文档：[VoIP Service Kit Message]。
	//
	// [VoIP Service Kit Message]: https://dev.vivo.com.cn/documentCenter/doc/979#s-kbzqsenz
	ExtensionExpireShow *bool `json:"extensionExpireShow,omitempty"`
}
