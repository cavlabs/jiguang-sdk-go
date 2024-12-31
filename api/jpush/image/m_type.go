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

package image

// 图片类型
type Type int

const (
	BigImage  Type = iota + 1 // 大图片 (1)
	BigIcon                   // 大图标 (2)
	SmallIcon                 // 小图标 (3)
)

func (t Type) IsValid() bool {
	switch t {
	case BigImage, BigIcon, SmallIcon:
		return true
	default:
		return false
	}
}
