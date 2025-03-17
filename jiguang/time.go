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
	"time"
)

type Time interface {
	Format() string
}

// ---------------------------------------------------------------------------------------------------------------------

type localTime struct {
	time.Time
}

type (
	LocalDate     localTime
	LocalTime     localTime
	LocalDateTime localTime
)

const (
	localDateFormat     = "2006-01-02"
	localTimeFormat     = "15:04:05"
	localDateTimeFormat = "2006-01-02 15:04:05"
)

var (
	zeroStdTime       = time.Time{}
	zeroLocalDate     = LocalDate{}
	zeroLocalTime     = LocalTime{}
	zeroLocalDateTime = LocalDateTime{}
)

// ---------------------------------------------------------------------------------------------------------------------

func LocalDateNow() LocalDate {
	return LocalDate{time.Now()}
}

func LocalTimeNow() LocalTime {
	return LocalTime{time.Now()}
}

func LocalDateTimeNow() LocalDateTime {
	return LocalDateTime{time.Now()}
}

// ---------------------------------------------------------------------------------------------------------------------

func BuildLocalDate(year, month, day int) LocalDate {
	return LocalDate{time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)}
}

func BuildLocalTime(hour, min, sec int) LocalTime {
	return LocalTime{time.Date(zeroStdTime.Year(), zeroStdTime.Month(), zeroStdTime.Day(), hour, min, sec, 0, time.Local)}
}

func BuildLocalDateTime(year, month, day, hour, min, sec int) LocalDateTime {
	return LocalDateTime{time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)}
}

// ---------------------------------------------------------------------------------------------------------------------

func ParseLocalDate(ds string) (LocalDate, error) {
	st, err := time.Parse(localDateFormat, ds)
	if err != nil {
		return zeroLocalDate, err
	}
	return LocalDate{st}, nil
}

func ParseLocalTime(ts string) (LocalTime, error) {
	st, err := time.Parse(localTimeFormat, ts)
	if err != nil {
		return zeroLocalTime, err
	}
	return LocalTime{st}, nil
}

func ParseLocalDateTime(dts string) (LocalDateTime, error) {
	st, err := time.Parse(localDateTimeFormat, dts)
	if err != nil {
		return zeroLocalDateTime, err
	}
	return LocalDateTime{st}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

func (t *LocalDate) UnmarshalJSON(data []byte) error {
	v := string(data)
	if v == "null" || v == `""` {
		*t = zeroLocalDate
		return nil
	}
	st, err := time.Parse(`"`+localDateFormat+`"`, v)
	if err != nil {
		return err
	}
	*t = LocalDate{st}
	return nil
}

func (t LocalDate) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format(localDateFormat))
}

func (t LocalDate) ToUnitTime() UnitTime {
	return UnitTime{t.Time, TimeUnitDay}
}

func (t LocalDate) Format() string {
	return t.Time.Format(localDateFormat)
}

func (t LocalDate) String() string {
	return t.Time.Format(localDateFormat)
}

// ---------------------------------------------------------------------------------------------------------------------

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	v := string(data)
	if v == "null" || v == `""` {
		*t = zeroLocalTime
		return nil
	}
	st, err := time.Parse(`"`+localTimeFormat+`"`, v)
	if err != nil {
		return err
	}
	*t = LocalTime{st}
	return nil
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	// Note: 与 LocalDate 和 LocalDateTime 不同，因为 00:00:00 也是一个有效的时间，所以这里不判断是否为零值，需要使用方自行判断！
	return json.Marshal(t.Time.Format(localTimeFormat))
}

func (t LocalTime) Format() string {
	return t.Time.Format(localTimeFormat)
}

func (t LocalTime) FormatUsingTimeUnit(_ TimeUnit) string {
	panic("not supported")
}

func (t LocalTime) String() string {
	return t.Time.Format(localTimeFormat)
}

// ---------------------------------------------------------------------------------------------------------------------

func (t *LocalDateTime) UnmarshalJSON(data []byte) error {
	v := string(data)
	if v == "null" || v == `""` {
		*t = zeroLocalDateTime
		return nil
	}
	st, err := time.Parse(`"`+localDateTimeFormat+`"`, v)
	if err != nil {
		return err
	}
	*t = LocalDateTime{st}
	return nil
}

func (t LocalDateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(t.Time.Format(localDateTimeFormat))
}

func (t LocalDateTime) Format() string {
	return t.Time.Format(localDateTimeFormat)
}

func (t LocalDateTime) String() string {
	return t.Time.Format(localDateTimeFormat)
}
