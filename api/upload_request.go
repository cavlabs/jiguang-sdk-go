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
	"errors"
	"mime/multipart"
)

type uploadField struct {
	name  string
	value string
}

type uploadFileEntry struct {
	field string
	file  *UploadFile
}

// UploadRequest 表示 multipart/form-data 请求体。
type UploadRequest struct {
	fields    []uploadField
	files     []uploadFileEntry
	validator *FileValidator
}

// NewUploadRequest 创建 UploadRequest。
func NewUploadRequest(opts ...UploadRequestOption) *UploadRequest {
	r := &UploadRequest{}
	for _, opt := range opts {
		if opt != nil {
			opt(r)
		}
	}
	return r
}

// WriteMultipart 将表单字段与文件写入 multipart.Writer。
func (r *UploadRequest) WriteMultipart(w *multipart.Writer) error {
	if r == nil {
		return errors.New("api: UploadRequest is nil")
	}
	if w == nil {
		return errors.New("api: multipart.Writer is nil")
	}

	for _, f := range r.fields {
		if err := w.WriteField(f.name, f.value); err != nil {
			return err
		}
	}

	for _, entry := range r.files {
		if entry.file == nil {
			return errors.New("api: UploadFile is nil")
		}
		if r.validator != nil {
			if err := r.validator.Validate(entry.file); err != nil {
				return err
			}
		}
		if err := entry.file.writeTo(w, entry.field); err != nil {
			return err
		}
	}
	return nil
}

// Close 关闭请求中由本对象托管的文件资源。可重复调用。
func (r *UploadRequest) Close() error {
	if r == nil {
		return nil
	}
	var firstErr error
	for _, entry := range r.files {
		if entry.file == nil {
			continue
		}
		if err := entry.file.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
