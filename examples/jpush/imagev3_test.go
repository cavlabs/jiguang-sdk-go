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
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/api/jpush/image"
)

func TestImageAPIv3_AddImageByUrl(t *testing.T) {
	param := &image.AddByUrlParam{
		ImageType:       image.BigImage,
		JiguangImageUrl: "https://img.jiguang.cn/jiguang/public/img/f237811.png",
	}
	result, err := imageAPIv3.AddImageByUrl(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ImageAPIv3_AddImageByUrl: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MediaID: %s, JiguangImageUrl: %s", result.MediaID, result.JiguangImageUrl)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestImageAPIv3_UpdateImageByUrl(t *testing.T) {
	mediaID := "jgmedia-1-c9f54726-6a06-4ed1-842d-81e32383ee5c"
	param := &image.UpdateByUrlParam{
		JiguangImageUrl: "https://img.jiguang.cn/jiguang/public/img/c866bd2.png",
	}
	result, err := imageAPIv3.UpdateImageByUrl(context.Background(), mediaID, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ImageAPIv3_UpdateImageByUrl: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MediaID: %s, JiguangImageUrl: %s", result.MediaID, result.JiguangImageUrl)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestImageAPIv3_AddImageByFile(t *testing.T) {
	param := &image.AddByFileParam{
		ImageType:     image.BigImage,
		OppoImageFile: "~/Desktop/xx.jpeg",
	}
	result, err := imageAPIv3.AddImageByFile(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ImageAPIv3_AddImageByFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MediaID: %s, OppoImageUrl: %s", result.MediaID, result.OppoImageUrl)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestImageAPIv3_UpdateImageByFile(t *testing.T) {
	mediaID := "jgmedia-1-c20d4b1f-e821-430d-b651-91c0c4bf1f60"
	param := &image.UpdateByFileParam{
		OppoImageFile: "~/Desktop/yy.jpg",
	}
	result, err := imageAPIv3.UpdateImageByFile(context.Background(), mediaID, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("ImageAPIv3_UpdateImageByFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> MediaID: %s, OppoImageUrl: %s", result.MediaID, result.OppoImageUrl)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
