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

package callback

import (
	"context"
	"encoding/json"

	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// 回调数据列表处理器。
type DataListProcessor interface {
	Process(list DataList)
}

// 回调数据项处理器。
type DataProcessor interface {
	Process(data Data)
}

// ---------------------------------------------------------------------------------------------------------------------

type loggingDataListProcessor struct {
	logger jiguang.Logger
}

func (p loggingDataListProcessor) Process(list DataList) {
	p.logger.Debugf(context.TODO(), "收到回调消息: %s", list.RawJSON)
}

// ---------------------------------------------------------------------------------------------------------------------

type loggingDataProcessor struct {
	logger jiguang.Logger
}

func (p loggingDataProcessor) Process(data Data) {
	s, _ := json.Marshal(data)
	if data.Type == Click {
		p.logger.Debugf(context.TODO(), "「　点击　」回调: %s", s) // Tip: 这里仅仅是为了打印对齐好看而已！
	} else {
		p.logger.Debugf(context.TODO(), "「%s」回调: %s", data.Type, s)
	}
}
