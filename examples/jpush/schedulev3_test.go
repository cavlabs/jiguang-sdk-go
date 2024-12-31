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
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/schedule"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestScheduleAPIv3_ScheduleSend(t *testing.T) {
	param := &schedule.SendParam{}

	// 推送唯一标识 (CID)。注：Push 中的 CID 即便被指定也会被忽略！
	param.CID = "7103a4c428a0b98974ec1849-711161d4-5f17-4d2f-b855-5e5a8909b26e"
	/*cidResult, _ := scheduleAPIv3.GetCidForSchedulePush(context.Background(), 1)
	if cidResult.IsSuccess() {
		param.CID = cidResult.CidList[0]
	} else {
		t.Fatalf("Failed! Error: %s", cidResult.Error)
	}*/
	param.Name = "定时普通推送示例"
	param.Enabled = true

	// 1. 定时任务，单次触发执行。
	trigger := &schedule.Trigger{}
	singleTime, _ := jiguang.ParseLocalDateTime("2014-09-17 12:00:00")
	trigger.Single = &schedule.Single{Time: singleTime}
	param.Trigger = trigger

	singlePush := &schedule.Push{}
	singlePush.Platform = platform.All       // 所有平台，all
	singlePush.Audience = push.BroadcastAuds // 广播推送，all
	singlePush.Notification = &push.Notification{Alert: "Hello, JPush!"}
	singlePush.CustomMessage = &push.CustomMessage{Content: "Message!"}
	singlePush.Options = &push.Options{TimeToLive: jiguang.Int64(60)}
	param.Push = singlePush

	singleResult, err := scheduleAPIv3.ScheduleSend(context.Background(), param)
	// singleResult, err := scheduleAPIv3.CustomScheduleSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_ScheduleSend(Single): StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		singleResult.StatusCode, singleResult.RateLimit(), singleResult.RateRemaining(), singleResult.RateReset())
	if singleResult.IsSuccess() {
		t.Log("ScheduleSend(Single) Success!")
		t.Logf(">>> RawBody: %s", singleResult.RawBody)
		t.Logf(">>> ScheduleID: %s", singleResult.ScheduleID)
		t.Logf(">>> Name: %s", singleResult.Name)
	} else {
		t.Errorf("Failed! Error: %s", singleResult.Error)
	}

	// 2. 定期任务，周期触发执行。
	param.Name = "定期文件推送示例"
	startTime, _ := jiguang.ParseLocalDateTime("2014-09-17 12:00:00")
	endTime := jiguang.BuildLocalDateTime(2014, 9, 18, 12, 0, 0)
	time, _ := jiguang.ParseLocalTime("12:00:00")
	periodical := &schedule.Periodical{
		StartTime: startTime,
		EndTime:   endTime,
		Time:      time,
		TimeUnit:  jiguang.TimeUnitWeek,
		Frequency: 1,
		Point:     []string{"WED", "FRI"},
	}
	param.Trigger = &schedule.Trigger{Periodical: periodical}

	periodicalPush := &schedule.Push{}
	periodicalPush.Platform = platform.All
	periodicalPush.Audience = &push.Audience{File: &push.FileAudience{FileID: "004f2b2c3d8fc649fa60ad2c-15563487-577d-4bb0-806b-56f1a19e06c1"}}
	periodicalPush.Notification = &push.Notification{Alert: "Hello, JPush!"}
	periodicalPush.CustomMessage = &push.CustomMessage{Content: "Hello, JPush!"}
	periodicalPush.Options = &push.Options{TimeToLive: jiguang.Int64(60), ApnsProduction: jiguang.Bool(false)}
	param.Push = periodicalPush

	periodicalResult, err := scheduleAPIv3.ScheduleSend(context.Background(), param)
	// periodicalResult, err := scheduleAPIv3.CustomScheduleSend(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_ScheduleSend(Periodical): StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		periodicalResult.StatusCode, periodicalResult.RateLimit(), periodicalResult.RateRemaining(), periodicalResult.RateReset())
	if periodicalResult.IsSuccess() {
		t.Log("ScheduleSend(Periodical) Success!")
		t.Logf(">>> RawBody: %s", periodicalResult.RawBody)
		t.Logf(">>> ScheduleID: %s", periodicalResult.ScheduleID)
		t.Logf(">>> Name: %s", periodicalResult.Name)
	} else {
		t.Errorf("Failed! Error: %s", periodicalResult.Error)
	}
}

