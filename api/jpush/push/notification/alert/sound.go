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

// iOS 消息通知声音。
type IosSound struct {
	// “重要警告” 标志。设置为 1 以启用 “重要警告”。
	Critical int `json:"critical"`
	// 声音文件的名称，文件应位于应用主包或应用容器目录的 Library/Sounds 文件夹中。指定字符串 "default" 可播放系统默认声音。
	Name string `json:"name,omitempty"`
	// 关键通知声音的音量。值应在 0（静音）到 1（最大音量）之间。
	Volume float64 `json:"volume"`
}
