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

// # 【短信】消息
//
// 包括：极光短信 (JSMS)、CMPP 短信 (CMPP)。
type SMS interface {
	getType() string // jsms: 极光短信；cmpp: CMPP 短信
}

// ---------------------------------------------------------------------------------------------------------------------

// # 极光【短信】消息
//
// 可参考 jsms.MessageSendParam。
type JSMS struct {
	SignID     int                    `json:"sign_id,omitempty"`   // 【可选】签名 ID，该字段为空则使用应用默认签名。
	TempID     int64                  `json:"temp_id"`             // 【必填】模板 ID。
	TempParams map[string]interface{} `json:"temp_para,omitempty"` // 【可选】模板参数，需要替换的参数名和参数值的键值对。
}

var _ SMS = &JSMS{}

func (m *JSMS) getType() string { return "jsms" }

// ---------------------------------------------------------------------------------------------------------------------

// # CMPP【短信】消息
type CMPP struct {
	Content string `json:"content"` // 【必填】消息内容。
}

var _ SMS = &CMPP{}

func (m *CMPP) getType() string { return "cmpp" }
