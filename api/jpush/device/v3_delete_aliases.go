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

package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 删除设备的别名
//   - 功能说明：批量解绑设备与别名之间的关系。
//   - 调用地址：POST `/v3/aliases/{alias}`，`alias` 为指定的别名值；`registrationIDs` 为必填参数，需要和该 `alias` 解除绑定的设备标识 Registration ID 值集合，最多支持 1000 个。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E8%AE%BE%E5%A4%87%E7%9A%84%E5%88%AB%E5%90%8D
func (d *apiv3) DeleteAliases(ctx context.Context, alias string, registrationIDs []string) (*AliasesDeleteResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if alias == "" {
		return nil, errors.New("`alias` cannot be empty")
	}
	l := len(registrationIDs)
	if l == 0 {
		return nil, errors.New("`registrationIDs` cannot be empty")
	}
	if l > 1000 {
		return nil, errors.New("`registrationIDs` cannot be more than 1000")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/aliases/" + alias,
		Auth:   d.auth,
		Body:   &aliasesDeleteParam{registrationIDsForAliasesDeleteParam{registrationIDs}},
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AliasesDeleteResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type aliasesDeleteParam struct {
	// 要删除别名的设备标识 Registration ID 集合，最多支持 1000 个。
	RegistrationIDs registrationIDsForAliasesDeleteParam `json:"registration_ids"`
}

// registrationIDsForAliasesDeleteParam 要删除别名的设备标识 Registration ID 集合，最多支持 1000 个。
type registrationIDsForAliasesDeleteParam struct {
	Remove []string `json:"remove"`
}

type AliasesDeleteResult struct {
	*api.Response `json:"-"`
	Error         *AliasesDeleteError `json:"error,omitempty"`
}

func (rs *AliasesDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && (rs.Error == nil || rs.Error.IsSuccess())
}

// 删除设备的别名失败错误
type AliasesDeleteError struct {
	api.CodeError
	IllegalRIDs []string `json:"illegal_rids,omitempty"` // 非法 Registration ID 集合
}

func (e *AliasesDeleteError) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}

func (e *AliasesDeleteError) String() string {
	if e == nil {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}
