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

package admin

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 极光 APP 证书上传
//   - 功能说明：上传开发或生产证书到对应的极光 APP。
//   - 调用地址：POST `/v1/app/{appKey}/certificate`，`appKey` 为 APP 的唯一标识。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#%E8%AF%81%E4%B9%A6%E4%B8%8A%E4%BC%A0
func (a *apiv1) UploadCertificate(ctx context.Context, appKey string, param *CertificateUploadParam) (*CertificateUploadResult, error) {
	if a == nil {
		return nil, api.ErrNilJPushAdminAPIv1
	}

	if appKey == "" {
		return nil, errors.New("`appKey` cannot be empty")
	}
	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	devCertPwd, devCertFile := param.DevCertificatePassword, param.DevCertificateFile
	proCertPwd, proCertFile := param.ProCertificatePassword, param.ProCertificateFile

	hasDevCertFile := devCertFile != nil && reflect.ValueOf(devCertFile).Kind() != reflect.Invalid && !reflect.ValueOf(devCertFile).IsZero()
	hasProCertFile := proCertFile != nil && reflect.ValueOf(proCertFile).Kind() != reflect.Invalid && !reflect.ValueOf(proCertFile).IsZero()

	if !hasDevCertFile && !hasProCertFile {
		return nil, errors.New("either `devCertificateFile` or `proCertificateFile` must be set")
	}

	var body api.MultipartFormDataBody

	if hasDevCertFile {
		if devCertPwd == "" {
			return nil, errors.New("`devCertificatePassword` is required when `devCertificateFile` is set")
		}
		body.Fields = append(body.Fields, api.FormField{
			Name:  "devCertificatePassword",
			Value: devCertPwd,
		})
		body.Files = append(body.Files, api.FormFile{
			FieldName: "devCertificateFile",
			FileData:  devCertFile,
		})
	}

	if hasProCertFile {
		if proCertPwd == "" {
			return nil, errors.New("`proCertificatePassword` is required when `proCertificateFile` is set")
		}
		body.Fields = append(body.Fields, api.FormField{
			Name:  "proCertificatePassword",
			Value: proCertPwd,
		})
		body.Files = append(body.Files, api.FormFile{
			FieldName: "proCertificateFile",
			FileData:  proCertFile,
		})
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  a.proto,
		URL:    a.host + "/v1/app/" + appKey + "/certificate",
		Auth:   a.auth,
		Body:   body,
	}
	resp, err := a.client.FormRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &CertificateUploadResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type CertificateUploadParam struct {
	DevCertificatePassword string      `json:"devCertificatePassword,omitempty"` // 「开发」证书密码
	DevCertificateFile     interface{} `json:"devCertificateFile,omitempty"`     // 「开发」证书文件
	ProCertificatePassword string      `json:"proCertificatePassword,omitempty"` // 「生产」证书密码
	ProCertificateFile     interface{} `json:"proCertificateFile,omitempty"`     // 「生产」证书文件
}

type CertificateUploadResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Success       string         `json:"success,omitempty"`
}

func (rs *CertificateUploadResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
