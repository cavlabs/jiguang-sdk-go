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
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 查询模板
//  - 功能说明：查询短信模板。
//  - 调用地址：GET `/v1/templates/{tempID}`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E6%9F%A5%E8%AF%A2%E6%A8%A1%E6%9D%BF-api
func (s *apiv1) GetTemplate(ctx context.Context, tempID int64) (*TemplateGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if tempID < 1 {
		return nil, errors.New("invalid `tempID`, required: >0")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/templates/" + strconv.FormatInt(tempID, 10),
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TemplateGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TemplateGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	TempID        int64          `json:"temp_id,omitempty"`  // 模板 ID
	Template      string         `json:"template,omitempty"` // 模板内容
	Type          int            `json:"type,omitempty"`     // 模板类型，1 为验证码类，2 为通知类，3 为营销类。
	TTL           int            `json:"ttl,omitempty"`      // 验证码有效期，单位为秒。
	Remark        string         `json:"remark,omitempty"`   // 正文模版的发送场景及发送对象的简略描述。
	Status        int            `json:"status,omitempty"`   // 模板审核状态：0 - 审核中；1 - 审核通过；2 - 审核不通过
}

func (rs *TemplateGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
