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

package platform

import "strings"

// 「极光推送」设备平台，当前支持 android、ios、hmos、quickapp 4 个平台。
type Platform string

const (
	All = Platform("all") // 所有平台，即 android、ios、hmos、quickapp 4 个平台。

	// ↓↓↓ 支持的设备平台枚举值 ↓↓↓

	Android  = Platform("android")  // Android 平台
	IOS      = Platform("ios")      // iOS 平台
	HMOS     = Platform("hmos")     // 鸿蒙平台
	QuickApp = Platform("quickapp") // 快应用平台
)

func (p *Platform) UnmarshalJSON(data []byte) error {
	val := string(data)
	switch val {
	case `"all"`:
		*p = All
	case `"android"`:
		*p = Android
	case `"ios"`:
		*p = IOS
	case `"hmos"`:
		*p = HMOS
	case `"quickapp"`:
		*p = QuickApp
	}
	return nil
}

func (p Platform) MarshalJSON() ([]byte, error) {
	switch p {
	case All:
		return []byte(`"all"`), nil
	case Android, IOS, HMOS, QuickApp:
		return []byte(`"` + string(p) + `"`), nil
	default:
		return nil, nil
	}
}

func (p Platform) String() string {
	return string(p)
}

// 将多个平台值连接为一个字符串，以 sep 分隔符分隔
func Concat(plats []Platform, sep string) string {
	size := len(plats)
	if size == 0 {
		return ""
	}
	s := make([]string, size)
	for i, p := range plats {
		if p == All {
			return string(All)
		}
		s[i] = string(p)
	}
	return strings.Join(s, sep)
}
