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
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 查询指定文件详情
//   - 功能说明：查询保存在极光服务器的，指定文件的详细信息。
//   - 调用地址：GET `/v3/files/{fileID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E6%9F%A5%E8%AF%A2%E6%8C%87%E5%AE%9A%E6%96%87%E4%BB%B6%E8%AF%A6%E6%83%85
func (f *apiv3) GetFile(ctx context.Context, fileID string) (*FileGetResult, error) {
	if f == nil {
		return nil, api.ErrNilJPushFileAPIv3
	}

	if fileID == "" {
		return nil, errors.New("`fileID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  f.proto,
		URL:    f.host + "/v3/files/" + fileID,
		Auth:   f.auth,
	}
	resp, err := f.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &FileGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type FileGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError         `json:"error,omitempty"`
	Type          string                 `json:"type,omitempty"`        // 文件类型，`alias` 或 `registration_id`
	FileID        string                 `json:"file_id,omitempty"`     // 文件 ID
	FileName      string                 `json:"file_name,omitempty"`   // 文件上传时的文件名，自 2021.10.14 日启用
	CreateTime    *jiguang.LocalDateTime `json:"create_time,omitempty"` // 文件创建（上传）时间
	ExpireTime    *jiguang.LocalDateTime `json:"expire_time,omitempty"` // 文件过期时间，当设置了文件有效期（即提交了 TTL 参数），则会携带该字段返回
}

func (rs *FileGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
