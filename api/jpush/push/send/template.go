/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

import "github.com/cavlabs/jiguang-sdk-go/api/jpush/push/options"

// # 模板参数
//
// 创建模板时，开发者设置的变量参数。
type TemplateParam struct {
	// 【可选】若创建模板时有设置变量，此处推送时则必须指定并设置变量名和变量值。
	Keys map[string]string `json:"keys,omitempty"`
	// 【必填】推送目标，支持 2 种类型值：
	//  - push.BroadcastAuds：发广播，给全部设备进行推送；
	//  - 推送设备对象 push.Audience，详见 [docs.jiguang.cn] 文档说明。
	// 注意事项：
	//  - 若指定为 push.BroadcastAuds，则 Params 模板参数列表仅允许配置一个；
	//  - 若指定为推送设备对象 push.Audience 的 RegistrationIDs 推送，则 Params 模板参数列表中的全部 Registration ID 个数最多为 1000 个；
	//  - 若指定为推送设备对象 push.Audience 的 Aliases 推送，则 Params 模板参数列表中的全部 Alias 个数最多为 1000 个；
	//  - 若指定为推送设备对象 push.Audience 的 Tags 推送，则 Params 模板参数列表仅允许配置一个；
	//  - 若指定为推送设备对象 push.Audience 的 Segments 推送，则 Params 模板参数列表仅允许配置一个，且只允许设置一个分群 ID 值。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#audience%EF%BC%9A%E6%8E%A8%E9%80%81%E7%9B%AE%E6%A0%87
	Audience interface{} `json:"audience"`
	// 【可选】推送可选项。
	//  - 详见 [docs.jiguang.cn] 文档说明。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#options%EF%BC%9A%E5%8F%AF%E9%80%89%E5%8F%82%E6%95%B0
	Options *options.Options `json:"options,omitempty"`
	// 【可选】客户自行指定的标识，API 会原样返回。
	TraceID string `json:"trace_id,omitempty"`
	// 【可选】地理围栏推送时，此处指定围栏相关参数。
	Geofence map[string]interface{} `json:"geofence,omitempty"`
}
