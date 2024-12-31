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

package schedule

import "context"

// Schedule API v3【极光推送 > REST API > 定时任务 API】
//   - 功能说明：API 层面支持定时功能。这是一个相对独立的任务执行模块，维护一个 Schedule 对象。
//   - 极光文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule
// 注意：调 API 创建的定时任务只能调 API 获取/修改/删除。
type APIv3 interface {
	// 定时推送（创建定时任务）
	//  - 功能说明：API 层面支持定时功能。这是一个相对独立的任务执行模块，维护一个 Schedule 对象。调 API 创建的定时任务只能调 API 获取/修改/删除。
	//	- 调用地址：POST `/v3/schedules`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E5%88%9B%E5%BB%BA%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1
	// 注意事项：
	//  1. 文件定时推送 API 接口频率规则和 File API v3 频率规则一样 20 次/min，且各个文件相关接口频率会互相消耗；
	//  2. 对于文件定时推送，创建定时任务成功后，若任务被执行前文件被删除，则任务执行时推送动作将会失败。
	ScheduleSend(ctx context.Context, param *SendParam) (*SendResult, error)

	// 定时模板推送（VIP） - 指定模板，指定时间，定时推送。
	//  - 功能说明：开发者可以将推送时所需参数通过极光 WebPortal 页面提前配置好，形成一套 “推送模板”，然后通过此 API 接口，指定 “模板 ID” 推送，大大降低开发难度，节省开发调试时间，节省服务器带宽成本。
	//	- 调用地址：POST `/v3/push/template/schedule`，`id` 为模板 ID（创建模板后，由极光服务器生成），`params` 为模板参数（创建模板时，开发者设置的变量参数）列表，`scheduleName` 为任务名称，`trigger` 为任务触发条件。
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_template#%E6%8C%87%E5%AE%9A%E6%A8%A1%E6%9D%BF%E3%80%90%E5%AE%9A%E6%97%B6%E6%8E%A8%E9%80%81%E3%80%91
	// 注意事项：不支持多种不同 Audience 类型的组合（Tags、AndTags、NotTags 除外）！
	//  1. params 数组里面，同一个推送里面的 Audience 不允许不同类型组合（Tags、AndTags、NotTags 除外）；
	//  2. params 数组里面，不同推送的 Audience 也不允许不同类型（Tags、AndTags、NotTags 除外）；
	//  3. 如果是文件方式推送，一次推送只支持指定一个文件；
	//  4. 如果是地理围栏方式推送，一次推送仅支持指定一个地理围栏。
	// 也就是说：假设数组长度是 3，那么就表示有 3 个推送，这 3 个推送要么都是 Registration ID 方式推送，要么都是 Alias 方式推送。
	ScheduleTemplateSend(ctx context.Context, id string, params []TemplateParam, scheduleName string, trigger *Trigger) (*TemplateSendResult, error)

	// 获取定时推送唯一标识 (CID)
	//  - 功能说明：CID 是用于防止 API 调用端重试造成服务端的重复推送而定义的一个推送参数。用户使用一个 CID 推送后，再次使用相同的 CID 进行推送，则会直接返回第一次成功推送的结果，不会再次进行推送。
	//  CID 的有效期为 1 天，格式为：{appkey}-{uuid}，在使用 CID 之前，必须通过接口获取你的 CID 池。
	//	- 调用地址：GET `/v3/push/cid?type=schedule&count={count}`，如 count < 1 则自动重置为 1。
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E8%8E%B7%E5%8F%96%E6%8E%A8%E9%80%81%E5%94%AF%E4%B8%80%E6%A0%87%E8%AF%86cid
	GetCidForSchedulePush(ctx context.Context, count int) (*CidGetResult, error)

	// 获取有效的定时任务列表
	//  - 功能说明：获取当前有效（EndTime 未过期）的定时任务列表。
	//	- 调用地址：GET `/v3/schedules?page={page}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E8%8E%B7%E5%8F%96%E6%9C%89%E6%95%88%E7%9A%84%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E5%88%97%E8%A1%A8
	// 注意事项：
	//  - 返回当前 page 请求页的详细的定时任务列表，如 page < 0 则自动重置为 1；
	//  - 结果列表按 “创建时间” 排序；
	//  - 如果请求页数大于总页数，则 page 为请求值，结果列表为空；
	//  - 每页最多返回 50 个定时任务，如请求页实际的个数小于 50，则按实际数量返回。
	GetSchedules(ctx context.Context, page int) (*SchedulesGetResult, error)

	// 获取定时任务详情
	//  - 功能说明：获取当前用户指定定时任务的详细信息。
	//	- 调用地址：GET `/v3/schedules/{scheduleID}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E8%8E%B7%E5%8F%96%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E8%AF%A6%E6%83%85
	GetSchedule(ctx context.Context, scheduleID string) (*ScheduleGetResult, error)

	// 获取某个定时任务的所有消息 ID
	//  - 功能说明：获取当前用户指定定时任务对应的所有消息 ID 列表。
	//	- 调用地址：GET `/v3/schedules/{scheduleID}/msg_ids`
	GetScheduleMsgIDs(ctx context.Context, scheduleID string) (*ScheduleMsgIDsResult, error)

	// 更新定时任务
	//  - 功能说明：更新当前用户指定定时任务的详细信息。
	//	- 调用地址：PUT `/v3/schedules/{scheduleID}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E6%9B%B4%E6%96%B0%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1
	// 注意事项：
	//  - 不支持部分更新；
	//  - 更新后的 Trigger 必须仍是有效合法的，否则即使 Trigger 整体更新也会失败。可以更新 Enabled = false 的任务；
	//  - 定时任务（Single）与定期任务（Periodical）之间不能进行相互更新，即，原先为 Single 类任务，则不能更新为 Periodical 任务，反之亦然；
	//  - 不能更新已过期的定时任务。
	UpdateSchedule(ctx context.Context, scheduleID string, param *UpdateParam) (*UpdateResult, error)

	// 删除定时任务
	//  - 功能说明：删除当前用户指定定时任务，如果 scheduleID 不合法，则 404。
	//	- 调用地址：DELETE `/v3/schedules/{scheduleID}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E5%88%A0%E9%99%A4%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1
	DeleteSchedule(ctx context.Context, scheduleID string) (*DeleteResult, error)

	// ********************* ↓↓↓ 如果遇到此 API 没有及时补充字段的情况，可以自行构建 JSON，调用下面的接口 ↓↓↓ *********************

	// 自定义定时推送：如果遇到 ScheduleSend 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
	CustomScheduleSend(ctx context.Context, param interface{}) (*SendResult, error)
}
