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

package file

import "context"

// File API v3【极光推送 > REST API > 文件管理 API】
//  - 功能说明：文件推送功能支持（调用文件推送接口推送前，必须先通过此模块接口上传文件，得到文件唯一标识（fileID）后方能推送），包括上传文件、查询文件、删除文件等相关 API。
//  - 极光文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file
type APIv3 interface {
	// 上传文件 (Alias)
	//  - 功能说明：可以将要推送的 alias 值先写入本地文件，然后将本地文件上传到极光服务器，后续就可以直接指定文件推送了。
	//	- 调用地址：POST `/v3/files/alias`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6
	UploadFileForAlias(ctx context.Context, param *FileUploadParam) (*FileUploadResult, error)

	// 上传文件 (Registration ID)
	//  - 功能说明：可以将要推送的 Registration ID 值先写入本地文件，然后将本地文件上传到极光服务器，后续就可以直接指定文件推送了。
	//	- 调用地址：POST `/v3/files/registration_id`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E4%B8%8A%E4%BC%A0%E6%96%87%E4%BB%B6
	UploadFileForRegistrationID(ctx context.Context, param *FileUploadParam) (*FileUploadResult, error)

	// 查询有效文件列表
	//  - 功能说明：获取当前保存在极光服务器的有效文件列表。
	//	- 调用地址：GET `/v3/files`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E6%9F%A5%E8%AF%A2%E6%9C%89%E6%95%88%E6%96%87%E4%BB%B6%E5%88%97%E8%A1%A8
	GetFiles(ctx context.Context) (*FilesGetResult, error)

	// 查询指定文件详情
	//  - 功能说明：查询保存在极光服务器的，指定文件的详细信息。
	//	- 调用地址：GET `/v3/files/{fileID}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E6%9F%A5%E8%AF%A2%E6%8C%87%E5%AE%9A%E6%96%87%E4%BB%B6%E8%AF%A6%E6%83%85
	GetFile(ctx context.Context, fileID string) (*FileGetResult, error)

	// 删除指定文件
	//  - 功能说明：删除保存在极光服务器的指定文件。
	//	- 调用地址：DELETE `/v3/files/{fileID}`
	//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E5%88%A0%E9%99%A4%E6%96%87%E4%BB%B6
	//
	// 注意事项：
	//  - fileID 不存在当成功处理。
	//  - 对于即时推送，建议创建推送任务 5 分钟后再执行文件删除操作，否则推送任务可能会失败；
	//  - 对于文件定时推送，创建定时任务成功后，若任务被执行前文件被删除，则任务执行时推送动作将会失败。
	DeleteFile(ctx context.Context, fileID string) (*FileDeleteResult, error)
}
