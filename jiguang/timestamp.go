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

package jiguang

import (
	"encoding/json"
	"strconv"
	"time"
)

type Timestamp localTime

var zeroTimestamp = Timestamp{}

func TimestampNow() Timestamp {
	return Timestamp{time.Now()}
}

func BuildTimestamp(year, month, day, hour, min, sec int) Timestamp {
	return Timestamp{time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)}
}

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	v := string(data)
	if v == "null" || v == `""` {
		*t = zeroTimestamp
		return
	}
	ts, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		*t = Timestamp{time.Unix(ts, 0)}
		if t.Year() > 3000 { // 处理时间戳为毫秒的情况
			*t = Timestamp{time.Unix(0, ts*1e6)}
		}
	} else {
		var st time.Time
		st, err = time.Parse(`"`+time.RFC3339+`"`, v)
		if err == nil {
			*t = Timestamp{st}
			return
		}
		st, err = time.Parse(`"`+localDateTimeFormat+`"`, v)
		if err == nil {
			*t = Timestamp{st}
			return
		}
	}
	return
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Unix())
}

func (t *Timestamp) GetTime() *time.Time {
	if t == nil {
		return nil
	}
	return &t.Time
}

func (t Timestamp) ToLocalDateTime() LocalDateTime {
	return LocalDateTime(t)
}

func (t Timestamp) Equal(u Timestamp) bool {
	return t.Time.Equal(u.Time)
}

func (t Timestamp) Format() string {
	return strconv.FormatInt(t.Unix(), 10)
}

func (t Timestamp) String() string {
	return strconv.FormatInt(t.Unix(), 10)
}
