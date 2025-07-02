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

package report

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 送达状态查询（VIP）
//   - 功能说明：查询已推送的一条消息在一组设备上的送达状态。
//   - 调用地址：POST `/v3/status/message`，`msgID` 为指定的消息 ID；`registrationIDs` 为必填参数，即需要获取在线状态的设备标识 Registration ID 集合，最多支持 1000 个。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report#%E9%80%81%E8%BE%BE%E7%8A%B6%E6%80%81%E6%9F%A5%E8%AF%A2%EF%BC%88vip%EF%BC%89
func (r *apiv3) GetMessageStatus(ctx context.Context, msgID string, registrationIDs []string, date *jiguang.LocalDate) (*MessageStatusGetResult, error) {
	if r == nil {
		return nil, api.ErrNilJPushReportAPIv3
	}

	if msgID == "" {
		return nil, errors.New("`msgID` cannot be empty")
	}
	l := len(registrationIDs)
	if l == 0 {
		return nil, errors.New("`registrationIDs` cannot be empty")
	}
	if l > 1000 {
		return nil, errors.New("`registrationIDs` cannot be more than 1000")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  r.proto,
		URL:    r.host + "/v3/status/message",
		Auth:   r.auth,
		Body: &messageStatusGetParam{
			MsgID:           msgID,
			RegistrationIDs: registrationIDs,
			Date:            date,
		},
	}
	resp, err := r.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &MessageStatusGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type messageStatusGetParam struct {
	MsgID           string             `json:"msg_id"`
	RegistrationIDs []string           `json:"registration_ids"`
	Date            *jiguang.LocalDate `json:"date,omitempty"`
}

type MessageStatusGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError           `json:"error,omitempty"`
	Status        map[string]MessageStatus `json:"-"` // key 为 Registration ID
}

func (rs *MessageStatusGetResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
	}

	if len(aux) > 0 {
		rs.Status = make(map[string]MessageStatus, len(aux))
		for k, v := range aux {
			// null
			if len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l' {
				continue
			}
			var message struct {
				Status *MessageStatus `json:"status,omitempty"`
			}
			if err := json.Unmarshal(v, &message); err != nil {
				return err
			}
			if message.Status != nil {
				rs.Status[k] = *message.Status
			}
		}
	}

	return nil
}

func (rs MessageStatusGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else if len(rs.Status) > 0 {
		type message struct {
			Status MessageStatus `json:"status"`
		}
		data := make(map[string]message, len(rs.Status))
		for k, v := range rs.Status {
			data[k] = message{v}
		}
		return json.Marshal(data)
	}
	return []byte("null"), nil
}

func (rs *MessageStatusGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
