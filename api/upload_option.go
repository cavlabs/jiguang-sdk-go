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

// UploadRequestOption 用于配置 UploadRequest。
type UploadRequestOption func(*UploadRequest)

// WithField 添加普通表单字段。
func WithField(name, value string) UploadRequestOption {
	return func(r *UploadRequest) {
		if r == nil || name == "" {
			return
		}
		r.fields = append(r.fields, uploadField{name: name, value: value})
	}
}

// WithFile 添加文件表单字段。
func WithFile(field string, file *UploadFile) UploadRequestOption {
	return func(r *UploadRequest) {
		if r == nil || field == "" || file == nil {
			return
		}
		r.files = append(r.files, uploadFileEntry{field: field, file: file})
	}
}

// WithFiles 向同一表单字段添加多个文件。
func WithFiles(field string, files ...*UploadFile) UploadRequestOption {
	return func(r *UploadRequest) {
		if r == nil || field == "" {
			return
		}
		for _, file := range files {
			if file == nil {
				continue
			}
			r.files = append(r.files, uploadFileEntry{field: field, file: file})
		}
	}
}

// WithValidator 设置文件校验器。传入 nil 表示不校验。
func WithValidator(v *FileValidator) UploadRequestOption {
	return func(r *UploadRequest) {
		if r == nil {
			return
		}
		r.validator = v
	}
}
