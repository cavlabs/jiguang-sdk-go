/*
 *
 * Copyright 2024 calvinit/jiguang-sdk-go authors.
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

package push

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/calvinit/jiguang-sdk-go/api"
)

// 厂商配额查询
//  - 功能说明：厂商为了控制应用推送消息的频率，会根据应用在厂商的日联网数计算每天推送数量上限。
//  目前已知小米、OPPO、vivo 都有每天的额度控制，额度可以在厂商后台查询，也可以通过极光 WebPortal 控制台查询；
//  为了进一步方便开发，极光汇总了三个平台的查询接口，供开发者调用查询。
//	- 调用地址：GET `/v3/push/quota`
//  - 接口文档：https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_advanced#%E5%8E%82%E5%95%86%E9%85%8D%E9%A2%9D%E6%9F%A5%E8%AF%A2-api
// 注意：厂商配额查询 API 接口频率和 Push API 接口频率共享，消耗 Push API 接口频率。
func (p *apiv3) GetQuota(ctx context.Context) (*QuotaGetResult, error) {
	if p == nil {
		return nil, api.ErrNilJPushPushAPIv3
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  p.proto,
		URL:    p.host + "/v3/push/quota",
		Auth:   p.auth,
	}
	resp, err := p.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &QuotaGetResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type QuotaGetResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Data          *QuotaData     `json:"data,omitempty"` // 厂商推送消息配额数据
}

// 厂商推送消息配额数据
type QuotaData struct {
	Xiaomi *MessageQuota `json:"xiaomi_quota,omitempty"` // 小米配额
	OPPO   *MessageQuota `json:"oppo_quota,omitempty"`   // OPPO 配额
	Vivo   *MessageQuota `json:"vivo_quota,omitempty"`   // vivo 配额
}

// 消息配额
type MessageQuota struct {
	System    *MessageQuotaDetail `json:"system,omitempty"`    // 系统消息
	Operation *MessageQuotaDetail `json:"operation,omitempty"` // 运营消息
}

// 消息配额详情
type MessageQuotaDetail struct {
	Total *int64 `json:"total,omitempty"` // 总配额，开通了不限量时，返回 -1
	Used  *int64 `json:"used,omitempty"`  // 已使用配额，开通了不限量时，返回 -1
}

func (rs *QuotaGetResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
		return nil
	}

	code, err := strconv.Atoi(string(aux["code"]))
	if err != nil {
		return err
	}
	delete(aux, "code")

	// 去掉字符串头尾的 `"`
	message := strings.Trim(string(aux["message"]), `"`)
	delete(aux, "message")

	rs.Error = &api.CodeError{Code: code, Message: message}

	if v, ok := aux["data"]; ok {
		// null
		if len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l' {
			rs.Data = nil
		} else {
			var quota QuotaData
			if err = json.Unmarshal(v, &quota); err != nil {
				return err
			}
			rs.Data = &quota
		}
	}

	return nil
}

func (rs QuotaGetResult) MarshalJSON() ([]byte, error) {
	rsmap := make(map[string]interface{}, 3)
	rsmap["code"] = rs.Error.Code
	rsmap["message"] = rs.Error.Message
	if rs.Data != nil {
		rsmap["data"] = rs.Data
	}
	return json.Marshal(rsmap)
}

func (rs *QuotaGetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}
