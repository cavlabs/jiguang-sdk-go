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
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

// # 文件校验器
//
// 用于校验上传文件的大小、扩展名和 MIME 类型。
type FileValidator struct {
	MaxSize      int64    // 最大文件大小（字节），<=0 表示不限制
	AllowedExts  []string // 允许的扩展名，如 ".png"、".txt"
	AllowedMimes []string // 允许的 MIME 类型，如 "image/png"、"text/plain"
}

// Validate 校验上传文件是否满足约束。
func (fv *FileValidator) Validate(f *UploadFile) error {
	if fv == nil || f == nil {
		return nil
	}

	name := f.Name
	if name == "" {
		name = "file"
	}

	// 校验文件大小
	if fv.MaxSize > 0 {
		if f.Size <= 0 {
			return fmt.Errorf("%q is empty", name)
		}
		if f.Size > fv.MaxSize {
			return fmt.Errorf("%q size %d exceeds limit %d", name, f.Size, fv.MaxSize)
		}
	}

	// 校验文件扩展名
	if len(fv.AllowedExts) > 0 {
		ext := strings.ToLower(filepath.Ext(name))
		if !stringInSlice(ext, fv.AllowedExts) {
			return fmt.Errorf("%q ext %q not allowed", name, ext)
		}
	}

	// 校验文件 MIME 类型
	if len(fv.AllowedMimes) > 0 {
		if f.Reader == nil {
			return fmt.Errorf("%q has nil Reader", name)
		}

		pos, err := f.Reader.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}

		buf := make([]byte, 512)
		n, err := f.Reader.Read(buf)
		if err != nil && err != io.EOF {
			_, _ = f.Reader.Seek(pos, io.SeekStart)
			return err
		}

		if _, err = f.Reader.Seek(pos, io.SeekStart); err != nil {
			return err
		}

		mimeType := http.DetectContentType(buf[:n])
		baseMime := getBaseMimeType(mimeType)
		if !stringInSlice(mimeType, fv.AllowedMimes) && !stringInSlice(baseMime, fv.AllowedMimes) {
			return fmt.Errorf("%q MIME %q not allowed", name, mimeType)
		}
	}

	return nil
}

// 提取 MIME 类型的基础部分，例如去掉 `;charset=utf-8`。
func getBaseMimeType(mimeType string) string {
	if i := strings.IndexByte(mimeType, ';'); i >= 0 {
		return strings.TrimSpace(mimeType[:i])
	}
	return mimeType
}

func stringInSlice(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
