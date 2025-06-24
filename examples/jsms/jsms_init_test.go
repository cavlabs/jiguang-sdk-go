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

package jsms

import (
	"os"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jsms"
)

var smsAPIv1 jsms.APIv1 // 极光短信 - JSMS API v1

func TestMain(m *testing.M) {
	/*client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 忽略 SSL 证书验证
			},
			ForceAttemptHTTP2: true,
		},
	}*/

	// logger := adapter.NewLogrusLogger("[JSMS] ")
	// logger := adapter.NewZapLogger("[JSMS]")
	// logger := adapter.NewZeroLogger("[JSMS] ")

	// appKey := os.Getenv("JSMS_APP_KEY")
	appKey := "de5b45e0de18292281cbb3cf"
	// masterSecret := os.Getenv("JSMS_MASTER_SECRET")
	masterSecret := "37e2077739009a532e71cce4"

	// devKey := os.Getenv("JSMS_DEV_KEY")
	devKey := "7e503edcb0cb725e331b0311"
	// devSecret := os.Getenv("JSMS_DEV_SECRET")
	devSecret := "7289516381dcdf1113730f2b"

	// 启用极光短信的回调接口服务，可以自定义服务的一些配置选项：
	//  - 1. 监听地址 (addr)：默认为 ":8088"；
	//  - 2. 回调路径 (path)：默认为 "/callback"；
	//  - 3. 日志打印器 (logger)：默认与 JSMS API v1 使用相同的日志打印器；
	//  - 4. HTTP Handler (handler)：默认为使用 net/http 实现的一个简单的 Handler；
	//  - 5. 「用户回复消息」SMS_REPLY 回执数据回调处理器，为 nil 时不处理；
	//  - 6. 「短信送达状态」SMS_REPORT 回执数据回调处理器，为 nil 时不处理；
	//  - 7. 「模板审核结果」SMS_TEMPLATE 回执数据回调处理器，为 nil 时不处理；
	//  - 8. 「签名审核结果」SMS_SIGN 回执数据回调处理器，为 nil 时不处理。
	/*callbackConfigOptions := []jsms.CallbackConfigOption{
		jsms.WithCallbackAddr(":8080"),
		jsms.WithCallbackPath("/jsms/callback"),
		// jsms.WithCallbackLogger(logger), 									// 这是默认行为，可以不设置！
		// jsms.WithCallbackHttpHandler(httpHandler),
		jsms.WithCallbackReplyDataProcessor(callbackReplyDataProcessor{}), 		// 自定义处理回调数据。
		// jsms.WithCallbackReportDataProcessor(callbackReportDataProcessor{}),
		jsms.WithCallbackReportDataProcessor(nil), 						        // 如果设置为 nil，则忽略不处理！
		jsms.WithCallbackTemplateDataProcessor(callbackTemplateDataProcessor{}),
		// jsms.WithCallbackSignDataProcessor(callbackSignDataProcessor{}),
	}*/

	// 极光短信 - JSMS API v1
	// `EnableCallback` 配置方法的参数是可选的，可以传入一些自定义的配置选项参数，当然也可以不传入任何参数。
	smsAPIv1, _ = jsms.NewAPIv1Builder().
		// SetClient(client).                        // 【可选】配置，如果不配置，则使用 api.DefaultClient。
		SetAppKey(appKey).                        // 【必填】配置。
		SetMasterSecret(masterSecret).            // 【必填】配置。
		SetDevKey(devKey).                        // 【可选】配置，但当需要同时使用 “账号余量查询” API 接口时，请务必同时设置 `devKey`。
		SetDevSecret(devSecret).                  // 【可选】配置，但当需要同时使用 “账号余量查询” API 接口时，请务必同时设置 `devSecret`。
		// SetLogger(logger).                        // 【可选】配置，如果不配置，则使用默认的 api.DefaultJSmsLogger。
		SetHttpLogLevel(api.HttpLogLevelFull).    // 【可选】配置，如果不配置，则使用 api.HttpLogLevelBasic。
		// EnableCallback(callbackConfigOptions...). // 【可选】启用回调接口服务（默认不启用），并且自定义服务的一些配置选项。
		Build()

	/*srv, err := smsAPIv1.Callback() // Note: 前面需要先启用回调接口服务 EnableCallback(...)！
	if err != nil {
		panic(fmt.Sprintf("Get JSMS Callback Server failed! Error: %s", err))
	}*/

	// ↓↓↓ 有两种方式可以启动回调接口服务：↓↓↓

	// 1、如果你的项目是一个 Web 服务，并且想要将 SDK 内置的默认的 HTTP Handler 添加到你的匹配路由中，可以这样做（以 gin 框架为例）：
	/*r := gin.Default()
	r.Match([]string{"GET", "POST"}, "/jsms/callback", func(c *gin.Context) {
		err = srv.Handle(c.Writer, c.Request)
		if err != nil {
			log.Panicf("Handle JSMS Callback Request failed! Error: %s", err)
		}
	})*/

	// 2、如果你希望直接启动为一个独立的 HTTP 服务，可以这样做：
	/*go func() { // Run() 是一个阻塞方法，如果后面还有其他的代码需要同时执行，可在单独的 goroutine 中运行它。
		err = srv.Run()
		if err != nil {
			panic(fmt.Sprintf("Run JSMS Callback Server failed! Error: %s", err))
		}
	}()*/

	code := m.Run()

	os.Exit(code)
}
