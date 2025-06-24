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

package device

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 修改测试设备（VIP）
//   - 功能说明：修改一个指定的测试设备。
//   - 调用地址：PUT `/v3/test/model/update`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E4%BF%AE%E6%94%B9%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87
func (d *apiv3) UpdateTestDevice(ctx context.Context, param *TestDeviceUpdateParam) (*TestDeviceUpdateResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPut,
		Proto:  d.proto,
		URL:    d.host + "/v3/test/model/update",
		Auth:   d.auth,
		Body:   param,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TestDeviceUpdateResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type (
	TestDeviceUpdateParam  = TestDeviceAddParam
	TestDeviceUpdateResult = TestDeviceAddResult
)
