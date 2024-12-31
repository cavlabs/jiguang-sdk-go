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

// 设置设备的标签、别名与手机号码
//  - 功能说明：更新当前设备的指定属性，当前支持 tags、alias 与 mobile；使用短信业务，请结合服务端【SMS_MESSAGE】字段：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) SetDevice(ctx context.Context, registrationID string, param *DeviceSetParam) (*DeviceSetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if registrationID == "" {
		return nil, errors.New("`registrationID` cannot be empty")
	}
	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/devices/" + registrationID,
		Auth:   d.auth,
		Body:   param,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &DeviceSetResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type DeviceSetParam struct {
	// 设备标签，可使用 TagsForDeviceSetParam 或 空字符串 设置；
	//  - 支持 add, remove 或者 空字符串 - 当设置为 空字符串 的时候，表示清空所有的 tags；add/remove 下是增加或删除指定的 tag；
	//  - 一次 add/remove tag 的上限均为 100 个，且总长度均不能超过 1000 字节；
	//  - 可以多次调用 API 设置，一个设备（Registration ID）能设置的 tag 上限为 1000 个，应用 tag 总数没有限制。
	Tags interface{} `json:"tags,omitempty"`
	// 更新设备的别名属性；当别名为 空字符串 时，删除指定设备的别名；
	//  - 注意：极光于 2020/03/10 对「别名设置」的上限进行限制，最多允许绑定 10 个设备，超过将报错 7015。
	Alias *string `json:"alias,omitempty"`
	// 设备关联的手机号码；当 mobile 为 空字符串 时，表示清空设备关联的手机号码。
	Mobile *string `json:"mobile,omitempty"`
}

// 设置设备标签，提供给 DeviceSetParam.Tags 字段使用。
type TagsForDeviceSetParam struct {
	Add    []string `json:"add,omitempty"`    // 增加的标签列表
	Remove []string `json:"remove,omitempty"` // 删除的标签列表
}

type DeviceSetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *DeviceSetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
