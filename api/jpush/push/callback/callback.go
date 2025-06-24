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

package callback

// # 回调参数
//   - 通过指定 Callback 参数，方便用户临时变更回调 URL 或者回调带上其自定义参数，满足其日常业务需求；
//   - 此功能仅针对极光 VIP 用户提供，主要提供消息送达、点击回执数据。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#callback%EF%BC%9A%E5%9B%9E%E8%B0%83%E5%8F%82%E6%95%B0
type Callback struct {
	// 【可选】数据临时回调地址。
	//  - 指定后以此处指定为准，仅针对这一次推送请求生效；
	//  - 不指定，则以极光后台配置为准。
	URL string `json:"url,omitempty"`
	// 【可选】需要回调给用户的自定义参数。
	Params map[string]interface{} `json:"params,omitempty"`
	// 【可选】回调数据类型。
	//  - 可选值为 {1, 2, 3, 8, 9, 10, 11}，包括 Received = 1, Clicked = 2, Push = 8 和它们的任意 or 组合。
	Type Type `json:"type,omitempty"`
}

// ---------------------------------------------------------------------------------------------------------------------

// # 回调数据类型
type Type int

const (
	Received Type = 1 << iota // 送达回执 (1)
	Clicked                   // 点击回执 (2)
	_                         // _
	Push                      // 推送成功回执 (8)

	// ↑↑↑ 可任意 “按位或” 组合以上回调数据类型 ↑↑↑
)
