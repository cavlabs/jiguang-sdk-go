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

package api

import (
	"net/http"
	"strconv"
)

const (
	headerRateLimit     = "X-Rate-Limit-Limit"
	headerRateRemaining = "X-Rate-Limit-Remaining"
	headerRateReset     = "X-Rate-Limit-Reset"
)

// API 频率控制信息。
//  - 参见文档：https://docs.jiguang.cn/jpush/server/push/server_overview#%E8%8E%B7%E5%8F%96%E9%A2%91%E7%8E%87%E4%BF%A1%E6%81%AF
type Rate struct {
	Limit     int   `json:"limit"`     // 当前在一个时间窗口内可调用次数。
	Remaining int   `json:"remaining"` // 当前的时间窗口剩余的可用次数。
	Reset     int64 `json:"reset"`     // 当前距离时间窗口重置剩余的秒数。
}

// 解析 HTTP 响应头以得到 API 频率控制信息。
func ParseRate(resp *http.Response) Rate {
	var rate Rate
	if limit := resp.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := resp.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := resp.Header.Get(headerRateReset); reset != "" {
		rate.Reset, _ = strconv.ParseInt(reset, 10, 64)
	}
	return rate
}
