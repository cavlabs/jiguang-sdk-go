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

package device

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 查询设备与标签的绑定关系
//  - 功能说明：查询某个设备是否在 tag 下。
//	- 调用地址：GET `/v3/tags/{tag}/registration_ids/{registrationID}`，`tag` 为指定的标签值；`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E8%AE%BE%E5%A4%87%E4%B8%8E%E6%A0%87%E7%AD%BE%E7%9A%84%E7%BB%91%E5%AE%9A%E5%85%B3%E7%B3%BB
func (d *apiv3) GetTag(ctx context.Context, tag string, registrationID string) (*TagGetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if tag == "" {
		return nil, errors.New("`tag` cannot be empty")
	}
	if registrationID == "" {
		return nil, errors.New("`registrationID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  d.proto,
		URL:    d.host + "/v3/tags/" + tag + "/registration_ids/" + registrationID,
		Auth:   d.auth,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TagGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type TagGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Result        *bool          `json:"result,omitempty"` // 设备是否与标签绑定
}

func (rs *TagGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
