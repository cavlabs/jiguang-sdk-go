/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

// # 设备平台
//
// 当前支持 android、ios、web、quickapp、hmos 5 个平台，其中 “web 平台” 在 JUMS 的【模板消息发送 API】的 [APP 参数] 中可用。
//
// [APP 参数]: https://docs.jiguang.cn/jums/server/rest_api_jums_template_message#app_para%EF%BC%9Aapp-%E5%8F%82%E6%95%B0
type Platform string

const (
	All = Platform("all") // 所有平台

	// ↓↓↓ 支持的设备平台枚举值 ↓↓↓

	Android  = Platform("android")  // Android 平台
	IOS      = Platform("ios")      // iOS 平台
	Web      = Platform("web")      // Web 平台
	QuickApp = Platform("quickapp") // 快应用平台
	HMOS     = Platform("hmos")     // 鸿蒙平台
)

func (p *Platform) UnmarshalJSON(data []byte) error {
	val := string(data)
	switch val {
	case `"all"`, `-1`:
		*p = All
	case `"android"`, `0`:
		*p = Android
	case `"ios"`, `1`:
		*p = IOS
	case `"web"`, `2`:
		*p = Web
	case `"quickapp"`, `3`:
		*p = QuickApp
	case `"hmos"`, `4`:
		*p = HMOS
	}
	return nil
}

func (p Platform) MarshalJSON() ([]byte, error) {
	switch p {
	case All:
		return []byte(`"all"`), nil
	case Android, IOS, Web, QuickApp, HMOS:
		return []byte(`"` + string(p) + `"`), nil
	default:
		return nil, nil
	}
}

func (p Platform) Index() int {
	switch p {
	case All:
		return -1
	case Android:
		return 0
	case IOS:
		return 1
	case Web:
		return 2
	case QuickApp:
		return 3
	case HMOS:
		return 4
	default:
		return -2
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
