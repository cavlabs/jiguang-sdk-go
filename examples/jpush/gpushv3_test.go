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

package jpush

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/gpush"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/alert"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/style"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestGroupPushAPIv3_Send(t *testing.T) {
	param := &gpush.SendParam{}

	// 目标人群
	audience := &push.Audience{}
	audience.RegistrationIDs = []string{"1104a89793af2cfc030", "1104a89793af2cfc031"}
	audience.Tags = []string{"深圳", "北京"}
	audience.LiveActivityID = "abcdef"
	// 指定目标
	// param.Audience = audience
	// 或者发送所有人
	param.Audience = push.BroadcastAuds // all

	// 指定平台
	param.Platform = []platform.Platform{platform.Android, platform.IOS} // android, ios
	// 或者发送所有平台
	// param.Platform = platform.All // all

	// Android 通知内容
	android := &push.AndroidNotification{}
	android.Alert = "Hi, JPush!"
	android.Title = "Group Send to Android"
	android.BuilderID = 1
	android.Style = style.BigText          // 大段文本样式 (1)
	android.AlertType = alert.DefaultSound // 提示音 (1)
	android.LargeIcon = "https://www.jiguang.cn/largeIcon.jpg"
	android.Intent = &push.NotificationIntent{URL: "intent:#Intent;component=com.jiguang.push/com.example.jpushdemo.SettingActivity;end"}
	android.Extras = map[string]interface{}{"newsid": 321}
	android.BadgeAddNum = jiguang.Int(1)
	showBeginTime, _ := jiguang.ParseLocalDateTime("2024-12-25 00:00:00")
	android.ShowBeginTime = &showBeginTime
	showEndTime := jiguang.BuildLocalDateTime(2024, 12, 31, 23, 59, 59)
	android.ShowEndTime = &showEndTime

	// iOS 通知内容
	ios := &push.IosNotification{}
	ios.Alert = "Hi, JPush!"
	ios.Sound = "default"
	ios.Badge = jiguang.String("+1")
	ios.ThreadID = "group001"
	ios.Extras = map[string]interface{}{"newsid": 123}
	ios.InterruptionLevel = alert.IosInterruptionLevelPassive // passive

	// HMOS 通知内容
	hmos := &push.HmosNotification{}
	hmos.Alert = "分组通知内容"
	hmos.Title = "分组推送个性化显示"
	hmos.Intent = &push.NotificationIntent{URL: "scheme://test-group?key1=val1&key2=val2"}
	hmos.BadgeAddNum = jiguang.Int(1)
	hmos.Extras = map[string]interface{}{"skuId": 54321, "useCoupon": false}
	hmos.Category = "MARKETING"
	hmos.TestMessage = jiguang.Bool(true)
	hmos.ReceiptID = "def3434"
	hmos.LargeIcon = "https://www.jiguang.cn/largeIcon.jpg"
	hmos.Style = style.Inbox // 多行文本样式 (2)
	hmos.Inbox = map[string]interface{}{"inboxContent": []string{"1. 通知栏消息样式", "2. 通知栏消息提醒方式和展示方式", "3. 通知栏消息语言本地化"}}
	hmos.PushType = push.HmosPushTypeAlert // 通知消息 (0)

	// QuickApp 通知内容
	quickApp := &push.QuickAppNotification{}
	quickApp.Alert = "Hi, JPush!"
	quickApp.Title = "Group Send to QuickApp"
	quickApp.Page = "/page1"

	// VoIP: 此功能需要 JPush iOS SDK v3.3.2 及以上版本支持
	voip := map[string]interface{}{"key": "value"} // 任意自定义 key/value 对，API 透传下去

	// 推送通知内容
	param.Notification = &push.Notification{
		Alert:    "Hello, {{content}}！",
		Android:  android,
		IOS:      ios,
		HMOS:     hmos,
		QuickApp: quickApp,
		VoIP:     voip,
	}

	// 自定义消息内容
	message := &push.CustomMessage{}
	message.Content = "Hi, JPush!"
	message.ContentType = "text"
	message.Title = "msg"
	message.Extras = map[string]interface{}{"key": "value"}
	param.CustomMessage = message

	// 应用内增强提醒
	param.InApp = &push.InAppMessage{Enabled: true}

	// 实时活动内容
	liveActivityIos := &push.LiveActivityIosMessage{}
	liveActivityIos.Event = push.LiveActivityEventStart // start 创建
	liveActivityIos.ContentState = map[string]interface{}{"progress": 0.1, "currentPositionInQueue": 8}
	liveActivityIos.AttributesType = "WaitTimeDemoAttributes"
	liveActivityIos.Attributes = map[string]interface{}{"waitlistName": "For Testing", "waitlistId": "", "bookingId": ""}
	liveActivityIos.Alert = &push.LiveActivityIosAlertMessage{
		Title: "Send to {{nickname}}",
		Body:  "Hi, {{nickname}}!",
		Sound: "default",
	}
	liveActivityIos.DismissalDate = 1168364460
	param.LiveActivity = &push.LiveActivityMessage{IOS: liveActivityIos}

	// 自定义消息转厂商通知内容（v1 版本）
	/*notification3rd := &push.ThirdNotification{} // Deprecated: 已过时，推荐使用 ThirdNotificationV2
	notification3rd.Content = "Hi, JPush!"
	notification3rd.Title = "msg"
	notification3rd.ChannelID = "channel001"
	notification3rd.UriActivity = "cn.jpush.android.ui.OpenClickActivity"
	notification3rd.UriAction = "cn.jpush.android.intent.CONNECTION"
	notification3rd.BadgeAddNum = jiguang.Int(1)
	notification3rd.BadgeSetNum = jiguang.Int(1)
	notification3rd.BadgeClass = "com.test.badge.MainActivity"
	notification3rd.Sound = "sound"
	notification3rd.Extras = map[string]interface{}{"news_id": 134, "my_key": "a value"}*/

	// 自定义消息转厂商通知内容（v2 版本）
	notification3rd := &push.ThirdNotificationV2{} // 推荐使用 v2 版本
	notification3rd.Android = android              // customize yourself, here only for example
	notification3rd.IOS = ios                      // customize yourself, here only for example
	notification3rd.HMOS = hmos                    // customize yourself, here only for example

	param.ThirdNotification = notification3rd

	// 短信
	sms := &push.SmsMessage{}
	sms.TempID = 1250
	sms.TempParams = map[string]interface{}{"code": "123456"}
	sms.DelayTime = 3600
	sms.ActiveFilter = jiguang.Bool(false)
	param.SmsMessage = sms

	// 可选项
	options := &push.Options{}
	options.TimeToLive = jiguang.Int64(60)
	options.ApnsProduction = jiguang.Bool(false)
	options.ApnsCollapseID = "jiguang_test_201706011100"
	huaweiChannelOptions := &push.ThirdPartyChannelOptions{
		Distribution: "first_ospush",
		Importance:   "NORMAL",
		Category:     "MARKETING",
	}
	options.ThirdPartyChannel = &push.ThirdPartyChannel{Huawei: huaweiChannelOptions}
	// options.PortalExtra = &push.PortalExtraOptions{Task: "group-task001"}
	options.Notification3rdVer = "v2" // v1 版本填 "v1"，v2 版本填 "v2"
	param.Options = options

	// 回调
	callback := &push.Callback{
		URL:    "https://www.bilibili.com",
		Params: map[string]interface{}{"name": "joe", "age": 26},
		Type:   push.CallbackTypeReceived | push.CallbackTypeClicked, // 送达和点击回执 (1 + 2 = 3)
	}
	param.Callback = callback

	result, err := gpushAPIv3.Send(context.Background(), param)
	// result, err := gpushAPIv3.CustomSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("GroupPushAPIv3_Send: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		t.Logf(">>> GroupMsgID: %s", result.GroupMsgID)
		for k, v := range result.Successes {
			success, _ := json.Marshal(v)
			t.Logf(">>> Success[%s]: %s", k, success)
		}
		for k, v := range result.Errors {
			t.Logf(">>> Failed.Error[%s]: %s", k, v.String())
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestGroupPushAPIv3_SendByFile(t *testing.T) {
	// 注意：调用文件上传接口获取 fileID 时，需要使用 devKey 和 devSecret 进行验证，详情参考【文件上传接口】：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file#%E8%B0%83%E7%94%A8%E9%AA%8C%E8%AF%81-1。
	param := &gpush.SendParam{
		Platform:      platform.All,
		Audience:      &push.FileAudience{FileID: "004f2b2c3d8fc649fa60ad2c-15563487-577d-4bb0-806b-56f1a19e06c1"},
		Notification:  &push.Notification{Alert: "Hello, JPush!"},
		CustomMessage: &push.CustomMessage{Content: "Hello, JPush!"},
		Options:       &push.Options{TimeToLive: jiguang.Int64(600), ApnsProduction: jiguang.Bool(false)},
	}

	result, err := gpushAPIv3.SendByFile(context.Background(), param)
	// result, err := gpushAPIv3.CustomSendByFile(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("GroupPushAPIv3_SendByFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		t.Logf(">>> GroupMsgID: %s", result.GroupMsgID)
		for k, v := range result.Successes {
			success, _ := json.Marshal(v)
			t.Logf(">>> Success[%s]: %s", k, success)
		}
		for k, v := range result.Errors {
			t.Logf(">>> Failed.Error[%s]: %s", k, v.String())
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// 注意：通过 Group Push API v3 同样可以调用以下接口：
// UploadFileForAlias、UploadFileForRegistrationID、GetFiles、GetFile、DeleteFile 等接口的使用示例请查看 examples/jpush/filev3_test.go。
