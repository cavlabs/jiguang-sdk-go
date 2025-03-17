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

package message

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// [支付宝生活号] 消息。
type AlipayLife struct {
	TemplateID string             `json:"template_id"`       // 【必填】消息模板 ID，最大长度 128。
	Context    *AlipayLifeContext `json:"context,omitempty"` // 【必填】消息模板上下文，即模板中定义的参数及参数值。
}

// [支付宝生活号] 消息 - 模板上下文。
type AlipayLifeContext struct {
	// 【必填】顶部色条的色值，最大长度 10。
	HeadColor string `json:"head_color"`
	// 【必填】点击消息后承接页的地址，最大长度 256。
	URL string `json:"url"`
	// 【必填】底部链接描述文字，如“查看详情”，最多能传 8 个汉字或 16 个英文字符。
	ActionName string `json:"action_name"`
	// 【可选】模板中占位符的值及文字颜色。
	//  - 将根据顺序组装成以 keyword1, keyword2, keyword3, ... 等为 key 的 JSON 格式传送；
	//  - 示例："keyword1":{"color":"#85be53","value":"HU7142"},"keyword2":{"color":"#85be53","value":"HU7142"},"keyword3":{"color":"#85be53","value":"HU7142"}。
	Keywords []AlipayLifeContextKeyword `json:"-,omitempty"`
	// 【可选】模板中占位符的值及文字颜色，First 一般为开头语的占位符。
	First *AlipayLifeContextKeyword `json:"first,omitempty"`
	// 【可选】模板中占位符的值及文字颜色，Remark 一般为结束语的占位符。
	Remark *AlipayLifeContextKeyword `json:"remark,omitempty"`
}

var alipayLifeContextKeywordsKeyRegex = regexp.MustCompile(`^keyword(\d+)$`)

func (rs *AlipayLifeContext) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if headColor, ok := aux["head_color"]; ok {
		// 去掉字符串头尾的 `"`
		rs.HeadColor = strings.Trim(string(headColor), `"`)
		delete(aux, "head_color")
	}
	if url, ok := aux["url"]; ok {
		rs.URL = strings.Trim(string(url), `"`)
		delete(aux, "url")
	}
	if actionName, ok := aux["action_name"]; ok {
		rs.ActionName = strings.Trim(string(actionName), `"`)
		delete(aux, "action_name")
	}
	if first, ok := aux["first"]; ok {
		var firstKeyword AlipayLifeContextKeyword
		if err := json.Unmarshal(first, &firstKeyword); err != nil {
			return err
		}
		rs.First = &firstKeyword
		delete(aux, "first")
	}
	if remark, ok := aux["remark"]; ok {
		var remarkKeyword AlipayLifeContextKeyword
		if err := json.Unmarshal(remark, &remarkKeyword); err != nil {
			return err
		}
		rs.Remark = &remarkKeyword
		delete(aux, "remark")
	}

	if len(aux) > 0 {
		var maxKeySno int
		keywords := make(map[int]json.RawMessage)
		for k, v := range aux {
			matches := alipayLifeContextKeywordsKeyRegex.FindStringSubmatch(k)
			if len(matches) > 1 {
				keySno, err := strconv.Atoi(matches[1])
				if err != nil {
					// return err
					continue
				}
				if keySno > maxKeySno {
					maxKeySno = keySno
				}
				keywords[keySno] = v
			}
		}
		for i := 1; i <= maxKeySno; i++ {
			var keyword AlipayLifeContextKeyword
			if v, ok := keywords[i]; ok {
				if err := json.Unmarshal(v, &keyword); err != nil {
					return err
				}
			}
			rs.Keywords = append(rs.Keywords, keyword)
		}
	}

	return nil
}

func (rs AlipayLifeContext) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{})
	data["head_color"] = rs.HeadColor
	data["url"] = rs.URL
	data["action_name"] = rs.ActionName
	for i, v := range rs.Keywords {
		data["keyword"+strconv.Itoa(i+1)] = v
	}
	if rs.First != nil {
		data["first"] = rs.First
	}
	if rs.Remark != nil {
		data["remark"] = rs.Remark
	}
	return json.Marshal(data)
}

// [支付宝生活号] 消息 - 模板上下文关键字。
type AlipayLifeContextKeyword struct {
	Color string `json:"color"` // 【必填】当前文字颜色。
	Value string `json:"value"` // 【必填】模板中占位符的值，最大长度 128。
}