func TestScheduleAPIv3_ScheduleTemplateSend(t *testing.T) {
	// 广播推送
	templateID := "c65b34b96512"
	templateParams := []schedule.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是广播标题", "content": "这是广播内容"},
			Audience: push.BroadcastAuds, // 广播推送，all
			TraceID:  "8de06fh-8djfgg",
		},
	}
	scheduleName := "定时模板推送示例（广播推送）"
	trigger := &schedule.Trigger{
		Periodical: &schedule.Periodical{
			StartTime: jiguang.BuildLocalDateTime(2023, 10, 19, 12, 0, 0),
			EndTime:   jiguang.BuildLocalDateTime(2023, 11, 19, 18, 30, 0),
			Time:      jiguang.BuildLocalTime(18, 5, 0),
			TimeUnit:  jiguang.TimeUnitDay,
			Frequency: 1,
		},
	}

	// Registration ID 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []schedule.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第⼀个regid的标题", "content": "这是第⼀个regid的内容"},
			Audience: &push.Audience{RegistrationIDs: []string{"13065ffa4ee8411496f"}},
			TraceID:  "98ode06fh-8dgg",
		},
		{
			Keys:     map[string]string{"title": "这是推送给第二个regid的标题", "content": "这是第二个regid的内容"},
			Audience: &push.Audience{RegistrationIDs: []string{"8914afdsa31"}},
			TraceID:  "8de06fh-8djfgg",
		},
	}
	scheduleName := "定时模板推送示例（Registration ID 方式推送）"
	trigger := &schedule.Trigger{
		Periodical: &schedule.Periodical{
			StartTime: jiguang.BuildLocalDateTime(2023, 10, 19, 12, 0, 0),
			EndTime:   jiguang.BuildLocalDateTime(2023, 11, 19, 18, 30, 0),
			Time:      jiguang.BuildLocalTime(18, 5, 0),
			TimeUnit:  jiguang.TimeUnitDay,
			Frequency: 1,
		},
	}*/

	// Alias 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []schedule.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第⼀个alias的标题", "content": "这是第⼀个alias的内容"},
			Audience: &push.Audience{Aliases: []string{"4312kjkla31"}},
			TraceID:  "98ode06fh-8dgg",
		},
		{
			Keys:     map[string]string{"title": "这是推送给第二个alias的标题", "content": "这是第二个alias的内容"},
			Audience: &push.Audience{Aliases: []string{"8914sa31"}},
			TraceID:  "8de06fh-8djfgg",
		},
	}
	scheduleName := "定时模板推送示例（Alias 方式推送）"
	trigger := &schedule.Trigger{
		Periodical: &schedule.Periodical{
			StartTime: jiguang.BuildLocalDateTime(2023, 10, 19, 12, 0, 0),
			EndTime:   jiguang.BuildLocalDateTime(2023, 11, 19, 18, 30, 0),
			Time:      jiguang.BuildLocalTime(18, 5, 0),
			TimeUnit:  jiguang.TimeUnitDay,
			Frequency: 1,
		},
	}*/

	// Tag 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []schedule.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第1个tag的标题", "content": "这是第1个tag的内容"},
			Audience: &push.Audience{Tags: []string{"tag1"}},
			TraceID:  "8de06fh-8djfgg",
		},
	}
	scheduleName := "定时模板推送示例（Tag 方式推送）"
	trigger := &schedule.Trigger{
		Periodical: &schedule.Periodical{
			StartTime: jiguang.BuildLocalDateTime(2023, 10, 19, 12, 0, 0),
			EndTime:   jiguang.BuildLocalDateTime(2023, 11, 19, 18, 30, 0),
			Time:      jiguang.BuildLocalTime(18, 5, 0),
			TimeUnit:  jiguang.TimeUnitDay,
			Frequency: 1,
		},
	}*/

	// Segment 方式推送
	/*templateID := "c65b34b96512"
	templateParams := []schedule.TemplateParam{
		{
			Keys:     map[string]string{"title": "这是推送给第1个segment的标题", "content": "这是第1个segment的内容"},
			Audience: &push.Audience{Segments: []string{"segmentid001"}},
			TraceID:  "8de06fh-8djfgg",
		},
	}
	scheduleName := "定时模板推送示例（Segment 方式推送）"
	trigger := &schedule.Trigger{
		Periodical: &schedule.Periodical{
			StartTime: jiguang.BuildLocalDateTime(2023, 10, 19, 12, 0, 0),
			EndTime:   jiguang.BuildLocalDateTime(2023, 11, 19, 18, 30, 0),
			Time:      jiguang.BuildLocalTime(18, 5, 0),
			TimeUnit:  jiguang.TimeUnitDay,
			Frequency: 1,
		},
	}*/

	result, err := scheduleAPIv3.ScheduleTemplateSend(context.Background(), templateID, templateParams, scheduleName, trigger)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_ScheduleTemplateSend: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		if result.Data != nil && len(result.Data.ScheduleList) > 0 {
			for i, v := range result.Data.ScheduleList {
				if v.IsSuccess() {
					t.Logf(">>> ScheduleList[%d]: Success! ScheduleID: %s, TraceID: %s", i, v.ScheduleID, v.TraceID)
				} else {
					t.Logf(">>> ScheduleList[%d]: Failed! ScheduleID: %s, TraceID: %s, Error: %s", i, v.ScheduleID, v.TraceID, v.Error)
				}
			}
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_GetCidForSchedulePush(t *testing.T) {
	result, err := scheduleAPIv3.GetCidForSchedulePush(context.Background(), 3)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_GetCidForSchedulePush: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> CidList: %s", result.CidList)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_GetSchedules(t *testing.T) {
	page := 1
	result, err := scheduleAPIv3.GetSchedules(context.Background(), page)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_GetSchedules(Page-%d): StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		page, result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! TotalCount: %d, TotalPages: %d, Page: %d", result.TotalCount, result.TotalPages, result.Page)
		for i, v := range result.Schedules {
			_trigger, _ := json.Marshal(v.Trigger)
			_push, _ := json.Marshal(v.Push)
			t.Logf(">>> Schedules[%02d]: ScheduleID: %s, Name: %s, Enabled: %t, Trigger: %s, Push: %s",
				i+1, v.ScheduleID, v.Name, v.Enabled, _trigger, _push)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_GetSchedule(t *testing.T) {
	result, err := scheduleAPIv3.GetSchedule(context.Background(), "0eac1b80-c2ac-4b69-948b-c65b34b96512")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_GetSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		_trigger, _ := json.Marshal(result.Trigger)
		_push, _ := json.Marshal(result.Push)
		t.Logf(">>> Schedule: ScheduleID: %s, Name: %s, Enabled: %t, Trigger: %s, Push: %s",
			result.ScheduleID, result.Name, result.Enabled, _trigger, _push)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_GetScheduleMsgIDs(t *testing.T) {
	result, err := scheduleAPIv3.GetScheduleMsgIDs(context.Background(), "0eac1b80-c2ac-4b69-948b-c65b34b96512")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_GetScheduleMsgIDs: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! Count: %d, MsgIDs: %s", *result.Count, result.MsgIDs)
		for i, v := range result.MsgIDs {
			r, ok := schedule.TryParseScheduleMsgIDFromString(v)
			if ok {
				t.Logf(">>> MsgIDs[%d]: MsgID: %s, NeedRetry: %t, Timestamp: %d, IsSuccess: %t",
					i, r.MsgID, r.NeedRetry, *r.Timestamp, r.IsSuccess())
			} else {
				t.Logf(">>> MsgIDs[%d]: MsgID: %s", i, v)
			}
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_UpdateSchedule(t *testing.T) {
	param := &schedule.UpdateParam{
		Name:    "定时推送示例（更新）",
		Enabled: jiguang.Bool(true),
		Trigger: &schedule.Trigger{
			Single: &schedule.Single{Time: jiguang.BuildLocalDateTime(2024, 1, 1, 0, 0, 0)},
		},
		Push: &schedule.Push{
			Platform:      []platform.Platform{platform.Android, platform.IOS},
			Audience:      push.BroadcastAuds,
			Notification:  &push.Notification{Alert: "Hi, JPush!"},
			CustomMessage: &push.CustomMessage{Content: "Hi, JPush!"},
			Options:       &push.Options{TimeToLive: jiguang.Int64(360), ApnsProduction: jiguang.Bool(true)},
		},
	}

	result, err := scheduleAPIv3.UpdateSchedule(context.Background(), "0eac1b80-c2ac-4b69-948b-c65b34b96512", param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_UpdateSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> RawBody: %s", result.RawBody)
		_trigger, _ := json.Marshal(result.Trigger)
		_push, _ := json.Marshal(result.Push)
		t.Logf(">>> Schedule: Name: %s, Enabled: %t, Trigger: %s, Push: %s", result.Name, result.Enabled, _trigger, _push)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestScheduleAPIv3_DeleteSchedule(t *testing.T) {
	result, err := scheduleAPIv3.DeleteSchedule(context.Background(), "0eac1b80-c2ac-4b69-948b-c65b34b96512")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ScheduleAPIv3_DeleteSchedule: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
