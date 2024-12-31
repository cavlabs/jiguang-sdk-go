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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 普通模板推送（VIP） - 指定模板 ID，模板参数（如有设置），进行立即推送。
//  - 功能说明：开发者可以将推送时所需参数通过极光 WebPortal 页面提前配置好，形成一套 “推送模板”，然后通过此 API 接口，指定 “模板 ID” 推送，大大降低开发难度，节省开发调试时间，节省服务器带宽成本。
//	- 调用地址：POST `/v3/push/template`，`id` 为模板 ID（创建模板后，由极光服务器生成），`params` 为模板参数（创建模板时，开发者设置的变量参数）列表。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_template#%E6%8C%87%E5%AE%9A%E6%A8%A1%E6%9D%BF%E3%80%90%E7%AB%8B%E5%8D%B3%E6%8E%A8%E9%80%81%E3%80%91
// 注意事项：不支持多种不同 Audience 类型的组合（Tags、AndTags、NotTags 除外）！
//  1. params 数组里面，同一个推送里面的 Audience 不允许不同类型组合（Tags、AndTags、NotTags 除外）；
//  2. params 数组里面，不同推送的 Audience 也不允许不同类型（Tags、AndTags、NotTags 除外）；
//  3. 如果是文件方式推送，一次推送只支持指定一个文件；
//  4. 如果是地理围栏方式推送，一次推送仅支持指定一个地理围栏。
// 也就是说：假设数组长度是 3，那么就表示有 3 个推送，这 3 个推送要么都是 Registration ID 方式推送，要么都是 Alias 方式推送。
func (p *apiv3) TemplateSend(ctx context.Context, id string, params []TemplateParam) (*TemplateSendResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if id == "" {
		return nil, errors.New("`id` cannot be empty")
	}
	if len(params) == 0 {
		return nil, errors.New("`params` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/template",
		Auth:   p.auth,
		Body:   &templateSendParam{ID: id, Params: params},
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TemplateSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type templateSendParam struct {
	// 【必填】模板 ID（创建模板后，由极光服务器生成）
	ID string `json:"id"`
	// 【必填】模板参数（创建模板时，开发者设置的变量参数）列表
	//
	// 注意事项：不支持多种不同 Audience 类型的组合（Tags、AndTags、NotTags 除外）！
	//  1. Params 数组里面，同一个推送里面的 Audience 不允许不同类型组合（Tags、AndTags、NotTags 除外）；
	//  2. Params 数组里面，不同推送的 Audience 也不允许不同类型（Tags、AndTags、NotTags 除外）；
	//  3. 如果是文件方式推送，一次推送只支持指定一个文件；
	//  4. 如果是地理围栏方式推送，一次推送仅支持指定一个地理围栏。
	// 也就是说：假设数组长度是 3，那么就表示有 3 个推送，这 3 个推送要么都是 Registration ID 方式推送，要么都是 Alias 方式推送。
	Params []TemplateParam `json:"params"`
}

type TemplateSendResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError    `json:"error,omitempty"`
	Data          *TemplateSendData `json:"data,omitempty"` // 模板推送结果数据
}

func (rs *TemplateSendResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
		return nil
	}

	code, err := strconv.Atoi(string(aux["code"]))
	if err != nil {
		return err
	}
	delete(aux, "code")

	// 去掉字符串头尾的 `"`
	message := strings.Trim(string(aux["message"]), `"`)
	delete(aux, "message")

	rs.Error = &api.CodeError{Code: code, Message: message}

	if v, ok := aux["data"]; ok {
		// null
		if len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l' {
			rs.Data = nil
		} else {
			var template TemplateSendData
			if err = json.Unmarshal(v, &template); err != nil {
				return err
			}
			rs.Data = &template
		}
	}

	return nil
}

func (rs TemplateSendResult) MarshalJSON() ([]byte, error) {
	rsmap := make(map[string]interface{}, 3)
	rsmap["code"] = rs.Error.Code
	rsmap["message"] = rs.Error.Message
	if rs.Data != nil {
		rsmap["data"] = rs.Data
	}
	return json.Marshal(rsmap)
}

func (rs *TemplateSendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 模板推送结果数据。
type TemplateSendData struct {
	PushList []TemplatePushDetail `json:"push_list"` // 推送列表结果。
}

// 模板推送结果明细。
type TemplatePushDetail struct {
	Error     *api.CodeError `json:"error,omitempty"`      // 推送失败时的错误信息。
	MessageID string         `json:"message_id,omitempty"` // 推送消息 ID。
	TraceID   string         `json:"trace_id,omitempty"`   // 请求时指定的标识。
}

func (rs *TemplatePushDetail) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	code, err := strconv.Atoi(string(aux["code"]))
	if err != nil {
		return err
	}
	delete(aux, "code")

	// 去掉字符串头尾的 `"`
	message := strings.Trim(string(aux["message"]), `"`)
	delete(aux, "message")

	rs.Error = &api.CodeError{Code: code, Message: message}

	if v, ok := aux["message_id"]; ok {
		// not null
		if !(len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l') {
			// 去掉字符串头尾的 `"`
			rs.MessageID = strings.Trim(string(v), `"`)
		}
	}

	if v, ok := aux["trace_id"]; ok {
		// not null
		if !(len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l') {
			// 去掉字符串头尾的 `"`
			rs.TraceID = strings.Trim(string(v), `"`)
		}
	}

	return nil
}

func (rs TemplatePushDetail) MarshalJSON() ([]byte, error) {
	rsmap := make(map[string]interface{}, 3)
	rsmap["code"] = rs.Error.Code
	if rs.Error.Message != "" {
		rsmap["message"] = rs.Error.Message
	}
	if rs.MessageID != "" {
		rsmap["message_id"] = rs.MessageID
	}
	if rs.TraceID != "" {
		rsmap["trace_id"] = rs.TraceID
	}
	return json.Marshal(rsmap)
}

func (rs *TemplatePushDetail) IsSuccess() bool {
	return rs != nil && rs.Error.IsSuccess()
}
