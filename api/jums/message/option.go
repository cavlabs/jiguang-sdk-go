/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
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

package message

// 消息对象可选参数。
type Option struct {
	SendNo   string `json:"sendno,omitempty"`   // 【可选】纯粹用来作为 API 调用标识，API 返回时被原样返回，以方便 API 调用方匹配请求与返回。
	Owner    string `json:"owner,omitempty"`    // 【可选】提交者用户名，当渠道开启了 API 消息审核时必填。
	BlackID  int    `json:"black_id,omitempty"` // 【可选】黑名单 ID，BlackID 和 WhiteID 不允许同时存在。
	WhiteID  int    `json:"white_id,omitempty"` // 【可选】白名单 ID，BlackID 和 WhiteID 不允许同时存在。
	Priority int    `json:"priority,omitempty"` // 【可选】消息优先级，取值：1（高）、2（中）、3（低）。
}
