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
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 查询别名
//   - 功能说明：获取指定 alias 下的设备，正常情况下最多输出 10 个，超过 10 个默认输出 10 个。
//   - 调用地址：GET `/v3/aliases/{alias}`，`alias` 为指定的别名值；`plats` 为可选参数，不填则默认为所有平台。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9F%A5%E8%AF%A2%E5%88%AB%E5%90%8D
func (d *apiv3) GetAlias(ctx context.Context, alias string, plats ...platform.Platform) (*AliasGetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if alias == "" {
		return nil, errors.New("`alias` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  d.proto,
		URL:    d.host + "/v3/aliases/" + alias + "?platform=" + platform.Concat(plats, ",") + "&new_format=true",
		Auth:   d.auth,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AliasGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type AliasGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Data          []AliasGetData `json:"data,omitempty"`
}

// 查询别名成功返回的数据结构
type AliasGetData struct {
	RegistrationID string             `json:"registration_id"`  // 设备标识 Registration ID
	Platform       platform.Platform  `json:"platform"`         // 平台
	LastOnlineDate *jiguang.LocalDate `json:"last_online_date"` // 最后一次上线日期
}

func (rs *AliasGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
