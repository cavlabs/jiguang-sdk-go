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

package report

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// # 用户统计（VIP）
//   - 功能说明：提供近 2 个月内某时间段的用户相关统计数据：新增用户、在线用户、活跃用户。
//     `start` 起始时间，它的时间单位支持：HOUR（小时，格式例：2014-06-11 09）、DAY（天，格式例：2014-06-11）、MONTH（月，格式例：2014-06）；
//     `duration` 持续时长，如果时间单位是天，则是持续的天数，以此类推；只支持查询 60 天以内的用户信息，对于时间单位为 HOUR 的，只支持输出当天的统计结果。
//   - 调用地址：GET `/v3/users`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E7%94%A8%E6%88%B7%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
func (r *apiv3) GetUserDetail(ctx context.Context, start jiguang.UnitTime, duration int) (*UserDetailGetResult, error) {
	if r == nil {
		return nil, api.ErrNilJPushReportAPIv3
	}

	tu := start.TimeUnit
	if tu == jiguang.TimeUnitWeek || !tu.IsValid() {
		return nil, errors.New("invalid `start` TimeUnit, only support HOUR, DAY, MONTH")
	}
	switch tu {
	case jiguang.TimeUnitHour:
		if duration < 1 || duration > 24 {
			return nil, errors.New("invalid `duration` value, only support 1-24")
		}
	case jiguang.TimeUnitDay:
		if duration < 1 || duration > 60 {
			return nil, errors.New("invalid `duration` value, only support 1-60")
		}
	case jiguang.TimeUnitMonth:
		if duration < 1 || duration > 2 {
			return nil, errors.New("invalid `duration` value, only support 1-2")
		}
	}

	query := "?time_unit=" + tu.String() + "&start=" + url.QueryEscape(start.Format()) + "&duration=" + strconv.Itoa(duration)
	req := &api.Request{
		Method: http.MethodGet,
		Proto:  r.proto,
		URL:    r.host + "/v3/users" + query,
		Auth:   r.auth,
	}
	resp, err := r.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &UserDetailGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UserDetailGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError   `json:"error,omitempty"`
	Start         jiguang.UnitTime `json:"start,omitempty"`    // 请求时的起始时间
	Duration      int              `json:"duration,omitempty"` // 请求时的持续时长
	Items         []UserStatsItem  `json:"items,omitempty"`    // 获取到的统计数据项
}

// 用户在不同平台的统计数据项
type UserStatsItem struct {
	Time     jiguang.UnitTime     `json:"time"`               // 统计时间
	Android  *UserStatsItemDetail `json:"android,omitempty"`  // Android 平台用户统计详情
	IOS      *UserStatsItemDetail `json:"ios,omitempty"`      // iOS 平台用户统计详情
	QuickApp *UserStatsItemDetail `json:"quickapp,omitempty"` // 快应用平台用户统计详情
	HMOS     *UserStatsItemDetail `json:"hmos,omitempty"`     // 鸿蒙平台用户统计详情
}

// 指定平台的用户统计详情
type UserStatsItemDetail struct {
	New    *uint64 `json:"new,omitempty"`    // 新增用户
	Online *uint64 `json:"online,omitempty"` // 在线用户
	Active *uint64 `json:"active,omitempty"` // 活跃用户
}

func (rs *UserDetailGetResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
	}

	var timeUnit jiguang.TimeUnit
	if tu, ok := aux["time_unit"]; ok {
		if err := json.Unmarshal(tu, &timeUnit); err != nil {
			return err
		}
		delete(aux, "time_unit")
	}

	if start, ok := aux["start"]; ok {
		if err := json.Unmarshal(start, &rs.Start); err != nil {
			return err
		}
		delete(aux, "start")
	}

	if rs.Start.TimeUnit != timeUnit {
		rs.Start.TimeUnit = timeUnit
	}

	if dur, ok := aux["duration"]; ok {
		duration, err := strconv.Atoi(string(dur))
		if err != nil {
			return err
		}
		rs.Duration = duration
		delete(aux, "duration")
	}

	if items, ok := aux["items"]; ok {
		if err := json.Unmarshal(items, &rs.Items); err != nil {
			return err
		}
		delete(aux, "items")
	}

	return nil
}

func (rs UserDetailGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	}
	data := make(map[string]interface{}, 4)
	data["time_unit"] = rs.Start.TimeUnit // 请求时的时间单位
	data["start"] = rs.Start.Format()     // 请求时的起始时间
	data["duration"] = rs.Duration        // 请求时的持续时长
	data["items"] = rs.Items              // 获取到的统计数据项
	return json.Marshal(data)
}

func (rs *UserDetailGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
