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
	"time"

	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/alert"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification/style"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestPushAPIv3_Send(t *testing.T) {
	param := &push.SendParam{}

	// 推送唯一标识 (CID)
	param.CID = "8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e"
	/*cidResult, _ := pushAPIv3.GetCidForPush(context.Background(), 1)
	if cidResult.IsSuccess() {
		param.CID = cidResult.CidList[0]
	} else {
		t.Fatalf("Failed! Error: %s", cidResult.Error)
	}*/

	// 目标人群
	audience := &push.Audience{}
	audience.RegistrationIDs = []string{"1104a89793af2cfc030", "1104a89793af2cfc031"}
	audience.Tags = []string{"深圳", "北京"}
	audience.LiveActivityID = "abcdef"
	// 指定目标
	param.Audience = audience
	// 或者发送所有人
	// param.Audience = push.BroadcastAuds // all

	// 指定平台
	// param.Platform = []platform.Platform{platform.Android, platform.IOS} // android, ios
	// 或者发送所有平台
	param.Platform = platform.All // all

	// Android 通知内容
	android := &push.AndroidNotification{}
	android.Alert = "Hi, JPush!"
	android.Title = "Send to Android"
	android.BuilderID = 1
	android.Style = style.BigPicture                             // 大图片样式 (3)
	android.AlertType = alert.DefaultSound | alert.DefaultLights // 提示音 (1) + 指示灯 (4) = 5
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
	// ios.Alert = "Hi, JPush!"
	ios.Alert = &alert.IosAlert{
		Title:        "Game Request",
		ActionLocKey: "PLAY",
		LocKey:       "GAME_PLAY_REQUEST_FORMAT",
		LocArgs:      []string{"Jenna", "Frank"},
	}
	// ios.Sound = "default"
	ios.Sound = &alert.IosSound{
		Critical: 0,
		Name:     "sound.caf",
		Volume:   0.5,
	}
	ios.Badge = jiguang.String("+1")
	ios.ThreadID = "default"
	ios.Extras = map[string]interface{}{"newsid": 321}
	ios.InterruptionLevel = alert.IosInterruptionLevelActive // active

	// HMOS 通知内容
	hmos := &push.HmosNotification{}
	hmos.Alert = "通知内容"
	hmos.Title = "推送个性化显示"
	hmos.Intent = &push.NotificationIntent{URL: "scheme://test?key1=val1&key2=val2"}
	hmos.BadgeAddNum = jiguang.Int(1)
	hmos.Extras = map[string]interface{}{"skuId": 12345, "useCoupon": true}
	hmos.Category = "MARKETING"
	hmos.TestMessage = jiguang.Bool(true)
	hmos.ReceiptID = "abc1212"
	hmos.LargeIcon = "https://www.jiguang.cn/largeIcon.jpg"
	hmos.Style = style.Inbox // 多行文本样式 (2)
	hmos.Inbox = map[string]interface{}{"inboxContent": []string{"1. 通知栏消息样式", "2. 通知栏消息提醒方式和展示方式", "3. 通知栏消息语言本地化"}}
	hmos.PushType = push.HmosPushTypeExtension // 通知扩展消息 (2)

	// QuickApp 通知内容
	quickApp := &push.QuickAppNotification{}
	quickApp.Alert = "Hi, JPush!"
	quickApp.Title = "Send to QuickApp"
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
	liveActivityIos.DismissalDate = time.Now().Add(6 * time.Hour).Unix() // 6 小时后过期
	param.LiveActivity = &push.LiveActivityMessage{IOS: liveActivityIos}

	// 自定义消息转厂商通知内容
	notification3rd := &push.ThirdNotification{}
	notification3rd.Content = "Hi, JPush!"
	notification3rd.Title = "msg"
	notification3rd.ChannelID = "channel001"
	notification3rd.UriActivity = "cn.jpush.android.ui.OpenClickActivity"
	notification3rd.UriAction = "cn.jpush.android.intent.CONNECTION"
	notification3rd.BadgeAddNum = jiguang.Int(1)
	notification3rd.BadgeSetNum = jiguang.Int(1)
	notification3rd.BadgeClass = "com.test.badge.MainActivity"
	notification3rd.Sound = "sound"
	notification3rd.Extras = map[string]interface{}{"news_id": 134, "my_key": "a value"}
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
	// options.PortalExtra = &push.PortalExtraOptions{Task: "task001"}
	param.Options = options

	// 回调
	callback := &push.Callback{
		URL:    "https://www.bilibili.com",
		Params: map[string]interface{}{"name": "joe", "age": 26},
		Type:   push.CallbackTypeReceived | push.CallbackTypeClicked, // 送达和点击回执 (1 + 2 = 3)
	}
	param.Callback = callback

	result, err := pushAPIv3.Send(context.Background(), param)
	// result, err := pushAPIv3.SendWithSM2(context.Background(), param)
	// result, err := pushAPIv3.CustomSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_Send: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		t.Logf(">>> MsgID: %s", result.MsgID)
		t.Logf(">>> SendNo: %s", result.SendNo)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_SendByFile(t *testing.T) {
	fileID := "004f2b2c3d8fc649fa60ad2c-15563487-577d-4bb0-806b-56f1a19e06c1"
	/*var fileID string
	uploadFileResult, _ := pushAPIv3.UploadFileForRegistrationID(context.Background(),
		&file.FileUploadParam{File: "~/Desktop/registrationId.txt"})
	if uploadFileResult.IsSuccess() {
		fileID = uploadFileResult.FileID
	} else {
		t.Fatalf("Failed! Error: %s", uploadFileResult.Error)
	}*/

	param := &push.SendParam{
		Platform:      platform.All,
		Audience:      &push.FileAudience{FileID: fileID},
		Notification:  &push.Notification{Alert: "Hello, JPush!"},
		CustomMessage: &push.CustomMessage{Content: "Hello, JPush!"},
		Options:       &push.Options{TimeToLive: jiguang.Int64(600), ApnsProduction: jiguang.Bool(false)},
	}

	result, err := pushAPIv3.SendByFile(context.Background(), param)
	// result, err := pushAPIv3.CustomSendByFile(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_SendByFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		t.Logf(">>> MsgID: %s", result.MsgID)
		t.Logf(">>> SendNo: %s", result.SendNo)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_ValidateSend(t *testing.T) {
	param := &push.SendParam{
		Audience: push.BroadcastAuds, // 发送所有人
		Platform: platform.All,       // 发送所有平台
	}
	// Android 通知内容
	android := &push.AndroidNotification{Extras: map[string]interface{}{"android-key1": "android-value1"}}
	ios := &push.IosNotification{
		Sound:  jiguang.String("sound.caf"),
		Badge:  jiguang.String("+1"),
		Extras: map[string]interface{}{"ios-key1": "ios-value1"},
	}
	// 推送通知内容
	param.Notification = &push.Notification{
		Alert:   "Hi, JPush!",
		Android: android,
		IOS:     ios,
	}

	result, err := pushAPIv3.ValidateSend(context.Background(), param)
	// result, err := pushAPIv3.ValidateCustomSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_ValidateSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		t.Logf(">>> MsgID: %s", result.MsgID)
		t.Logf(">>> SendNo: %s", result.SendNo)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_BatchSend(t *testing.T) {
	pushList := map[string]push.BatchPushParam{
		"8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e": {
			Platform:      platform.All,
			Target:        "1104a89793af2cfc030",
			Notification:  &push.Notification{Alert: "Hello, JPush!"},
			CustomMessage: &push.CustomMessage{Content: "Hello, JPush!"},
			SmsMessage:    &push.SmsMessage{TempID: 1250, TempParams: map[string]interface{}{"code": "123456"}},
			Options:       &push.Options{TimeToLive: jiguang.Int64(600), ApnsProduction: jiguang.Bool(false)},
		},
		"7103a4c428a0b98974ec1849-711161d4-5f17-4d2f-b855-5e5a8909b26e": {
			Platform:      platform.All,
			Target:        "1104a89793af2cfc031",
			Notification:  &push.Notification{Alert: "Hi, JPush!"},
			CustomMessage: &push.CustomMessage{Content: "Hi, JPush!"},
			SmsMessage:    &push.SmsMessage{TempID: 1250, TempParams: map[string]interface{}{"code": "123456"}},
			Options:       &push.Options{TimeToLive: jiguang.Int64(600), ApnsProduction: jiguang.Bool(false)},
		},
	}

	result, err := pushAPIv3.BatchSendByRegistrationID(context.Background(), pushList)
	// result, err := pushAPIv3.BatchSendByAlias(context.Background(), pushList)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_BatchSend(by Registration ID or Alias): StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		for cid, prs := range result.SendResult {
			t.Logf(">>> CID: %s, MsgID: %s, Error: %s", cid, prs.MsgID, prs.Error)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_TemplateSend(t *testing.T) {
	// 广播推送
	templateID := "c65b34b96512"
	templateParams := []push.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是广播标题", "content": "这是广播内容"},
			Audience: push.BroadcastAuds,
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "8de06fh-8djfgg",
		},
	}

	// Registration ID 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []push.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第⼀个regid的标题", "content": "这是第⼀个regid的内容"},
			Audience: &push.Audience{RegistrationIDs: []string{"13065ffa4ee8411496f"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "98ode06fh-8dgg",
		},
		{
			Keys:     map[string]string{"title": "这是推送给第二个regid的标题", "content": "这是第二个regid的内容"},
			Audience: &push.Audience{RegistrationIDs: []string{"8914afdsa31"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "8de06fh-8djfgg",
		},
	}*/

	// Alias 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []push.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第⼀个alias的标题", "content": "这是第⼀个alias的内容"},
			Audience: &push.Audience{Aliases: []string{"4312kjkla31"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "98ode06fh-8dgg",
		},
		{
			Keys:     map[string]string{"title": "这是推送给第二个alias的标题", "content": "这是第二个alias的内容"},
			Audience: &push.Audience{Aliases: []string{"8914sa31"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "8de06fh-8djfgg",
		},
	}*/

	// Tag 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []push.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第1个tag的标题", "content": "这是第1个tag的内容"},
			Audience: &push.Audience{Tags: []string{"tag1"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "8de06fh-8djfgg",
		},
	}*/

	// Segment 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []push.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第1个segment的标题", "content": "这是第1个segment的内容"},
			Audience: &push.Audience{Segments: []string{"segmentid001"}},
			Options:  &push.Options{TestMode: jiguang.Bool(true)},
			TraceID:  "8de06fh-8djfgg",
		},
	}*/

	result, err := pushAPIv3.TemplateSend(context.Background(), templateID, templateParams)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_TemplateSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		if result.Data != nil && len(result.Data.PushList) > 0 {
			for i, v := range result.Data.PushList {
				if v.IsSuccess() {
					t.Logf(">>> PushList[%d]: Success! MessageID: %s, TraceID: %s", i, v.MessageID, v.TraceID)
				} else {
					t.Logf(">>> PushList[%d]: Failed! MessageID: %s, TraceID: %s, Error: %s", i, v.MessageID, v.TraceID, v.Error)
				}
			}
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_GetCidForPush(t *testing.T) {
	result, err := pushAPIv3.GetCidForPush(context.Background(), 10)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_GetCidForPush: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> CidList: %s", result.CidList)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_WithdrawMessage(t *testing.T) {
	msgID := "18101213529672826"
	result, err := pushAPIv3.WithdrawMessage(context.Background(), msgID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_WithdrawMessage: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestPushAPIv3_GetQuota(t *testing.T) {
	result, err := pushAPIv3.GetQuota(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("PushAPIv3_GetQuota: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		xiaomi, _ := json.Marshal(result.Data.Xiaomi)
		t.Logf(">>> Quota of xiaomi: %s", xiaomi)
		oppo, _ := json.Marshal(result.Data.OPPO)
		t.Logf(">>> Quota of OPPO: %s", oppo)
		vivo, _ := json.Marshal(result.Data.Vivo)
		t.Logf(">>> Quota of vivo: %s", vivo)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// 注意：通过 Push API v3 同样可以调用以下接口：
// UploadFileForAlias、UploadFileForRegistrationID、GetFiles、GetFile、DeleteFile 等接口的使用示例请查看 examples/jpush/filev3_test.go；
// AddImageByUrl、UpdateImageByUrl、AddImageByFile、UpdateImageByFile 等接口的使用示例请查看 examples/jpush/imagev3_test.go；
// ScheduleSend、ScheduleTemplateSend、GetCidForSchedulePush、GetSchedules、GetSchedule、GetScheduleMsgIDs、UpdateSchedule、DeleteSchedule、CustomScheduleSend 等接口的使用示例请查看 examples/jpush/schedulev3_test.go。
