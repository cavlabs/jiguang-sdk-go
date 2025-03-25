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

	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
)

// # Device API v3
//
// 【极光推送 > REST API > 标签别名 API】
//   - 功能说明：用于在服务器端查询、设置、更新、删除设备的 tag、alias、mobile 等信息。
//   - 详见 [docs.jiguang.cn] 文档说明。
//
// 使用时需要注意不要让服务端设置的标签又被客户端给覆盖了：
//   - 如果不是很熟悉 tag、alias 的逻辑的话，建议只使用客户端或服务端二者中的一种；
//   - 如果是两边同时使用，请确认自己的应用可以处理好标签和别名的同步。
//
// 需要了解 tag、alias 的详细信息，请参考对应客户端平台的 API 说明：[Android]、[iOS]、[HarmonyOS]。
//
// 包含了 device、tag 和 alias 三组 API，其中：
//   - device 用于查询/设置设备的各种属性，包含 tags 和 alias，手机号码 mobile；
//   - tag 用于查询/设置/删除设备的标签；
//   - alias 用于查询/设置/删除设备的别名。
//
// 需要用到的关键信息还有 Registration ID：
//   - 设备的 Registration ID 在客户端集成后获取，详情查看 [Android、iOS、HarmonyOS] 的 API 文档。
//   - 服务端未提供 API 去获取设备的 Registration ID 值，需要开发者在客户端获取到 Registration ID 后上传给开发者服务器保存。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device
// [Android]: https://docs.jiguang.cn/jpush/client/Android/android_api#%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE-api
// [iOS]: https://docs.jiguang.cn/jpush/client/iOS/ios_api#%E6%A0%87%E7%AD%BE%E4%B8%8E%E5%88%AB%E5%90%8D-api%EF%BC%88ios%EF%BC%89
// [HarmonyOS]: https://docs.jiguang.cn/jpush/client/HarmonyOS/hmos_api#%E6%A0%87%E7%AD%BE%E4%B8%8E%E5%88%AB%E5%90%8D-api
// [Android、iOS、HarmonyOS]: https://docs.jiguang.cn/jpush/client/HarmonyOS/hmos_api#%E8%8E%B7%E5%8F%96rid%EF%BC%88getregistrationid%EF%BC%89
type APIv3 interface {
	// # 查询设备的标签、别名与手机号码
	//  - 功能说明：获取当前设备的所有属性，包含 tags、alias 与 mobile。
	//	- 调用地址：GET `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	GetDevice(ctx context.Context, registrationID string) (*DeviceGetResult, error)

	// # 设置设备的标签、别名与手机号码
	//  - 功能说明：更新当前设备的指定属性，当前支持 tags、alias 与 mobile；使用短信业务，请结合服务端 [SMS_MESSAGE] 字段。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	// [SMS_MESSAGE]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push#sms_message%EF%BC%9A%E7%9F%AD%E4%BF%A1
	SetDevice(ctx context.Context, registrationID string, param *DeviceSetParam) (*DeviceSetResult, error)

	// # 清空设备的标签
	//  - 功能说明：清空当前设备的 tags 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceTags(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的别名
	//  - 功能说明：清空当前设备的 alias 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceAlias(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的手机号码
	//  - 功能说明：清空当前设备的 mobile 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的标签与别名
	//  - 功能说明：清空当前设备的 tags 与 alias 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceTagsAndAlias(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的标签与手机号码
	//  - 功能说明：清空当前设备的 tags 和 mobile 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceTagsAndMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的别名与手机号码
	//  - 功能说明：清空当前设备的 alias 与 mobile 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceAliasAndMobile(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 清空设备的标签、别名与手机号码
	//  - 功能说明：清空当前设备的 tags、alias 与 mobile 属性。
	//	- 调用地址：POST `/v3/devices/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%AE%BE%E7%BD%AE%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D%E4%B8%8E%E6%A0%87%E7%AD%BE
	ClearDeviceAll(ctx context.Context, registrationID string) (*DeviceClearResult, error)

	// # 获取用户在线状态（VIP）
	//  - 功能说明：查询用户是否在线。
	//	- 调用地址：POST `/v3/devices/status`，`registrationIDs` 为必填参数，需要获取在线状态的设备标识 Registration ID 集合，最多支持 1000 个。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81%EF%BC%88vip%EF%BC%89
	GetDeviceStatus(ctx context.Context, registrationIDs []string) (*DeviceStatusGetResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 新增测试设备（VIP）
	//  - 功能说明：新增一个测试设备，确保测试模式下的每次推送仅触达测试用户。
	//	- 调用地址：POST `/v3/test/model/add`
	//  - 接口文档：[docs.jiguang.cn]
	// 详细功能逻辑可参考文档：[测试模式]。
	//
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%96%B0%E5%A2%9E%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87
	// [测试模式]: https://docs.jiguang.cn/jpush/console/push_manage/testmode
	AddTestDevice(ctx context.Context, param *TestDeviceAddParam) (*TestDeviceAddResult, error)

	// # 修改测试设备（VIP）
	//  - 功能说明：修改一个指定的测试设备。
	//	- 调用地址：PUT `/v3/test/model/update`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E4%BF%AE%E6%94%B9%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87
	UpdateTestDevice(ctx context.Context, param *TestDeviceUpdateParam) (*TestDeviceUpdateResult, error)

	// # 删除测试设备（VIP）
	//  - 功能说明：删除一个指定的测试设备。
	//	- 调用地址：DELETE `/v3/test/model/delete/{registrationID}`，`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87
	DeleteTestDevice(ctx context.Context, registrationID string) (*TestDeviceDeleteResult, error)

	// # 获取测试设备列表（VIP）
	//  - 功能说明：分页获取测试设备列表。
	//	- 调用地址：GET `/v3/test/model/list?page={page}&page_size={pageSize}&device_name={deviceName}&registration_id={registrationID}`；
	//  `page` 为查询页码，`pageSize` 为每页记录条数，`deviceName` 为开发者自定义的设备名称（模糊查询），`registrationID` 为设备标识 Registration ID（精确查询）；
	//	`page` 和 `pageSize` 不传（为 0）则默认返回所有数据（默认为 1 和 200），二者要么都传（都不为 0），要么两者都不传（都为 0）；
	//	`deviceName` 和 `registrationID` 只能同时存在其一，如果两者同时存在（都不为空），只会传 `deviceName`。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E8%8E%B7%E5%8F%96%E6%B5%8B%E8%AF%95%E8%AE%BE%E5%A4%87%E5%88%97%E8%A1%A8
	ListTestDevices(ctx context.Context, page, pageSize int, deviceName, registrationID string) (*TestDevicesListResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 查询标签列表
	//  - 功能说明：获取当前应用的所有标签列表，每个平台最多返回 100 个。
	//	- 调用地址：GET `/v3/tags`
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E6%A0%87%E7%AD%BE%E5%88%97%E8%A1%A8
	GetTags(ctx context.Context) (*TagsGetResult, error)

	// # 查询设备与标签的绑定关系
	//  - 功能说明：查询某个设备是否在 tag 下。
	//	- 调用地址：GET `/v3/tags/{tag}/registration_ids/{registrationID}`，`tag` 为指定的标签值；`registrationID` 为设备标识 Registration ID。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E8%AE%BE%E5%A4%87%E4%B8%8E%E6%A0%87%E7%AD%BE%E7%9A%84%E7%BB%91%E5%AE%9A%E5%85%B3%E7%B3%BB
	GetTag(ctx context.Context, tag string, registrationID string) (*TagGetResult, error)

	// # 更新标签
	//  - 功能说明：为一个标签添加或者删除设备。
	//	- 调用地址：POST `/v3/tags/{tag}`，`tag` 为指定的标签值；`adds`/`removes` 为增加或删除的设备标识 Registration ID 集合，最多各支持 1000 个，不能同时为空。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9B%B4%E6%96%B0%E6%A0%87%E7%AD%BE
	SetTag(ctx context.Context, tag string, adds, removes []string) (*TagSetResult, error)

	// # 删除标签
	//  - 功能说明：删除一个标签，以及标签与设备之间的关联关系。
	//	- 调用地址：DELETE `/v3/tags/{tag}`，`tag` 为指定的标签值；`plats` 为可选参数，不填则默认为所有平台。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E6%A0%87%E7%AD%BE
	DeleteTag(ctx context.Context, tag string, plats ...platform.Platform) (*TagDeleteResult, error)

	// -----------------------------------------------------------------------------------------------------------------

	// # 查询别名
	//  - 功能说明：获取指定 alias 下的设备，正常情况下最多输出 10 个，超过 10 个默认输出 10 个。
	//	- 调用地址：GET `/v3/aliases/{alias}`，`alias` 为指定的别名值；`plats` 为可选参数，不填则默认为所有平台。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E5%88%AB%E5%90%8D
	GetAlias(ctx context.Context, alias string, plats ...platform.Platform) (*AliasGetResult, error)

	// # 删除别名
	//  - 功能说明：删除一个别名，以及该别名与设备的绑定关系。
	//	- 调用地址：DELETE `/v3/aliases/{alias}`，`alias` 为指定的别名值；`plats` 为可选参数，不填则默认为所有平台。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E5%88%AB%E5%90%8D
	DeleteAlias(ctx context.Context, alias string, plats ...platform.Platform) (*AliasDeleteResult, error)

	// # 删除设备的别名
	//  - 功能说明：批量解绑设备与别名之间的关系。
	//	- 调用地址：POST `/v3/aliases/{alias}`，`alias` 为指定的别名值；`registrationIDs` 为必填参数，需要和该 `alias` 解除绑定的设备标识 Registration ID 值集合，最多支持 1000 个。
	//  - 接口文档：[docs.jiguang.cn]
	// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D
	DeleteAliases(ctx context.Context, alias string, registrationIDs []string) (*AliasesDeleteResult, error)
}
