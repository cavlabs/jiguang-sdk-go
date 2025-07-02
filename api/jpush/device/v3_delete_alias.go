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
)

// # 删除别名
//   - 功能说明：删除一个别名，以及该别名与设备的绑定关系。
//   - 调用地址：DELETE `/v3/aliases/{alias}`，`alias` 为指定的别名值；`plats` 为可选参数，不填则默认为所有平台。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E5%88%A0%E9%99%A4%E5%88%AB%E5%90%8D
func (d *apiv3) DeleteAlias(ctx context.Context, alias string, plats ...platform.Platform) (*AliasDeleteResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if alias == "" {
		return nil, errors.New("`alias` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodDelete,
		Proto:  d.proto,
		URL:    d.host + "/v3/aliases/" + alias + "?platform=" + platform.Concat(plats, ","),
		Auth:   d.auth,
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &AliasDeleteResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type AliasDeleteResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
}

func (rs *AliasDeleteResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
