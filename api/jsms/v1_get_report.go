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

package jsms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jsms/inquire"
)

// # 获取短信发送状态
//
// SMS_REPORT - 短信送达状态。
//   - 功能说明：获取 appKey 下短信的发送状态报告；API 调用频率为 5 秒/次，每次调用最多可获取 500 条数据；每条短信发送状态报告数据最多保留七天，即极光收到运营商返回状态报告后的七天。
//   - 调用地址：GET `/v1/report`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire#%E8%8E%B7%E5%8F%96%E7%9F%AD%E4%BF%A1%E5%8F%91%E9%80%81%E7%8A%B6%E6%80%81
func (s *apiv1) GetReport(ctx context.Context) (*ReportGetResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v1/report",
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ReportGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ReportGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError       `json:"error,omitempty"`
	Data          []inquire.ReportData `json:"-"` // 短信送达状态报告数据
}

func (rs *ReportGetResult) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '[' && data[len(data)-1] == ']' {
		var reports []inquire.ReportData
		if err := json.Unmarshal(data, &reports); err != nil {
			return err
		}
		rs.Data = reports
		return nil
	}
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `ReportGetResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias ReportGetResult
	aux := (*Alias)(rs)
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (rs ReportGetResult) MarshalJSON() ([]byte, error) {
	if rs.Error != nil {
		data := make(map[string]*api.CodeError, 1)
		data["error"] = rs.Error
		return json.Marshal(data)
	}
	return json.Marshal(rs.Data)
}

func (rs *ReportGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
