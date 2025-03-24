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

package send

import (
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/callback"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/liveactivity"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/message"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/options"
)

// # 推送参数
type Param struct {
	// 【可选】用于防止 API 调用端重试造成服务端的重复推送而定义的一个标识符，可通过 GetCidForPush 接口获取。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E8%8E%B7%E5%8F%96%E6%8E%A8%E9%80%81%E5%94%AF%E4%B8%80%E6%A0%87%E8%AF%86cid
	CID string `json:"cid,omitempty"`
	// 【必填】推送平台，支持 2 种类型值：
	//  - platform.All：推送到所有平台；
	//  - platform.Android、platform.IOS、platform.QuickApp、platform.HMOS 的组合列表：指定特定推送平台。
	// 注意事项：
	//  - 如果目标平台为 iOS 平台，推送 Notification 时需要在 Options 中通过 ApnsProduction 字段来设定推送环境；
	//  - true 表示推送生产环境，false 表示要推送开发环境，如果不指定则为推送生产环境，一次只能推送给一个环境。
	// 详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#platform%EF%BC%9A%E6%8E%A8%E9%80%81%E5%B9%B3%E5%8F%B0。
	Platform interface{} `json:"platform"`
	// 【必填】推送目标，支持 2 种类型值：
	//  - push.BroadcastAuds：发广播，给全部设备进行推送；
	//  - 推送设备对象 push.Audience，详见 [docs.jiguang.cn] 文档说明。
	// 备注：
	//  - 基于业务优化的需求，极光于 2020 年 3 月 10 日对「广播推送」的频率进行限制，调整为 10 次每天，超过调用限制时将返回报错码 2008，官网控制台将与 Push API 同步调整。
	//  - 本次调整仅限制广播，对广播外的推送不影响。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#audience%EF%BC%9A%E6%8E%A8%E9%80%81%E7%9B%AE%E6%A0%87
	Audience interface{} `json:"audience,omitempty"`
	// 【可选】推送可选项。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#options%EF%BC%9A%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0
	Options *options.Options `json:"options,omitempty"`
	// 【可选】通知内容，是被推送到客户端的内容；与 CustomMessage 一起二者必须有其一，可以二者并存。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification%EF%BC%9A%E9%80%9A%E7%9F%A5
	Notification *notification.Notification `json:"notification,omitempty"`
	// 【可选】自定义消息内容，是被推送到客户端的内容；与 Notification 一起二者必须有其一，可以二者并存。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#message%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF
	CustomMessage *message.Custom `json:"message,omitempty"`
	// 【可选】实时活动内容。不可与 Notification 或 CustomMessage 等并存。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#live_activity%EF%BC%9A%E5%AE%9E%E6%97%B6%E6%B4%BB%E5%8A%A8%E6%B6%88%E6%81%AF
	LiveActivity *liveactivity.Message `json:"live_activity,omitempty"`
	// 【可选】应用内增强提醒，面向于通知栏消息类型，需搭配 Notification 参数一起使用，对于通知权限关闭的用户可设置启用此功能。不可与 CustomMessage 同时并存。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#inapp_message%EF%BC%9A%E5%BA%94%E7%94%A8%E5%86%85%E5%A2%9E%E5%BC%BA%E6%8F%90%E9%86%92
	InApp *message.InApp `json:"inapp_message,omitempty"`
	// 【可选】自定义消息转厂商通知内容。与 CustomMessage 一起使用。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#notification_3rd%EF%BC%9A%E8%87%AA%E5%AE%9A%E4%B9%89%E6%B6%88%E6%81%AF%E8%BD%AC%E5%8E%82%E5%95%86%E9%80%9A%E7%9F%A5
	ThirdNotification *notification.Third `json:"notification_3rd,omitempty"`
	// 【可选】短信渠道补充送达内容。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1
	SmsMessage *message.SMS `json:"sms_message,omitempty"`
	// 【可选】回调参数。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#callback%EF%BC%9A%E5%9B%9E%E8%B0%83%E5%8F%82%E6%95%B0
	Callback *callback.Callback `json:"callback,omitempty"`
}
