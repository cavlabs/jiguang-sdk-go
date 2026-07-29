// Copyright 2026 cavlabs/jiguang-sdk-go authors.
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

package jpush

import (
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/device/platform"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push/notification"
)

// 常用推送载荷类型别名；完整结构仍见各子包。

type (
	SendParam     = push.SendParam
	SendResult    = push.SendResult
	Notification  = push.Notification
	Platform      = platform.Platform
	AndroidNotify = notification.Android
	IosNotify     = notification.IOS
	HmosNotify    = notification.HMOS
)
