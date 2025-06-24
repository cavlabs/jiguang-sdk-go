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
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 获取用户在线状态（VIP）
//   - 功能说明：查询用户是否在线。
//   - 调用地址：POST `/v3/devices/status`，`registrationIDs` 为必填参数，需要获取在线状态的设备标识 Registration ID 集合，最多支持 1000 个。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81%EF%BC%88vip%EF%BC%89
func (d *apiv3) GetDeviceStatus(ctx context.Context, registrationIDs []string) (*DeviceStatusGetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	l := len(registrationIDs)
	if l == 0 {
		return nil, errors.New("`registrationIDs` cannot be empty")
	}
	if l > 1000 {
		return nil, errors.New("`registrationIDs` cannot be more than 1000")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/devices/status",
		Auth:   d.auth,
		Body:   &deviceStatusGetParam{RegistrationIDs: registrationIDs},
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &DeviceStatusGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type deviceStatusGetParam struct {
	RegistrationIDs []string `json:"registration_ids"` // 需要在线状态的用户 Registration ID 集合，最多支持 1000 个
}

type DeviceStatusGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	// 用户在线状态结果
	//  - key 为 Registration ID；
	//  - 对于无效的或者不属于请求该 API 的认证 appKey 的 Registration ID，它的返回的结果为空。
	Result map[string]DeviceStatusResult `json:"-"`
}

// 用户在线状态结果
type DeviceStatusResult struct {
	// 10 分钟之内是否在线
	Online bool `json:"online"`
	// 最后一次在线时间
	//  - 10 分钟之内在线时，不返回该字段；
	//  - 10 分钟之内不在线，且该字段不返回时，则表示最后一次在线时间是在两天之前。
	LastOnlineTime *jiguang.LocalDateTime `json:"last_online_time,omitempty"`
}

func (rs *DeviceStatusGetResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
	}

	if len(aux) > 0 {
		rs.Result = make(map[string]DeviceStatusResult)
		for k, v := range aux {
			// not null
			if len(v) != 4 || v[0] != 'n' || v[1] != 'u' || v[2] != 'l' || v[3] != 'l' {
				var result DeviceStatusResult
				if err := json.Unmarshal(v, &result); err != nil {
					return err
				}
				rs.Result[k] = result
			}
		}
	}

	return nil
}

func (rs DeviceStatusGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	} else if len(rs.Result) > 0 {
		data := make(map[string]DeviceStatusResult, len(rs.Result))
		for k, v := range rs.Result {
			data[k] = v
		}
		return json.Marshal(data)
	}
	return []byte("null"), nil
}

func (rs *DeviceStatusGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
