/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

package jsms

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/api/jsms"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

func TestAPIv1_CreateSign(t *testing.T) {
	param := &jsms.SignCreateParam{
		Sign: "SDK6",
		Type: 1,
		Images: []string{
			"~/Downloads/1.jpg",
			"~/Downloads/2.jpeg",
			"~/Downloads/3.jpeg",
		},
		Remark: "SDK测试",
	}
	result, err := smsAPIv1.CreateSign(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_CreateSign: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SignID: %d", result.SignID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_UpdateSign(t *testing.T) {
	/*image, err := os.Open("/Users/z/Desktop/1-sign-info.png")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	defer image.Close()*/

	param := &jsms.SignCreateParam{
		Sign:   "SDK6",
		Type:   1,
		Images: "~/Downloads/3.jpeg",
		Remark: "SDK测试-修改",
	}
	result, err := smsAPIv1.UpdateSign(context.Background(), 37582, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_UpdateSign: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SignID: %d", result.SignID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_GetSign(t *testing.T) {
	result, err := smsAPIv1.GetSign(context.Background(), 37582)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetSign: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SignID: %d, Sign: %s, Status: %d, IsDefault: %t, UseStatus: %t",
			result.SignID, result.Sign, result.Status, result.IsDefault, result.UseStatus)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_DeleteSign(t *testing.T) {
	result, err := smsAPIv1.DeleteSign(context.Background(), 37582)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_DeleteSign: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAPIv1_CreateTemplate(t *testing.T) {
	param := &jsms.TemplateCreateParam{
		Template: "您好，您的验证码是 {{code}}，2 分钟内有效！",
		Type:     1,
		TTL:      120,
		Remark:   "此模板用于用户注册",
	}
	result, err := smsAPIv1.CreateTemplate(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_CreateTemplate: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> TempID: %d", result.TempID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_UpdateTemplate(t *testing.T) {
	param := &jsms.TemplateUpdateParam{
		Template: "【极光短信】您好，您的注册验证码是 {{code}}，2 分钟内有效！",
		Type:     1,
		TTL:      120,
		Remark:   "此模板用于用户注册",
	}
	result, err := smsAPIv1.UpdateTemplate(context.Background(), 144923, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_UpdateTemplate: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> TempID: %d", result.TempID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_GetTemplate(t *testing.T) {
	result, err := smsAPIv1.GetTemplate(context.Background(), 144923)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetTemplate: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> TempID: %d, Template: %s, Type: %d, TTL: %d, Remark: %s, Status: %d",
			result.TempID, result.Template, result.Type, result.TTL, result.Remark, result.Status)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_DeleteTemplate(t *testing.T) {
	result, err := smsAPIv1.DeleteTemplate(context.Background(), 144923)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_DeleteTemplate: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAPIv1_SendCode(t *testing.T) {
	param := &jsms.CodeSendParam{
		Mobile: "13800138000",
		TempID: 1,
	}
	result, err := smsAPIv1.SendCode(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendCode: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MsgID: %s", result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_SendVoiceCode(t *testing.T) {
	param := &jsms.VoiceCodeSendParam{
		Mobile:    "13800138000",
		Code:      "123456",
		VoiceLang: jsms.VoiceLangChinese,
		TTL:       90,
	}
	result, err := smsAPIv1.SendVoiceCode(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendVoiceCode: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MsgID: %s", result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_VerifyCode(t *testing.T) {
	result, err := smsAPIv1.VerifyCode(context.Background(), "01658697-45d9-4644-996d-69a1b14e2bb8", "556618")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_VerifyCode: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success!")
		t.Logf(">>> Is code valid? %t", result.IsValid)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_SendMessage(t *testing.T) {
	param := &jsms.MessageSendParam{
		Mobile:     "13800138000",
		TempID:     1,
		TempParams: map[string]interface{}{"test": "jpush"},
	}
	result, err := smsAPIv1.SendMessage(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendMessage: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MsgID: %s", result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_SendBatchMessage(t *testing.T) {
	param := &jsms.BatchMessageSendParam{
		SignID: 0,
		TempID: 1,
		Recipients: []jsms.BatchSendRecipientParam{
			{
				Mobile:     "13812345678",
				TempParams: map[string]interface{}{"code": "638938"},
			},
			{
				Mobile:     "18603050709",
				TempParams: map[string]interface{}{"code": "147721"},
			},
		},
	}
	result, err := smsAPIv1.SendBatchMessage(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendBatchMessage: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		recipients, _ := json.Marshal(result.Recipients)
		t.Logf(">>> SuccessCount: %d, FailureCount: %d, Recipients: %s",
			*result.SuccessCount, *result.FailureCount, recipients)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAPIv1_SendSchedule(t *testing.T) {
	sendTime, _ := jiguang.ParseLocalDateTime("2017-07-01 09:00:00")
	param := &jsms.ScheduleSendParam{
		SendTime:   sendTime,
		Mobile:     "13812345678",
		SignID:     37582,
		TempID:     1250,
		TempParams: map[string]interface{}{"number": "741627"},
	}
	result, err := smsAPIv1.SendSchedule(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> ScheduleID: %s", result.ScheduleID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_UpdateSchedule(t *testing.T) {
	scheduleID := "1a886e7c-a267-49e6-9970-0d396ca5bb1e"
	sendTime, _ := jiguang.ParseLocalDateTime("2017-07-01 09:00:00")
	param := &jsms.ScheduleSendParam{
		SendTime:   sendTime,
		Mobile:     "13812345678",
		SignID:     37582,
		TempID:     1250,
		TempParams: map[string]interface{}{"number": "741627"},
	}
	result, err := smsAPIv1.UpdateSchedule(context.Background(), scheduleID, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_UpdateSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> ScheduleID: %s", result.ScheduleID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_GetSchedule(t *testing.T) {
	scheduleID := "1a886e7c-a267-49e6-9970-0d396ca5bb1e"
	result, err := smsAPIv1.GetSchedule(context.Background(), scheduleID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! ScheduleID: %s, SendTime: %s, TempID: %d", result.ScheduleID, result.SendTime, result.TempID)
		for i, recipient := range result.Recipients {
			t.Logf(">>> %d. MsgID: %s, Mobile: %s, TempParams: %s", i+1, recipient.MsgID, recipient.Mobile, recipient.TempParams)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_DeleteSchedule(t *testing.T) {
	scheduleID := "1a886e7c-a267-49e6-9970-0d396ca5bb1e"
	result, err := smsAPIv1.DeleteSchedule(context.Background(), scheduleID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_DeleteSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_SendBatchSchedule(t *testing.T) {
	sendTime, _ := jiguang.ParseLocalDateTime("2017-07-01 09:00:00")
	param := &jsms.BatchScheduleSendParam{
		SendTime: sendTime,
		SignID:   37582,
		TempID:   1250,
		Tag:      "标签",
		Recipients: []jsms.BatchScheduleRecipientParam{
			{
				Mobile:     "13812345678",
				TempParams: map[string]interface{}{"number": "741627"},
			},
			{
				Mobile:     "18603050709",
				TempParams: map[string]interface{}{"number": "147721"},
			},
		},
	}
	result, err := smsAPIv1.SendBatchSchedule(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_SendBatchSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		failureRecipients, _ := json.Marshal(result.FailureRecipients)
		t.Logf("Success! SuccessCount: %d, FailureCount: %d, FailureRecipients: %s",
			*result.SuccessCount, *result.FailureCount, failureRecipients)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_UpdateBatchSchedule(t *testing.T) {
	scheduleID := "1a886e7c-a267-49e6-9970-0d396ca5bb1e"
	sendTime, _ := jiguang.ParseLocalDateTime("2017-07-01 09:00:00")
	param := &jsms.BatchScheduleSendParam{
		SendTime: sendTime,
		SignID:   37582,
		TempID:   1250,
		Tag:      "标签",
		Recipients: []jsms.BatchScheduleRecipientParam{
			{
				Mobile:     "13812345678",
				TempParams: map[string]interface{}{"number": "741627"},
			},
			{
				Mobile:     "18603050709",
				TempParams: map[string]interface{}{"number": "147721"},
			},
		},
	}
	result, err := smsAPIv1.UpdateBatchSchedule(context.Background(), scheduleID, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_UpdateBatchSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		failureRecipients, _ := json.Marshal(result.FailureRecipients)
		t.Logf("Success! SuccessCount: %d, FailureCount: %d, FailureRecipients: %s",
			*result.SuccessCount, *result.FailureCount, failureRecipients)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAPIv1_GetDevBalance(t *testing.T) {
	result, err := smsAPIv1.GetDevBalance(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetDevBalance: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> DevBalance: %d, DevVoice: %d, DevIndustry: %d, DevMarket: %d",
			*result.DevBalance, *result.DevVoice, *result.DevIndustry, *result.DevMarket)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_GetAppBalance(t *testing.T) {
	result, err := smsAPIv1.GetAppBalance(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetAppBalance: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> AppBalance: %d, AppVoice: %d, AppIndustry: %d, AppMarket: %d",
			*result.AppBalance, *result.AppVoice, *result.AppIndustry, *result.AppMarket)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func TestAPIv1_GetReport(t *testing.T) {
	result, err := smsAPIv1.GetReport(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetReport: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for i, report := range result.Data {
			t.Logf(">>> %d. MsgID: %s, Status: %d, Phone: %s, ReceiveTime: %s",
				i+1, report.MsgID, report.Status, report.Phone, report.ReceiveTime.Format("2006-01-02 15:04:05"))
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAPIv1_GetReply(t *testing.T) {
	result, err := smsAPIv1.GetReply(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetReply: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for i, reply := range result.Data {
			t.Logf(">>> %d. Phone: %s, ReplyTime: %s, Content: %s",
				i+1, reply.Phone, reply.ReplyTime.Format("2006-01-02 15:04:05"), reply.Content)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
