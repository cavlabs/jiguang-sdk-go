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

package report

import (
	"context"

	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # Report API v3
//
// 【极光推送 > REST API > 统计 API】
//   - 功能说明：提供各类统计数据查询功能。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report
type APIv3 interface {
	// # 送达统计详情
	//  - 功能说明：获取指定 msgIDs（最多支持 100 个）的送达统计数据。如果一次 API 调用推送有很多对象（比如广播推送），则此 API 返回的统计数据会因为持续有客户端送达而持续增加。
	//  此接口会根据消息是通过极光自有通道下发、Android 厂商通道下发进行数据统计区分。
	//  每条推送消息的送达统计数据最多保留一个月。即发起推送请求后从最后一个推送送达记录时间点开始保留一个月，如果保留期间有新的送达，将在这个新送达的时间点起再往后保留一个月。
	//  - 调用地址：GET `/v3/received/detail`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E9%80%81%E8%BE%BE%E7%BB%9F%E8%AE%A1%E8%AF%A6%E6%83%85
	GetReceivedDetail(ctx context.Context, msgIDs []string) (*ReceivedDetailGetResult, error)

	// # 送达状态查询（VIP）
	//  - 功能说明：查询已推送的一条消息在一组设备上的送达状态。
	//  - 调用地址：POST `/v3/status/message`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E9%80%81%E8%BE%BE%E7%8A%B6%E6%80%81%E6%9F%A5%E8%AF%A2%EF%BC%88vip%EF%BC%89
	GetMessageStatus(ctx context.Context, msgID string, registrationIDs []string, date *jiguang.LocalDate) (*MessageStatusGetResult, error)

	// # 消息统计详情（VIP-新）
	//  - 功能说明：获取指定 msgIDs（最多支持 100 个）的消息统计数据。与 “送达统计详情” 不同的是，该 API 提供了更多的针对单个 msgID 的统计数据。
	//  - 调用地址：GET `/v3/messages/detail`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E6%B6%88%E6%81%AF%E7%BB%9F%E8%AE%A1%E8%AF%A6%E6%83%85%EF%BC%88vip-%E6%96%B0%EF%BC%89
	GetMessageDetail(ctx context.Context, msgIDs []string) (*MessageDetailGetResult, error)

	// # 用户统计（VIP）
	//  - 功能说明：提供近 2 个月内某时间段的用户相关统计数据：新增用户、在线用户、活跃用户。
	//  `start` 起始时间，它的时间单位支持：HOUR（小时，格式例：2014-06-11 09）、DAY（天，格式例：2014-06-11）、MONTH（月，格式例：2014-06）；
	//  `duration` 持续时长，如果时间单位是天，则是持续的天数，以此类推；只支持查询 60 天以内的用户信息，对于时间单位为 HOUR 的，只支持输出当天的统计结果。
	//  - 调用地址：GET `/v3/users`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E7%94%A8%E6%88%B7%E7%BB%9F%E8%AE%A1%EF%BC%88vip%EF%BC%89
	GetUserDetail(ctx context.Context, start jiguang.UnitTime, duration int) (*UserDetailGetResult, error)
}
