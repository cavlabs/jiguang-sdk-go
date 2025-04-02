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

	"github.com/calvinit/jiguang-sdk-go/api/jpush/device"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func TestDeviceAPIv3_GetDevice(t *testing.T) {
	registrationID := "1507bfd3f6f7aaf8781"
	result, err := deviceAPIv3.GetDevice(context.Background(), registrationID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_GetDevice: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Tags: %v", result.Tags)
		t.Logf(">>> Alias: %v", result.Alias)
		t.Logf(">>> Mobile: %v", result.Mobile)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_SetDevice(t *testing.T) {
	registrationID, alias, mobile := "1507bfd3f6f7aaf8781", "13111111111", "13111111111"
	param := &device.DeviceSetParam{
		Tags: &device.TagsForDeviceSetParam{
			Add:    []string{"13111111111", "13222222222"},
			Remove: []string{"13333333333", "13444444444"},
		},
		Alias:  &alias,
		Mobile: &mobile,
	}
	result, err := deviceAPIv3.SetDevice(context.Background(), registrationID, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_SetDevice: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}

	// 「空字符串」在这里有特殊含义，比如参数的 Tags 字段是空字符串时，表示清空所有的 tags，如下示例：
	clearAll := ""
	param = &device.DeviceSetParam{
		Tags:   jiguang.String(""), // or &clearAll, only for demonstration here.
		Alias:  &clearAll,
		Mobile: &mobile,
	}
	clearAllParam, _ := json.Marshal(param)
	t.Logf("Demo of CLEAR ALL param: %s", clearAllParam)
}

func TestDeviceAPIv3_ClearDevice(t *testing.T) {
	registrationID := "1507bfd3f6f7aaf8781"
	result, err := deviceAPIv3.ClearDeviceTags(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceAlias(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceMobile(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceTagsAndAlias(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceTagsAndMobile(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceAliasAndMobile(context.Background(), registrationID)
	// result, err := deviceAPIv3.ClearDeviceAll(context.Background(), registrationID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_ClearDevice*: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_GetDeviceStatus(t *testing.T) {
	registrationIDs := []string{"1104a89793af2cfc030", "1507bfd3f6f7aaf8781", "4n6mwxh23l72l0i1l32", "s6ua4j0jwdj3dsrax0w", "hgahpgu2aehv933pli1", "hhg2kmsjxqzqf5d1jt5"}
	result, err := deviceAPIv3.GetDeviceStatus(context.Background(), registrationIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_GetDeviceStatus: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, registrationID := range registrationIDs {
			if r, exists := result.Result[registrationID]; exists {
				t.Logf(">>> RegistrationID(%s) => Online: %t, LastOnlineTime: %s", registrationID, r.Online, r.LastOnlineTime)
			} else {
				t.Logf(">>> RegistrationID(%s) => ERROR: Invalid or not belongs to this appKey.", registrationID)
			}
		}
	} else {
		t.Errorf("Failed! Error: %v", result.Error)
	}
}

func TestDeviceAPIv3_GetTags(t *testing.T) {
	result, err := deviceAPIv3.GetTags(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_GetTags: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Tags: %s", result.Tags)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_GetTag(t *testing.T) {
	tag := "tag1"
	registrationID := "registration_id1"
	result, err := deviceAPIv3.GetTag(context.Background(), tag, registrationID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_GetTag: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Result: %t", *result.Result)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_SetTag(t *testing.T) {
	tag := "tag1"
	adds := []string{"registration_id1", "registration_id2"}
	removes := []string{"registration_id3", "registration_id4"}
	result, err := deviceAPIv3.SetTag(context.Background(), tag, adds, removes)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_SetTag: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error! StatusCode: %d", result.StatusCode)
		t.Errorf(">>> Code: %d", result.Error.Code)
		t.Errorf(">>> IllegalRIDs: %s", result.Error.IllegalRIDs)
	}
}

func TestDeviceAPIv3_DeleteTag(t *testing.T) {
	tag := "tag1"
	// plats := []platform.Platform{platform.Android, platform.IOS, platform.QuickApp, platform.HMOS}
	// result, err := deviceAPIv3.DeleteTag(context.Background(), tag, plats...)
	result, err := deviceAPIv3.DeleteTag(context.Background(), tag, platform.Android, platform.IOS)
	// result, err := deviceAPIv3.DeleteTag(context.Background(), tag)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_DeleteTag: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_GetAlias(t *testing.T) {
	alias := "alias1"
	// plats := []platform.Platform{platform.Android, platform.IOS, platform.QuickApp, platform.HMOS}
	// result, err := deviceAPIv3.GetAlias(context.Background(), alias, plats...)
	result, err := deviceAPIv3.GetAlias(context.Background(), alias, platform.Android, platform.IOS)
	// result, err := deviceAPIv3.GetAlias(context.Background(), alias)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_GetAlias: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		for _, reg := range result.Data {
			t.Logf(">>> RegistrationID(%s): Platform: %s, LastOnlineDate: %s", reg.RegistrationID, reg.Platform, reg.LastOnlineDate)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_DeleteAlias(t *testing.T) {
	alias := "alias1"
	// plats := []platform.Platform{platform.Android, platform.IOS, platform.QuickApp, platform.HMOS}
	// result, err := deviceAPIv3.DeleteAlias(context.Background(), alias, plats...)
	result, err := deviceAPIv3.DeleteAlias(context.Background(), alias, platform.Android, platform.IOS)
	// result, err := deviceAPIv3.DeleteAlias(context.Background(), alias)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_DeleteAlias: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_DeleteAliases(t *testing.T) {
	alias := "alias1"
	registrationIDs := []string{"registration_id1", "registration_id2"}
	result, err := deviceAPIv3.DeleteAliases(context.Background(), alias, registrationIDs)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_DeleteAliases: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error! StatusCode: %d", result.StatusCode)
		t.Errorf(">>> Code: %d", result.Error.Code)
		t.Errorf(">>> IllegalRIDs: %s", result.Error.IllegalRIDs)
	}
}

func TestDeviceAPIv3_AddTestDevice(t *testing.T) {
	deviceName, registrationID := "device_name_1", "registration_id_1"
	param := &device.TestDeviceAddParam{
		DeviceName:     deviceName,
		RegistrationID: registrationID,
	}
	result, err := deviceAPIv3.AddTestDevice(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_AddTestDevice: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_UpdateTestDevice(t *testing.T) {
	deviceName, registrationID := "device_name_1", "registration_id_1"
	param := &device.TestDeviceUpdateParam{
		DeviceName:     deviceName,
		RegistrationID: registrationID,
	}
	result, err := deviceAPIv3.UpdateTestDevice(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_UpdateTestDevice: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_DeleteTestDevice(t *testing.T) {
	registrationID := "registration_id_1"
	result, err := deviceAPIv3.DeleteTestDevice(context.Background(), registrationID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_DeleteTestDevice: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestDeviceAPIv3_ListTestDevices(t *testing.T) {
	page, pageSize := 1, 10
	deviceName, registrationID := "", ""
	result, err := deviceAPIv3.ListTestDevices(context.Background(), page, pageSize, deviceName, registrationID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("DeviceAPIv3_ListTestDevices: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Total: %d", result.Total)
		t.Logf(">>> Page: %d", result.Page)
		t.Logf(">>> PageSize: %d", result.PageSize)
		for i, detail := range result.Detail {
			t.Logf(">>> %02d、DeviceName: %s, DeviceModel: %s, RegistrationID: %s, RegistrationTime: %s, CreateTime: %s",
				i+1, detail.DeviceName, detail.DeviceModel, detail.RegistrationID, detail.RegistrationTime, detail.CreateTime)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
