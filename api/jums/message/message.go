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

package message

// # 消息的内容
//
// 在请求体中填写的 msg_xxx 决定本次消息将发送给哪几种通道：
//   - 如果使用自定义通道注册 ID 发送，aud_xxx 和 msg_xxx 需要一一对应；
//   - 如果使用了发送策略，策略中包含的通道和 msg_xxx 中的通道信息需要一致；
//   - 向【企业微信互联企业】通道发送时，与【企业微信】通道共用 msg_wechatwk。
type Message struct {
	// 【可选】【APP】消息内容。
	//  - 当前 UMS 默认对接极光推送，因此 APP 的消息内容可参考 push.SendParam。
	//  - JPush 支持发送的大部分参数均可在此传递，不支持的参数有：Callback、Audience、VoIP。
	Apps []App `json:"msg_app,omitempty"`
	// 【可选】【微信公众号】消息内容。
	//  - 微信公众号支持 [模板消息](WechatoaTemplate)、[订阅通知](WechatoaSubscription) 2 种，通过 Type 区分；
	//  - 除了接收者 openID 被 UMS 的目标字段替代外，官方文档中的相关参数均可在此传递。
	// [模板消息]: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html#5
	// [订阅通知]: https://developers.weixin.qq.com/doc/offiaccount/Subscription_Messages/api.html#send%E5%8F%91%E9%80%81%E8%AE%A2%E9%98%85%E9%80%9A%E7%9F%A5
	Wechatoa []Wechatoa `json:"msg_wechatoa,omitempty"`
	// 【可选】【微信小程序】消息内容。
	//  - 微信小程序支持 [模板消息]，除了接收者 openID 被 UMS 的目标字段替代外，官方文档中的相关参数均可在此传递。
	// [模板消息]: https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	Wechatmp []Wechatmp `json:"msg_wechatmp,omitempty"`
	// 【可选】【邮件】消息内容。
	//  - 邮件消息支持传递邮件标题、邮件内容。
	Email []Email `json:"msg_email,omitempty"`
	// 【可选】【短信】消息内容。
	//  - 当前 UMS 默认对接极光短信，因此使用极光短信时的消息内容可参考 [短信发送 API 中的单条模板消息](JSMS)；
	//  - 使用 CMPP 对接的短信平台只需传递短信内容 (CMPP.Content)。
	// [短信发送 API 中的单条模板消息]: https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E5%8D%95%E6%9D%A1%E6%A8%A1%E6%9D%BF%E7%9F%AD%E4%BF%A1-api
	SMS []SMS `json:"msg_sms,omitempty"`
	// 【可选】【钉钉】工作通知消息内容。
	//  - 当前钉钉工作通知支持发送 [文本消息](DingtalkCCMsgText)。
	// [文本消息]: https://developers.dingtalk.com/document/app/message-types-and-data-format
	DingtalkCC []DingtalkCC `json:"msg_dingtalk_cc,omitempty"`
	// 【可选】【支付宝生活号】消息内容。
	//  - 当前支付宝生活号支持发送 [模板消息]。
	// [模板消息]: https://opendocs.alipay.com/apis/api_6/alipay.open.public.message.single.send
	AlipayLife []AlipayLife `json:"msg_alipay_life,omitempty"`
	// 【可选】【企业微信】/【企业微信互联企业】消息内容。
	//  - 企业微信支持的消息类型包含：[文本消息]、[图片消息]、[文件消息]、[外链图文]、[图文消息]、[小程序通知消息]；
	//  - 企业微信同时支持向企微互联用户发送，消息格式与本企业的发送是一致的，参考 [官方文档]，因此你只需传 1 次消息内容，UMS 系统将会根据 aud_xxx 的设置自动路由到互联企业用户或者本企业用户。
	// [文本消息]: https://developer.work.weixin.qq.com/document/path/90372#%E6%96%87%E6%9C%AC%E6%B6%88%E6%81%AF
	// [图片消息]: https://developer.work.weixin.qq.com/document/path/90372#%E5%9B%BE%E7%89%87%E6%B6%88%E6%81%AF
	// [文件消息]: https://developer.work.weixin.qq.com/document/path/90372#%E6%96%87%E4%BB%B6%E6%B6%88%E6%81%AF
	// [外链图文]: https://developer.work.weixin.qq.com/document/path/90372#%E5%9B%BE%E6%96%87%E6%B6%88%E6%81%AF
	// [图文消息]: https://developer.work.weixin.qq.com/document/path/90372#%E5%9B%BE%E6%96%87%E6%B6%88%E6%81%AF%EF%BC%88mpnews%EF%BC%89
	// [小程序通知消息]: https://developer.work.weixin.qq.com/document/path/90372#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E9%80%9A%E7%9F%A5%E6%B6%88%E6%81%AF
	// [官方文档]: https://developer.work.weixin.qq.com/document/path/90250
	Wechatwk []Wechatwk `json:"msg_wechatwk,omitempty"`
}
