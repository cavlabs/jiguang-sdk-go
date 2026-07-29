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

package gpush

import (
	"io"
	"os"

	"github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/file"
)

// 上传文件相关类型与构造函数。
type UploadFile = file.UploadFile

func UploadFileFromPath(path string) (*UploadFile, error) {
	return file.UploadFileFromPath(path)
}

func UploadFileFromFile(f *os.File) (*UploadFile, error) {
	return file.UploadFileFromFile(f)
}

func UploadFileFromBytes(name string, data []byte) *UploadFile {
	return file.UploadFileFromBytes(name, data)
}

func UploadFileFromReadSeeker(name string, r io.ReadSeeker, size int64) (*UploadFile, error) {
	return file.UploadFileFromReadSeeker(name, r, size)
}
