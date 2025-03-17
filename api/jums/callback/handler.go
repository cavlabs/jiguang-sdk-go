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
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

// 默认回调请求处理器。
type defaultHandler struct {
	channelKey    string
	masterSecret  string
	checkAuth     bool
	targetValid   DataProcessor
	targetInvalid DataProcessor
	sentSucc      DataProcessor
	sentFail      DataProcessor
	receivedSucc  DataProcessor
	receivedFail  DataProcessor
	click         DataProcessor
	retractedSucc DataProcessor
	retractedFail DataProcessor
	unified       DataListProcessor
}

func (h defaultHandler) Callback(w http.ResponseWriter, r *http.Request) {
	if h.checkAuth {
		channelKey, masterSecret, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "invalid auth channel", http.StatusUnauthorized)
			return
		}
		if channelKey != h.channelKey {
			http.Error(w, "channel key mismatch", http.StatusForbidden)
			return
		}
		if masterSecret != h.masterSecret {
			http.Error(w, "master secret mismatch", http.StatusForbidden)
			return
		}
	}

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
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		list := DataList{RawJSON: string(body)}
		if err = json.Unmarshal(body, &list.Items); err != nil {
			http.Error(w, "invalid callback data", http.StatusBadRequest)
			return
		}

		if h.unified != nil {
			h.unified.Process(list)
			w.WriteHeader(http.StatusOK)
			return
		}

		var wg sync.WaitGroup
		for _, data := range list.Items {
			wg.Add(1)
			go h.process(data, &wg)
		}
		wg.Wait()

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "request method '"+r.Method+"' not supported", http.StatusMethodNotAllowed)
	}
}

func (h defaultHandler) process(data Data, wg *sync.WaitGroup) {
	defer wg.Done()

	switch data.Type {
	case TargetValid: // 目标有效 (0)
		if h.targetValid != nil {
			h.targetValid.Process(data)
		}
	case TargetInvalid: // 目标无效 (1)
		if h.targetInvalid != nil {
			h.targetInvalid.Process(data)
		}
	case SentSucc: // 提交成功 (2)
		if h.sentSucc != nil {
			h.sentSucc.Process(data)
		}
	case SentFail: // 提交失败 (3)
		if h.sentFail != nil {
			h.sentFail.Process(data)
		}
	case ReceivedSucc: // 送达成功 (4)
		if h.receivedSucc != nil {
			h.receivedSucc.Process(data)
		}
	case ReceivedFail: // 送达失败 (5)
		if h.receivedFail != nil {
			h.receivedFail.Process(data)
		}
	case Click: // 点击 (6)
		if h.click != nil {
			h.click.Process(data)
		}
	case RetractedSucc: // 撤回成功 (7)
		if h.retractedSucc != nil {
			h.retractedSucc.Process(data)
		}
	case RetractedFail: // 撤回失败 (8)
		if h.retractedFail != nil {
			h.retractedFail.Process(data)
		}
	}
}
