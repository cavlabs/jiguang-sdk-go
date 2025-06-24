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

package schedule

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/cavlabs/jiguang-sdk-go/api"
	"github.com/cavlabs/jiguang-sdk-go/jiguang"
)

// # 获取某个定时任务的所有消息 ID
//   - 功能说明：获取当前用户指定定时任务对应的所有消息 ID 列表。
//   - 调用地址：GET `/v3/schedules/{scheduleID}/msg_ids`
//   - 接口文档：[docs.jiguang.cn]
//
// [docs.jiguang.cn]: https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule#%E8%8E%B7%E5%8F%96%E6%9F%90%E4%B8%AA%E5%AE%9A%E6%97%B6%E4%BB%BB%E5%8A%A1%E7%9A%84%E6%89%80%E6%9C%89%E6%B6%88%E6%81%AF-id
func (s *apiv3) GetScheduleMsgIDs(ctx context.Context, scheduleID string) (*ScheduleMsgIDsResult, error) {
	if s == nil {
		return nil, api.ErrNilJPushScheduleAPIv3
	}

	if scheduleID == "" {
		return nil, errors.New("`scheduleID` cannot be empty")
	}

	req := &api.Request{
		Method: http.MethodGet,
		Proto:  s.proto,
		URL:    s.host + "/v3/schedules/" + scheduleID + "/msg_ids",
		Auth:   s.auth,
	}
	resp, err := s.client.Request(ctx, req)
	if err != nil {
		return nil, err
	}

	result := &ScheduleMsgIDsResult{Response: resp}
	err = json.Unmarshal(resp.RawBody, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type ScheduleMsgIDsResult struct {
	*api.Response `json:"-"`
	Error         *api.CodeError `json:"error,omitempty"`
	Count         *int           `json:"count,omitempty"`  // 消息 ID 数量
	MsgIDs        []string       `json:"msgids,omitempty"` // 消息 ID 列表
}

func (rs *ScheduleMsgIDsResult) UnmarshalJSON(data []byte) error {
	var aux map[string]json.RawMessage

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if codeError, ok := aux["error"]; ok {
		if err := json.Unmarshal(codeError, &rs.Error); err != nil {
			return err
		}
		delete(aux, "error")
	}

	if v, ok := aux["count"]; ok {
		// null
		if len(v) == 4 && v[0] == 'n' && v[1] == 'u' && v[2] == 'l' && v[3] == 'l' {
			rs.Count = nil
		} else {
			count, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}
			rs.Count = jiguang.Int(count)
		}
	}

	if v, ok := aux["msgids"]; ok {
		if err := json.Unmarshal(v, &rs.MsgIDs); err != nil {
			return err
		}
	}

	return nil
}

func (rs *ScheduleMsgIDsResult) IsSuccess() bool {
	return rs != nil && rs.StatusCode/100 == 2 && rs.Error.IsSuccess()
}

// 定时任务 GetScheduleMsgIDs 接口的响应消息 ID 结构，2018-09-13 后的新数据格式。
type ScheduleMsgID struct {
	MsgID     string         `json:"msg_id,omitempty"` // 消息 ID
	Error     *api.CodeError `json:"error,omitempty"`  // 错误信息
	NeedRetry bool           `json:"needRetry"`        // 是否需要重试
	Timestamp *int64         `json:"ts,omitempty"`     // 定时任务执行成功的时间戳，精确到毫秒，对于 2022-03-17 后创建的定时任务，返回的内容中将增加该字段。
}

func (m *ScheduleMsgID) IsSuccess() bool {
	return m != nil && m.Error.IsSuccess()
}

func TryParseScheduleMsgIDFromString(s string) (msgID ScheduleMsgID, ok bool) {
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		ok = json.Unmarshal([]byte(s), &msgID) == nil
	}
	return
}
