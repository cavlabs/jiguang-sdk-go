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

package image

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 更新图片（URL 方式）
//   - 功能说明：通过指定网络图片资源的 URL 形式来修改或更新适配结果。
//   - 调用地址：PUT `/v3/images/byurls/{mediaID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%9B%B4%E6%96%B0%E5%9B%BE%E7%89%87%EF%BC%88url%E6%96%B9%E5%BC%8F%EF%BC%89
func (i *apiv3) UpdateImageByUrl(ctx context.Context, mediaID string, param *UpdateByUrlParam) (*UpdateByUrlResult, error) {
	if i == nil {
		return nil, api.ErrNilJPushImageAPIv3
	}

	if mediaID == "" {
		return nil, errors.New("`mediaID` cannot be empty")
	}
	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	req := &api.Request{
		Method: http.MethodPut,
		Proto:  i.proto,
		URL:    i.host + "/v3/images/byurls/" + mediaID,
		Auth:   i.auth,
		Body:   param,
	}
	resp, err := i.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &UpdateByUrlResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 更新图片（URL 方式）请求参数正文。
//   - 各通道的字段值不能同时为空；
//   - 若未指定更新配置的厂商地址，则该厂商的配置结果保留更新前的值而不会被置空；
//   - 若指定更新配置的厂商不支持对应类型，则该厂商的配置结果将保留更新前的值而不会被覆盖。
type UpdateByUrlParam struct {
	// 【可选】配置极光通道的图片地址。
	JiguangImageUrl string `json:"jiguang_image_url,omitempty"`
	// 【可选】配置小米通道的图片地址。
	//  - 如果您的应用没有通过极光平台开通小米厂商通道，则不对该字段做处理；
	//  - 本接口将会对该地址的图片资源自动进行尺寸的调整和压缩，以适配小米对该图片的要求，其要求具体见参考；
	//  - 若提交的图片资源尺寸不符，则进行等比例缩放且以留白的形式进行调整，而不会对图片进行剪裁，若适配调整失败，将返回错误；
	//  - 建议提交与要求相符的图片资源，小米厂商对该图片要求请参考 [小米-图片上传]；
	//  - 小米从 2023.08 开始，官方在新设备/系统已经不再支持推送时动态设置小图标、右侧图标、大图片功能，对于历史设备和应用也在逐步覆盖，等于不再支持推送时动态设置小图标、右侧图标、大图片功能。
	// [小米-图片上传]: https://dev.mi.com/console/doc/detail?pId=1278#4_4_2
	XiaomiImageUrl string `json:"xiaomi_image_url,omitempty"`
	// 【可选】配置适配 OPPO 通道的图片地址。
	//  - 如果您的应用没有通过极光平台开通 OPPO 厂商通道，则不对该字段做处理；
	//  - 本接口将会对该地址的图片资源自动进行尺寸的调整和压缩，以适配 OPPO 对该图片的要求，其要求具体见参考；
	//  - 若提交的图片资源尺寸不符，则进行等比例缩放且以留白的形式进行调整，而不会对图片进行剪裁，若适配调整失败，将返回错误；
	//  - 建议提交与要求相符的图片资源，OPPO 厂商对该图片要求请参考 [OPPO-图片上传]。
	// [OPPO-图片上传]: https://open.oppomobile.com/new/developmentDoc/info?id=11241
	OppoImageUrl string `json:"oppo_image_url,omitempty"`
	// 【可选】配置华为通道的图片地址。
	//  - 如果您的应用没有通过极光平台开通华为厂商通道，则不对该字段做处理。
	HuaweiImageUrl string `json:"huawei_image_url,omitempty"`
	// 【可选】配置荣耀通道的图片地址。
	//  - 如果您的应用没有通过极光平台开通荣耀厂商通道，则不对该字段做处理。
	HonorImageUrl string `json:"honor_image_url,omitempty"`
	// 【可选】配置 FCM 通道的图片地址。
	//  - 如果您的应用没有通过极光平台开通 FCM 厂商通道，则不对该字段做处理。
	FcmImageUrl string `json:"fcm_image_url,omitempty"`
	// 【可选】配置鸿蒙厂商的图片地址，2024.08.13 新增。
	//  - 如果您的应用没有通过极光平台开通鸿蒙厂商，则不对该字段做处理。
	HmosImageUrl string `json:"hmos_image_url,omitempty"`
}

// 更新图片（URL 方式）响应结果正文。
type UpdateByUrlResult = AddByUrlResult
