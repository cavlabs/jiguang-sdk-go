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
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/notification/hmos"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/notification/style"
)

// # 鸿蒙（HarmonyOS）平台上的通知
type HMOS struct {
	// 【必填】通知内容。
	//  - 这里指定后会覆盖上级统一指定的 Alert 信息；
	//  - 内容不可以是空字符串，否则推送厂商会返回失败。
	Alert string `json:"alert"`
	// 【可选】通知标题。
	//  - 如果指定了，则通知里原来展示 APP 名称的地方，将展示 Title。否则使用 WebPortal 配置的默认 Title。
	Title string `json:"title,omitempty"`
	// 【必填】通知栏消息分类条目。
	//  - 此字段由于厂商为必填字段，效果也完全依赖 ROM 厂商对 Category 的处理策略，请开发者务必填写。极光内部对此字段实际未进行必填校验，请开发者按照必填处理；
	//  - 此字段值对应官方「云端 Category」取值，开发者通过极光服务发起推送时如果传递了此字段值，请务必按照官方要求传递，官方 Category 分类取值规则也可参考 [鸿蒙消息分类标准]。
	// [鸿蒙消息分类标准]: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-noti-classification-0000001727885246#section1521814368537
	Category string `json:"category"`
	// 【可选】通知栏大图标。
	//  - 要求传递网络地址，使用 https 协议，取值样例：https://example.com/image.png；
	//  - 图标大小不超过 30 k，图片 长 × 宽 < 12800 像素。
	LargeIcon string `json:"large_icon,omitempty"`
	// 【可选】指定跳转页面。
	// - 支持跳转到应用首页、deeplink 地址和 action 跳转三种类型：
	//  1. 跳转应用首页：固定 action.system.home
	// 	2. 跳转到 deeplink 地址: scheme://test?key1=val1&key2=val2
	// 	3. 跳转到 action 地址: com.test.action
	Intent *Intent `json:"intent,omitempty"`
	// 【可选】设置角标数字累加值。
	//  - 此字段如果不填，表示不改变角标数字；
	//  - 取值范围为：1～99，若设置了取值范围内的数字，下一条通知栏消息配置的 BadgeAddNum 数据会和原角标数量进行相加，建议取值为 1。
	//  举例：BadgeAddNum 取值为 1，原角标数为 2，发送此角标消息后，应用角标数显示为 3。
	BadgeAddNum *int `json:"badge_add_num,omitempty"`
	// 【可选】设置角标数字为固定值。
	//  - 此字段如果不填，表示不改变角标数字；
	//  - 取值范围为：0～99，若设置了取值范围内的数字，对应下一条通知栏消息配置的 BadgeSetNum 数字则为角标数值。
	//  举例：BadgeSetNum 取值为 1，无论应用之前角标数为多少，发送此角标消息后，应用角标数均显示为 1。
	BadgeSetNum *int `json:"badge_set_num,omitempty"`
	// 【可选】测试消息标识。
	//  - false：正常消息（默认值）；
	//  - true：测试消息。
	TestMessage *bool `json:"test_message,omitempty"`
	// 【可选】华为回执 ID。
	//  - 输入一个唯一的回执 ID 指定本次下行消息的回执地址及配置，该回执 ID 可以在 [鸿蒙回执参数配置] 中查看。
	// [鸿蒙回执参数配置]: https://docs.jiguang.cn/jpush/client/HarmonyOS/hmos_3rd_param#%E9%B8%BF%E8%92%99%E9%80%9A%E9%81%93%E5%9B%9E%E6%89%A7%E9%85%8D%E7%BD%AE%E6%8C%87%E5%8D%97
	ReceiptID string `json:"receipt_id,omitempty"`
	// 【可选】扩展字段。
	//  - 这里自定义 JSON 格式的 key/value 信息，以供业务使用。
	Extras map[string]interface{} `json:"extras,omitempty"`
	// 【可选】通知栏样式类型。
	//  - 默认为 0：0-普通样式，2-多行文本样式 (style.Inbox)。
	Style style.Style `json:"style,omitempty"`
	// 【可选】多行文本样式。
	//  - 对应 Style 的取值类型 style.Inbox。
	Inbox map[string]interface{} `json:"inbox,omitempty"`
	// 【可选】推送类型。
	//
	// 对应华为 push-type 字段，默认值 0 (hmos.PushTypeAlert)，目前仅支持：
	//  - 0: 通知消息 (hmos.PushTypeAlert)
	//  - 1: 卡片刷新消息 (hmos.PushTypeFormUpdate)
	//  - 2: 通知拓展消息 (hmos.PushTypeExtension)
	//  - 7: 实况窗消息 (hmos.PushTypeLiveView)
	//  - 10: VoIP 呼叫消息 (hmos.PushTypeVoIPCall)
	// 其它值报错，VoIP 消息与通知消息互斥，不可同时下发。
	//
	// 详细取值说明可参考官方 [场景消息介绍]。
	//
	// 补充说明：官方 PushType = hmos.PushTypeBackground 的后台消息，对应 [极光自定义消息]，不属于通知消息范畴，所以此处传值有限制不支持 hmos.PushTypeBackground。
	//
	// [场景消息介绍]: https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-intro
	// [极光自定义消息]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#message%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
	PushType hmos.PushType `json:"push_type"`
	// 【可选】附加数据。
	//  - 对应华为 extraData 字段，当 PushType = hmos.PushTypeExtension 或 PushType = hmos.PushTypeVoIPCall 时生效，此时是必填的，
	//  PushType = hmos.PushTypeAlert 时忽略此字段。
	ExtraData string `json:"extra_data,omitempty"`
	// 【可选】APP 在前台，通知是否展示。
	//  - 值为 "1" 时，APP 在前台会弹出/展示通知栏消息；
	//  - 值为 "0" 时，APP 在前台不会弹出/展示通知栏消息；
	//  - 默认情况下 APP 在前台会弹出/展示通知栏消息。
	DisplayForeground string `json:"display_foreground,omitempty"`
	// 【可选】自定义通知铃声。
	//  - 该字段对推送走极光通道 (JPush HarmonyOS SDK ≥ v1.3.0 版本) 和鸿蒙厂商通道通知均生效；
	//  - 此处设置的铃声文件必须放在应用的 /resources/rawfile 路径下，详见官方 [Sound 字段说明]；
	//  - 涉及到开发者需要向鸿蒙官方 [申请自定义铃声权益]。
	// [Sound 字段说明]: https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section17371529101117
	// [申请自定义铃声权益]: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-apply-right#section623416398277
	Sound string `json:"sound,omitempty"`
	// 【可选】自定义通知铃声时长，需要配合 Sound 字段使用。
	//  - 该字段只对推送走鸿蒙厂商通道生效；
	//  - 只有当请求同时携带 Sound 字段，SoundDuration 字段才会生效。仅支持数字，单位为秒，取值范围 [1, 60]，详见官方 [SoundDuration 字段说明]；
	//  - 涉及到开发者需要向鸿蒙官方 [申请自定义铃声权益]。
	// [SoundDuration 字段说明]: https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-param#section17371529101117
	// [申请自定义铃声权益]: https://developer.huawei.com/consumer/cn/doc/harmonyos-guides/push-apply-right#section623416398277
	SoundDuration int `json:"sound_duration,omitempty"`
	// 【可选】鸿蒙卡片刷新消息、实况窗消息内容体。
	//  - 仅对 PushType = hmos.PushTypeFormUpdate（卡片刷新消息）、PushType = hmos.PushTypeLiveView（实况窗消息）生效，且 PushType = hmos.PushTypeFormUpdate 或 PushType = hmos.PushTypeLiveView 时要求此字段必填，上述其它必填字段可忽略不填；
	//  - 极光 HmPayload 字段值会直接透传给鸿蒙 payload 字段，极光不对消息体完整性和正确性进行校验，请开发者对照官方文档说明进行传参，详见：[鸿蒙场景化消息请求示例]；
	//  - 当您推送 “卡片刷新” 或 “实况窗” 消息时，推送仅支持通过鸿蒙通道下发，推送下发策略仅支持 ospush（下发策略默认值 ospush，如自定义其它下发策略请求报错）。
	// [鸿蒙场景化消息请求示例]: https://developer.huawei.com/consumer/cn/doc/harmonyos-references/push-scenariozed-api-request-example
	HmPayload map[string]interface{} `json:"hm_payload,omitempty"`
}
