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

package jsms

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

var voiceCodeRegex = regexp.MustCompile(`^\d{4,8}$`)

// 发送语音验证码短信
//   - 功能说明：发送语音验证码短信。
//   - 调用地址：POST `/v1/voice_codes`
//   - 接口文档：https://docs.jiguang.cn/jsms/server/rest_api_jsms#%E5%8F%91%E9%80%81%E8%AF%AD%E9%9F%B3%E9%AA%8C%E8%AF%81%E7%A0%81%E7%9F%AD%E4%BF%A1-api
func (s *apiv1) SendVoiceCode(ctx context.Context, param *VoiceCodeSendParam) (*VoiceCodeSendResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}
	code := param.Code
	if code != "" && !voiceCodeRegex.MatchString(code) {
		return nil, errors.New("voice `code` must be 4-8 digits")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  s.proto,
		URL:    s.host + "/v1/voice_codes",
		Auth:   s.auth,
		Body:   param,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &VoiceCodeSendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 语音验证码发送参数
type VoiceCodeSendParam struct {
	Mobile    string    `json:"mobile"`               // 【必填】手机号码
	Code      string    `json:"code,omitempty"`       // 【可选】语音验证码的值，验证码仅支持 4-8 个数字
	VoiceLang VoiceLang `json:"voice_lang,omitempty"` // 【可选】语音验证码播报语言选择，0：中文播报，1：英文播报，2：中英混合播报
	TTL       int       `json:"ttl,omitempty"`        // 【可选】验证码有效期，默认为 60 秒
}

// 语音验证码播报语言
type VoiceLang string

const (
	VoiceLangChinese        VoiceLang = "0" // 中文播报
	VoiceLangEnglish        VoiceLang = "1" // 英文播报
	VoiceLangChineseEnglish VoiceLang = "2" // 中英混合播报
)

func (vl VoiceLang) String() string {
	switch vl {
	case VoiceLangChinese:
		return "Chinese"
	case VoiceLangEnglish:
		return "English"
	case VoiceLangChineseEnglish:
		return "Chinese+English"
	default:
		return "Default"
	}
}

func (p *VoiceCodeSendParam) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `VoiceCodeSendParam` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias VoiceCodeSendParam
	aux := &struct {
		VoiceLang *int `json:"voice_lang,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.VoiceLang != nil {
		p.VoiceLang = VoiceLang(strconv.Itoa(*aux.VoiceLang))
	}
	return nil
}

func (p VoiceCodeSendParam) MarshalJSON() ([]byte, error) {
	param := make(map[string]interface{}, 4)
	param["mobile"] = p.Mobile
	if p.Code != "" {
		param["code"] = p.Code
	}
	if p.VoiceLang != "" {
		param["voice_lang"], _ = strconv.Atoi(string(p.VoiceLang))
	}
	if p.TTL > 0 {
		param["ttl"] = p.TTL
	}
	return json.Marshal(param)
}

// 语音验证码发送结果
type VoiceCodeSendResult = sendResult
