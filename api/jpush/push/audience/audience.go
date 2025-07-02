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

package audience

const All = "all" // 广播推送，表示推送给所有设备 (all)。

// # 推送设备对象
//
// 表示一条推送可以被推送到哪些设备列表。
//
// 确认推送设备对象，JPush 提供了多种方式，比如：别名、标签、注册 ID、分群、广播等，详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#audience%EF%BC%9A%E6%8E%A8%E9%80%81%E7%9B%AE%E6%A0%87
type Audience struct {
	// 【可选】注册 ID 列表，多个注册 ID 之间是 OR 的关系，即取并集。
	//  - 设备标识 Registration ID，客户端集成 SDK 后可获取到该值；
	//  - 一次推送最多 1000 个；
	//  - 如果您一次推送的 Registration ID 值超过 1000 个，可以直接使用 “文件推送” 功能。
	RegistrationIDs []string `json:"registration_id,omitempty"`
	// 【可选】标签列表，多个标签之间是 OR 的关系，即取并集。
	//  - 用标签来进行大规模的设备属性、用户属性分群，此功能为 VIP 用户功能；
	//  - 一次推送最多 20 个；
	//  - 有效的标签组成：字母（区分大小写）、数字、下划线、汉字、特殊字符 @!#$&*+=.|￥；
	//  - 限制：每一个标签的长度限制为 40 字节（判断长度需采用 UTF-8 编码）。
	Tags []string `json:"tag,omitempty"`
	// 【可选】标签 AND 列表，多个标签之间是 AND 的关系，即取交集。
	//  - 此功能为 VIP 用户功能，注意与 TagOrList 区分；
	//  - 一次推送最多 20 个。
	AndTags []string `json:"tag_and,omitempty"`
	// 【可选】标签 NOT 列表，多个标签之间，先取多标签的并集，再对该结果取补集。
	//  - 此功能为 VIP 用户功能；
	//  - 一次推送最多 20 个。
	NotTags []string `json:"tag_not,omitempty"`
	// 【可选】别名列表，多个别名之间是 OR 的关系，即取并集。
	//  - 用别名来标识一个用户，一个设备只能绑定一个别名，但多个设备可以绑定同一个别名；
	//  - 一次推送最多 1000 个；
	//  - 有效的别名组成：字母（区分大小写）、数字、下划线、汉字、特殊字符 @!#$&*+=.|￥；
	//  - 限制：每一个别名的长度限制为 40 字节（判断长度需采用 UTF-8 编码）。
	Aliases []string `json:"alias,omitempty"`
	// 【可选】用户分群 ID 列表。
	//  - 在页面创建的用户分群的 ID。定义为数组，但目前限制一次只能推送一个。
	Segments []string `json:"segment,omitempty"`
	// 【可选】A/B Test ID 列表。
	//  - 在页面创建的 A/B 测试的 ID。定义为数组，但目前限制是一次只能推送一个。
	AbTests []string `json:"abtest,omitempty"`
	// 【可选】实时活动标识。
	//  - 对应 iOS SDK liveActivityId 的值，参考客户端：registerLiveActivity；
	//  - 这种方式不能和其它 Audience 方式组合使用，比如不允许（LiveActivityID + Tags 组合）。
	LiveActivityID string `json:"live_activity_id,omitempty"`

	// 【可选】指定文件推送。
	//  - 可用于包括 SendByFile（文件立即推送）和 ScheduleSend（文件定时推送）等相关接口。
	//  - 详见 [文件推送 API] 文档说明。
	// [文件推送 API]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%96%87%E4%BB%B6%E6%8E%A8%E9%80%81-api
	File *File `json:"file,omitempty"`
}
