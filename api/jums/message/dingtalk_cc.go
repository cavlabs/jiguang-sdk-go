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

import "errors"

// # 【钉钉】工作通知消息类型
type DingtalkCCMsgType string

const DingtalkCCMsgTypeText DingtalkCCMsgType = "text" // 文本消息

// # 【钉钉】工作通知
type DingtalkCC struct {
	// 【必填】消息。
	Msg *DingtalkCCMsg `json:"msg"`
}

// # 【钉钉】工作通知 - 消息
type DingtalkCCMsg struct {
	// 【必填】消息类型。
	//  - [文本消息] 类型为: DingtalkCCMsgTypeText，此时 Text 字段必填。
	MsgType DingtalkCCMsgType `json:"msgtype"`
	// 【可选】文本消息内容 (DingtalkCCTypeText)。
	Text *DingtalkCCMsgText `json:"text,omitempty"`
}

// 验证消息参数。
func (d *DingtalkCCMsg) Validate() error {
	switch d.MsgType {
	case DingtalkCCMsgTypeText:
		if d.Text == nil {
			return errors.New("msg_dingtalk_cc.[*].msg.text is required when msgtype is `text`")
		}
	default:
		return errors.New(string("unsupported msg_dingtalk_cc.[*].msg.msgtype " + d.MsgType))
	}
	return nil
}

// # 【钉钉】工作通知 - 文本消息
type DingtalkCCMsgText struct {
	Content string `json:"content"` // 【必填】消息内容，建议 500 字符以内。
}
