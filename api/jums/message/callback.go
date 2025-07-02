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

// # 消息回调参数
type Callback struct {
	URL    string                 `json:"url,omitempty"`    // 【可选】数据临时回调地址，仅针对这一次消息发送请求生效，该地址必须在极光后台有校验通过方可使用；不指定则以极光后台配置的默认地址为准。
	Params map[string]interface{} `json:"params,omitempty"` // 【可选】需要回调给用户的自定义参数。
}
