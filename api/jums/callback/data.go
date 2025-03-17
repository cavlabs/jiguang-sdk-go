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

package callback

import (
	"encoding/json"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/api/jums/message"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 回调数据列表。
type DataList struct {
	RawJSON string `json:"-"`     // 原始 JSON 数据
	Items   []Data `json:"items"` // 回调参数数据项列表
}

// 回调数据项。
type Data struct {
	Type           Type                            `json:"type"`                      // 回调数据类型
	ChannelKey     string                          `json:"channelkey,omitempty"`      // 该条记录所对应的渠道的唯一标识
	MsgID          string                          `json:"msgid,omitempty"`           // 该条记录所对应的消息唯一 ID
	UserID         string                          `json:"userid,omitempty"`          // 该条记录所对应的用户 UserID，使用自定义通道 ID 发送时值为空
	SenderID       string                          `json:"sender_id,omitempty"`       // 该条记录所对应的通道注册 ID。向通道下发消息后产生的所有记录，以及使用自定义通道 ID 发送时有值。举例：如果是 APP 通道，则 SenderID 就是 APP 的 RegistrationID
	Sender         *message.Sender                 `json:"sender,omitempty"`          // 该条记录所对应的下发通道/处理模块
	WechatwkSender *message.SenderWechatwkCategory `json:"wechatwk_sender,omitempty"` // 企业微信通道的通道类别，Sender = message.SenderWechatwk 时有值
	AppPlatform    *platform.Platform              `json:"app_platform,omitempty"`    // APP 消息下发所对应的设备的操作系统，取值：Sender = message.SenderApp 时有值
	AppSender      *message.SenderAppCategory      `json:"app_sender,omitempty"`      // APP 厂商通道，Sender = message.SenderApp 时有值
	SentTime       *jiguang.Timestamp              `json:"sent_time,omitempty"`       // 产生该记录的时间点，秒级时间戳
	Error          *api.CodeError                  `json:"error,omitempty"`           // 目标无效/提交失败/送达失败/撤回失败的原因，当 Type = TargetInvalid / SentFail / ReceivedFail / RetractedFail 时有值
	Params         map[string]interface{}          `json:"params,omitempty"`          // 自定义参数，用户在发送消息请求时在 message.Callback 里自行指定的参数 Params
}

func (d Data) MarshalJSON() ([]byte, error) {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `Data` 的方法，包括 `MarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias Data
	aux := struct {
		AppPlatform *int `json:"app_platform,omitempty"`
		Alias
	}{
		Alias: (Alias)(d),
	}
	if d.AppPlatform != nil {
		aux.AppPlatform = jiguang.Int(d.AppPlatform.Index())
	}
	return json.Marshal(&aux)
}
