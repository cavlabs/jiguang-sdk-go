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

// 创建模板
//  - 功能说明：创建短信模板。
//  - 调用地址：POST `/v1/templates`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E5%88%9B%E5%BB%BA%E6%A8%A1%E6%9D%BF-api
func (s *apiv1) CreateTemplate(ctx context.Context, param *TemplateCreateParam) (*TemplateCreateResult, error) {
	return s.template(ctx, 0, param, "create")
}

// 修改模板
//  - 功能说明：修改审核不通过的模板，并再次提交审核。
//  - 调用地址：PUT `/v1/templates/{tempID}`
//  - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E4%BF%AE%E6%94%B9%E6%A8%A1%E6%9D%BF-api
func (s *apiv1) UpdateTemplate(ctx context.Context, tempID int64, param *TemplateUpdateParam) (*TemplateUpdateResult, error) {
	return s.template(ctx, tempID, param, "update")
}

// 创建或修改模板，createOrUpdate 可选值：create、update。
func (s *apiv1) template(ctx context.Context, tempID int64, param *TemplateCreateParam, createOrUpdate string) (*TemplateCreateResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}
	l := len(param.Template)
	if l < 1 || l > 500 {
		return nil, errors.New("`template` must be 1-500 characters")
	}
	t := param.Type
	if t < 1 || t > 3 {
		return nil, errors.New("invalid `type`, must be 1-3")
	}
	if t == 1 && param.TTL < 1 {
		return nil, errors.New("`ttl` is required when `type` is 1")
	}
	if param.TTL > 86400 {
		return nil, errors.New("`ttl` must be ≤86400 seconds")
	}
	l = len(param.Remark)
	if l > 100 {
		return nil, errors.New("`remark` too long, must be ≤100 characters")
	}

	var body interface{} = param

	method, url := http.MethodPost, s.host+"/v1/templates"
	if createOrUpdate == "update" {
		method = http.MethodPut
		url += "/" + strconv.FormatInt(tempID, 10)

		body = &templateUpdateParam{
			TemplateCreateParam: *param,
			TempID:              tempID,
		}
	}

	req := &api.Request{
		Method: method,
		Proto:  s.proto,
		URL:    url,
		Auth:   s.auth,
		Body:   body,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TemplateCreateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TemplateCreateParam struct {
	// 【必填】模板内容。
	//  - 1. 短信内容不超过 500 个字，短信内容包括：签名、正文、退订方式（仅营销短信），创建模版时请预留签名等文字的字数；
	//  短信计费规则：70 字按 1 条短信计费，超过 70 字，按 67 字 / 条逐条计费。单个汉字、标点、英文都算一个字、回车记两个字。
	//  若短信中包含变量，计费条数以实际发送的字数为准；
	//  - 2. 验证码模版仅支持设置一个变量，且变量名必须为 code；
	//  - 3. 通知、营销短信中，变量名仅支持英文及数字，若含有链接变量，变量名必须为 url ，为避免短信发送时因进入人工审核而导致发送延迟，
	//  请在 Remark 参数中填写链接以报备，支持不设置参数。
	Template string `json:"template"`
	// 【必填】模板类型，1 为验证码类，2 为通知类，3 为营销类。
	Type int `json:"type,omitempty"`
	// 【可选】验证码有效期，必须大于 0 且不超过 86400 ，单位为秒（当模板类型为 1 时必传）。
	TTL int `json:"ttl,omitempty"`
	// 【可选】简略描述正文模版的发送场景及发送对象，不超过 100 字（包含链接的情况下需要将链接地址放置在最前面，请求才会成功）。
	Remark string `json:"remark,omitempty"`
}

type TemplateCreateResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	TempID        int64          `json:"temp_id,omitempty"` // 模板 ID
}

func (rs *TemplateCreateResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

type templateUpdateParam struct {
	TemplateCreateParam
	TempID int64 `json:"temp_id"` // 【必填】模板 ID
}

type (
	TemplateUpdateParam  = TemplateCreateParam
	TemplateUpdateResult = TemplateCreateResult
)
