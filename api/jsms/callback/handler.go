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

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

// 默认回调请求处理器。
type defaultHandler struct {
	appKey       string
	masterSecret string
	reply        ReplyDataProcessor
	report       ReportDataProcessor
	template     TemplateDataProcessor
	sign         SignDataProcessor
}

func (h defaultHandler) Callback(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		echostr := r.URL.Query().Get("echostr")
		if echostr == "" {
			http.Error(w, "missing 'echostr' parameter", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(echostr))
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "failed to parse form", http.StatusBadRequest)
			return
		}

		nonce := r.Form.Get("nonce")         // 随机长整数
		timestamp := r.Form.Get("timestamp") // 当前时间戳，毫秒值
		signature := r.Form.Get("signature") // 签名，结合 appKey、masterSecret、nonce、timestamp 生成
		dataType := r.Form.Get("type")       // 通知类型
		rawData := r.Form.Get("data")        // 通知内容，JSON 字符串，开发者可以根据 type 反序列化 data

		if sha1Sign(h.appKey, h.masterSecret, nonce, timestamp) != signature {
			http.Error(w, "signature not match", http.StatusForbidden)
			return
		}

		switch dataType {
		case "SMS_REPLY": // 「用户回复消息」回调通知类型
			if h.reply != nil {
				data := ReplyData{RawJSON: rawData}
				if err := json.Unmarshal([]byte(rawData), &data); err != nil {
					http.Error(w, "invalid SMS_REPLY data", http.StatusBadRequest)
					return
				}
				h.reply.Process(data)
			}
		case "SMS_REPORT": // 「短信送达状态」回调通知类型
			if h.report != nil {
				data := ReportData{RawJSON: rawData}
				if err := json.Unmarshal([]byte(rawData), &data); err != nil {
					http.Error(w, "invalid SMS_REPORT data", http.StatusBadRequest)
					return
				}
				h.report.Process(data)
			}
		case "SMS_TEMPLATE": // 「模板审核结果」回调通知类型
			if h.template != nil {
				data := TemplateData{RawJSON: rawData}
				if err := json.Unmarshal([]byte(rawData), &data); err != nil {
					http.Error(w, "invalid SMS_TEMPLATE data", http.StatusBadRequest)
					return
				}
				h.template.Process(data)
			}
		case "SMS_SIGN": // 「签名审核结果」回调通知类型
			if h.sign != nil {
				data := SignData{RawJSON: rawData}
				if err := json.Unmarshal([]byte(rawData), &data); err != nil {
					http.Error(w, "invalid SMS_SIGN data", http.StatusBadRequest)
					return
				}
				h.sign.Process(data)
			}
		default:
			http.Error(w, fmt.Sprintf("unsupported callback type %q", dataType), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "request method '"+r.Method+"' not supported", http.StatusMethodNotAllowed)
	}
}

func sha1Sign(appKey, masterSecret, nonce, timestamp string) string {
	s := fmt.Sprintf("appKey=%s&appMasterSecret=%s&nonce=%s&timestamp=%s", appKey, masterSecret, nonce, timestamp)
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
