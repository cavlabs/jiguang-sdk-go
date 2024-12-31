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

package jiguang

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

// UnitTime 组合了 time.Time 及 jiguang.TimeUnit。
type UnitTime struct {
	time.Time
	TimeUnit
}

// 用于匹配输入日期时间字符串的正则表达式。
var unitTimePattern = regexp.MustCompile(`^(\d{4})-(\d{2})(?:-(\d{2})(?: (\w{3}|\d{2}))?)?$`)

// 尝试使用正则表达式匹配输入的日期时间字符串 ts 以识别有效的时间单位。
//  - 如果匹配成功，则返回对应的时间单位，ok 返回 true；否则 ok 返回 false。
func tryParseTimeUnit(ts string) (tu TimeUnit, ok bool) {
	matches := unitTimePattern.FindStringSubmatch(ts)
	if matches == nil {
		ok = false
		return
	}

	switch {
	case matches[3] == "" && matches[4] == "":
		tu, ok = TimeUnitMonth, true
	case matches[3] != "" && matches[4] == "":
		tu, ok = TimeUnitDay, true
	case matches[3] != "" && len(matches[4]) == 3:
		tu, ok = TimeUnitWeek, true
	case matches[3] != "" && len(matches[4]) == 2:
		tu, ok = TimeUnitHour, true
	default:
		ok = false
	}

	return
}

func (ut *UnitTime) UnmarshalJSON(data []byte) error {
	v := string(data)
	if v == "null" || v == `""` {
		return nil
	}

	v = strings.Trim(v, `"`)
	// 尝试从时间字符串中解析出时间单位
	if tu, ok := tryParseTimeUnit(v); ok {
		t, err := time.Parse(tu.Layout(), v)
		if err != nil {
			return nil
		}
		*ut = UnitTime{t, tu}
	}
	return nil
}

func (ut UnitTime) MarshalJSON() ([]byte, error) {
	if ut.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ut.Format())
}

func (ut UnitTime) ToLocalDate() LocalDate {
	return LocalDate{ut.Time}
}

func (ut UnitTime) ToLocalTime() LocalTime {
	return LocalTime{ut.Time}
}

func (ut UnitTime) ToLocalDateTime() LocalDateTime {
	return LocalDateTime{ut.Time}
}

func (ut UnitTime) Format() string {
	return ut.Time.Format(ut.Layout())
}

func (ut UnitTime) String() string {
	return ut.Format() + " " + ut.TimeUnit.String()
}
