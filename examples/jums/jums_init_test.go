/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
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

package jums

import (
	"os"
	"testing"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jums"
)

var umsAPIv1 jums.APIv1 // 极光统一消息 - JUMS API v1

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

	// logger := adapter.NewLogrusLogger("[JUMS] ")
	// logger := adapter.NewZapLogger("[JUMS]")
	// logger := adapter.NewZeroLogger("[JUMS] ")

	// channelKey := os.Getenv("JUMS_CHANNEL_KEY")
	channelKey := "7d431e42dfa6a6d693ac2d04"
	// masterSecret := os.Getenv("JUMS_MASTER_SECRET")
	masterSecret := "5e987ac6d2e04d95a9d8f0d1"

	// accessKey := os.Getenv("JUMS_ACCESS_KEY")
	accessKey := "3f9a1b82cde476f5a8b913d2"
	// accessMasterSecret := os.Getenv("JUMS_ACCESS_MASTER_SECRET")
	accessMasterSecret := "8e2d47a5fcb301d69c74e8f1"

	// 启用极光统一消息的回调接口服务，可以自定义服务的一些配置选项：
	//  -  1. 监听地址 (addr)：默认为 ":8089"；
	//  -  2. 回调路径 (path)：默认为 "/callback"；
	//  -  3. 日志打印器 (logger)：默认与 JUMS API v1 使用相同的日志打印器；
	//  -  4. 是否开启安全校验 (checkAuth)：默认为 true，开启安全校验；
	//  -  5. HTTP Handler (handler)：默认为使用 net/http 实现的一个简单的 Handler；
	//  -  6. 目标有效 (0) 回调数据处理器 (targetValid)，为 nil 时不处理；
	//  -  7. 目标无效 (1) 回调数据处理器 (targetInvalid)，为 nil 时不处理；
	//  -  8. 提交成功 (2) 回调数据处理器 (sentSucc)，为 nil 时不处理；
	//  -  9. 提交失败 (3) 回调数据处理器 (sentFail)，为 nil 时不处理；
	//  - 10. 送达成功 (4) 回调数据处理器 (receivedSucc)，为 nil 时不处理；
	//  - 11. 送达失败 (5) 回调数据处理器 (receivedFail)，为 nil 时不处理；
	//  - 12. 点击 (6) 回调数据处理器 (click)，为 nil 时不处理；
	//  - 13. 撤回成功 (7) 回调数据处理器 (retractedSucc)，为 nil 时不处理；
	//  - 14. 撤回失败 (8) 回调数据处理器 (retractedFail)，为 nil 时不处理；
	//  - 15. 统一的回调数据列表处理器（所有类型的回调数据汇总到此处理器中进行统一处理）(unified)，为 nil 时不处理。
	/*callbackConfigOptions := []jums.CallbackConfigOption{
		jums.WithCallbackAddr(":8081"),
		jums.WithCallbackPath("/jums/callback"),
		// jums.WithCallbackLogger(logger), 									              // 这是默认行为，可以不设置！
		// jums.WithCallbackCheckAuth(false),
		// jums.WithCallbackHttpHandler(httpHandler),
		// jums.WithCallbackTargetValidDataProcessor(callbackTargetValidDataProcessor{}),     // 自定义处理回调数据。
		// jums.WithCallbackTargetInvalidDataProcessor(callbackTargetInvalidDataProcessor{}),
		// jums.WithCallbackSentSuccDataProcessor(callbackSentSuccDataProcessor{}),
		jums.WithCallbackSentSuccDataProcessor(nil),                                          // 如果设置为 nil，则忽略不处理！
		// jums.WithCallbackSentFailDataProcessor(callbackSentFailDataProcessor{}),
		jums.WithCallbackReceivedSuccDataProcessor(callbackReceivedSuccDataProcessor{}),
		// jums.WithCallbackReceivedFailDataProcessor(callbackReceivedFailDataProcessor{}),
		// jums.WithCallbackClickDataProcessor(callbackClickDataProcessor{}),
		jums.WithCallbackRetractedSuccDataProcessor(callbackRetractedSuccDataProcessor{}),
		// jums.WithCallbackRetractedFailDataProcessor(callbackRetractedFailDataProcessor{}),
		jums.WithCallbackDataListProcessor(callbackDataListProcessor{}),                      // 如果你不希望根据不同的回调数据类型使用不同的处理器，可以仅配置使用此选项（6-14 选项不用配置）。
	}*/

	// 极光统一消息 - JUMS API v1
	// `EnableCallback` 配置方法的参数是可选的，可以传入一些自定义的配置选项参数，当然也可以不传入任何参数。
	umsAPIv1, _ = jums.NewAPIv1Builder().
		// SetClient(client).                         // 【可选】配置，如果不配置，则使用 api.DefaultClient。
		SetChannelKey(channelKey).                 // 【必填】配置。
		SetMasterSecret(masterSecret).             // 【必填】配置。
		SetAccessKey(accessKey).                   // 【可选】配置，但当需要同时使用 “用户管理” API 接口时，请务必同时设置 `accessKey`。
		SetAccessMasterSecret(accessMasterSecret). // 【可选】配置，但当需要同时使用 “用户管理” API 接口时，请务必同时设置 `accessMasterSecret`。
		// SetLogger(logger).                         // 【可选】配置，如果不配置，则使用默认的 api.DefaultJUmsLogger。
		SetHttpLogLevel(api.HttpLogLevelFull).     // 【可选】配置，如果不配置，则使用 api.HttpLogLevelBasic。
		// EnableCallback(callbackConfigOptions...).  // 【可选】启用回调接口服务（默认不启用），并且自定义服务的一些配置选项。
		Build()

	/*srv, err := umsAPIv1.Callback() // Note: 前面需要先启用回调接口服务 EnableCallback(...)！
	if err != nil {
		panic(fmt.Sprintf("Get JUMS Callback Server failed! Error: %s", err))
	}*/

	// ↓↓↓ 有两种方式可以启动回调接口服务：↓↓↓

	// 1、如果你的项目是一个 Web 服务，并且想要将 SDK 内置的默认的 HTTP Handler 添加到你的匹配路由中，可以这样做（以 gin 框架为例）：
	/*r := gin.Default()
	r.Match([]string{"GET", "POST"}, "/jums/callback", func(c *gin.Context) {
		err = srv.Handle(c.Writer, c.Request)
		if err != nil {
			log.Panicf("Handle JUMS Callback Request failed! Error: %s", err)
		}
	})*/

	// 2、如果你希望直接启动为一个独立的 HTTP 服务，可以这样做：
	/*go func() { // Run() 是一个阻塞方法，如果后面还有其他的代码需要同时执行，可在单独的 goroutine 中运行它。
		err = srv.Run()
		if err != nil {
			panic(fmt.Sprintf("Run JUMS Callback Server failed! Error: %s", err))
		}
	}()*/

	code := m.Run()

	os.Exit(code)
}
