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

package message

// # 【邮件】消息
type Email struct {
	// 【必填】邮件标题。
	Subject string `json:"subject"`
	// 【必填】邮件内容，支持 HTML 格式。
	Text string `json:"text"`
	// 【可选】邮件附件。
	//  - 使用 UploadMaterial 上传附件，然后将获得的 URL 传值在此；
	//  - 将任意 URL 地址传值在此，需保证该地址可被 UMS 访问。
	Files []string `json:"files,omitempty"`
}
