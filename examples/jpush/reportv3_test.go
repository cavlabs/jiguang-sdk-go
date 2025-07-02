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

package jpush

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

func TestReportAPIv3_GetReceivedDetail(t *testing.T) {
	msgIDs := []string{"1613113584", "1229760629"}
	result, err := reportAPIv3.GetReceivedDetail(context.Background(), msgIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ReportAPIv3_GetReceivedDetail: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, detail := range result.ReceivedDetails {
			recvDetail, _ := json.Marshal(detail)
			t.Logf(">>> %s: %s", detail.MsgID, recvDetail)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestReportAPIv3_GetMessageStatus(t *testing.T) {
	msgID := "327640176"
	registrationIDs := []string{"1507bfd3a7c568d4761", "02078f0f1b8", "0207870a9b8"}
	// date := jiguang.BuildLocalDate(2024, 12, 25)
	// result, err := reportAPIv3.GetMessageStatus(context.Background(), msgID, registrationIDs, &date)
	result, err := reportAPIv3.GetMessageStatus(context.Background(), msgID, registrationIDs, nil)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ReportAPIv3_GetMessageStatus: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, registrationID := range registrationIDs {
			if status, exists := result.Status[registrationID]; exists {
				t.Logf(">>> %s: %s", registrationID, status)
			} else {
				t.Errorf(">>> %s: Not Found", registrationID)
			}
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestReportAPIv3_GetMessageDetail(t *testing.T) {
	msgIDs := []string{"269978303", "1613113584"}
	result, err := reportAPIv3.GetMessageDetail(context.Background(), msgIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ReportAPIv3_GetMessageDetail: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, detail := range result.MessageDetails {
			msgDetail, _ := json.Marshal(detail)
			t.Logf(">>> %s: %s", detail.MsgID, msgDetail)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestReportAPIv3_GetUserDetail(t *testing.T) {
	/*start := jiguang.UnitTime{
		Time:     time.Now().Add(-30 * 24 * time.Hour),
		TimeUnit: jiguang.TimeUnitHour,
	}*/
	start := jiguang.BuildLocalDate(2014, 6, 10).ToUnitTime()
	duration := 30
	result, err := reportAPIv3.GetUserDetail(context.Background(), start, duration)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ReportAPIv3_GetUserDetail: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! TimeUnit: %s, Start: %s, Duration: %d.", result.Start.TimeUnit, result.Start, result.Duration)
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

func uint64Ptr2str(uint64Ptr *uint64) string {
	if uint64Ptr == nil {
		return "-"
	}
	return strconv.FormatUint(*uint64Ptr, 10)
}
