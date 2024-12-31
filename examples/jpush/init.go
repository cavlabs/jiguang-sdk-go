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

package jpush

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/admin"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/device"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/file"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/gpush"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/greport"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/image"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/report"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/schedule"
	"github.com/calvinit/jiguang-sdk-go/examples/adapter"
)

var (
	_client_ = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略 SSL 证书验证
			},
			ForceAttemptHTTP2: true,
		},
	}
	// _client_ = adapter.NewRestyClient()

	// _logger_ = adapter.NewLogrusLogger()
	// _logger_ = adapter.NewZapLogger()
	_logger_ = adapter.NewZeroLogger()

	// ###################### ↓↓↓ 此为演示数据，请替换成真实数据 ↓↓↓ ######################

	// _devKey_ = os.Getenv("JPUSH_DEV_KEY")
	_devKey_ = "c2dc75e97486529205528b23"
	// _devSecret_ = os.Getenv("JPUSH_DEV_SECRET")
	_devSecret_ = "d219e2f001df2fe4f08b6754"

	// _appKey_ = os.Getenv("JPUSH_APP_KEY")
	_appKey_ = "de5b45e0de18292281cbb3cf"
	// _masterSecret_ = os.Getenv("JPUSH_MASTER_SECRET")
	_masterSecret_ = "37e2077739009a532e71cce4"

	// _groupKey_ = os.Getenv("JPUSH_GROUP_KEY")
	_groupKey_ = "2ed1465b94aab3f03f6778e0"
	// _groupMasterSecret_ = os.Getenv("JPUSH_GROUP_MASTER_SECRET")
	_groupMasterSecret_ = "5b69dce5f9dc0dabe70aa33a"

	// ###################### ↑↑↑ 此为演示数据，请替换成真实数据 ↑↑↑ ######################

	// 应用管理 - Admin API v1
	adminAPIv1, _ = admin.NewAPIv1Builder().
			SetClient(_client_).                   // 【可选】配置，如果不配置，则使用默认的 api.DefaultClient。
			SetDevKey(_devKey_).                   // 【必填】配置。
			SetDevSecret(_devSecret_).             // 【必填】配置。
			SetLogger(_logger_).                   // 【可选】配置，如果不配置，则使用默认的 api.DefaultLogger。
			SetHttpLogLevel(api.HttpLogLevelFull). // 【可选】配置，如果不配置，则使用 api.HttpLogLevelBasic。
			Build()

	// 设备/标签/别名 - Device API v3
	deviceAPIv3, _ = device.NewAPIv3Builder().
			SetClient(_client_).
			SetAppKey(_appKey_).             // 【必填】配置。
			SetMasterSecret(_masterSecret_). // 【必填】配置。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 推送 - Push API v3
	pushAPIv3, _ = push.NewAPIv3Builder().
			SetClient(_client_).
			SetAppKey(_appKey_).             // 【必填】配置。
			SetMasterSecret(_masterSecret_). // 【必填】配置。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 分组推送 - Group Push API v3
	gpushAPIv3, _ = gpush.NewAPIv3Builder().
			SetClient(_client_).
			SetGroupKey(_groupKey_).                   // 【必填】配置。
			SetGroupMasterSecret(_groupMasterSecret_). // 【必填】配置。
			SetDevKey(_devKey_).                       // 【可选】配置，但当需要同时使用 “上传文件” 等相关「文件管理」的 API 接口时，请务必同时设置 `devKey`。
			SetDevSecret(_devSecret_).                 // 【可选】配置，但当需要同时使用 “上传文件” 等相关「文件管理」的 API 接口时，请务必同时设置 `devSecret`。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 定时任务 - Schedule API v3
	scheduleAPIv3, _ = schedule.NewAPIv3Builder().
				SetClient(_client_).
				SetAppKey(_appKey_).             // 【必填】配置。
				SetMasterSecret(_masterSecret_). // 【必填】配置。
				SetLogger(_logger_).
				SetHttpLogLevel(api.HttpLogLevelFull).
				Build()

	// 文件管理 - File API v3
	fileAPIv3, _ = file.NewAPIv3Builder().
			SetClient(_client_).
			SetAuthKey(_appKey_).          // 【必填】配置，可使用 `appKey` 或者 `devKey`。
			SetAuthSecret(_masterSecret_). // 【必填】配置，可使用 `masterSecret` 或者 `devSecret`。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 图片管理 - Image API v3
	imageAPIv3, _ = image.NewAPIv3Builder().
			SetClient(_client_).
			SetAppKey(_appKey_).             // 【必填】配置。
			SetMasterSecret(_masterSecret_). // 【必填】配置。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 推送统计 - Report API v3
	reportAPIv3, _ = report.NewAPIv3Builder().
			SetClient(_client_).
			SetAppKey(_appKey_).             // 【必填】配置。
			SetMasterSecret(_masterSecret_). // 【必填】配置。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()

	// 分组推送统计 - Group Report API v3
	greportAPIv3, _ = greport.NewAPIv3Builder().
			SetClient(_client_).
			SetGroupKey(_groupKey_).                   // 【必填】配置。
			SetGroupMasterSecret(_groupMasterSecret_). // 【必填】配置。
			SetLogger(_logger_).
			SetHttpLogLevel(api.HttpLogLevelFull).
			Build()
)
