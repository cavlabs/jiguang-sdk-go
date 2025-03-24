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

package gpush

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 分组文件推送（VIP）
//   - 功能说明：该 API 用于为开发者在 portal 端创建的应用分组进行文件推送，推送参数和格式跟文件推送一样。
//   - 调用地址：POST `/v3/grouppush/file`
//   - 接口文档：[docs.jiguang.cn]
//
// 注意事项：
//   - 此接口只对已经开通权限对客户支持，未开通权限客户使用将会返回错误返回码 2007。
//   - 调用文件上传接口获取 fileID 时，需要使用 devKey 和 devSecret 进行验证，详情参考 [文件上传接口]。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E5%BA%94%E7%94%A8%E5%88%86%E7%BB%84%E6%96%87%E4%BB%B6%E6%8E%A8%E9%80%81-api%EF%BC%88vip%EF%BC%89
// [文件上传接口]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E8%B0%83%E7%94%A8%E9%AA%8C%E8%AF%81-1
func (gp *apiv3) SendByFile(ctx context.Context, param *SendParam) (*SendResult, error) {
	return gp.CustomSendByFile(ctx, param)
}

// # 自定义分组文件推送
//
// 如果遇到 SendByFile 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
func (gp *apiv3) CustomSendByFile(ctx context.Context, param interface{}) (*SendResult, error) {
	if gp == nil {
		return nil, api.ErrNilJPushGroupPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  gp.proto,
		URL:    gp.host + "/v3/grouppush/file",
		Auth:   gp.auth,
		Body:   param,
	}
	resp, err := gp.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
