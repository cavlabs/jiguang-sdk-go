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

package jsms

import "context"

// # JSMS API v1
//
// 【极光短信 > REST API】
//
// 1.【短信签名 API】
//   - 支持创建、修改、查询和删除短信签名。
//
// 2.【短信模板 API】
//   - 支持创建、修改、查询和删除短信模板。
//
// 3.【短信发送 API】
//   - 支持发送文本、语音验证码短信；
//   - 支持验证验证码；
//   - 支持发送单条、批量模板短信。
//
// 4.【短信定时发送 API】
//   - 支持提交、修改、查询和删除模板短信定时发送任务。
//
// 5.【短信余量查询 API】
//   - 支持查询开发者账号余量；
//   - 支持查询应用余量。
//
// 6.【短信回执 API】
//   - 支持 “获取短信发送状态” 和 “获取上行短信消息” 接口；
//   - 注意：短信发送状态、上行短信消息已提供回调的获取方式（7），若开发者已使用了回调的获取方式，则不支持使用短信回执 API 获取。
//
// 7.【回调接口】
//   - 支持使用默认或定制回调接口服务，包括 “发送状态回调”、“回复内容回调”、“模板审核回调” 和 “签名审核回调”。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms
type APIv1 interface {
	// # 创建签名
	//  - 功能说明：创建短信签名。
	//  - 调用地址：POST `/v1/sign`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E5%88%9B%E5%BB%BA%E7%AD%BE%E5%90%8D-api
	CreateSign(ctx context.Context, param *SignCreateParam) (*SignCreateResult, error)

	// # 修改签名
	//  - 功能说明：修改审核不通过的签名，并再次提交审核。
	//  - 调用地址：POST `/v1/sign/{signID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E4%BF%AE%E6%94%B9%E7%AD%BE%E5%90%8Dapi
	UpdateSign(ctx context.Context, signID int, param *SignUpdateParam) (*SignUpdateResult, error)

	// # 查询签名
	//  - 功能说明：查询短信签名。
	//  - 调用地址：GET `/v1/sign/{signID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E6%9F%A5%E8%AF%A2%E7%AD%BE%E5%90%8Dapi
	GetSign(ctx context.Context, signID int) (*SignGetResult, error)

	// # 删除签名
	//  - 功能说明：删除短信签名。
	//  - 调用地址：DELETE `/v1/sign/{signID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign#%E5%88%A0%E9%99%A4%E7%AD%BE%E5%90%8D-api
	DeleteSign(ctx context.Context, signID int) (*SignDeleteResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 创建模板
	//  - 功能说明：创建短信模板。
	//  - 调用地址：POST `/v1/templates`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E5%88%9B%E5%BB%BA%E6%A8%A1%E6%9D%BF-api
	CreateTemplate(ctx context.Context, param *TemplateCreateParam) (*TemplateCreateResult, error)

	// # 修改模板
	//  - 功能说明：修改审核不通过的模板，并再次提交审核。
	//  - 调用地址：PUT `/v1/templates/{tempID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E4%BF%AE%E6%94%B9%E6%A8%A1%E6%9D%BF-api
	UpdateTemplate(ctx context.Context, tempID int64, param *TemplateUpdateParam) (*TemplateUpdateResult, error)

	// # 查询模板
	//  - 功能说明：查询短信模板。
	//  - 调用地址：GET `/v1/templates/{tempID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E6%9F%A5%E8%AF%A2%E6%A8%A1%E6%9D%BF-api
	GetTemplate(ctx context.Context, tempID int64) (*TemplateGetResult, error)

	// # 删除模板
	//  - 功能说明：删除短信模板。
	//  - 调用地址：DELETE `/v1/templates/{tempID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E5%88%A0%E9%99%A4%E6%A8%A1%E6%9D%BF-api
	DeleteTemplate(ctx context.Context, tempID int64) (*TemplateDeleteResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 发送文本验证码短信
	//  - 功能说明：发送文本验证码短信。
	//  - 调用地址：POST `/v1/codes`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E6%96%87%E6%9C%AC%E9%AA%8C%E8%AF%81%E7%A0%81%E7%9F%AD%E4%BF%A1-api
	SendCode(ctx context.Context, param *CodeSendParam) (*CodeSendResult, error)

	// # 发送语音验证码短信
	//  - 功能说明：发送语音验证码短信。
	//  - 调用地址：POST `/v1/voice_codes`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E8%AF%AD%E9%9F%B3%E9%AA%8C%E8%AF%81%E7%A0%81%E7%9F%AD%E4%BF%A1-api
	SendVoiceCode(ctx context.Context, param *VoiceCodeSendParam) (*VoiceCodeSendResult, error)

	// # 验证码验证
	//  - 功能说明：验证验证码是否有效。
	//  - 调用地址：POST `/v1/codes/{msgID}/valid`，其中 `msgID` 为发送验证码短信时返回的消息 ID，`code` 为用户输入的验证码。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E9%AA%8C%E8%AF%81%E7%A0%81%E9%AA%8C%E8%AF%81-api
	VerifyCode(ctx context.Context, msgID, code string) (*CodeVerifyResult, error)

	// # 发送单条模板短信
	//  - 功能说明：发送单条模板短信。
	//  - 调用地址：POST `/v1/messages`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E5%8D%95%E6%9D%A1%E6%A8%A1%E6%9D%BF%E7%9F%AD%E4%BF%A1-api
	SendMessage(ctx context.Context, param *MessageSendParam) (*MessageSendResult, error)

	// # 发送批量模板短信
	//  - 功能说明：发送批量模板短信。
	//  - 调用地址：POST `/v1/messages/batch`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E6%89%B9%E9%87%8F%E6%A8%A1%E6%9D%BF%E7%9F%AD%E4%BF%A1-api
	SendBatchMessage(ctx context.Context, param *BatchMessageSendParam) (*BatchMessageSendResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 单条定时短信提交
	//  - 功能说明：提交单条模板短信定时发送任务。
	//  - 调用地址：POST `/v1/schedule`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%8D%95%E6%9D%A1%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%8F%90%E4%BA%A4-api
	SendSchedule(ctx context.Context, param *ScheduleSendParam) (*ScheduleSendResult, error)

	// # 单条定时短信修改
	//  - 功能说明：修改单条模版短信定时发送任务。
	//  - 调用地址：PUT `/v1/schedule/{scheduleID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%8D%95%E6%9D%A1%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E4%BF%AE%E6%94%B9-api
	UpdateSchedule(ctx context.Context, scheduleID string, param *ScheduleUpdateParam) (*ScheduleUpdateResult, error)

	// # 定时短信查询
	//  - 功能说明：查询模板短信定时发送任务。
	//  - 调用地址：GET `/v1/schedule/{scheduleID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%9F%A5%E8%AF%A2api
	GetSchedule(ctx context.Context, scheduleID string) (*ScheduleGetResult, error)

	// # 定时短信删除
	//  - 功能说明：删除模板短信定时发送任务。
	//  - 调用地址：DELETE `/v1/schedule/{scheduleID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E5%88%A0%E9%99%A4-api
	DeleteSchedule(ctx context.Context, scheduleID string) (*ScheduleDeleteResult, error)

	// # 批量定时短信提交
	//  - 功能说明：提交批量模板短信定时发送任务。
	//  - 调用地址：POST `/v1/schedule/batch`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E6%89%B9%E9%87%8F%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E6%8F%90%E4%BA%A4-api
	SendBatchSchedule(ctx context.Context, param *BatchScheduleSendParam) (*BatchScheduleSendResult, error)

	// # 批量定时短信修改
	//  - 功能说明：修改批量模板短信定时发送任务。
	//  - 调用地址：PUT `/v1/schedule/batch/{scheduleID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule#%E6%89%B9%E9%87%8F%E5%AE%9A%E6%97%B6%E7%9F%AD%E4%BF%A1%E4%BF%AE%E6%94%B9-api
	UpdateBatchSchedule(ctx context.Context, scheduleID string, param *BatchScheduleUpdateParam) (*BatchScheduleUpdateResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 账号余量查询
	//
	// 使用 devKey:devSecret 认证。
	//  - 功能说明：查询账号余量，账号余量指未分配给某个应用，属于账号共享的短信余量。
	//  - 调用地址：GET `/v1/accounts/dev`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#%E8%B4%A6%E5%8F%B7%E4%BD%99%E9%87%8F%E6%9F%A5%E8%AF%A2-api
	GetDevBalance(ctx context.Context) (*DevBalanceGetResult, error)

	// # 应用余量查询
	//
	// 使用 appKey:masterSecret 认证。
	//  - 功能说明：查询应用余量，应用余量指分配给某个应用单独使用的短信余量。
	//  - 调用地址：GET `/v1/accounts/app`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_jsms_api_account#%E5%BA%94%E7%94%A8%E4%BD%99%E9%87%8F%E6%9F%A5%E8%AF%A2-api
	GetAppBalance(ctx context.Context) (*AppBalanceGetResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 获取短信发送状态
	//
	// SMS_REPORT - 短信送达状态。
	//  - 功能说明：获取 appKey 下短信的发送状态报告；API 调用频率为 5 秒/次，每次调用最多可获取 500 条数据；每条短信发送状态报告数据最多保留七天，即极光收到运营商返回状态报告后的七天。
	//  - 调用地址：GET `/v1/report`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire#%E8%8E%B7%E5%8F%96%E7%9F%AD%E4%BF%A1%E5%8F%91%E9%80%81%E7%8A%B6%E6%80%81
	GetReport(ctx context.Context) (*ReportGetResult, error)

	// # 获取上行短信消息
	//
	// SMS_REPLY - 用户回复消息。
	//  - 功能说明：获取 appKey 下的上行短信消息；API 调用频率为 5 秒/次，每次调用最多可获取 500 条数据；每条上行短信消息数据最多保留七天，即极光收到运营商返回上行消息后的七天。
	//  - 调用地址：GET `/v1/reply`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire#%E8%8E%B7%E5%8F%96%E4%B8%8A%E8%A1%8C%E7%9F%AD%E4%BF%A1%E6%B6%88%E6%81%AF
	GetReply(ctx context.Context) (*ReplyGetResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 获取配置好的回调接口服务
	Callback() (*Callback, error)
}
