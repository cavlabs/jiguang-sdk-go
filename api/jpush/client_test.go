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
	"context"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/v2/api"
	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push"
)

func TestNew_WithoutCredentials(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	_, err = client.Push.Send(context.Background(), &push.SendParam{})
	if err != api.ErrNilJPushPushAPIv3 {
		t.Fatalf("Push.Send error = %v, want %v", err, api.ErrNilJPushPushAPIv3)
	}
	_, err = client.Device.GetDevice(context.Background(), "reg-id")
	if err != api.ErrNilJPushDeviceAPIv3 {
		t.Fatalf("Device.GetDevice error = %v, want %v", err, api.ErrNilJPushDeviceAPIv3)
	}
	_, err = client.GPush.Send(context.Background(), &push.SendParam{})
	if err != api.ErrNilJPushGroupPushAPIv3 {
		t.Fatalf("GPush.Send error = %v, want %v", err, api.ErrNilJPushGroupPushAPIv3)
	}
	_, err = client.Admin.CreateApp(context.Background(), nil)
	if err != api.ErrNilJPushAdminAPIv1 {
		t.Fatalf("Admin.CreateApp error = %v, want %v", err, api.ErrNilJPushAdminAPIv1)
	}
}

func TestNew_WithApp(t *testing.T) {
	client, err := New(WithApp("app-key", "master-secret"))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if client.Push.APIv3 == nil || client.Push.V3 == nil {
		t.Fatal("expected Push API to be ready")
	}
	if client.Push.APIv3 != client.Push.V3 {
		t.Fatal("expected Push default and V3 to be the same instance")
	}
	_, err = client.GPush.Send(context.Background(), &push.SendParam{})
	if err != api.ErrNilJPushGroupPushAPIv3 {
		t.Fatalf("GPush.Send error = %v, want %v", err, api.ErrNilJPushGroupPushAPIv3)
	}
}
