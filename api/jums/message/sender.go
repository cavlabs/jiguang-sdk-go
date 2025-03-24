/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
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

package message

// # 消息下发通道 / 处理模块
type Sender int

const (
	SenderUms         Sender = iota // UMS 系统 (0)
	SenderApp                       // APP (1)
	SenderWechatoa                  // 微信公众号 (2)
	SenderWechatmp                  // 微信小程序 (3)
	SenderEmail                     // 邮件 (4)
	SenderSms                       // 短信 (5)
	SenderDingtalkCC                // 钉钉 (6)
	SenderAlipayLife                // 支付宝生活号 (7)
	_                               //
	Sender5GMessaging               // 5G 消息 (8)
	SenderWechatwk                  // 企业微信 (9)
)

var senderDescs = [...]string{
	"UMS 系统",
	"APP",
	"微信公众号",
	"微信小程序",
	"邮件",
	"短信",
	"钉钉",
	"支付宝生活号",
	"-",
	"5G 消息",
	"企业微信",
}

func (s Sender) String() string {
	return senderDescs[s]
}

// ---------------------------------------------------------------------------------------------------------------------

// # 【企业微信】通道的通道类别
type SenderWechatwkCategory int

const (
	SenderWechatwkInternal   SenderWechatwkCategory = iota // 企业微信内部接口 (0)
	SenderWechatwkLinkedCorp                               // 企业微信互联企业发送接口 (1)
)

// ---------------------------------------------------------------------------------------------------------------------

// # 【APP】通道的通道类别
type SenderAppCategory int

const (
	SenderAppJiguang SenderAppCategory = iota // 极光 (0)
	SenderAppHuawei                           // 华为 (1)
	SenderAppOppo                             // OPPO (2)
	SenderAppVivo                             // vivo (3)
	SenderAppXiaomi                           // 小米 (4)
	SenderAppMeizu                            // 魅族 (5)
	SenderAppFcm                              // FCM (6)
	_                                         // -
	SenderAppApns                             // APNs (8)
)
