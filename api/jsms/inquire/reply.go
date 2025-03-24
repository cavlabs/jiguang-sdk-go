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

package inquire

import (
	"encoding/json"
	"time"
)

// # 用户回复消息
//
// SMS_REPLY 回执数据。
type ReplyData struct {
	RawJSON   string    `json:"-"`         // 原始 JSON 数据
	Phone     string    `json:"phone"`     // 主叫号码（用户手机号码）
	ReplyTime time.Time `json:"replyTime"` // 消息送达到极光平台的时间
	Content   string    `json:"content"`   // 用户回复的消息内容
}

func (r *ReplyData) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ReplyData` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ReplyData
	aux := &struct {
		ReplyTime interface{} `json:"replyTime"` // 兼容时间戳和时间字符串类型的 `replyTime`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch v := aux.ReplyTime.(type) {
	case float64: // 时间戳，毫秒值
		ts := int64(v)
		r.ReplyTime = time.Unix(ts/1000, (ts%1000)*int64(time.Millisecond))
	case string: // 时间字符串
		if t, err := normalizeAndParseTime(v); err == nil {
			r.ReplyTime = t
		}
	}

	return nil
}

func (r ReplyData) MarshalJSON() ([]byte, error) {
	return []byte(r.RawJSON), nil
}
