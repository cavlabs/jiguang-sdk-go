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

package liveactivity

// # 实时活动内容
//
// 实时活动消息要求使用 iOS P8 证书，对应 [极光 WebPortal 集成设置中 iOS 鉴权方式需要选择「Token Authentication 配置」] 方式。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [极光 WebPortal 集成设置中 iOS 鉴权方式需要选择「Token Authentication 配置」]: https://docs.jiguang.cn/jpush/console/push_setting/integration_set#ios
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#live_activity%EF%BC%9A%E5%AE%9E%E6%97%B6%E6%B4%BB%E5%8A%A8%E6%B6%88%E6%81%AF
type Message struct {
	// 【必填】iOS 的实时活动消息。
	//   - iOS 实时活动消息（Live Activity）JPush 要转发给苹果服务器。苹果要求实时活动消息（ActivityKit）远程推送的动态更新数据大小不超过 4096 字节；
	//   - JPush 因为需要重新组包，并且考虑一点安全冗余，要求 "ios":{} 及大括号内的总体长度不超过：3584 个字节。JPush 使用 UTF-8 编码，所以一个汉字占用 3 个字节长度。
	IOS *IosMessage `json:"ios"`
}

// # iOS 的实时活动消息
type IosMessage struct {
	// 【必填】实时活动事件类型。
	Event Event `json:"event"`
	// 【必填】实时活动动态内容，需与客户端 SDK 值匹配（对应 Apple 官方的 [content-state 字段]）。
	//
	// [content-state 字段]: https://developer.apple.com/documentation/activitykit/updating-and-ending-your-live-activity-with-activitykit-push-notifications
	ContentState map[string]interface{} `json:"content-state"`
	// 【可选】实时活动属性类型，开发者自定义值，当 Event 为 EventStart 时该参数必填。
	AttributesType string `json:"attributes-type,omitempty"`
	// 【可选】实时活动属性。
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// 【可选】实时活动通知内容。
	Alert *IosAlertMessage `json:"alert,omitempty"`
	// 【可选】实时活动结束展示时间。
	DismissalDate int64 `json:"dismissal-date,omitempty"`
	// 【可选】实时活动显示过期时间，如果该时间小于当前时间，实时活动将不会更新。
	StaleDate int64 `json:"stale-date,omitempty"`
	// 【可选】实时活动在灵动岛上展示的优先级，取值范围为 [1, 100]，该值和实时活动的重要性呈正相关，不填默认为最高。
	RelevanceScore int `json:"relevance-score,omitempty"`
	// 【可选】为 5 或 10 ，不填默认为 10。因为时候活动通知每小时是有频控限制的，ApnsPriority = 5 的通知将不消耗苹果厂商频控配额，当超出频控上限，推送通知将被限制。
	ApnsPriority int `json:"apns-priority,omitempty"`
}

// # iOS 实时活动通知内容
type IosAlertMessage struct {
	// 【可选】显示到 Apple Watch 的消息标题。
	Title string `json:"title,omitempty"`
	// 【可选】显示到 Apple Watch 的消息内容。
	Body string `json:"body,omitempty"`
	// 【可选】提示音。
	Sound string `json:"sound,omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------

// # 实时活动事件类型
type Event string

const (
	EventStart  Event = "start"  // 创建
	EventUpdate Event = "update" // 更新
	EventEnd    Event = "end"    // 结束
)
