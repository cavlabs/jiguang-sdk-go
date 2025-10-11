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

// # 快应用平台上的通知
//
// 自 2025.10.30 开始，此功能升级为 VIP 用户功能。
type QuickApp struct {
	// 【必填】通知内容。
	//  - 这里指定了，则会覆盖上级统一指定的 Alert 信息。
	Alert string `json:"alert"`
	// 【必填】通知标题。
	//  - 快应用推送通知的标题。
	Title string `json:"title"`
	// 【必填】跳转页面。
	//  - 快应用通知跳转地址。
	Page string `json:"page"`
	// 【可选】附加字段。
	//  - 这里自定义 key/value 信息，以供业务使用。
	Extras map[string]interface{} `json:"extras,omitempty"`
}
