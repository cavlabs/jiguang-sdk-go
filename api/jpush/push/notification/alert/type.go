/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
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

package alert

// 通知提醒方式。
type Type int

const (
	DefaultSound   Type = 1 << iota // 提示音 (DEFAULT_SOUND)，1 (0x00000001)
	DefaultVibrate                  // 震动 (DEFAULT_VIBRATE)，2 (0x00000002)
	DefaultLights                   // 指示灯 (DEFAULT_LIGHTS)，4 (0x00000004)

	// ↑↑↑ 可任意 “按位或” 组合以上默认通知提醒方式 ↑↑↑

	DefaultAll Type = -1 // 启用所有 (DEFAULT_ALL)，-1 (0xffffffff)
)
