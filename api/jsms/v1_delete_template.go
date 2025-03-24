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
	"strconv"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// # 删除模板
//   - 功能说明：删除短信模板。
//   - 调用地址：DELETE `/v1/templates/{tempID}`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates#%E5%88%A0%E9%99%A4%E6%A8%A1%E6%9D%BF-api
func (s *apiv1) DeleteTemplate(ctx context.Context, tempID int64) (*TemplateDeleteResult, error) {
	if s == nil {
		return nil, api.ErrNilJSmsAPIv1
	}

	if tempID < 1 {
		return nil, errors.New("invalid `tempID`, required: >0")
	}

	req := &api.Request{
		Method: http.MethodDelete,
		Proto:  s.proto,
		URL:    s.host + "/v1/templates/" + strconv.FormatInt(tempID, 10),
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TemplateDeleteResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type TemplateDeleteResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *TemplateDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
