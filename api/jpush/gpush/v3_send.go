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

package gpush

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
)

// 分组推送
//  - 功能说明：该 API 用于为开发者在 portal 端创建的应用分组创建推送。
//	- 调用地址：POST `/v3/grouppush`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_grouppush
// 注意：暂不支持 Options 中 OverrideMsgID 的属性；分组推送仅在官网支持设置定时，调 Schedule API 时不支持。
func (gp *apiv3) Send(ctx context.Context, param *SendParam) (*SendResult, error) {
	return gp.CustomSend(ctx, param)
}

// 自定义分组推送：如果遇到 Send 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
func (gp *apiv3) CustomSend(ctx context.Context, param interface{}) (*SendResult, error) {
	if gp == nil {
		return nil, api.ErrNilJPushGroupPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  gp.proto,
		URL:    gp.host + "/v3/grouppush",
		Auth:   gp.auth,
		Body:   param,
	}
	resp, err := gp.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SendParam = push.SendParam

type SendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError         `json:"error,omitempty"`
	GroupMsgID    string                 `json:"group_msgid,omitempty"` // 分组推送消息 ID，用于唯一标识本次分组推送请求，后续可用于分组推送统计。2020.08.11 新增
	Successes     map[string]SendSuccess `json:"-"`                     // 推送成功集合
	Errors        map[string]SendError   `json:"-"`                     // 推送失败错误集合
}

func (rs *SendResult) UnmarshalJSON(data []byte) error {
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

	if groupMsgID, ok := aux["group_msgid"]; ok {
		// 去掉字符串头尾的 `"`
		rs.GroupMsgID = strings.Trim(string(groupMsgID), `"`)
		delete(aux, "group_msgid")
	}

	if len(aux) > 0 {
		rs.Successes = make(map[string]SendSuccess)
		rs.Errors = make(map[string]SendError)
		for k, v := range aux {
			var sendSuccess SendSuccess
			if err := json.Unmarshal(v, &sendSuccess); err != nil {
				return err
			} else if sendSuccess.MsgID != "" {
				rs.Successes[k] = sendSuccess
				continue
			}

			var sendFail struct {
				Error *SendError `json:"error,omitempty"`
			}
			if err := json.Unmarshal(v, &sendFail); err != nil {
				return err
			} else if sendFail.Error != nil {
				rs.Errors[k] = *sendFail.Error
			}
		}
	}

	return nil
}

func (rs SendResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	}

	data := make(map[string]interface{})

	if len(rs.Successes) > 0 {
		for k, v := range rs.Successes {
			data[k] = v
		}
	}

	if len(rs.Errors) > 0 {
		type sendFail struct {
			Error SendError `json:"error"`
		}
		for k, v := range rs.Errors {
			data[k] = sendFail{v}
		}
	}

	data["group_msgid"] = rs.GroupMsgID

	return json.Marshal(data)
}

func (rs *SendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 推送成功
type SendSuccess struct {
	MsgID  string `json:"msg_id,omitempty"` // 推送消息 ID
	SendNo string `json:"sendno,omitempty"` // 推送序号
}

// 推送失败错误
type SendError struct {
	api.CodeError
	IllegalRIDs []string `json:"illegal_rids,omitempty"` // 非法 Registration ID 集合
}

func (e *SendError) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}

func (e *SendError) String() string {
	if e == nil {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}
