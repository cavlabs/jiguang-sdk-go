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

package inquire

// # 模板审核结果
//
// SMS_TEMPLATE 回执数据。
type TemplateData struct {
	RawJSON      string `json:"-"`                      // 原始 JSON 数据
	TempID       int64  `json:"tempId"`                 // 模板 ID
	Status       int    `json:"status"`                 // 模板状态：1 - 审核通过；2 - 审核不通过
	RefuseReason string `json:"refuseReason,omitempty"` // 审核不通过的原因
}

func (t TemplateData) MarshalJSON() ([]byte, error) {
	return []byte(t.RawJSON), nil
}
