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

package push

import (
	"context"

	"github.com/calvinit/jiguang-sdk-go/api/jpush/file"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/image"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/schedule"
)

type (
	fileAPIv3     = file.APIv3
	imageAPIv3    = image.APIv3
	scheduleAPIv3 = schedule.APIv3
)

// # Push API v3
//
// 【极光推送 > REST API > 推送 API】
//   - 功能说明：包括普通推送、定时推送、文件推送、批量单推、模板推送、自定义推送等相关 API。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/push
type APIv3 interface {
	fileAPIv3
	imageAPIv3
	scheduleAPIv3

	// # 普通推送
	//  - 功能说明：向某单个设备或者某设备列表推送一条通知或者消息。推送的内容只能是 JSON 表示的一个推送对象。
	//	- 调用地址：POST `/v3/push`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push
	Send(ctx context.Context, param *SendParam) (*SendResult, error)

	// # 普通推送（SM2 加密）
	//  - 功能说明：向某单个设备或者某设备列表推送一条通知或者消息。推送的内容只能是 JSON 表示的一个推送对象。
	//	- 调用地址：POST `/v3/push`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push
	SendWithSM2(ctx context.Context, param *SendParam) (*SendResult, error)

	// # 文件推送（文件立即推送）
	//  - 功能说明：支持指定文件唯一标识（fileID）进行推送，文件唯一标识（fileID）可以参考 File API v3 的 文件上传接口 获得。
	//	- 调用地址：POST `/v3/push/file`
	//  - 接口文档：[docs.jiguang.cn]
	// 注意事项：
	//  1. “文件定时推送” 请使用 ScheduleSend 接口；
	//  2. 厂商配额查询 API 接口频率和推送 API 接口频率共享，消耗推送 API 接口频率；
	//  3. 建议创建推送任务 5 分钟后再执行文件删除操作，否则推送任务可能会失败。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%96%87%E4%BB%B6%E6%8E%A8%E9%80%81-api
	SendByFile(ctx context.Context, param *SendParam) (*SendResult, error)

	// # 推送校验
	//  - 功能说明：该 API 只用于验证推送调用是否能够成功，与推送 API 的区别在于：不向用户发送任何消息。
	//	- 调用地址：POST `/v3/push/validate`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%8E%A8%E9%80%81%E6%A0%A1%E9%AA%8C-api
	ValidateSend(ctx context.Context, param *SendParam) (*SendResult, error)

	// # 批量单推（Registration ID 方式）
	//  - 功能说明：如果您在给每个用户的推送内容都不同的情况下，可以使用此接口。使用此接口前，您需要配合使用 GetCidForPush 接口提前获取到 CID 池。
	//	- 调用地址：POST `/v3/push/batch/regid/single`，`pushList` 的 key 为 CID 值，最多支持填写 1000 个。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_single
	BatchSendByRegistrationID(ctx context.Context, pushList map[string]BatchPushParam) (*BatchSendResult, error)

	// # 批量单推（Alias 方式）
	//  - 功能说明：如果您在给每个用户的推送内容都不同的情况下，可以使用此接口。使用此接口前，您需要配合使用 GetCidForPush 接口提前获取到 CID 池。
	//	- 调用地址：POST `/v3/push/batch/alias/single`，`pushList` 的 key 为 CID 值，最多支持填写 1000 个。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_single
	BatchSendByAlias(ctx context.Context, pushList map[string]BatchPushParam) (*BatchSendResult, error)

	// # 普通模板推送（VIP）
	//
	// 指定模板 ID，模板参数（如有设置），进行立即推送。
	//  - 功能说明：开发者可以将推送时所需参数通过极光 WebPortal 页面提前配置好，形成一套 “推送模板”，然后通过此 API 接口，指定 “模板 ID” 推送，大大降低开发难度，节省开发调试时间，节省服务器带宽成本。
	//	- 调用地址：POST `/v3/push/template`，`id` 为模板 ID（创建模板后，由极光服务器生成），`params` 为模板参数（创建模板时，开发者设置的变量参数）列表。
	//  - 接口文档：[docs.jiguang.cn]
	// 注意事项：不支持多种不同 Audience 类型的组合（Tags、AndTags、NotTags 除外）！
	//  1. params 数组里面，同一个推送里面的 Audience 不允许不同类型组合（Tags、AndTags、NotTags 除外）；
	//  2. params 数组里面，不同推送的 Audience 也不允许不同类型（Tags、AndTags、NotTags 除外）；
	//  3. 如果是文件方式推送，一次推送只支持指定一个文件；
	//  4. 如果是地理围栏方式推送，一次推送仅支持指定一个地理围栏。
	// 也就是说：假设数组长度是 3，那么就表示有 3 个推送，这 3 个推送要么都是 Registration ID 方式推送，要么都是 Alias 方式推送。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_template#%E6%8C%87%E5%AE%9A%E6%A8%A1%E6%9D%BF%E3%80%90%E7%AB%8B%E5%8D%B3%E6%8E%A8%E9%80%81%E3%80%91
	TemplateSend(ctx context.Context, id string, params []TemplateParam) (*TemplateSendResult, error)

	// # 获取推送唯一标识 (CID)
	//  - 功能说明：CID 是用于防止 API 调用端重试造成服务端的重复推送而定义的一个推送参数。用户使用一个 CID 推送后，再次使用相同的 CID 进行推送，则会直接返回第一次成功推送的结果，不会再次进行推送。
	//  CID 的有效期为 1 天，格式为：{appkey}-{uuid}，在使用 CID 之前，必须通过接口获取你的 CID 池。
	//	- 调用地址：GET `/v3/push/cid?type=push&count={count}`，如 count < 1 则自动重置为 1。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E8%8E%B7%E5%8F%96%E6%8E%A8%E9%80%81%E5%94%AF%E4%B8%80%E6%A0%87%E8%AF%86cid
	GetCidForPush(ctx context.Context, count int) (*CidGetResult, error)

	// # 推送撤销
	//  - 功能说明：撤销操作首先会从服务端尝试撤销（Android 消息，排队中/发送中状态可以服务端撤销；
	//  iOS 消息，排队中状态可以服务端撤销）；其次，针对 Push SDK（JPush Android SDK v3.5.0 及以上和 JPush iOS SDK v3.2.8 及以上），
	//  会尝试从设备端撤销已展示但未被点击的消息。
	//	- 调用地址：DELETE `/v3/push/{msgID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%8E%A8%E9%80%81%E6%92%A4%E9%94%80-api
	WithdrawMessage(ctx context.Context, msgID string) (*WithdrawMessageResult, error)

	// # 厂商配额查询
	//  - 功能说明：厂商为了控制应用推送消息的频率，会根据应用在厂商的日联网数计算每天推送数量上限。
	//  目前已知小米、OPPO、vivo 都有每天的额度控制，额度可以在厂商后台查询，也可以通过极光 WebPortal 控制台查询；
	//  为了进一步方便开发，极光汇总了三个平台的查询接口，供开发者调用查询。
	//	- 调用地址：GET `/v3/push/quota`
	//  - 接口文档：[docs.jiguang.cn]
	// 注意：厂商配额查询 API 接口频率和 Push API 接口频率共享，消耗 Push API 接口频率。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E5%8E%82%E5%95%86%E9%85%8D%E9%A2%9D%E6%9F%A5%E8%AF%A2-api
	GetQuota(ctx context.Context) (*QuotaGetResult, error)

	// ********************* ↓↓↓ 如果遇到此 API 没有及时补充字段的情况，可以自行构建 JSON，调用下面的接口 ↓↓↓ *********************

	// # 自定义推送
	//
	// 如果遇到 Send 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
	CustomSend(ctx context.Context, param interface{}) (*SendResult, error)

	// # 自定义文件推送
	//
	// 如果遇到 SendByFile 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
	CustomSendByFile(ctx context.Context, param interface{}) (*SendResult, error)

	// # 自定义推送校验
	//
	// 如果遇到 ValidateSend 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
	ValidateCustomSend(ctx context.Context, param interface{}) (*SendResult, error)
}
