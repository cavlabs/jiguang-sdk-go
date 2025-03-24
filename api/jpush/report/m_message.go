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

package report

import "fmt"

// # 消息推送送达状态
type MessageStatus int

const (
	MessageStatusReceived       MessageStatus = iota // 送达
	MessageStatusNotReceived                         // 未送达
	MessageStatusInvalidRegID                        // 设备标识 Registration ID 不属于该应用
	MessageStatusRegIDNotTarget                      // 设备标识 Registration ID 属于该应用，但不是该条消息的推送目标
	MessageStatusSystemError                         // 系统异常
)

var messageStatusDescs = map[MessageStatus]string{
	MessageStatusReceived:       "Received",
	MessageStatusNotReceived:    "Not Received",
	MessageStatusInvalidRegID:   "Invalid Registration ID",
	MessageStatusRegIDNotTarget: "Registration ID Not Target",
	MessageStatusSystemError:    "System Error",
}

func (s MessageStatus) String() string {
	if desc, exists := messageStatusDescs[s]; exists {
		return desc
	}
	return fmt.Sprintf("Unknown Status [%d]", s)
}
