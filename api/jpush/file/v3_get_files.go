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

package file

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 查询有效文件列表
//  - 功能说明：获取当前保存在极光服务器的有效文件列表。
//	- 调用地址：GET `/v3/files`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E6%9F%A5%E8%AF%A2%E6%9C%89%E6%95%88%E6%96%87%E4%BB%B6%E5%88%97%E8%A1%A8
func (f *apiv3) GetFiles(ctx context.Context) (*FilesGetResult, error) {
	if f == nil {
		return nil, api.ErrNilJPushFileAPIv3
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  f.proto,
		URL:    f.host + "/v3/files",
		Auth:   f.auth,
	}
	resp, err := f.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &FilesGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type FilesGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError  `json:"error,omitempty"`
	TotalCount    *int            `json:"total_count,omitempty"` // 当前有效的文件数
	Files         []FileGetResult `json:"files,omitempty"`       // 文件列表详情
}

func (rs *FilesGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
