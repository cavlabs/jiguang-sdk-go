/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
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

package jums

import (
	"context"
	"strings"
	"testing"

	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/api/jums"
	"github.com/calvinit/jiguang-sdk-go/api/jums/audience"
	"github.com/calvinit/jiguang-sdk-go/api/jums/message"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestAPIv1_BroadcastSend(t *testing.T) {
	param := &jums.BroadcastSendParam{
		Message: mockMessage(),
		RuleID:  1001,
		Option: &message.Option{
			SendNo:   "test",
			Owner:    "admin",
			BlackID:  1234,
			Priority: 2,
		},
		Callback: &message.Callback{
			URL: "https://www.jiguang.cn/ums-portal",
			Params: map[string]interface{}{
				"name": "joe",
				"age":  26,
			},
		},
	}
	result, err := umsAPIv1.BroadcastSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_BroadcastSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SendNo: %s, MsgID: %s", result.SendNo, result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_Send(t *testing.T) {
	param := &jums.SendParam{
		Audience: mockAudience(),
		Message:  mockMessage(),
		RuleID:   1001,
		Option: &message.Option{
			SendNo:   "test",
			Owner:    "admin",
			BlackID:  1234,
			Priority: 2,
		},
		Callback: &message.Callback{
			URL: "https://www.jiguang.cn/ums-portal",
			Params: map[string]interface{}{
				"name": "joe",
				"age":  26,
			},
		},
	}
	result, err := umsAPIv1.Send(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_Send: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SendNo: %s, MsgID: %s", result.SendNo, result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_TemplateBroadcastSend(t *testing.T) {
	param := &jums.TemplateBroadcastSendParam{
		TemplateID: 10001,
		TemplateParams: map[string]interface{}{
			"user": "xxx",
			"url":  "xxx",
		},
		AppParams: &message.AppParams{
			Platform:       platform.IOS,
			TimeToLive:     jiguang.Int64(9999),
			ApnsProduction: jiguang.Bool(true),
		},
		RuleID: 1001,
		Option: &message.Option{
			SendNo:   "test",
			Owner:    "admin",
			BlackID:  1234,
			Priority: 2,
		},
		Callback: &message.Callback{
			URL: "https://www.jiguang.cn/ums-portal",
			Params: map[string]interface{}{
				"name": "joe",
				"age":  26,
			},
		},
	}
	result, err := umsAPIv1.TemplateBroadcastSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_TemplateBroadcastSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SendNo: %s, MsgID: %s", result.SendNo, result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_TemplateSend(t *testing.T) {
	param := &jums.TemplateSendParam{
		Audience: &audience.Audience{
			Tags:     []string{"tag1", "tag2"},
			UserIDs:  []string{"user1", "user2"},
			Segments: []string{"1000"},
		},
		TemplateID: 10001,
		TemplateParams: map[string]interface{}{
			"user": "xxx",
			"url":  "xxx",
		},
		AppParams: &message.AppParams{
			Platform:       platform.IOS,
			TimeToLive:     jiguang.Int64(9999),
			ApnsProduction: jiguang.Bool(true),
		},
		RuleID: 1001,
		Option: &message.Option{
			SendNo:   "test",
			Owner:    "admin",
			BlackID:  1234,
			Priority: 2,
		},
	}
	result, err := umsAPIv1.TemplateSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_TemplateSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> SendNo: %s, MsgID: %s", result.SendNo, result.MsgID)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_Retract(t *testing.T) {
	msgID := "v:1t:1603707455p:27133n:2"
	result, err := umsAPIv1.Retract(context.Background(), msgID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_Retract: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_BatchAddOrUpdateUsers(t *testing.T) {
	param := []jums.UsersBatchAddOrUpdateParam{
		{
			UserID: "具体userid",
			Add: &jums.UsersBatchOpts{
				Tag: map[string][]string{
					"具体channel_key": {"具体tag_value"},
				},
				APP: map[string][]string{
					"具体sender_code": {"具体app_regid"},
				},
				WechatwkLinkedCorp: []string{
					"具体wechatwk_linkedcorp_userid1",
					"具体wechatwk_linkedcorp_userid2",
				},
			},
			Set: &jums.UsersBatchOpts{
				Tag: map[string][]string{
					"具体channel_key": {"具体tag_value"},
				},
				Phone: map[string]string{
					"具体channel_key": "具体phone_number",
				},
				Email: map[string]string{
					"具体channel_key": "具体email_value",
				},
				APP: map[string][]string{
					"具体sender_code": {"具体app_regid"},
				},
				Wechatoa: map[string]string{
					"具体sender_code": "具体wechatoa_openid",
				},
				Wechatmp: map[string]string{
					"具体sender_code": "具体wechatmp_openid",
				},
				AlipayLife: map[string]string{
					"具体sender_code": "具体alipaylife_userid",
				},
				DingtalkCC: "具体dingtalk_userid",
				Wechatwk:   "具体wechatwk_userid",
				WechatwkLinkedCorp: []string{
					"具体wechatwk_linkedcorp_userid1",
					"具体wechatwk_linkedcorp_userid2",
				},
			},
			Del: &jums.UsersBatchOpts{
				Tag: map[string][]string{
					"具体channel_key": {"具体tag_value"},
				},
				Phone: map[string]string{
					"具体channel_key": "具体phone_number",
				},
				Email: map[string]string{
					"具体channel_key": "具体email_value",
				},
				APP: map[string][]string{
					"具体sender_code": {"具体app_regid"},
				},
				Wechatoa: map[string]string{
					"具体sender_code": "具体wechatoa_openid",
				},
				Wechatmp: map[string]string{
					"具体sender_code": "具体wechatmp_openid",
				},
				AlipayLife: map[string]string{
					"具体sender_code": "具体alipaylife_userid",
				},
				DingtalkCC: "具体dingtalk_userid",
				Wechatwk:   "具体wechatwk_userid",
				WechatwkLinkedCorp: []string{
					"具体wechatwk_linkedcorp_userid1",
					"具体wechatwk_linkedcorp_userid2",
				},
			},
		},
	}
	result, err := umsAPIv1.BatchAddOrUpdateUsers(context.Background(), param, true)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_BatchAddOrUpdateUsers: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		if data, ok := result.Data.(*jums.UsersBatchAddOrUpdateData); ok {
			t.Logf(">>> 1. Success: %s", strings.Join(data.Success, ", "))
			for i, failData := range data.Fail {
				t.Logf(">>> 2. Fail_%d: (%s, %s - %s)", i+1, failData.UserID, failData.ErrCode, failData.ErrMsg)
			}
		} else {
			t.Log("WARN!!! Invalid Result Data.")
		}
	} else {
		t.Errorf("Failed! Error: %s, Data: %s", result.CodeError, result.Data.(string))
	}
}

func TestAPIv1_BatchDeleteUsers(t *testing.T) {
	userIDs := []string{"userid1", "userid2"}
	result, err := umsAPIv1.BatchDeleteUsers(context.Background(), userIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_BatchDeleteUsers: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		if data, ok := result.Data.(*jums.UsersBatchDeleteData); ok {
			t.Logf(">>> 1. Success: %s", strings.Join(data.Success, ", "))
			for i, failData := range data.Fail {
				t.Logf(">>> 2. Fail_%d: (%s, %s - %s)", i+1, failData.UserID, failData.ErrCode, failData.ErrMsg)
			}
		} else {
			t.Log("WARN!!! Invalid Result Data.")
		}
	} else {
		t.Errorf("Failed! Error: %s, Data: %s", result.CodeError, result.Data.(string))
	}
}

func TestAPIv1_UploadMaterial(t *testing.T) {
	param := &jums.MaterialUploadParam{
		Type:       "attachment",
		File:       "D:/ums.pdf",
		TimeToLive: 169,
	}
	result, err := umsAPIv1.UploadMaterial(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_UploadMaterial: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Material URL: %s", result.URL)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}

func TestAPIv1_GetChannelToken(t *testing.T) {
	channelType := "wechatoa"
	result, err := umsAPIv1.GetChannelToken(context.Background(), channelType)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("APIv1_GetChannelToken: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Channel Token: %s, ExpiresIn: %d", result.Token, result.ExpiresIn)
	} else {
		t.Errorf("Failed! Error: %s", result.CodeError)
	}
}
