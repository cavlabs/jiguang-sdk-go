// Copyright 2025 cavlabs/jiguang-sdk-go authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 删除指定文件
//   - 功能说明：删除保存在极光服务器的指定文件。
//   - 调用地址：DELETE `/v3/files/{fileID}`
//   - 接口文档：[docs.jiguang.cn]
//
// 注意事项：
//   - fileID 不存在当成功处理。
//   - 对于即时推送，建议创建推送任务 5 分钟后再执行文件删除操作，否则推送任务可能会失败；
//   - 对于文件定时推送，创建定时任务成功后，若任务被执行前文件被删除，则任务执行时推送动作将会失败。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E5%88%A0%E9%99%A4%E6%96%87%E4%BB%B6
func (f *apiv3) DeleteFile(ctx context.Context, fileID string) (*FileDeleteResult, error) {
	if f == nil {
		return nil, api.ErrNilJPushFileAPIv3
	}

	if fileID == "" {
		return nil, errors.New("`fileID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodDelete,
		Proto:  f.proto,
		URL:    f.host + "/v3/files/" + fileID,
		Auth:   f.auth,
	}
	resp, err := f.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &FileDeleteResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type FileDeleteResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *FileDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
