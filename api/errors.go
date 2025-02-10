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

import "fmt"

// API 访问客户端未初始化错误哨兵。

var (
	ErrNilJPushAdminAPIv1       = newNilAPIError("jpush", "admin", 1)
	ErrNilJPushDeviceAPIv3      = newNilAPIError("jpush", "device", 3)
	ErrNilJPushPushAPIv3        = newNilAPIError("jpush", "push", 3)
	ErrNilJPushGroupPushAPIv3   = newNilAPIError("jpush", "group push", 3)
	ErrNilJPushScheduleAPIv3    = newNilAPIError("jpush", "schedule", 3)
	ErrNilJPushFileAPIv3        = newNilAPIError("jpush", "file", 3)
	ErrNilJPushImageAPIv3       = newNilAPIError("jpush", "image", 3)
	ErrNilJPushReportAPIv3      = newNilAPIError("jpush", "report", 3)
	ErrNilJPushGroupReportAPIv3 = newNilAPIError("jpush", "group report", 3)
	ErrNilJSmsAPIv1             = newNilAPIError("jsms", "", 1)
)

// ---------------------------------------------------------------------------------------------------------------------

// 极光 API 访问客户端为空错误。
type nilAPIError struct {
	category string // API 类别
	name     string // API 名称
	version  int    // API 版本
}

func newNilAPIError(category, name string, version int) *nilAPIError {
	return &nilAPIError{
		category: category,
		name:     name,
		version:  version,
	}
}

func (e *nilAPIError) Error() string {
	if e == nil {
		return "nil api error"
	}
	name := ""
	if e.name != "" {
		name = e.name + " "
	}
	return fmt.Sprintf("%s: nil %sapiv%d", e.category, name, e.version)
}

func (e *nilAPIError) String() string {
	if e == nil {
		return ""
	}
	name := ""
	if e.name != "" {
		name = e.name + " "
	}
	return fmt.Sprintf("%s: nil %sapiv%d", e.category, name, e.version)
}

// ---------------------------------------------------------------------------------------------------------------------

// 极光 REST API 的通用错误响应封装。
type CodeError struct {
	Code    int    `json:"code"`              // 错误返回码，0 表示成功，非 0 表示失败。
	Message string `json:"message,omitempty"` // 错误描述说明。
}

func (e *CodeError) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *CodeError) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *CodeError) IsSuccess() bool {
	return e == nil || e.Code == 0
}
