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

package jums

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 上传素材
//   - 功能说明：可以上传图片、视频、文件等素材。
//   - 调用地址：POST `/v1/material`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_material#%E4%B8%8A%E4%BC%A0%E7%B4%A0%E6%9D%90
func (u *apiv1) UploadMaterial(ctx context.Context, param *MaterialUploadParam) (*MaterialUploadResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}
	if param.Type != "attachment" {
		// return nil, errors.New("`type` must be `attachment`")
		param.Type = "attachment"
	}
	if param.TimeToLive < 1 || param.TimeToLive > 168 {
		// return nil, errors.New("`time_to_live` must be between 1 and 168")
		param.TimeToLive = 0 // default 24 hours
	}

	uploadFile := param.File
	hasUploadFile := uploadFile != nil && reflect.ValueOf(uploadFile).Kind() != reflect.Invalid && !reflect.ValueOf(uploadFile).IsZero()
	if !hasUploadFile {
		return nil, errors.New("empty attachment file")
	}

	body := api.MultipartFormDataBody{
		Fields: []api.FormField{{Name: "type", Value: param.Type}},
		Files:  []api.FormFile{{FieldName: "file", FileData: uploadFile}},
	}
	if param.TimeToLive != 0 {
		body.Fields = append(body.Fields, api.FormField{Name: "time_to_live", Value: strconv.Itoa(param.TimeToLive)})
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  u.proto,
		URL:    u.host + "/v1/material",
		Auth:   u.auth,
		Body:   body,
	}
	resp, err := u.client.FormRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &MaterialUploadResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type MaterialUploadParam struct {
	Type       string      `json:"type"`                   // 【必填】素材类型，当前仅支持 attachment 1 种类型，代表 “附件”。
	File       interface{} `json:"file"`                   // 【必填】一次仅允许上传 1 个文件。
	TimeToLive int         `json:"time_to_live,omitempty"` // 【可选】素材有效期，默认值是 24 小时。对于 “附件” 类型，有效期可传的数值范围在 [1, 168] 间，即 1 小时至 168 小时（7 天）。
}

type MaterialUploadResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	URL string `json:"url,omitempty"` // 上传成功时，将返回文件的地址。对于 “附件” 来说：该地址仅可用于在消息发送中传值，无法直接请求下载。
}

func (rs *MaterialUploadResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
