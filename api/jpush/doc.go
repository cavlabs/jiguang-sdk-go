// Copyright 2026 cavlabs/jiguang-sdk-go authors.
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

// Package jpush 提供极光推送（JPush）产品级入口。
//
//	client, err := jpush.New(
//		jpush.WithApp(appKey, masterSecret),
//		jpush.WithLogger(logger),
//	)
//	result, err := client.Push.Send(ctx, param)    // 默认 REST v3
//	result, err := client.Push.V3.Send(ctx, param) // 显式版本
//
// 未配置对应凭证时仍可 New 成功；调用未就绪的子 API 会返回 api.ErrNil*。
// 进阶场景仍可直接使用各子包 NewAPIv*Builder。
package jpush
