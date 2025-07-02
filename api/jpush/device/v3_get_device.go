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

package device

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 查询设备的标签、别名与手机号码
//   - 功能说明：获取当前设备的所有属性，包含 tags、alias 与 mobile。
//   - 调用地址：GET `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) GetDevice(ctx context.Context, registrationID string) (*DeviceGetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if registrationID == "" {
		return nil, errors.New("`registrationID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  d.proto,
		URL:    d.host + "/v3/devices/" + registrationID,
		Auth:   d.auth,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &DeviceGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type DeviceGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Tags          []string       `json:"tags,omitempty"`   // 设备标签
	Alias         string         `json:"alias,omitempty"`  // 设备别名
	Mobile        string         `json:"mobile,omitempty"` // 设备手机号码
}

func (rs *DeviceGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
