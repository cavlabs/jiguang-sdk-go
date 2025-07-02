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

package audience

// # 消息的发送目标
//
// 在调用广播发送的 API 时无需传递这些字段；详细的说明文档参见 [aud_xxx：发送目标]。
//
// UMS 当前支持「 广播所有人、标签、用户 ID、用户分群、自定义通道注册 ID 」共 5 种目标。
//
// 用户类目标：「广播、标签、用户 ID、用户分群」均基于 UMS 中的用户体系，需要先使用 [用户管理 API] 上传用户、各通道注册 ID、用户与各通道注册 ID 的绑定关系。
//
// 使用此种方式发送消息，UMS 在向各通道下发消息前，将根据特定的规则筛选出通道注册 ID，说明如下：
//   - 标签发送时，先筛选出设置了该标签的 userID（对该标签优先选择绑定标识为本渠道 ChannelKey 的 userID，若本渠道无此标签，才会选择绑定标识为全局 all 的 userID），再筛选对应的通道注册 ID 进行下发；
//   - 对于 APP、微信公众号、微信小程序、支付宝生活号通道，userID 是通过通道编码与通道注册 ID 绑定，因此在用户信息中选择「本渠道授权的通道」绑定的通道注册 ID 进行下发；
//   - 对于短信和邮件通道，在用户信息中优先选择与本渠道 ChannelKey 绑定的手机号码、邮箱，没有的情况下选择全局绑定的手机号码、邮箱；
//   - 对于钉钉通道，ID 全局唯一，因此在用户信息中选择全局绑定的钉钉注册 ID 进行下发。
//
// 注意：
//   - 在一条消息中，自定义通道注册 ID 和用户类目标（标签、用户 ID、用户分群）不允许同时存在；
//   - 可以同时给多个通道发送，每个通道一次发送最多传 1000 个 ID。
//
// [aud_xxx：发送目标]: https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message#aud_xxx%EF%BC%9A%E5%8F%91%E9%80%81%E7%9B%AE%E6%A0%87
// [用户管理 API]: https://docs.jiguang.cn/jums/server/rest_api_jums_user
type Audience struct {
	// 【可选】标签列表，一次发送最多 20 个。有效性说明：
	//  - 中英文、数字、下划线、特殊字符 @!#$&*+=.|￥；
	//  - 长度不超过 40 字节（UTF-8 编码）。
	Tags []string `json:"aud_tag,omitempty"`
	// 【可选】用户列表，一次发送最多 1000 个。有效性说明：
	//  - 大小写字母、数字、下划线、特殊字符 @!#$&*+=.|￥；
	//  - 长度不超过 64 字符。
	UserIDs []string `json:"aud_userid,omitempty"`
	// 【可选】在页面创建的用户分群的 ID。定义为数组，但目前限制一次只能发送一个。
	Segments []string `json:"aud_segment,omitempty"`

	// 【可选】[APP] 自定义通道注册 ID 列表。
	Apps []CustomChannel `json:"aud_app,omitempty"`
	// 【可选】[微信公众号] 自定义通道注册 ID 列表。
	Wechatoa []CustomChannel `json:"aud_wechatoa,omitempty"`
	// 【可选】[微信小程序] 自定义通道注册 ID 列表。
	Wechatmp []CustomChannel `json:"aud_wechatmp,omitempty"`
	// 【可选】[短信] 自定义通道注册 ID 列表。
	SMS []CustomChannel `json:"aud_sms,omitempty"`
	// 【可选】[邮件] 自定义通道注册 ID 列表。
	Email []CustomChannel `json:"aud_email,omitempty"`
	// 【可选】[支付宝生活号] 自定义通道注册 ID 列表。
	AlipayLife []CustomChannel `json:"aud_alipay_life,omitempty"`
	// 【可选】[钉钉] 自定义通道注册 ID 列表。
	DingtalkCC []CustomChannel `json:"aud_dingtalk_cc,omitempty"`
	// 【可选】[企业微信] 自定义通道注册 ID 列表。
	Wechatwk []CustomChannel `json:"aud_wechatwk,omitempty"`
	// 【可选】[企业微信互联企业] 自定义通道注册 ID 列表。
	WechatwkLinkedCorp []CustomChannel `json:"aud_wechatwk_linkedcorp,omitempty"`
}
