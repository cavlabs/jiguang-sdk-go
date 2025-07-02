// Copyright 2025 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package greport

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/report"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 用户统计（VIP）
//   - 功能说明：针对分组应用，提供近 1 个月内某时间段的用户相关统计数据：新增用户、在线用户、活跃用户。
//     `start` 起始时间，它的时间单位支持：HOUR（小时，格式例：2020-08-11 09）、DAY（天，格式例：2020-08-11）、MONTH（月，格式例：2020-08）；
//     `duration` 持续时长，如果时间单位是天，则是持续的天数，以此类推，只支持查询 30 天以内的用户信息，对于时间单位为 HOUR 的，只支持输出当天的统计结果。
//   - 调用地址：GET `/v3/group/users`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E5%88%86%E7%BB%84%E7%BB%9F%E8%AE%A1-%E7%94%A8%E6%88%B7%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
func (gr *apiv3) GetUserDetail(ctx context.Context, start jiguang.UnitTime, duration int) (*UserDetailGetResult, error) {
	if gr == nil {
		return nil, api.ErrNilJPushGroupReportAPIv3
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
		if duration < 1 || duration > 30 {
			return nil, errors.New("invalid `duration` value, only support 1-30")
		}
	case jiguang.TimeUnitMonth:
		if duration != 1 {
			return nil, errors.New("invalid `duration` value, only support 1")
		}
	}

	query := "?time_unit=" + tu.String() + "&start=" + url.QueryEscape(start.Format()) + "&duration=" + strconv.Itoa(duration)
	req := &api.Request{
		Method: http.MethodGet,
		Proto:  gr.proto,
		URL:    gr.host + "/v3/group/users" + query,
		Auth:   gr.auth,
	}
	resp, err := gr.client.Request(ctx, req)
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

type UserDetailGetResult = report.UserDetailGetResult
