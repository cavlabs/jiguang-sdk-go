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

package send

import (
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/notification"
	"github.com/cavlabs/jiguang-sdk-go/api/jpush/push/options"
)

// # 默认值设置选项
type DefaultOption interface {
	apply(*Param)
}

// # APNS 生产环境默认值选项
type apnsProductionOption bool

func (o apnsProductionOption) apply(p *Param) {
	if p.Options == nil {
		p.Options = &options.Options{}
	}
	if p.Options.ApnsProduction == nil {
		production := bool(o)
		p.Options.ApnsProduction = &production
	}
}

// WithDefaultApnsProduction 设置默认 APNS 推送环境，当 Options.ApnsProduction 字段为空时生效
func WithDefaultApnsProduction(production bool) DefaultOption {
	return apnsProductionOption(production)
}

// # iOS Badge 默认值选项
type iosBadgeOption string

func (o iosBadgeOption) apply(p *Param) {
	// 确保 Notification 字段存在
	if p.Notification == nil {
		p.Notification = &notification.Notification{}
	}

	// 确保 IOS 字段存在
	if p.Notification.IOS == nil {
		p.Notification.IOS = &notification.IOS{}
	}

	// 只有当 Badge 字段为空时才设置默认值
	if p.Notification.IOS.Badge == nil {
		badge := string(o)
		p.Notification.IOS.Badge = &badge
	}
}

// WithDefaultIOSBadge 设置默认 iOS 角标值，当 Notification.IOS.Badge 字段为空时生效
func WithDefaultIOSBadge(badge string) DefaultOption {
	return iosBadgeOption(badge)
}

// # 应用默认值到推送参数
//
// 功能说明：为推送参数应用默认值，只会设置原本为空的字段，不会覆盖已有值。
//
// 使用示例：
//
//	param := &send.Param{
//	    Audience: audience.All,
//	    Notification: &notification.Notification{Alert: "test"},
//	}
//	send.ApplyDefaults(param,
//	    send.WithDefaultApnsProduction(false),
//	    send.WithDefaultIOSBadge("+1"),
//	)
func ApplyDefaults(param *Param, opts ...DefaultOption) {
	if param == nil {
		return
	}

	for _, opt := range opts {
		opt.apply(param)
	}
}
