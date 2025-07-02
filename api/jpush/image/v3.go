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

import "context"

// # Image API v3
//
// 【极光推送 > REST API > 图片 API】
//   - 功能说明：需要结合 Push Android SDK v3.9.0 及其以上版本配套使用。开发者提交图片或者网络资源 URL 至极光服务器，极光会根据开发者需求对接适配各个厂商，同时开发者将从极光得到对应资源的 MediaID，该 MediaID 可于 Push API v3 中使用，达到统一推送大图片、大图标和小图标的需求。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// 调用限制：
//   - 上传图片大小不超过 2M；
//   - 图片默认最多保存 30 天（从 2023.08.31 开始实施，请开发者关注）；
//   - API 请求频率限制与 Push API v3 接口共享。
//
// 各通道支持及要求情况：
//   - 该接口可适配的厂商通道有：极光、OPPO、华为、荣耀以及 FCM；
//   - 小米从 2023.08 开始，官方在新设备/系统已经不再支持推送时动态设置小图标、右侧图标、大图片功能，对于历史设备和应用也在逐步覆盖，等于不再支持推送时动态设置小图标、右侧图标、大图片功能；
//   - 目前该接口暂不提供存储文件服务，故请确保提交的图片资源可被访问且符合要求；
//   - 不同通道对不同类型的图片支持情况不同，具体参考 [各通道支持及要求情况]。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image
// [各通道支持及要求情况]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E5%90%84%E9%80%9A%E9%81%93%E6%94%AF%E6%8C%81%E5%8F%8A%E8%A6%81%E6%B1%82%E6%83%85%E5%86%B5
type APIv3 interface {
	// # 新增图片（URL 方式）
	//  - 功能说明：通过指定网络图片资源的 URL 形式来新增一个适配。
	//	- 调用地址：POST `/v3/images/byurls`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%96%B0%E5%A2%9E%E5%9B%BE%E7%89%87%EF%BC%88url%E6%96%B9%E5%BC%8F%EF%BC%89
	AddImageByUrl(ctx context.Context, param *AddByUrlParam) (*AddByUrlResult, error)

	// # 更新图片（URL 方式）
	//  - 功能说明：通过指定网络图片资源的 URL 形式来修改或更新适配结果。
	//	- 调用地址：PUT `/v3/images/byurls/{mediaID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%9B%B4%E6%96%B0%E5%9B%BE%E7%89%87%EF%BC%88url%E6%96%B9%E5%BC%8F%EF%BC%89
	UpdateImageByUrl(ctx context.Context, mediaID string, param *UpdateByUrlParam) (*UpdateByUrlResult, error)

	// # 新增图片（文件方式）
	//  - 功能说明：通过上传图片文件形式来新增一个适配，该接口目前仅支持小米和 OPPO。不过从 2023.08 开始，小米官方在新设备/系统已经不再支持推送时动态设置小图标、右侧图标、大图片功能，对于历史设备和应用也在逐步覆盖，等于不再支持推送时动态设置小图标、右侧图标、大图片功能；
	//	- 调用地址：POST `/v3/images/byfiles`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%96%B0%E5%A2%9E%E5%9B%BE%E7%89%87%EF%BC%88%E6%96%87%E4%BB%B6%E6%96%B9%E5%BC%8F%EF%BC%89
	AddImageByFile(ctx context.Context, param *AddByFileParam) (*AddByFileResult, error)

	// # 更新图片（文件方式）
	//  - 功能说明：通过上传图片文件形式来修改或更新适配结果，该接口目前仅支持 OPPO。
	//	- 调用地址：PUT `/v3/images/byfiles/{mediaID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image#%E6%9B%B4%E6%96%B0%E5%9B%BE%E7%89%87%EF%BC%88%E6%96%87%E4%BB%B6%E6%96%B9%E5%BC%8F%EF%BC%89
	UpdateImageByFile(ctx context.Context, mediaID string, param *UpdateByFileParam) (*UpdateByFileResult, error)
}
