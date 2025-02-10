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

package api

import (
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

var (
	DefaultJPushLogger jiguang.Logger = jiguang.NewStdLogger(jiguang.WithLogPrefix("[JPush] "))
	DefaultJSMSLogger  jiguang.Logger = jiguang.NewStdLogger(jiguang.WithLogPrefix("[JSMS] "))
)

// HTTP 请求和响应的日志记录器。
type httpLogger struct {
	jiguang.Logger              // 底层使用的日志记录器。
	Level          HttpLogLevel // HTTP 日志记录级别。
}

func newHttpLogger(logger jiguang.Logger, level HttpLogLevel) *httpLogger {
	return &httpLogger{
		Logger: logger,
		Level:  level,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

// HTTP 日志记录级别：用于指定记录 API 的 HTTP 请求和响应的日志信息的详细程度。
type HttpLogLevel uint8

const (
	HttpLogLevelNone    HttpLogLevel = iota // 不记录任何日志。
	HttpLogLevelBasic                       // 仅记录请求方法、URL、响应状态和耗时。
	HttpLogLevelHeaders                     // 记录请求和响应的基本信息以及头部。
	HttpLogLevelFull                        // 记录请求和响应的头部、正文和元数据。
)

// 判断 level 是否是有效的 HTTP 日志记录级别。
func (level HttpLogLevel) IsValid() bool {
	switch level {
	case HttpLogLevelNone, HttpLogLevelBasic, HttpLogLevelHeaders, HttpLogLevelFull:
		return true
	default:
		return false
	}
}

func (level HttpLogLevel) String() string {
	switch level {
	case HttpLogLevelNone:
		return "None"
	case HttpLogLevelBasic:
		return "Basic"
	case HttpLogLevelHeaders:
		return "Headers"
	case HttpLogLevelFull:
		return "Full"
	default:
		return strconv.Itoa(int(level))
	}
}
