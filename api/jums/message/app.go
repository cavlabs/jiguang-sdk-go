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

package message

import "github.com/cavlabs/jiguang-sdk-go/api/jpush/push"

// # 【APP】消息
type App = push.SendParam

// ---------------------------------------------------------------------------------------------------------------------

// # APP 参数
type AppParams struct {
	// 【必填】APP 推送平台设置，支持 2 种类型值：
	//  - platform.All：推送到所有平台；
	//  - platform.Android、platform.IOS、platform.QuickApp、platform.Web 的组合列表：指定特定推送平台。
	Platform interface{} `json:"platform"`
	// 【可选】离线消息保留时长（单位：秒）。
	//  - 默认 86400（1 天），最长 10 天。
	TimeToLive *int64 `json:"time_to_live,omitempty"`
	// 【可选】APNs 是否生产环境。
	//  - true 表示推送生产环境，false 表示要推送开发环境；
	//  - 如果不指定则为推送生产环境。
	ApnsProduction *bool `json:"apns_production,omitempty"`
}
