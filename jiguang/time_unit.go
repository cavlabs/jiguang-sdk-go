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
	"strings"
	"time"
)

// 时间单位
type TimeUnit string

const (
	TimeUnitHour  TimeUnit = "HOUR"  // 小时 HOUR，任务执行点 Point 无效！
	TimeUnitDay   TimeUnit = "DAY"   // 天 DAY，任务执行点 Point 无效！
	TimeUnitWeek  TimeUnit = "WEEK"  // 周 WEEK，任务执行点 Point 的取值范围为 [MON, TUE, WED, THU, FRI, SAT, SUN]。
	TimeUnitMonth TimeUnit = "MONTH" // 月 MONTH，任务执行点 Point 的取值范围为 [01, 02, 03, ......, 31]。
)

var timeUnitLayouts = map[TimeUnit]string{
	TimeUnitHour:  "2006-01-02 15",
	TimeUnitDay:   "2006-01-02",
	TimeUnitWeek:  "2006-01-02 Mon",
	TimeUnitMonth: "2006-01",
}

func (tu TimeUnit) IsValid() bool {
	_, exists := timeUnitLayouts[tu]
	return exists
}

func (tu TimeUnit) Layout() string {
	layout, _ := timeUnitLayouts[tu]
	return layout
}

func (tu TimeUnit) Format(st time.Time) string {
	return st.Format(tu.Layout())
}

func (tu *TimeUnit) UnmarshalJSON(data []byte) error {
	val := strings.ToUpper(string(data))
	switch val {
	case `"HOUR"`:
		*tu = TimeUnitHour
	case `"DAY"`:
		*tu = TimeUnitDay
	case `"WEEK"`:
		*tu = TimeUnitWeek
	case `"MONTH"`:
		*tu = TimeUnitMonth
	}
	return nil
}

func (tu TimeUnit) MarshalJSON() ([]byte, error) {
	if tu.IsValid() {
		return []byte(`"` + string(tu) + `"`), nil
	}
	return []byte("null"), nil
}

func (tu TimeUnit) String() string {
	return string(tu)
}
