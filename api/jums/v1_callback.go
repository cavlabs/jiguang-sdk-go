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
	"errors"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jums/callback"
)

// # 获取配置好的回调接口服务
func (u *apiv1) Callback() (*Callback, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}
	if u.callback == nil {
		return nil, errors.New("JUMS callback server is disabled")
	}
	return u.callback, nil
}

// ↓↓↓ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↓↓↓

type (
	Callback                  = callback.Server            // 回调接口服务核心结构。
	CallbackType              = callback.Type              // 回调数据类型。
	CallbackConfigOption      = callback.ConfigOption      // 回调接口服务配置选项。
	CallbackDataList          = callback.DataList          // 回调数据列表。
	CallbackDataListProcessor = callback.DataListProcessor // 回调数据列表处理器。
	CallbackData              = callback.Data              // 回调数据项。
	CallbackDataProcessor     = callback.DataProcessor     // 回调数据项处理器。
)

var (
	// 自定义配置回调接口服务监听地址，默认为 ":8089"。
	WithCallbackAddr = callback.WithAddr
	// 自定义配置回调接口服务回调路径，默认为 "/callback"。
	WithCallbackPath = callback.WithPath
	// 自定义配置回调接口服务的日志打印器，默认与 JUMS API v1 使用相同的日志打印器。
	WithCallbackLogger = callback.WithLogger
	// 自定义配置回调接口服务是否开启安全校验，默认开启。
	WithCallbackCheckAuth = callback.WithCheckAuth
	// 自定义配置回调接口服务的 HTTP Handler，默认为使用 net/http 实现的一个简单的 Handler。
	WithCallbackHttpHandler = callback.WithHttpHandler
	// 自定义配置 目标有效 (0) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackTargetValidDataProcessor = callback.WithTargetValidDataProcessor
	// 自定义配置 目标无效 (1) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackTargetInvalidDataProcessor = callback.WithTargetInvalidDataProcessor
	// 自定义配置 短信提交成功 (2) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackSentSuccDataProcessor = callback.WithSentSuccDataProcessor
	// 自定义配置 短信提交失败 (3) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackSentFailDataProcessor = callback.WithSentFailDataProcessor
	// 自定义配置 短信送达成功 (4) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackReceivedSuccDataProcessor = callback.WithReceivedSuccDataProcessor
	// 自定义配置 短信送达失败 (5) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackReceivedFailDataProcessor = callback.WithReceivedFailDataProcessor
	// 自定义配置 短信点击 (6) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackClickDataProcessor = callback.WithClickDataProcessor
	// 自定义配置 短信撤回成功 (7) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackRetractedSuccDataProcessor = callback.WithRetractedSuccDataProcessor
	// 自定义配置 短信撤回失败 (8) 回调数据处理器。注：你的自定义处理器需要实现 CallbackDataProcessor 接口。
	WithCallbackRetractedFailDataProcessor = callback.WithRetractedFailDataProcessor
	// 统一的回调数据列表处理器（所有类型的回调数据汇总到此处理器中进行统一处理）。注：你的自定义处理器需要实现 CallbackDataListProcessor 接口。
	//  - 如果你不希望根据不同的回调数据类型使用不同的处理器，可以仅配置使用此选项。
	WithCallbackDataListProcessor = callback.WithDataListProcessor
)

const (
	CallbackTypeTargetValid   = callback.TargetValid   // 目标有效 (0)
	CallbackTypeTargetInvalid = callback.TargetInvalid // 目标无效 (1)
	CallbackTypeSentSucc      = callback.SentSucc      // 提交成功 (2)
	CallbackTypeSentFail      = callback.SentFail      // 提交失败 (3)
	CallbackTypeReceivedSucc  = callback.ReceivedSucc  // 送达成功 (4)
	CallbackTypeReceivedFail  = callback.ReceivedFail  // 送达失败 (5)
	CallbackTypeClick         = callback.Click         // 点击 (6)
	CallbackTypeRetractedSucc = callback.RetractedSucc // 撤回成功 (7)
	CallbackTypeRetractedFail = callback.RetractedFail // 撤回失败 (8)
)

// ↑↑↑ 这是为了方便 SDK 的使用者，提供了一些共享模型的别名定义。↑↑↑
