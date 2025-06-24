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

// # 新增测试设备（VIP）
//   - 功能说明：新增一个测试设备，确保测试模式下的每次推送仅触达测试用户。
//   - 调用地址：POST `/v3/test/model/add`
//   - 接口文档：[docs.jiguang.cn]
//
// 详细功能逻辑可参考文档：[测试模式]。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%96%B0%E5%A2%9E%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87
// [测试模式]: https://docs.jiguang.cn/jpush/console/push_manage/testmode
func (d *apiv3) AddTestDevice(ctx context.Context, param *TestDeviceAddParam) (*TestDeviceAddResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/test/model/add",
		Auth:   d.auth,
		Body:   param,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TestDeviceAddResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type TestDeviceAddParam struct {
	DeviceName     string `json:"device_name"`     // 开发者自定义的设备名称
	RegistrationID string `json:"registration_id"` // 极光生成的设备唯一标识
}

type TestDeviceAddResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *TestDeviceAddResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
