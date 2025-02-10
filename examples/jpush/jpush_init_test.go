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
	"os"
	"testing"
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
	adminAPIv1    admin.APIv1    // 应用管理 - Admin API v1
	deviceAPIv3   device.APIv3   // 设备/标签/别名 - Device API v3
	pushAPIv3     push.APIv3     // 推送 - Push API v3
	gpushAPIv3    gpush.APIv3    // 分组推送 - Group Push API v3
	scheduleAPIv3 schedule.APIv3 // 定时任务 - Schedule API v3
	fileAPIv3     file.APIv3     // 文件管理 - File API v3
	imageAPIv3    image.APIv3    // 图片管理 - Image API v3
	reportAPIv3   report.APIv3   // 推送统计 - Report API v3
	greportAPIv3  greport.APIv3  // 分组推送统计 - Group Report API v3
)

func TestMain(m *testing.M) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略 SSL 证书验证
			},
			ForceAttemptHTTP2: true,
		},
	}
	// client := adapter.NewRestyClient()

	// logger := adapter.NewLogrusLogger("[JPush] ")
	// logger := adapter.NewZapLogger("[JPush]")
	logger := adapter.NewZeroLogger("[JPush] ")

	// ###################### ↓↓↓ 此为演示数据，请替换成真实数据 ↓↓↓ ######################

	// devKey := os.Getenv("JPUSH_DEV_KEY")
	devKey := "c2dc75e97486529205528b23"
	// devSecret := os.Getenv("JPUSH_DEV_SECRET")
	devSecret := "d219e2f001df2fe4f08b6754"

	// appKey := os.Getenv("JPUSH_APP_KEY")
	appKey := "de5b45e0de18292281cbb3cf"
	// masterSecret := os.Getenv("JPUSH_MASTER_SECRET")
	masterSecret := "37e2077739009a532e71cce4"

	// groupKey := os.Getenv("JPUSH_GROUP_KEY")
	groupKey := "2ed1465b94aab3f03f6778e0"
	// groupMasterSecret := os.Getenv("JPUSH_GROUP_MASTER_SECRET")
	groupMasterSecret := "5b69dce5f9dc0dabe70aa33a"

	// ###################### ↑↑↑ 此为演示数据，请替换成真实数据 ↑↑↑ ######################

	// 应用管理 - Admin API v1
	adminAPIv1, _ = admin.NewAPIv1Builder().
		SetClient(client).                     // 【可选】配置，如果不配置，则使用默认的 api.DefaultClient。
		SetDevKey(devKey).                     // 【必填】配置。
		SetDevSecret(devSecret).               // 【必填】配置。
		SetLogger(logger).                     // 【可选】配置，如果不配置，则使用默认的 api.DefaultJPushLogger。
		SetHttpLogLevel(api.HttpLogLevelFull). // 【可选】配置，如果不配置，则使用 api.HttpLogLevelBasic。
		Build()

	// 设备/标签/别名 - Device API v3
	deviceAPIv3, _ = device.NewAPIv3Builder().
		SetClient(client).
		SetAppKey(appKey).             // 【必填】配置。
		SetMasterSecret(masterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 推送 - Push API v3
	pushAPIv3, _ = push.NewAPIv3Builder().
		SetClient(client).
		SetAppKey(appKey).             // 【必填】配置。
		SetMasterSecret(masterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 分组推送 - Group Push API v3
	gpushAPIv3, _ = gpush.NewAPIv3Builder().
		SetClient(client).
		SetGroupKey(groupKey).                   // 【必填】配置。
		SetGroupMasterSecret(groupMasterSecret). // 【必填】配置。
		SetDevKey(devKey).                       // 【可选】配置，但当需要同时使用 “上传文件” 等相关「文件管理」的 API 接口时，请务必同时设置 `devKey`。
		SetDevSecret(devSecret).                 // 【可选】配置，但当需要同时使用 “上传文件” 等相关「文件管理」的 API 接口时，请务必同时设置 `devSecret`。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 定时任务 - Schedule API v3
	scheduleAPIv3, _ = schedule.NewAPIv3Builder().
		SetClient(client).
		SetAppKey(appKey).             // 【必填】配置。
		SetMasterSecret(masterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 文件管理 - File API v3
	fileAPIv3, _ = file.NewAPIv3Builder().
		SetClient(client).
		SetAuthKey(appKey).          // 【必填】配置，可使用 `appKey` 或者 `devKey`。
		SetAuthSecret(masterSecret). // 【必填】配置，可使用 `masterSecret` 或者 `devSecret`。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 图片管理 - Image API v3
	imageAPIv3, _ = image.NewAPIv3Builder().
		SetClient(client).
		SetAppKey(appKey).             // 【必填】配置。
		SetMasterSecret(masterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 推送统计 - Report API v3
	reportAPIv3, _ = report.NewAPIv3Builder().
		SetClient(client).
		SetAppKey(appKey).             // 【必填】配置。
		SetMasterSecret(masterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	// 分组推送统计 - Group Report API v3
	greportAPIv3, _ = greport.NewAPIv3Builder().
		SetClient(client).
		SetGroupKey(groupKey).                   // 【必填】配置。
		SetGroupMasterSecret(groupMasterSecret). // 【必填】配置。
		SetLogger(logger).
		SetHttpLogLevel(api.HttpLogLevelFull).
		Build()

	code := m.Run()

	os.Exit(code)
}
