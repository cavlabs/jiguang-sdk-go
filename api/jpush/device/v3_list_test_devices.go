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
	"fmt"
	"net/http"
	"net/url"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 获取测试设备列表（VIP）
//   - 功能说明：分页获取测试设备列表。
//   - 调用地址：GET `/v3/test/model/list?page={page}&page_size={pageSize}&device_name={deviceName}&registration_id={registrationID}`；
//     `page` 为查询页码，`pageSize` 为每页记录条数，`deviceName` 为开发者自定义的设备名称（模糊查询），`registrationID` 为设备标识 Registration ID（精确查询）；
//     `page` 和 `pageSize` 不传（为 0）则默认返回所有数据（默认为 1 和 200），二者要么都传（都不为 0），要么两者都不传（都为 0）；
//     `deviceName` 和 `registrationID` 只能同时存在其一，如果两者同时存在（都不为空），只会传 `deviceName`。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%8E%B7%E5%8F%96%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87%E5%88%97%E8%A1%A8
func (d *apiv3) ListTestDevices(ctx context.Context, page, pageSize int, deviceName, registrationID string) (*TestDevicesListResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	var query string
	if page > 0 {
		if pageSize <= 0 {
			pageSize = 200
		}
		query += fmt.Sprintf("page=%d&page_size=%d", page, pageSize)
	} else {
		if pageSize > 0 {
			// page = 1
			query += fmt.Sprintf("page=1&page_size=%d", pageSize)
		}
	}
	if deviceName != "" && registrationID != "" {
		// registrationID = ""
		query += "&device_name=" + url.QueryEscape(deviceName)
	} else {
		if deviceName != "" {
			query += "&device_name=" + url.QueryEscape(deviceName)
		}
		if registrationID != "" {
			query += "&registration_id=" + url.QueryEscape(registrationID)
		}
	}

	if query != "" {
		query = "?" + query
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  d.proto,
		URL:    d.host + "/v3/test/model/list" + query,
		Auth:   d.auth,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TestDevicesListResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TestDevicesListResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError     `json:"error,omitempty"`
	Total         int                `json:"total,omitempty"`     // 总记录数
	Page          int                `json:"page,omitempty"`      // 当前页码
	PageSize      int                `json:"page_size,omitempty"` // 每页记录条数
	Detail        []TestDeviceDetail `json:"detail,omitempty"`    // 测试设备详情列表
}

// 测试设备详情
type TestDeviceDetail struct {
	DeviceName       string                `json:"device_name,omitempty"`       // 设备名称
	DeviceModel      string                `json:"device_model,omitempty"`      // 设备型号
	RegistrationID   string                `json:"registration_id,omitempty"`   // 极光生成的设备唯一标识
	RegistrationTime jiguang.LocalDateTime `json:"registration_time,omitempty"` // 设备注册时间，格式：2006-01-02 15:04:05
	CreateTime       jiguang.LocalDateTime `json:"create_time,omitempty"`       // 信息添加时间，格式：2006-01-02 15:04:05
}

func (rs *TestDevicesListResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
