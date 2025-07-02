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

package audience

// # 文件推送对象
//
// 用于指定文件推送。
type File struct {
	// 【必填】文件唯一标识，可通过文件上传接口 UploadFileForAlias 或 UploadFileForRegistrationID 获得。
	FileID string `json:"file_id"`
}
