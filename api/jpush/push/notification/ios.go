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

package notification

import "github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/alert"

// iOS 平台上的 APNs 通知。
//  - 参考文档：https://developer.apple.com/documentation/usernotifications/generating-a-remote-notification#Payload-key-reference
type IOS struct {
	// 【必填】通知内容。
	//  - 这里指定内容将会覆盖上级统一指定的 Alert 信息；
	//  - 内容为空则不展示到通知栏；
	//  - 支持字符串形式（string）也支持官方定义的 alert payload 结构（alert.IosAlert）。
	// 文档链接：
	//  1. 《Payload Key Reference》：https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/PayloadKeyReference.html。
	Alert interface{} `json:"alert"`
	// 【可选】通知提示声音或警告通知。
	//  - 普通通知：string 类型，如果无此字段，则此消息无声音提示；有此字段，如果找到了指定的声音就播放该声音，否则播放默认声音，如果此字段为空字符串，iOS 7 为默认声音，iOS 8 及以上系统为无声音。
	//
	//  说明：JPush 官方 SDK 会默认填充声音字段，提供另外的方法关闭声音，详情查看各 SDK 的源码。
	//  - 告警通知：JSON Object，支持官方定义的 payload 结构（alert.IosSound），详见文档：https://developer.apple.com/documentation/usernotifications/generating-a-remote-notification#2990112。
	//  - 自定义铃声说明：格式必须是 Linear PCM、MA4（IMA/ADPCM）、alaw，μLaw 的一种，将声频文件放到项目 bundle 目录中，且时长要求 30s 以下，否则就是系统默认的铃声，详见【自定义铃声】：https://docs.jiguang.cn/jpush/practice/custom_ringtone#apns-%E9%80%9A%E9%81%93%E9%80%9A%E7%9F%A5%E5%AE%9E%E7%8E%B0。
	Sound interface{} `json:"sound,omitempty"`
	// 【可选】应用角标。
	//  - 可设置为 N、+N、-N，N 的取值范围为 [0, 99]。若上传的角标值 value 为 10，表示角标会设置为 N、10+N、10-N（值小于 0 时默认清除角标）；
	//  - 为 0 或空字符串，则表示清除角标；
	//  - 如果不填，表示不改变角标数字。
	Badge *string `json:"badge,omitempty"`
	// 【可选】推送唤醒。
	//  - 推送的时候携带 "content-available":true，说明是 Background Remote Notification，如果不携带此字段则是普通的 Remote Notification，详情参考【Background Remote Notification】：https://docs.jiguang.cn/jpush/client/iOS/ios_new_fetures#ios-7-background-remote-notification。
	ContentAvailable *bool `json:"content-available,omitempty"`
	// 【可选】通知扩展。
	//  - iOS 10 新增的 Notification Service Extension 功能，用于上报每条 APNs 信息的送达状态，使用该功能需要客户端实现 Service Extension 接口，并在服务端使用 "mutable-content" 字段完成设置；
	//  - 设置为 true 说明支持 iOS 10 的 UNNotificationServiceExtension 功能；
	//  - 如果不携带此字段则是普通的 Remote Notification，无法统计抵达数据。
	MutableContent *bool `json:"mutable-content,omitempty"`
	// 【可选】通知分类。
	//  - iOS 8 开始支持，设置 APNs payload 中的 "category" 字段值。
	Category string `json:"category,omitempty"`
	// 【可选】附加字段。
	//  - 这里自定义 key/value 信息，以供业务使用。
	// 详情参考：
	//  1. 【如何设置右侧图标/大图片】：https://docs.jiguang.cn/jpush/practice/set_icon#%E5%8F%B3%E4%BE%A7%E5%9B%BE%E6%A0%87--%E5%A4%A7%E5%9B%BE%E7%89%87；
	//  2. 【iOS 通知点击跳转】：https://docs.jiguang.cn/jpush/practice/intent_ios。
	Extras map[string]interface{} `json:"extras,omitempty"`
	// 【可选】通知分组。
	//  - iOS 的远程通知通过该属性来对通知进行分组，同一个 ThreadID 的通知归为一组。
	ThreadID string `json:"thread-id,omitempty"`
	// 【可选】通知优先级和投递时间的中断级别。
	//  - iOS 15 的通知级别，取值只能是 alert.IosInterruptionLevelActive、alert.IosInterruptionLevelCritical、alert.IosInterruptionLevelPassive、alert.IosInterruptionLevelTimeSensitive 中的一个，
	//  详情参考【UNNotificationInterruptionLevel】：https://developer.apple.com/documentation/usernotifications/unnotificationinterruptionlevel。
	InterruptionLevel alert.IosInterruptionLevel `json:"interruption-level,omitempty"`
}
