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

package audience

// # 【自定义通道】注册 ID 发送目标
type CustomChannel struct {
	Instance string   `json:"instance,omitempty"` // 【可选】目标标识，预留字段，目前无效。
	Data     []string `json:"data"`               // 【必填】通道注册 ID 列表。一次发送最多 1000 个。有效性遵循各通道的要求即可。
}
