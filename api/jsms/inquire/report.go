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

package inquire

import (
	"encoding/json"
	"strconv"
	"time"
)

// # 短信送达状态
//
// SMS_REPORT 回执数据。
type ReportData struct {
	// 原始 JSON 数据
	RawJSON string `json:"-"`
	// API 调用的时候返回的 msg_id
	MsgID string `json:"msgId"`
	// 发送状态返回码，详见 [docs.jiguang.cn] 文档说明。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/callback#%E5%8F%91%E9%80%81%E7%8A%B6%E6%80%81%E8%BF%94%E5%9B%9E%E7%A0%81
	Status int `json:"status"`
	// 短信送达手机号
	Phone string `json:"phone"`
	// 短信送达时间
	ReceiveTime time.Time `json:"receiveTime"`
}

func (r *ReportData) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ReportData` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ReportData
	aux := &struct {
		MsgID       json.RawMessage `json:"msgId"`       // 兼容字符串和数字类型的 `msgId`
		ReceiveTime interface{}     `json:"receiveTime"` // 兼容时间戳和时间字符串类型的 `receiveTime`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var strMsgID string
	if err := json.Unmarshal(aux.MsgID, &strMsgID); err == nil {
		r.MsgID = strMsgID
	} else {
		var intMsgID int64
		if err = json.Unmarshal(aux.MsgID, &intMsgID); err == nil {
			r.MsgID = strconv.FormatInt(intMsgID, 10)
		} else {
			r.MsgID = string(aux.MsgID)
		}
	}

	switch v := aux.ReceiveTime.(type) {
	case float64: // 时间戳，毫秒值
		ts := int64(v)
		r.ReceiveTime = time.Unix(ts/1000, (ts%1000)*int64(time.Millisecond))
	case string: // 时间字符串
		if t, err := normalizeAndParseTime(v); err == nil {
			r.ReceiveTime = t
		}
	}

	return nil
}

func (r ReportData) MarshalJSON() ([]byte, error) {
	return []byte(r.RawJSON), nil
}
