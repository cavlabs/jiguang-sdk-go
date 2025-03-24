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

package message

import "errors"

// # 【企业微信】/【企业微信互联企业】消息类型
type WechatwkMsgType string

const (
	WechatwkMsgTypeText              WechatwkMsgType = "text"               // 文本消息
	WechatwkMsgTypeImage             WechatwkMsgType = "image"              // 图片消息
	WechatwkMsgTypeFile              WechatwkMsgType = "file"               // 文件消息
	WechatwkMsgTypeNews              WechatwkMsgType = "news"               // 外链图文消息
	WechatwkMsgTypeMpnews            WechatwkMsgType = "mpnews"             // 图文消息
	WechatwkMsgTypeMiniProgramNotice WechatwkMsgType = "miniprogram_notice" // 小程序通知消息
)

// ---------------------------------------------------------------------------------------------------------------------

// # 【企业微信】/【企业微信互联企业】消息
type Wechatwk struct {
	// 【必填】消息类型。
	//  - [文本消息] 类型为: WechatwkMsgTypeText，此时 Text 字段必填；
	//  - [图片消息] 类型为: WechatwkMsgTypeImage，此时 Image 字段必填；
	//  - [文件消息] 类型为: WechatwkMsgTypeFile，此时 File 字段必填；
	//  - [外链图文消息] 类型为: WechatwkMsgTypeNews，此时 News 字段必填；
	//  - [图文消息] 类型为: WechatwkMsgTypeMpnews，此时 Mpnews 字段必填；
	//  - [小程序通知消息] 类型为: WechatwkMsgTypeMiniProgramNotice，此时 MiniProgramNotice 字段必填。
	MsgType WechatwkMsgType `json:"msgtype,omitempty"`
	// 【可选】文本消息内容 (WechatwkTypeText)。
	Text *WechatwkText `json:"text,omitempty"`
	// 【可选】图片媒体文件 (WechatwkTypeImage)。
	Image *WechatwkImage `json:"image,omitempty"`
	// 【可选】素材媒体文件 (WechatwkTypeFile)。
	File *WechatwkFile `json:"file,omitempty"`
	// 【可选】外链图文消息 (WechatwkTypeNews)。
	News *WechatwkNews `json:"news,omitempty"`
	// 【可选】图文消息 (WechatwkTypeMpnews)。
	Mpnews *WechatwkMpnews `json:"mpnews,omitempty"`
	// 【可选】小程序通知消息 (WechatwkTypeMiniProgramNotice)。
	MiniProgramNotice *WechatwkMiniProgramNotice `json:"miniprogram_notice,omitempty"`
	// 【可选】是否是保密消息，0 表示可对外分享；1 表示不能分享且内容显示水印。默认为 0。
	Safe int `json:"safe"`
	// 【可选】是否开启重复消息检查，0 表示否；1 表示是。默认为 0。
	EnableDuplicateCheck int `json:"enable_duplicate_check"`
	// 【可选】表示是否重复消息检查的时间间隔，默认 1800 秒，最大不超过 4 小时。
	DuplicateCheckInterval int64 `json:"duplicate_check_interval,omitempty"`
}

