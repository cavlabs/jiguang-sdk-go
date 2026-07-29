// Copyright 2026 cavlabs/jiguang-sdk-go authors.
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
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/admin"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/device"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/file"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/gpush"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/greport"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/image"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/report"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/schedule"
)

// Push 默认指向 Push API v3；可通过 V3 显式访问。
//
// 注：历史原因下 push.APIv3 也嵌入了 File/Image/Schedule 方法；新代码推荐使用并列的
// Client.File / Image / Schedule 命名空间。
type Push struct {
	push.APIv3
	V3 push.APIv3
}

// Device 默认指向 Device API v3。
type Device struct {
	device.APIv3
	V3 device.APIv3
}

// File 默认指向 File API v3。
type File struct {
	file.APIv3
	V3 file.APIv3
}

// Image 默认指向 Image API v3。
type Image struct {
	image.APIv3
	V3 image.APIv3
}

// Schedule 默认指向 Schedule API v3。
type Schedule struct {
	schedule.APIv3
	V3 schedule.APIv3
}

// Report 默认指向 Report API v3。
type Report struct {
	report.APIv3
	V3 report.APIv3
}

// GPush 默认指向 Group Push API v3。
type GPush struct {
	gpush.APIv3
	V3 gpush.APIv3
}

// GReport 默认指向 Group Report API v3。
type GReport struct {
	greport.APIv3
	V3 greport.APIv3
}

// Admin 默认指向 Admin API v1。
type Admin struct {
	admin.APIv1
	V1 admin.APIv1
}

func newPush(o *options) Push {
	b := push.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasApp() {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	a, _ := b.Build()
	return Push{APIv3: a, V3: a}
}

func newDevice(o *options) Device {
	b := device.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasApp() {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	a, _ := b.Build()
	return Device{APIv3: a, V3: a}
}

func newFile(o *options) File {
	b := file.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	switch {
	case o.hasApp():
		b.SetAuthKey(o.appKey).SetAuthSecret(o.masterSecret)
	case o.hasDev():
		b.SetAuthKey(o.devKey).SetAuthSecret(o.devSecret)
	}
	a, _ := b.Build()
	return File{APIv3: a, V3: a}
}

func newImage(o *options) Image {
	b := image.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasApp() {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	a, _ := b.Build()
	return Image{APIv3: a, V3: a}
}

func newSchedule(o *options) Schedule {
	b := schedule.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasApp() {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	a, _ := b.Build()
	return Schedule{APIv3: a, V3: a}
}

func newReport(o *options) Report {
	b := report.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasApp() {
		b.SetAppKey(o.appKey).SetMasterSecret(o.masterSecret)
	}
	a, _ := b.Build()
	return Report{APIv3: a, V3: a}
}

func newGPush(o *options) GPush {
	b := gpush.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasGroup() {
		b.SetGroupKey(o.groupKey).SetGroupMasterSecret(o.groupMasterSecret)
	}
	if o.hasDev() {
		b.SetDevKey(o.devKey).SetDevSecret(o.devSecret)
	}
	a, _ := b.Build()
	return GPush{APIv3: a, V3: a}
}

func newGReport(o *options) GReport {
	b := greport.NewAPIv3Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasGroup() {
		b.SetGroupKey(o.groupKey).SetGroupMasterSecret(o.groupMasterSecret)
	}
	a, _ := b.Build()
	return GReport{APIv3: a, V3: a}
}

func newAdmin(o *options) Admin {
	b := admin.NewAPIv1Builder().SetClient(o.httpClient)
	if o.logger != nil {
		b.SetLogger(o.logger)
	}
	if o.httpLogSet {
		b.SetHttpLogLevel(o.httpLogLevel)
	}
	if o.hasDev() {
		b.SetDevKey(o.devKey).SetDevSecret(o.devSecret)
	}
	a, _ := b.Build()
	return Admin{APIv1: a, V1: a}
}
