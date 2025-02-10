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

package jsms

import (
	"encoding/json"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

type sendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	MsgID         string         `json:"msg_id,omitempty"` // 消息 ID
}

func (rs *sendResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `alias` 不继承原始类型 `sendResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type alias sendResult
	aux := struct {
		MsgID json.RawMessage `json:"msg_id"`
		*alias
	}{
		alias: (*alias)(rs),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var strMsgID string
	if err := json.Unmarshal(aux.MsgID, &strMsgID); err == nil {
		if strMsgID != "" { // omitempty
			rs.MsgID = strMsgID
		}
		return nil
	} else {
		var intMsgID int64 // 兼容数字类型的 `msg_id`
		if err = json.Unmarshal(aux.MsgID, &intMsgID); err == nil {
			if intMsgID != 0 { // omitempty
				rs.MsgID = strconv.FormatInt(intMsgID, 10)
			}
		} else {
			// return err
			rs.MsgID = string(aux.MsgID)
		}
	}

	return nil
}

func (rs *sendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
