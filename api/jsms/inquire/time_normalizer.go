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

package inquire

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// 时间字符串格式
const (
	timeWithMicrosecsLayout        = "2006-01-02 15:04:05.000000"   // 普通格式，带 6 位微秒值
	timeWithMicrosecsISO8601Layout = "2006-01-02T15:04:05.000000"   // ISO 8601 格式，带 6 位微秒值
	timeCSTLayout                  = "Mon Jan 02 15:04:05 MST 2006" // CST 格式，不带 6 位微秒值
)

// 匹配时间字符串中的毫秒/微秒部分，即形如 .123 或 .123456 的部分
var timeMilliRegex = regexp.MustCompile(`\.\d+`)

func normalizeAndParseTime(input string) (t time.Time, err error) {
	matches := timeMilliRegex.FindString(input)

	normalizedInput := input
	if matches != "" {
		// 补全到 6 位
		microsecs := matches[1:] // 去掉前面的点 "."
		if len(microsecs) < 6 {
			microsecs += strings.Repeat("0", 6-len(microsecs))
		}
		// 替换原字符串中的毫秒部分
		normalizedInput = strings.Replace(input, matches, fmt.Sprintf(".%s", microsecs), 1)
	} else {
		// 如果没有毫秒部分，补全为 .000000
		normalizedInput += ".000000"
	}

	if t, err = time.Parse(timeWithMicrosecsLayout, normalizedInput); err == nil {
		return
	}
	if t, err = time.Parse(timeWithMicrosecsISO8601Layout, normalizedInput); err == nil {
		return
	}
	if t, err = time.Parse(timeCSTLayout, input); err == nil {
		return
	}
	return t, fmt.Errorf("invalid time input: %s", input)
}
