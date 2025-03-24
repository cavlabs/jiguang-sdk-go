/*
 *
 * Copyright 2025 calvinit/jiguang-sdk-go authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package callback

// # 回调数据类型
type Type int

const (
	TargetValid   Type = iota // 目标有效 (0)
	TargetInvalid             // 目标无效 (1)
	SentSucc                  // 提交成功 (2)
	SentFail                  // 提交失败 (3)
	ReceivedSucc              // 送达成功 (4)
	ReceivedFail              // 送达失败 (5)
	Click                     // 点击 (6)
	RetractedSucc             // 撤回成功 (7)
	RetractedFail             // 撤回失败 (8)
)

var typeDescs = [...]string{
	"目标有效",
	"目标无效",
	"提交成功",
	"提交失败",
	"送达成功",
	"送达失败",
	"点击",
	"撤回成功",
	"撤回失败",
}

func (t Type) String() string {
	return typeDescs[t]
}
