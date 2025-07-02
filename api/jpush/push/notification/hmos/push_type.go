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

package hmos

// # 华为场景化消息类型
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://developer.huawei.com/consumer/cn/doc/harmonyos-references-V5/push-scenariozed-api-request-struct-V5
type PushType int

const (
	PushTypeAlert      PushType = 0  // 通知消息
	PushTypeSubscribe  PushType = 0  // 授权订阅消息
	PushTypeFormUpdate PushType = 1  // 卡片刷新消息
	PushTypeExtension  PushType = 2  // 通知扩展消息
	PushTypeBackground PushType = 6  // 后台消息
	PushTypeLiveView   PushType = 7  // 实况窗消息
	PushTypeVoIPCall   PushType = 10 // 应用内通话消息
)
