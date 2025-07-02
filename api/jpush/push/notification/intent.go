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

// # 指定通知点击跳转页面
type Intent struct {
	// 【必填】指定跳转页面 URL（指定点击通知栏后跳转的目标页面）。
	//
	// 支持以下三种类型：
	//  1. 跳转到目标页: intent:#Intent;action=action 路径;component= 包名 /Activity 全名;end
	//
	// 	注：OPPO 和 FCM 通道必须传 "action 路径"，其他厂商必须传 "Activity 全名", 否则将出现对应厂商无法跳转问题！
	//  2. 跳转到 deeplink 地址：scheme://test?key1=val1&key2=val2
	//  3. 应用首页: intent:#Intent;action=android.intent.action.MAIN;end（固定为此地址）
	URL string `json:"url"`
}
