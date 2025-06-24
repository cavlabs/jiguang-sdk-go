/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

package file

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 上传文件 (Alias)
//   - 功能说明：可以将要推送的 alias 值先写入本地文件，然后将本地文件上传到极光服务器，后续就可以直接指定文件推送了。
//   - 调用地址：POST `/v3/files/alias`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6
func (f *apiv3) UploadFileForAlias(ctx context.Context, param *FileUploadParam) (*FileUploadResult, error) {
	return f.uploadFile(ctx, "alias", param)
}

// # 上传文件 (Registration ID)
//   - 功能说明：可以将要推送的 Registration ID 值先写入本地文件，然后将本地文件上传到极光服务器，后续就可以直接指定文件推送了。
//   - 调用地址：POST `/v3/files/registration_id`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6
func (f *apiv3) UploadFileForRegistrationID(ctx context.Context, param *FileUploadParam) (*FileUploadResult, error) {
	return f.uploadFile(ctx, "registration_id", param)
}

func (f *apiv3) uploadFile(ctx context.Context, forType string, param *FileUploadParam) (*FileUploadResult, error) {
	if f == nil {
		return nil, api.ErrNilJPushFileAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	uploadFile := param.File
	hasUploadFile := uploadFile != nil && reflect.ValueOf(uploadFile).Kind() != reflect.Invalid && !reflect.ValueOf(uploadFile).IsZero()
	if !hasUploadFile {
		return nil, fmt.Errorf("empty %q file", forType)
	}

	body := api.MultipartFormDataBody{
		Files: []api.FormFile{{FieldName: "filename", FileData: uploadFile}},
		FileValidator: &api.FileValidator{
			MaxSize:      10 * 1024 * 1024, // 10MB
			AllowedMimes: []string{"text/plain", "application/octet-stream"},
			AllowedExts:  []string{".txt"},
		},
	}

	if param.TTL != nil && (*param.TTL >= 1 && *param.TTL <= 720) {
		body.Fields = []api.FormField{{Name: "ttl", Value: strconv.Itoa(*param.TTL)}}
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  f.proto,
		URL:    f.host + "/v3/files/" + forType,
		Auth:   f.auth,
		Body:   body,
	}
	resp, err := f.client.FormRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &FileUploadResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type FileUploadParam struct {
	// 【必填】文件名。
	//  - 文件一行一个内容，忽略每行的前后空格和换行符作为实际的 Registration ID 值或者 Alias 值。
	//  - 文件只支持 txt 格式，要求文件内容必须是 UTF-8 编码。
	//  - 文件不超过 10M。
	//  - 文件自创建起，若不指定 TTL 参数，则服务器会默认保存 720 小时（即 30 天），超过有效期，服务器自动将文件删除。
	//  - 有效期内的文件不允许超过 20 个。
	File interface{} `json:"filename"`
	// 【可选】文件有效期，单位：小时，默认值：720，取值范围：1~720，即 1 小时至 720 小时（30天）。
	//  - 超过有效期，服务器自动将文件删除。
	TTL *int `json:"ttl,omitempty"`
}

type FileUploadResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	FileID        string         `json:"file_id,omitempty"` // 文件 ID，后续供删除、查询、推送等时使用。
}

func (rs *FileUploadResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
