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
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

// 清空设备的标签
//  - 功能说明：清空当前设备的 tags 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceTags(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, true, false, false)
}

// 清空设备的别名
//  - 功能说明：清空当前设备的 alias 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceAlias(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, false, true, false)
}

// 清空设备的手机号码
//  - 功能说明：清空当前设备的 mobile 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, false, false, true)
}

// 清空设备的标签与别名
//  - 功能说明：清空当前设备的 tags 与 alias 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceTagsAndAlias(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, true, true, false)
}

// 清空设备的标签与手机号码
//  - 功能说明：清空当前设备的 tags 和 mobile 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceTagsAndMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, true, false, true)
}

// 清空设备的别名与手机号码
//  - 功能说明：清空当前设备的 alias 与 mobile 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceAliasAndMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, false, true, true)
}

// 清空设备的标签、别名与手机号码
//  - 功能说明：清空当前设备的 tags、alias 与 mobile 属性。
//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
func (d *apiv3) ClearDeviceAll(ctx context.Context, registrationID string) (*DeviceClearResult, error) {
	return d.clearDevice(ctx, registrationID, true, true, true)
}

// clearDevice 清空设备的 tags、alias 或 mobile 属性，该接口与 SetDevice 是同一个接口。
func (d *apiv3) clearDevice(ctx context.Context, registrationID string, clearTags, clearAlias, clearMobile bool) (*DeviceClearResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if registrationID == "" {
		return nil, errors.New("`registrationID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/devices/" + registrationID,
		Auth:   d.auth,
		Body:   newDeviceClearParam(clearTags, clearAlias, clearMobile),
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &DeviceClearResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type deviceClearParam struct {
	Tags   *string `json:"tags,omitempty"`   // 如果要清空 tags，传入空字符串
	Alias  *string `json:"alias,omitempty"`  // 如果要清空 alias，传入空字符串
	Mobile *string `json:"mobile,omitempty"` // 如果要清空 mobile，传入空字符串
}

// newDeviceClearParam 创建一个新的 deviceClearParam 实例。
//  - clearTags: 是否清空 tags；
//  - clearAlias: 是否清空 alias；
//  - clearMobile: 是否清空 mobile。
func newDeviceClearParam(clearTags, clearAlias, clearMobile bool) *deviceClearParam {
	param := deviceClearParam{}
	if clearTags {
		param.Tags = jiguang.String("")
	}
	if clearAlias {
		param.Alias = jiguang.String("")
	}
	if clearMobile {
		param.Mobile = jiguang.String("")
	}
	return &param
}

type DeviceClearResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *DeviceClearResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
