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

package api

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUploadFileFromBytesAndWriteMultipart(t *testing.T) {
	uf := UploadFileFromBytes("hello.txt", []byte("hello jiguang"))
	if uf.Name != "hello.txt" || uf.Size != 13 {
		t.Fatalf("unexpected upload file: name=%q size=%d", uf.Name, uf.Size)
	}

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	req := NewUploadRequest(
		WithField("type", "attachment"),
		WithFile("file", uf),
	)
	if err := req.WriteMultipart(w); err != nil {
		t.Fatalf("WriteMultipart failed: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("close writer failed: %v", err)
	}

	body := buf.String()
	if !strings.Contains(body, `name="type"`) || !strings.Contains(body, "attachment") {
		t.Fatalf("missing form field in body: %s", body)
	}
	if !strings.Contains(body, `name="file"`) || !strings.Contains(body, `filename="hello.txt"`) {
		t.Fatalf("missing file part in body: %s", body)
	}
	if !strings.Contains(body, "hello jiguang") {
		t.Fatalf("missing file content in body: %s", body)
	}
}

func TestUploadFileFromPathAndValidator(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "alias.txt")
	if err := os.WriteFile(path, []byte("rid-1\nrid-2\n"), 0o644); err != nil {
		t.Fatalf("write temp file: %v", err)
	}

	uf, err := UploadFileFromPath(path)
	if err != nil {
		t.Fatalf("UploadFileFromPath failed: %v", err)
	}
	defer uf.Close()

	validator := &FileValidator{
		MaxSize:      10 * 1024,
		AllowedExts:  []string{".txt"},
		AllowedMimes: []string{"text/plain"},
	}
	if err := validator.Validate(uf); err != nil {
		t.Fatalf("Validate failed: %v", err)
	}

	// After validation, reader should still be readable from start.
	data, err := io.ReadAll(uf.Reader)
	if err != nil {
		t.Fatalf("read after validate failed: %v", err)
	}
	if string(data) != "rid-1\nrid-2\n" {
		t.Fatalf("unexpected content after validate: %q", data)
	}
}

func TestFileValidatorRejectsInvalidExt(t *testing.T) {
	uf := UploadFileFromBytes("bad.exe", []byte("x"))
	err := (&FileValidator{AllowedExts: []string{".txt"}}).Validate(uf)
	if err == nil {
		t.Fatal("expected extension validation error")
	}
}

func TestUploadRequestRejectsNilBodyType(t *testing.T) {
	_, err := newMultipartFormDataRequest(nil, &Request{
		Method: "POST",
		URL:    "http://example.com",
		Body:   "not-multipart",
	})
	if err == nil {
		t.Fatal("expected type assertion error")
	}
}
