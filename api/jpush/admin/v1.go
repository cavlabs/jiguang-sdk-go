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

package admin

import "context"

// # Admin API v1
//
// 极光推送 > REST API > 应用管理 API
//   - 功能说明：提供给开发者操作创建或删除 APP，上传证书等管理功能。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1
type APIv1 interface {
	// # 创建极光 APP
	//  - 功能说明：在开发者账号下创建一个 APP。
	//  - 调用地址：POST `/v1/app`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#%E5%88%9B%E5%BB%BA%E6%9E%81%E5%85%89-app
	CreateApp(ctx context.Context, param *AppCreateParam) (*AppCreateResult, error)

	// # 删除极光 APP
	//  - 功能说明：删除开发者账号下的指定 APP。
	//  - 调用地址：POST `/v1/app/{appKey}/delete`，`appKey` 为 APP 的唯一标识。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#app-%E5%88%A0%E9%99%A4
	DeleteApp(ctx context.Context, appKey string) (*AppDeleteResult, error)

	// # 极光 APP 证书上传
	//  - 功能说明：上传开发或生产证书到对应的极光 APP。
	//  - 调用地址：POST `/v1/app/{appKey}/certificate`，`appKey` 为 APP 的唯一标识。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1#%E8%AF%81%E4%B9%A6%E4%B8%8A%E4%BC%A0
	UploadCertificate(ctx context.Context, appKey string, param *CertificateUploadParam) (*CertificateUploadResult, error)
}
