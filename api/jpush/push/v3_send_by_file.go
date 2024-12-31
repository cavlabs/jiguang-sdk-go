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

package push

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 文件推送（文件立即推送）
//  - 功能说明：支持指定文件唯一标识（fileID）进行推送，文件唯一标识（fileID）可以参考 File API v3 的 文件上传接口 获得。
//	- 调用地址：POST `/v3/push/file`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E6%96%87%E4%BB%B6%E6%8E%A8%E9%80%81-api
// 注意事项：
//  1. “文件定时推送” 请使用 ScheduleSend 接口；
//  2. 厂商配额查询 API 接口频率和推送 API 接口频率共享，消耗推送 API 接口频率；
//  3. 建议创建推送任务 5 分钟后再执行文件删除操作，否则推送任务可能会失败。
func (p *apiv3) SendByFile(ctx context.Context, param *SendParam) (*SendResult, error) {
	return p.CustomSendByFile(ctx, param)
}

// 自定义文件推送：如果遇到 SendByFile 接口没有及时补充字段的情况，可以自行构建 JSON，调用此接口。
func (p *apiv3) CustomSendByFile(ctx context.Context, param interface{}) (*SendResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/file",
		Auth:   p.auth,
		Body:   param,
	}
	resp, err := p.client.Request(ctx, req)
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
