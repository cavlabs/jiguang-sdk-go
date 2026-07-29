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

package jums

import (
	"context"
	"testing"

	"github.com/cavlabs/jiguang-sdk-go/v2/api"
)

func TestNew_WithoutCredentials(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	_, err = client.Send(context.Background(), nil)
	if err != api.ErrNilJUmsAPIv1 {
		t.Fatalf("Send error = %v, want %v", err, api.ErrNilJUmsAPIv1)
	}
}

func TestNew_WithChannel(t *testing.T) {
	client, err := New(WithChannel("channel-key", "master-secret"))
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	if client.APIv1 == nil || client.V1 == nil {
		t.Fatal("expected JUMS API to be ready")
	}
	if client.APIv1 != client.V1 {
		t.Fatal("expected default and V1 to be the same instance")
	}
}