// 验证消息参数。
func (w *Wechatwk) Validate() error {
	switch w.MsgType {
	case WechatwkMsgTypeText:
		if w.Text == nil {
			return errors.New("msg_wechatwk.[*].text is required when msgtype is `text`")
		}
	case WechatwkMsgTypeImage:
		if w.Image == nil {
			return errors.New("msg_wechatwk.[*].image is required when msgtype is `image`")
		}
	case WechatwkMsgTypeFile:
		if w.File == nil {
			return errors.New("msg_wechatwk.[*].file is required when msgtype is `file`")
		}
	case WechatwkMsgTypeNews:
		if w.News == nil {
			return errors.New("msg_wechatwk.[*].news is required when msgtype is `news`")
		}
	case WechatwkMsgTypeMpnews:
		if w.Mpnews == nil {
			return errors.New("msg_wechatwk.[*].mpnews is required when msgtype is `mpnews`")
		}
	case WechatwkMsgTypeMiniProgramNotice:
		if w.MiniProgramNotice == nil {
			return errors.New("msg_wechatwk.[*].miniprogram_notice is required when msgtype is `miniprogram_notice`")
		}
	default:
		return errors.New(string("unsupported msg_wechatwk.[*].msgtype " + w.MsgType))
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// # 【企业微信】/【企业微信互联企业】消息 - 文本消息
type WechatwkText struct {
	Content string `json:"content"` // 【必填】消息内容，最长不超过 2048 个字节。
}

// # 【企业微信】/【企业微信互联企业】消息 - 图片消息
type WechatwkImage struct {
	// 【必填】图片媒体文件 ID，可以调用 [上传临时素材] 接口获取。
	//
	// [上传临时素材]: https://developer.work.weixin.qq.com/document/path/90389
	MediaID string `json:"media_id"`
}

// # 【企业微信】/【企业微信互联企业】消息 - 文件消息
type WechatwkFile struct {
	// 【必填】媒体文件 ID，可以调用 [上传临时素材] 接口获取。
	//
	// [上传临时素材]: https://developer.work.weixin.qq.com/document/path/90389
	MediaID string `json:"media_id"`
}

// # 【企业微信】/【企业微信互联企业】消息 - 外链图文消息
type WechatwkNews struct {
	Articles []WechatwkNewsArticle `json:"articles"` // 【必填】[外链图文消息] 列表，一个图文消息支持 1 到 8 条图文。
}

// # 【企业微信】/【企业微信互联企业】消息 - 外链图文消息 - 明细
type WechatwkNewsArticle struct {
	Title       string `json:"title"`                 // 【必填】标题，不超过 128 个字节。
	Description string `json:"description,omitempty"` // 【可选】描述，不超过 512 个字节。
	URL         string `json:"url,omitempty"`         // 【可选】点击后跳转的链接，最长 2048 字节。请确保包含了协议头（http/https），小程序或者 URL 必须填写一个。
	PicURL      string `json:"picurl,omitempty"`      // 【可选】图文消息的封面图片链接，支持 JPG、PNG 格式，较好的效果为大图 1068×455，小图 150×150。
	AppID       string `json:"appid,omitempty"`       // 【可选】小程序 AppID，必须是与当前应用关联的小程序，AppID 和 PagePath 必须同时填写，填写后会忽略 URL 字段。
	PagePath    string `json:"pagepath,omitempty"`    // 【可选】点击消息卡片后的小程序页面，仅限本小程序内的页面。AppID 和 PagePath 必须同时填写，填写后会忽略 URL 字段。
}

// # 【企业微信】/【企业微信互联企业】消息 - 图文消息
type WechatwkMpnews struct {
	Articles []WechatwkMpnewsArticle `json:"articles"` // 【必填】[图文消息] 列表，一个图文消息支持 1 到 8 条图文。
}

// # 【企业微信】/【企业微信互联企业】消息 - 图文消息 - 明细
type WechatwkMpnewsArticle struct {
	// 【必填】图文消息缩略图的 MediaID，可以通过 [素材管理] 接口获得。此处 ThumbMediaID 即上传接口返回的 media_id。
	//
	// [素材管理]: https://developer.work.weixin.qq.com/document/path/90389
	ThumbMediaID string `json:"thumb_media_id"`
	// 【必填】标题，不超过 128 个字节。
	Title string `json:"title"`
	// 【必填】图文消息的内容，支持 HTML 标签，不超过 666K 个字节。
	Content string `json:"content"`
	// 【可选】图文消息的描述，不超过 512 个字节。
	Digest string `json:"digest,omitempty"`
	// 【可选】图文消息的作者，不超过 64 个字节。
	Author string `json:"author,omitempty"`
	// 【可选】图文消息点击 “阅读原文” 之后的页面链接。
	ContentSourceURL string `json:"content_source_url,omitempty"`
}

// # 【企业微信】/【企业微信互联企业】消息 - 小程序通知消息
type WechatwkMiniProgramNotice struct {
	AppID             string                                 `json:"appid"`                         // 【必填】小程序 AppID，必须是与当前应用关联的小程序。
	Page              string                                 `json:"page,omitempty"`                // 【可选】点击消息卡片后的小程序页面，仅限本小程序内的页面。该字段不填则消息点击后不跳转。
	Title             string                                 `json:"title"`                         // 【必填】消息标题，长度限制 4-12 个汉字。
	Description       string                                 `json:"description,omitempty"`         // 【可选】消息描述，长度限制 4-12 个汉字。
	EmphasisFirstItem bool                                   `json:"emphasis_first_item,omitempty"` // 【可选】是否放大第一个消息内容键值对项。
	ContentItems      []WechatwkMiniProgramNoticeContentItem `json:"content_item,omitempty"`        // 【可选】消息内容键值对，最多允许 10 个项。
}

// # 【企业微信】/【企业微信互联企业】消息 - 小程序通知消息 - 内容键值对项
type WechatwkMiniProgramNoticeContentItem struct {
	Key   string `json:"key"`   // 【必填】消息内容键，长度 10 个汉字以内。
	Value string `json:"value"` // 【必填】消息内容值，长度 30 个汉字以内。
}
