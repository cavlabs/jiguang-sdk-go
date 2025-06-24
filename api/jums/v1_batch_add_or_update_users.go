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

package jums

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/cavlabs/jiguang-sdk-go/api"
)

// # 批量添加、更新用户信息
//   - 功能说明：批量添加、更新用户信息。
//   - 调用地址：POST `/v1/user/opt`
//   - 接口文档：[docs.jiguang.cn]
//
// 支持批量添加与更新用户信息：如果 UserID 已存在，则为更新操作；如不存在，则为添加操作。
//
// 参数 accessAuth 为 true 时，使用全局 accessKey 和 accessMasterSecret 鉴权方式；否则为默认的渠道 channelKey 和 masterSecret 鉴权方式。
//
// 在导入 UserID 与各个通道注册 ID 的对应关系时，有一个重要的关键字段是 [bind_id]（绑定标识），它可以取值 all、通道编码、渠道 Key。
//   - all 即全局设置，所有渠道、通道通用，对标签、手机号码、邮箱地址、钉钉 ID、企业微信 ID、企业微信互联企业 ID 有效，注意：钉钉 ID 的关联标识仅允许传 all；
//   - 具体的 [Channel Key]，在渠道信息中可获取到该值，对标签、手机号码、邮箱地址有效，如果需要设置某一渠道的专属信息，则可以用 Channel Key 做标识，如不需要，则传 all；
//   - 具体的 [通道编码]，在通道详情中可获取到该值（在创建通道时由运营人员自行设定），在为 UserID 绑定 APP 注册 ID、微信公众号 ID、微信小程序 ID、支付宝生活号 ID 时必须指定通道编码。
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jums/server/rest_api_jums_user#%E6%89%B9%E9%87%8F%E6%B7%BB%E5%8A%A0%E4%B8%8E%E6%9B%B4%E6%96%B0-api
// [bind_id]: https://img.jiguang.cn/docs/2021/jums/server/image/user1.png
// [Channel Key]: https://img.jiguang.cn/docs/2023/jums/server/assets/16825055309596.jpg
// [通道编码]: https://img.jiguang.cn/docs/2023/jums/server/assets/16825055860019.jpg
func (u *apiv1) BatchAddOrUpdateUsers(ctx context.Context, param []UsersBatchAddOrUpdateParam, accessAuth bool) (*UsersBatchAddOrUpdateResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	var auth string
	if accessAuth {
		if auth = u.accessAuth; auth == "" {
			return nil, errors.New("please set the `accessKey` and `accessMasterSecret` required for this API")
		}
	} else {
		auth = u.auth
	}
	req := &api.Request{
		Method: http.MethodPost,
		Proto:  u.proto,
		URL:    u.host + "/v1/user/opt",
		Auth:   auth,
		Body:   param,
	}
	resp, err := u.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &UsersBatchAddOrUpdateResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UsersBatchAddOrUpdateParam struct {
	UserID string          `json:"userid"`        // 【必填】用户的唯一标识。
	Add    *UsersBatchOpts `json:"add,omitempty"` // 【可选】对标签、APP 注册 ID 、企业微信互联企业 ID 进行增加操作。
	Set    *UsersBatchOpts `json:"set,omitempty"` // 【可选】对标签、APP 注册 ID 、企业微信互联企业 ID 进行设置或覆盖更新。
	Del    *UsersBatchOpts `json:"del,omitempty"` // 【可选】对标签、各通道注册 ID 进行删除操作。
}

// 用户信息批量 Add、Set 和 Del 的关联参数
type UsersBatchOpts struct {
	Tag                map[string][]string `json:"tag,omitempty"`                 // 【可选】标签，在同一个关联标识下可以设置多个。
	Phone              map[string]string   `json:"phone,omitempty"`               // 【可选】手机号码，在同一个关联标识下仅允许设置 1 个。
	Email              map[string]string   `json:"email,omitempty"`               // 【可选】邮箱，在同一个关联标识下仅允许设置 1 个。
	APP                map[string][]string `json:"app,omitempty"`                 // 【可选】APP 通道的注册 ID，在同一个关联标识下可以设置多个。
	Wechatoa           map[string]string   `json:"wechatoa,omitempty"`            // 【可选】微信公众号通道的注册 ID，在同一个关联标识下仅允许设置 1 个。
	Wechatmp           map[string]string   `json:"wechatmp,omitempty"`            // 【可选】微信小程序通道的注册 ID，在同一个关联标识下仅允许设置 1 个。
	AlipayLife         map[string]string   `json:"alipaylife,omitempty"`          // 【可选】支付宝生活号通道的注册 ID，在同一个关联标识下仅允许设置 1 个。
	DingtalkCC         string              `json:"dingtalkcc,omitempty"`          // 【可选】钉钉通道的注册 ID，仅允许设置 1 个。
	Wechatwk           string              `json:"wechatwk,omitempty"`            // 【可选】企业微信通道的注册 ID，仅允许设置 1 个。
	WechatwkLinkedCorp []string            `json:"wechatwk_linkedcorp,omitempty"` // 【可选】企业微信互联企业的用户 ID，可以设置多个值。请注意，如果该用户是互联企业名下的，用户 ID 需要拼接企业 ID 后上传，格式为 CorpId1/userid1。
}

type UsersBatchAddOrUpdateResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	Data interface{} `json:"data,omitempty"` // [string / UsersBatchAddOrUpdateData] 操作成功/失败的详细数据，当请求失败时，数据为空。
}

// 当操作成功或部分成功时，将返回成功/失败的详细数据
type UsersBatchAddOrUpdateData struct {
	Success []string                        `json:"success"` // 操作成功的数据
	Fail    []UsersBatchAddOrUpdateFailData `json:"fail"`    // 操作失败的数据
}

// 操作失败的详细数据
type UsersBatchAddOrUpdateFailData struct {
	UserID  string `json:"userid,omitempty"`  // 用户的唯一标识。
	ErrCode string `json:"errcode,omitempty"` // 当有失败数据时，返回具体的失败错误码。
	ErrMsg  string `json:"errmsg,omitempty"`  // 当有失败数据时，返回具体的失败错误原因。
}

func (rs *UsersBatchAddOrUpdateResult) UnmarshalJSON(data []byte) error {
	// 为了避免 "fatal error: stack overflow"，因为类型 `Alias` 不继承原始类型 `UsersBatchAddOrUpdateResult` 的方法，包括 `UnmarshalJSON` 方法。
	// 参见：https://choly.ca/post/go-json-marshalling/
	type Alias UsersBatchAddOrUpdateResult
	aux := &struct {
		Data json.RawMessage `json:"data,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(rs),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	rs.CodeError = aux.CodeError

	if aux.Data == nil {
		return nil
	}

	var strData string
	if err := json.Unmarshal(aux.Data, &strData); err == nil {
		rs.Data = strData
		return nil
	} else {
		objData := &UsersBatchAddOrUpdateData{}
		if err = json.Unmarshal(aux.Data, &objData); err == nil {
			if objData != nil { // omitempty
				rs.Data = objData
			}
		} else {
			// return err
			rs.Data = string(aux.Data)
		}
	}

	return nil
}

func (rs *UsersBatchAddOrUpdateResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
