/*
 *
 * Copyright 2025 cavlabs/jiguang-sdk-go authors.
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

// # 更新标签
//   - 功能说明：为一个标签添加或者删除设备。
//   - 调用地址：POST `/v3/tags/{tag}`，`tag` 为指定的标签值；`adds`/`removes` 为增加或删除的设备标识 Registration ID 集合，最多各支持 1000 个，不能同时为空。
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device#%E6%9B%B4%E6%96%B0%E6%A0%87%E7%AD%BE
func (d *apiv3) SetTag(ctx context.Context, tag string, adds, removes []string) (*TagSetResult, error) {
	if d == nil {
		return nil, api.ErrNilJPushDeviceAPIv3
	}

	if tag == "" {
		return nil, errors.New("`tag` cannot be empty")
	}
	al, rl := len(adds), len(removes)
	if al == 0 && rl == 0 {
		return nil, errors.New("`adds` and `removes` cannot both be empty")
	}
	if al > 1000 {
		return nil, errors.New("`adds` cannot be more than 1000")
	}
	if rl > 1000 {
		return nil, errors.New("`removes` cannot be more than 1000")
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  d.proto,
		URL:    d.host + "/v3/tags/" + tag,
		Auth:   d.auth,
		Body:   &tagSetParam{RegistrationIDs: registrationIDsForTagSetParam{Add: adds, Remove: removes}},
	}
	resp, err := d.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &TagSetResult{Response: resp}
	if !resp.IsNoContent() {
		err = json.Unmarshal(resp.RawBody, result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

type tagSetParam struct {
	RegistrationIDs registrationIDsForTagSetParam `json:"registration_ids"` // 更新标签的设备标识 Registration ID 集合
}

// 更新标签的设备标识 Registration ID 集合，提供给 tagSetParam.RegistrationIDs 使用
type registrationIDsForTagSetParam struct {
	Add    []string `json:"add,omitempty"`    // 增加的设备标识 Registration ID 集合，最多支持 1000 个
	Remove []string `json:"remove,omitempty"` // 删除的设备标识 Registration ID 集合，最多支持 1000 个
}

type TagSetResult struct {
	*api.Response `json:"-"`
	Error         *TagSetError `json:"error,omitempty"`
}

func (rs *TagSetResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && (rs.Error == nil || rs.Error.IsSuccess())
}

// 更新标签失败错误
type TagSetError struct {
	api.CodeError
	IllegalRIDs []string `json:"illegal_rids,omitempty"` // 非法 Registration ID 集合
}

func (e *TagSetError) Error() string {
	if e == nil || e.Code == 0 {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}

func (e *TagSetError) String() string {
	if e == nil {
		return ""
	}
	if e.Message == "" {
		return fmt.Sprintf("[%d] %s", e.Code, strings.Join(e.IllegalRIDs, ","))
	}
	return fmt.Sprintf("[%d] %s %s", e.Code, e.Message, strings.Join(e.IllegalRIDs, ","))
}
