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

package alert

// # iOS 通知优先级和投递时间的中断级别
type IosInterruptionLevel string

const (
	IosInterruptionLevelActive        IosInterruptionLevel = "active"         // 系统立即展示通知，点亮屏幕，并可播放声音。
	IosInterruptionLevelCritical      IosInterruptionLevel = "critical"       // 系统立即展示通知，点亮屏幕，并绕过静音开关播放声音。
	IosInterruptionLevelPassive       IosInterruptionLevel = "passive"        // 系统将通知添加到通知列表中，但不会点亮屏幕或播放声音。
	IosInterruptionLevelTimeSensitive IosInterruptionLevel = "time-sensitive" // 系统立即展示通知，点亮屏幕，可播放声音，并突破系统的通知控制。
)
