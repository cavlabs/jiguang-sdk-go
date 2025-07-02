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

	"github.com/cavlabs/jiguang-sdk-go/api/jpush/file"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

func TestFileAPIv3_UploadFileForAlias(t *testing.T) {
	param := &file.FileUploadParam{
		File: "~/Desktop/alias.txt",
		TTL:  jiguang.Int(360),
	}
	result, err := fileAPIv3.UploadFileForAlias(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("FileAPIv3_UploadFileForAlias: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> FileID: %s", result.FileID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestFileAPIv3_UploadFileForRegistrationID(t *testing.T) {
	param := &file.FileUploadParam{File: "~/Desktop/registrationId.txt"}
	result, err := fileAPIv3.UploadFileForRegistrationID(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("FileAPIv3_UploadFileForRegistrationID: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> FileID: %s", result.FileID)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestFileAPIv3_GetFiles(t *testing.T) {
	result, err := fileAPIv3.GetFiles(context.Background())
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("FileAPIv3_GetFiles: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Logf("Success! TotalCount: %d", result.TotalCount)
		for i, f := range result.Files {
			t.Logf(">>> %02d、Type: %s, FileID: %s, FileName: %s, CreateTime: %s, ExpireTime: %s",
				i+1, f.Type, f.FileID, f.FileName, f.CreateTime, f.ExpireTime)
		}
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestFileAPIv3_GetFile(t *testing.T) {
	fileID := "b266cd5c8544ba09b23733e3-6f82b892-206a-4b91-a8f4-1aaa4471c918"
	result, err := fileAPIv3.GetFile(context.Background(), fileID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("FileAPIv3_GetFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Type: %s, FileID: %s, FileName: %s, CreateTime: %s, ExpireTime: %s",
			result.Type, result.FileID, result.FileName, result.CreateTime, result.ExpireTime)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestFileAPIv3_DeleteFile(t *testing.T) {
	fileID := "b266cd5c8544ba09b23733e3-6f82b892-206a-4b91-a8f4-1aaa4471c918"
	result, err := fileAPIv3.DeleteFile(context.Background(), fileID)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("FileAPIv3_DeleteFile: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
