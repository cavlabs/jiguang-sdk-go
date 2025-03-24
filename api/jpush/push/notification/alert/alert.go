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

package alert

// # iOS 消息通知内容
type IosAlert struct {
	// 【可选】通知的标题。Apple Watch 会在简短通知界面中显示该字符串。应指定一个用户可以快速理解的简短文本。
	Title string `json:"title,omitempty"`
	// 【可选】通知的副标题，说明通知目的的其他信息。
	Subtitle string `json:"subtitle,omitempty"`
	// 【可选】通知的消息文本。
	Body string `json:"body,omitempty"`
	// 【可选】应用程序包中图片文件的名称（可包含或不包含扩展名）。当用户点击操作按钮或滑动操作条时，使用此图片作为启动画面。
	// 如果未指定，系统将使用之前的快照、Info.plist 文件中由 UILaunchImageFile 键标识的图片，或回退到默认的 Default.png。
	LaunchImage string `json:"launch-image,omitempty"`
	// 【可选】设置此字段后，系统会显示包含 “关闭” 和 “查看” 按钮的弹窗，此键用于指定 “查看” 按钮的本地化标题，而非默认的 “View”。
	ActionLocKey string `json:"action-loc-key,omitempty"`
	// 【可选】本地化标题的键。使用此键（而非 Title）从应用的 Localizable.strings 文件中检索标题。值必须是 Localizable.strings 文件中定义的键名称。
	TitleLocKey string `json:"title-loc-key,omitempty"`
	// 【可选】用于替换本地化标题中变量的值。
	// TitleLocKey 指定的字符串内容中，每个 %@ 字符将被此数组中的值替换。数组的第一个元素替换字符串中第一个 %@，第二个元素替换第二个 %@，以此类推。
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	// 【可选】本地化副标题的键。使用此键（而非 Subtitle）从应用的 Localizable.strings 文件中检索副标题。值必须是 Localizable.strings 文件中定义的键名称。
	SubtitleLocKey string `json:"subtitle-loc-key,omitempty"`
	// 【可选】用于替换本地化副标题中变量的值。
	// SubtitleLocKey 指定的字符串内容中，每个 %@ 字符将被此数组中的值替换。数组的第一个元素替换字符串中第一个 %@，第二个元素替换第二个 %@，以此类推。
	SubtitleLocArgs []string `json:"subtitle-loc-args,omitempty"`
	// 【可选】本地化消息文本的键。使用此键（而非 Body）从应用的 Localizable.strings 文件中检索消息文本。值必须是 Localizable.strings 文件中定义的键名称。
	LocKey string `json:"loc-key,omitempty"`
	// 【可选】用于替换本地化消息文本中变量的值。
	// LocKey 指定的字符串内容中，每个 %@ 字符将被此数组中的值替换。数组的第一个元素替换字符串中第一个 %@，第二个元素替换第二个 %@，以此类推。
	LocArgs []string `json:"loc-args,omitempty"`
}
