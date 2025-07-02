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

package jums

import "context"

// # JUMS API v1
//
// 【极光统一消息 > REST API】
//
// 1.【普通消息发送 API】
//   - 可以自由选择发送通道，填写消息内容，每个通道的消息字段完全遵循通道本身的要求；
//   - 可以选择向单个用户或用户列表发送一条消息。
//
// 2.【模板消息发送 API】
//   - 需要先在官网控制台创建一个通用模板，才可以进行发送；
//   - 发送的通道由创建模板时选择的通道决定；
//   - 一次变量传值对所有通道生效。
//
// 3.【消息撤回 API】
//   - 支持撤回 24 小时内发出的消息；
//   - 调用 API 后，UMS 将中止系统内未下发的消息，并调通道撤回接口进行撤回（如果通道支持）。
//
// 4.【用户管理 API】
//   - 用于在服务器端对用户信息进行增加、删除、修改。
//
// 5.【素材管理 API】
//   - 对 UMS 系统的内部素材和第三方平台的素材进行管理的 API，可以上传图片、视频、文件等素材；
//   - 当前仅支持上传附件，用于邮件带附件下发的功能。
//
// 6.【获取通道 Token API】
//   - 对于使用 AppID 秘钥的方式进行鉴权的通道，提供 API 获取已获取的 Token，开发者业务系统可以使用 UMS 系统做 Token 池，向 UMS 获取 Token，即可避免 Token 冲突问题；
//   - 如果企业的业务系统已经自行接入了微信公众号/小程序业务，希望 UMS 在使用微信公众号/小程序时不要与业务系统产生冲突（Token、获取统计数据），则可以在创建通道时，选择使用扫码授权的方式进行。
//
// 7.【消息回调设置】
//   - 如果需要获得一条消息对于所有目标的发送状态，可以在 高级功能-回调设置 中配置一个回调地址，勾选所需回调的状态类型，开启回调即可；
//   - 此功能仅针对极光 VIP 用户提供，提供「目标有效/无效、提交成功/失败、送达成功/失败、点击、撤回成功/失败」9 种消息状态，需在官网控制台设置所需回调的状态。
//
// 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/ums_server_overview
type APIv1 interface {
	// # 普通消息 - 广播发送
	//  - 功能说明：普通消息广播发送。
	//  - 调用地址：POST `/v1/broadcast`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message
	BroadcastSend(ctx context.Context, param *BroadcastSendParam) (*BroadcastSendResult, error)

	// # 普通消息 - 其他方式发送
	//  - 功能说明：普通消息发送。
	//  - 调用地址：POST `/v1/sent`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message
	Send(ctx context.Context, param *SendParam) (*SendResult, error)

	// # 模板消息 - 广播发送
	//  - 功能说明：模板消息广播发送。
	//  - 调用地址：POST `/v1/template/broadcast`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_template_message
	TemplateBroadcastSend(ctx context.Context, param *TemplateBroadcastSendParam) (*TemplateBroadcastSendResult, error)

	// # 模板消息 - 其他方式发送
	//  - 功能说明：模板消息发送。
	//  - 调用地址：POST `/v1/template/sent`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_template_message
	TemplateSend(ctx context.Context, param *TemplateSendParam) (*TemplateSendResult, error)

	// # 撤回消息
	//  - 功能说明：撤回消息发送，[撤回效果说明]。
	//  - 调用地址：POST `/v1/retract/{msgID}`，`msgID` 为请求发送消息成功后返回的 ID 值，一次仅允许传 1 个 `msgID`。
	//  - 接口文档：[docs.jiguang.cn]
	// [撤回效果说明]: https://docs.jiguang.cn/jums/server/rest_api_jums_retract_message#%E6%92%A4%E5%9B%9E%E6%95%88%E6%9E%9C%E8%AF%B4%E6%98%8E
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_template_message
	Retract(ctx context.Context, msgID string) (*RetractResult, error)

	// # 批量添加、更新用户信息
	//  - 功能说明：批量添加、更新用户信息。
	//  - 调用地址：POST `/v1/user/opt`
	//  - 接口文档：[docs.jiguang.cn]
	// 支持批量添加与更新用户信息：如果 UserID 已存在，则为更新操作；如不存在，则为添加操作。
	//
	// 参数 accessAuth 为 true 时，使用全局 accessKey 和 accessMasterSecret 鉴权方式；否则为默认的渠道 channelKey 和 masterSecret 鉴权方式。
	//
	// 在导入 UserID 与各个通道注册 ID 的对应关系时，有一个重要的关键字段是 [bind_id]（绑定标识），它可以取值 all、通道编码、渠道 Key。
	//  - all 即全局设置，所有渠道、通道通用，对标签、手机号码、邮箱地址、钉钉 ID、企业微信 ID、企业微信互联企业 ID 有效，注意：钉钉 ID 的关联标识仅允许传 all；
	//  - 具体的 [Channel Key]，在渠道信息中可获取到该值，对标签、手机号码、邮箱地址有效，如果需要设置某一渠道的专属信息，则可以用 Channel Key 做标识，如不需要，则传 all；
	//  - 具体的 [通道编码]，在通道详情中可获取到该值（在创建通道时由运营人员自行设定），在为 UserID 绑定 APP 注册 ID、微信公众号 ID、微信小程序 ID、支付宝生活号 ID 时必须指定通道编码。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_user#%E6%89%B9%E9%87%8F%E6%B7%BB%E5%8A%A0%E4%B8%8E%E6%9B%B4%E6%96%B0-api
	// [bind_id]: https://img.jiguang.cn/docs/2021/jums/server/image/user1.png
	// [Channel Key]: https://img.jiguang.cn/docs/2023/jums/server/assets/16825055309596.jpg
	// [通道编码]: https://img.jiguang.cn/docs/2023/jums/server/assets/16825055860019.jpg
	BatchAddOrUpdateUsers(ctx context.Context, param []UsersBatchAddOrUpdateParam, accessAuth bool) (*UsersBatchAddOrUpdateResult, error)

	// # 批量删除用户信息
	//  - 功能说明：本 API 将删除用户的唯一 ID 及其所绑定的所有信息，请谨慎操作。该操作必须使用全局 accessKey 和 accessMasterSecret 进行鉴权。
	//  - 调用地址：POST `/v1/user/delete`，`userIDs` 为要批量删除的用户的唯一标识列表。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_user#%E6%89%B9%E9%87%8F%E5%88%A0%E9%99%A4-api
	BatchDeleteUsers(ctx context.Context, userIDs []string) (*UsersBatchDeleteResult, error)

	// # 上传素材
	//  - 功能说明：可以上传图片、视频、文件等素材。
	//  - 调用地址：POST `/v1/material`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_material#%E4%B8%8A%E4%BC%A0%E7%B4%A0%E6%9D%90
	UploadMaterial(ctx context.Context, param *MaterialUploadParam) (*MaterialUploadResult, error)

	// # 获取通道 Token
	//  - 功能说明：获取通道 Token。
	//  - 调用地址：GET `/v1/token?type={channelType}`
	//  - 接口文档：[docs.jiguang.cn]
	// channelType 为所需要获取的通道的类型，通过渠道鉴权，可以取到本渠道授权的通道的 Token 值，它的取值说明如下：
	//  - wechatoa：微信公众号；
	//  - wechatmp：微信小程序；
	//  - dingtalkcc：钉钉；
	//  - wechatwk：企业微信。
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_token
	GetChannelToken(ctx context.Context, channelType string) (*ChannelTokenGetResult, error)

	// # 获取配置好的回调接口服务
	Callback() (*Callback, error)
}
