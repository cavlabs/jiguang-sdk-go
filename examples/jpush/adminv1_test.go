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

	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/admin"
)

func TestAdminAPIv1_CreateApp(t *testing.T) {
	param := &admin.AppCreateParam{
		AppName:     "jiguang-test",
		PackageName: "cn.jiguang.test",
	}
	result, err := adminAPIv1.CreateApp(context.Background(), param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("AdminAPIv1_CreateApp: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> AppKey: %v", result.AppKey)
		t.Logf(">>> PackageName: %v", result.PackageName)
		if result.IsNewCreated == nil {
			t.Log(">>> IsNewCreated: <nil>")
		} else {
			t.Logf(">>> IsNewCreated: %t", *result.IsNewCreated)
		}
		t.Logf(">>> MasterSecret: %s", result.MasterSecret)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAdminAPIv1_DeleteApp(t *testing.T) {
	appKey := "fbd202ec599b20ad95da9907"
	result, err := adminAPIv1.DeleteApp(context.Background(), appKey)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("AdminAPIv1_DeleteApp: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Success: %s", result.Success)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}

func TestAdminAPIv1_UploadCertificate(t *testing.T) {
	appKey := "ffbb0932c267d938941e470b"

	devCertFile, err := admin.UploadFileFromPath("~/Desktop/dev-cert.pem")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	proCertFile, err := admin.UploadFileFromPath("~/Desktop/pro-cert.pem")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}

	param := &admin.CertificateUploadParam{
		DevCertificatePassword: "1234567890abcdefghijklmnopqrstuvwxyz",
		DevCertificateFile:     devCertFile,
		ProCertificatePassword: "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		ProCertificateFile:     proCertFile,
	}

	/*devCertificateFile, err := os.Open("/Users/z/Desktop/dev-cert.pem")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	defer devCertificateFile.Close()
	proCertificateFile, err := os.Open("/Users/z/Desktop/pro-cert.pem")
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	defer proCertificateFile.Close()

	devUploadFile, err := admin.UploadFileFromFile(devCertificateFile)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	proUploadFile, err := admin.UploadFileFromFile(proCertificateFile)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	param := &admin.CertificateUploadParam{
		DevCertificatePassword: "1234567890abcdefghijklmnopqrstuvwxyz",
		DevCertificateFile:     devUploadFile,
		ProCertificatePassword: "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		ProCertificateFile:     proUploadFile,
	}*/

	result, err := adminAPIv1.UploadCertificate(context.Background(), appKey, param)
	if err != nil {
		t.Fatalf("Failed! Error: %s", err)
	}
	t.Logf("AdminAPIv1_UploadCertificate: StatusCode: %d, RateLimit: %d, RateRemaining: %d, RateReset: %d.",
		result.StatusCode, result.RateLimit(), result.RateRemaining(), result.RateReset())
	if result.IsSuccess() {
		t.Log("Success!")
		t.Logf(">>> Success: %s", result.Success)
	} else {
		t.Errorf("Failed! Error: %s", result.Error)
	}
}
