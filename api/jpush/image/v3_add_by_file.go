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

package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 新增图片（文件方式）
//   - 功能说明：通过上传图片文件形式来新增一个适配，该接口目前仅支持小米和 OPPO。不过从 2023.08 开始，小米官方在新设备/系统已经不再支持推送时动态设置小图标、右侧图标、大图片功能，对于历史设备和应用也在逐步覆盖，等于不再支持推送时动态设置小图标、右侧图标、大图片功能；
//   - 调用地址：POST `/v3/images/byfiles`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%96%B0%E5%A2%9E%E5%9B%BE%E7%89%87%EF%BC%88%E6%96%87%E4%BB%B6%E6%96%B9%E5%BC%8F%EF%BC%89
func (i *apiv3) AddImageByFile(ctx context.Context, param *AddByFileParam) (*AddByFileResult, error) {
	if i == nil {
		return nil, api.ErrNilJPushImageAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}
	if !param.ImageType.IsValid() {
		return nil, errors.New("invalid image type " + strconv.Itoa(int(param.ImageType)))
	}

	xiaomiImageFile, oppoImageFile := param.XiaomiImageFile, param.OppoImageFile

	hasXiaomiImageFile := xiaomiImageFile != nil && reflect.ValueOf(xiaomiImageFile).Kind() != reflect.Invalid && !reflect.ValueOf(xiaomiImageFile).IsZero()
	hasOppoImageFile := oppoImageFile != nil && reflect.ValueOf(oppoImageFile).Kind() != reflect.Invalid && !reflect.ValueOf(oppoImageFile).IsZero()

	if !hasXiaomiImageFile && !hasOppoImageFile {
		return nil, errors.New("at least one of `xiaomi_file` and `oppo_file` must be set")
	}

	var body = api.MultipartFormDataBody{
		Fields: []api.FormField{{Name: "image_type", Value: strconv.Itoa(int(param.ImageType))}},
		FileValidator: &api.FileValidator{
			MaxSize:      1 * 1024 * 1024, // 1MB
			AllowedMimes: []string{"image/jpeg", "image/png"},
			AllowedExts:  []string{".jpg", ".jpeg", ".png"},
		},
	}

	if hasXiaomiImageFile {
		body.Files = append(body.Files, api.FormFile{
			FieldName: "xiaomi_file",
			FileData:  xiaomiImageFile,
		})
	}

	if hasOppoImageFile {
		body.Files = append(body.Files, api.FormFile{
			FieldName: "oppo_file",
			FileData:  oppoImageFile,
		})
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  i.proto,
		URL:    i.host + "/v3/images/byfiles",
		Auth:   i.auth,
		Body:   body,
	}
	resp, err := i.client.FormRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AddByFileResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 新增图片（文件方式）请求参数正文。
//   - XiaomiImageFile 和 OppoImageFile 不能同时为空。
type AddByFileParam struct {
	// 【必填】指定适配的图片类型。
	ImageType Type `json:"image_type"`
	// 【可选】上传配置小米通道的图片文件。
	//  - 本接口将会对该图片文件大小进行校验，若不适配小米对该图片的要求，则返回错误，小米对该图片对要求参考 [小米-图片上传]；
	//  - 小米从 2023.08 开始不再支持推送时动态设置小图标、右侧图标、大图片功能，开发者可不再设置此字段值。
	// [小米-图片上传]: https://dev.mi.com/console/doc/detail?pId=1278#4_4_2
	XiaomiImageFile interface{} `json:"xiaomi_file,omitempty"`
	// 【可选】上传配置 OPPO 通道的图片文件；
	//  - 本接口将会对该图片文件大小进行校验，若不适配 OPPO 对该图片的要求，则返回错误，OPPO 对该图片对要求参考 [OPPO-图片上传]。
	// [OPPO-图片上传]: https://open.oppomobile.com/new/developmentDoc/info?id=11241
	OppoImageFile interface{} `json:"oppo_file,omitempty"`
}

// 新增图片（文件方式）响应结果正文。
type AddByFileResult = AddByUrlResult
