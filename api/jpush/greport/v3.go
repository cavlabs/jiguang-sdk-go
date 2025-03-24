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

package greport

import (
	"context"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// # Group Report API v3
//
// 【极光推送 > REST API > 统计 API】
//   - 功能说明：提供分组应用的各类统计数据查询功能。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report
type APIv3 interface {
	// # 消息统计详情（VIP）
	//  - 功能说明：针对分组应用，获取指定 groupMsgIDs（最多支持 10 个）的消息统计数据。只支持查询 30 天以内的推送信息。
	//  - 调用地址：GET `/v3/group/messages/detail`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E5%88%86%E7%BB%84%E7%BB%9F%E8%AE%A1-%E6%B6%88%E6%81%AF%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
	GetMessageDetail(ctx context.Context, groupMsgIDs []string) (*MessageDetailGetResult, error)

	// # 用户统计（VIP）
	//  - 功能说明：针对分组应用，提供近 1 个月内某时间段的用户相关统计数据：新增用户、在线用户、活跃用户。
	//  `start` 起始时间，它的时间单位支持：HOUR（小时，格式例：2020-08-11 09）、DAY（天，格式例：2020-08-11）、MONTH（月，格式例：2020-08）；
	//  `duration` 持续时长，如果时间单位是天，则是持续的天数，以此类推，只支持查询 30 天以内的用户信息，对于时间单位为 HOUR 的，只支持输出当天的统计结果。
	//  - 调用地址：GET `/v3/group/users`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E5%88%86%E7%BB%84%E7%BB%9F%E8%AE%A1-%E7%94%A8%E6%88%B7%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
	GetUserDetail(ctx context.Context, start jiguang.UnitTime, duration int) (*UserDetailGetResult, error)
}
