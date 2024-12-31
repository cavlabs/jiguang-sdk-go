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

	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestGroupReportAPIv3_GetMessageDetail(t *testing.T) {
	groupMsgIDs := []string{"bsp54sn8bpn8hc3etq40"}
	result, err := greportAPIv3.GetMessageDetail(context.Background(), groupMsgIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("GroupReportAPIv3_GetMessageDetail: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, detail := range result.MessageDetails {
			msgDetail, _ := json.Marshal(detail)
			t.Logf(">>> %s: %s", detail.GroupMsgID, msgDetail)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestGroupReportAPIv3_GetUserDetail(t *testing.T) {
	/*start := jiguang.UnitTime{
		Time:     time.Now().Add(-30 * 24 * time.Hour),
		TimeUnit: jiguang.TimeUnitHour,
	}*/
	start := jiguang.BuildLocalDate(2020, 8, 6).ToUnitTime()
	duration := 2
	result, err := greportAPIv3.GetUserDetail(context.Background(), start, duration)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("GroupReportAPIv3_GetUserDetail: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! TimeUnit: %s, Start: %v, Duration: %d.", result.Start.TimeUnit, result.Start, result.Duration)
		for _, item := range result.Items {
			if android := item.Android; android != nil {
				t.Logf(">>> %s: Android -> New: %s, Online: %s, Active: %s.",
					item.Time, uint64Ptr2str(android.New), uint64Ptr2str(android.Online), uint64Ptr2str(android.Active))
			}
			if ios := item.IOS; ios != nil {
				t.Logf(">>> %s: IOS -> New: %s, Online: %s, Active: %s.",
					item.Time, uint64Ptr2str(ios.New), uint64Ptr2str(ios.Online), uint64Ptr2str(ios.Active))
			}
			if quickApp := item.QuickApp; quickApp != nil {
				t.Logf(">>> %s: QuickApp -> New: %s, Online: %s, Active: %s.",
					item.Time, uint64Ptr2str(quickApp.New), uint64Ptr2str(quickApp.Online), uint64Ptr2str(quickApp.Active))
			}
			if hmos := item.HMOS; hmos != nil {
				t.Logf(">>> %s: HMOS -> New: %s, Online: %s, Active: %s.",
					item.Time, uint64Ptr2str(hmos.New), uint64Ptr2str(hmos.Online), uint64Ptr2str(hmos.Active))
			}
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
