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

package jums

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/calvinit/jiguang-sdk-go/api"
	"github.com/calvinit/jiguang-sdk-go/api/jums/audience"
	"github.com/calvinit/jiguang-sdk-go/api/jums/message"
)

// 普通消息 - 其他方式发送
//  - 功能说明：普通消息发送。
//  - 调用地址：POST `/v1/sent`
//  - 接口文档：https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message
func (u *apiv1) Send(ctx context.Context, param *SendParam) (*SendResult, error) {
	if u == nil {
		return nil, api.ErrNilJUmsAPIv1
	}

	if param == nil {
		return nil, errors.New("`param` cannot be nil")
	}

	if len(param.Message.DingtalkCC) > 0 {
		for i, dingtalkCC := range param.Message.DingtalkCC {
			dingtalkCCMsg := dingtalkCC.Msg
			if dingtalkCCMsg == nil {
				return nil, fmt.Errorf("msg_dingtalk_cc.[%d].msg cannot be nil", i)
			}
			if err := dingtalkCCMsg.Validate(); err != nil {
				return nil, err
			}
		}
	}

	if len(param.Message.Wechatwk) > 0 {
		for _, wechatwk := range param.Message.Wechatwk {
			if err := wechatwk.Validate(); err != nil {
				return nil, err
			}
		}
	}

	req := &api.Request{
		Method: http.MethodPost,
		Proto:  u.proto,
		URL:    u.host + "/v1/sent",
		Auth:   u.auth,
		Body:   param,
	}
	resp, err := u.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &SendResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type SendParam struct {
	*audience.Audience // 消息的发送目标
	*message.Message   // 消息的内容

	// 【可选】发送策略 ID，如果是同时发送可传 0 或不传。当使用自定义通道 ID 发送时，该字段无效。
	//
	// 1. 如不需要进行补发，仅进行单通道、多通道同时发送，则不需填写策略 ID，或者设置为 0。
	//
	// 2. 在 官网控制台-渠道-发送策略 中创建一个补发策略后，调 API 时可使用策略 ID 进行指定。
	//
	//  - 在使用自定义通道注册 ID 发送时，发送策略不生效；
	//  - 如果使用了发送策略，策略中包含的通道和 msg_xxx 中的通道信息需要一致。
	RuleID int `json:"rule_id,omitempty"`
	// 【可选】可选参数，用于黑白名单 ID、提交人等信息的填写。
	Option *message.Option `json:"option,omitempty"`
	// 【可选】回调参数。
	//
	// 调 API 发送消息时，可以指定 Callback 参数，方便用户临时变更回调 URL 或者回调带上其自定义参数，满足其日常业务需求。详细使用说明请阅读 [消息回调设置]。
	//
	// 此功能仅针对极光 VIP 用户提供，提供「目标有效/无效、提交成功/失败、送达成功/失败、点击、撤回成功/失败」9 种消息状态，需在官网控制台设置所需回调的状态。
	//
	// [消息回调设置]: https://docs.jiguang.cn/jums/advanced/callback
	Callback *message.Callback `json:"callback,omitempty"`
}

type SendResult struct {
	*api.Response `json:"-"`
	*api.CodeError
	MsgID  string `json:"msgid,omitempty"`  // 消息唯一 ID，当失败时不返回。
	SendNo string `json:"sendno,omitempty"` // API 调用时传递的标识，在请求成功时将被原样返回。
}

func (rs *SendResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.CodeError.IsSuccess()
}
