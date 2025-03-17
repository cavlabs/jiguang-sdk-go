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
	"github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push"
	"github.com/calvinit/jiguang-sdk-go/api/jpush/push/notification"
	"github.com/calvinit/jiguang-sdk-go/api/jums/audience"
	"github.com/calvinit/jiguang-sdk-go/api/jums/message"
	"github.com/calvinit/jiguang-sdk-go/jiguang"
)

func mockAudience() *audience.Audience {
	return &audience.Audience{
		// Tags:     []string{"tag1", "tag2"},
		// UserIDs:  []string{"user1", "user2"},
		// Segments: []string{"1000"},
		Apps: []audience.CustomChannel{
			{
				Instance: "app",
				Data:     []string{"170976fa8a9277fac6e"},
			},
		},
		Wechatoa: []audience.CustomChannel{
			{
				Instance: "wechatoa",
				Data:     []string{"oMtZu6kApZYEPJJWwyIHpWQ2L_DI"},
			},
		},
		Wechatmp: []audience.CustomChannel{
			{
				Instance: "wechatmp",
				Data:     []string{"oXNQs5B3LGA3xkU7-g2SdK3SsUaw"},
			},
		},
		SMS: []audience.CustomChannel{
			{
				Instance: "sms",
				Data:     []string{"18866007799"},
			},
		},
		Email: []audience.CustomChannel{
			{
				Instance: "email",
				Data:     []string{"support@jiguang.cn"},
			},
		},
		AlipayLife: []audience.CustomChannel{
			{
				Instance: "alipaylife",
				Data:     []string{"2088102733318286"},
			},
		},
		DingtalkCC: []audience.CustomChannel{
			{
				Instance: "dingtalkcc",
				Data:     []string{"a3c213779d163837895b30f47aaa94c3"},
			},
		},
		Wechatwk: []audience.CustomChannel{
			{
				Instance: "wechatwk",
				Data:     []string{"zhangsan"},
			},
		},
		WechatwkLinkedCorp: []audience.CustomChannel{
			{
				Instance: "wechatwk_linkedcorp",
				Data:     []string{"CorpId1/userid1", "CorpId2/userid2"},
			},
		},
	}
}

func mockMessage() *message.Message {
	return &message.Message{
		Apps: []message.App{
			{
				CID:      "8103a4c628a0b98974ec1949-711261d4-5f17-4d2f-a855-5e5a8909b26e",
				Platform: platform.All,
				Notification: &notification.Notification{
					Android: &notification.Android{
						Alert:     "Hi, JPush from JUMS!",
						Title:     "Send to Android",
						BuilderID: 1,
						LargeIcon: "http://www.jiguang.cn/largeIcon.jpg",
						Intent: &notification.Intent{
							URL: "intent:#Intent;component=com.jiguang.push/com.example.jpushdemo.SettingActivity;end",
						},
						Extras: map[string]interface{}{"newsid": 321},
					},
					IOS: &notification.IOS{
						Alert:    "Hi, JPush from JUMS!",
						Sound:    "default",
						Badge:    jiguang.String("+1"),
						ThreadID: "default",
						Extras:   map[string]interface{}{"newsid": 321},
					},
					QuickApp: &notification.QuickApp{
						Alert: "Hi, JPush from JUMS!",
						Title: "Send to QuickApp",
						Page:  "/page1",
					},
				},
				CustomMessage: &push.CustomMessage{
					Content:     "Hi, JPush from JUMS!",
					ContentType: "text",
					Title:       "msg",
					Extras:      map[string]interface{}{"key": "value"},
				},
				Options: &push.Options{
					TimeToLive:     jiguang.Int64(60),
					ApnsProduction: jiguang.Bool(false),
					ApnsCollapseID: "jiguang_test_201706011100",
				},
			},
		},
		Wechatoa: []message.Wechatoa{
			&message.WechatoaTemplate{
				Type:       0,
				TemplateID: "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
				URL:        "http://weixin.qq.com/download",
				MiniProgram: map[string]interface{}{
					"appid":    "xiaochengxuappid12345",
					"pagepath": "index?foo=bar",
				},
				Data: map[string]interface{}{
					"first": map[string]interface{}{
						"value": "恭喜你购买成功！",
						"color": "#173177",
					},
					"keyword1": map[string]interface{}{
						"value": "巧克力",
						"color": "#173177",
					},
					"keyword2": map[string]interface{}{
						"value": "39.8元",
						"color": "#173177",
					},
					"keyword3": map[string]interface{}{
						"value": "2014年9月22日",
						"color": "#173177",
					},
					"remark": map[string]interface{}{
						"value": "欢迎再次购买！",
						"color": "#173177",
					},
				},
			},
			&message.WechatoaSubscription{
				Type:       1,
				TemplateID: "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
				Page:       "http://weixin.qq.com/download",
				MiniProgram: map[string]interface{}{
					"appid":    "xiaochengxuappid12345",
					"pagepath": "index?foo=bar",
				},
				Data: map[string]interface{}{
					"name1": map[string]interface{}{
						"value": "广州腾讯科技有限公司",
					},
					"thing8": map[string]interface{}{
						"value": "广州腾讯科技有限公司",
					},
					"time7": map[string]interface{}{
						"value": "2019年8月8日",
					},
				},
			},
		},
		Wechatmp: []message.Wechatmp{
			{
				TemplateID:       "TEMPLATE_ID",
				Page:             "index",
				MiniProgramState: "developer",
				Lang:             "zh_CN",
				Data: map[string]interface{}{
					"number01": map[string]interface{}{
						"value": "339208499",
					},
					"date01": map[string]interface{}{
						"value": "2015年01月05日",
					},
					"site01": map[string]interface{}{
						"value": "TIT创意园",
					},
					"site02": map[string]interface{}{
						"value": "广州市新港中路397号",
					},
				},
			},
		},
		SMS: []message.SMS{
			&message.JSMS{
				SignID: 1,
				TempID: 1,
				TempParams: map[string]interface{}{
					"code": "1234",
				},
			},
			&message.CMPP{Content: "您登录系统的动态码为：111 ，动态码有效时间为1分钟，请注意保密。"},
		},
		Email: []message.Email{
			{
				Subject: "hello, ums email!",
				Text:    "ums email test. reply",
				Files: []string{
					"填写url1",
					"填写url2",
				},
			},
		},
		AlipayLife: []message.AlipayLife{
			{
				TemplateID: "e5326c1d5c71419893646ad9571a93e8",
				Context: &message.AlipayLifeContext{
					HeadColor:  "#85be53",
					URL:        "http://www.baidu.com",
					ActionName: "查看详情",
					Keywords: []message.AlipayLifeContextKeyword{
						{
							Color: "#85be53",
							Value: "你已经激活了上火星指令",
						},
						{
							Color: "#85be53",
							Value: "2020 09-29",
						},
						{
							Color: "#85be53",
							Value: "HI.真的 已经激活了指令",
						},
						{
							Color: "#85be53",
							Value: "指令已经通过",
						},
					},
				},
			},
		},
		DingtalkCC: []message.DingtalkCC{
			{
				Msg: &message.DingtalkCCMsg{
					MsgType: message.DingtalkCCMsgTypeText,
					Text: &message.DingtalkCCMsgText{
						Content: "test",
					},
				},
			},
		},
		Wechatwk: []message.Wechatwk{
			{
				MsgType: message.WechatwkMsgTypeText,
				Text: &message.WechatwkText{
					Content: "你的快递已到，请携带工卡前往邮件中心领取,聪明避开排队。",
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
			{
				MsgType: message.WechatwkMsgTypeImage,
				Image: &message.WechatwkImage{
					MediaID: "MEDIA_ID",
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
			{
				MsgType: message.WechatwkMsgTypeFile,
				File: &message.WechatwkFile{
					MediaID: "1Yv-zXfHjSjU-7LH-GwtYqDGS-zz6w22KmWAT5COgP7o",
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
			{
				MsgType: message.WechatwkMsgTypeNews,
				News: &message.WechatwkNews{
					Articles: []message.WechatwkNewsArticle{
						{
							Title:       "中秋节礼品领取",
							Description: "今年中秋节公司有豪礼相送",
							URL:         "URL",
							PicURL:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
						},
					},
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
			{
				MsgType: message.WechatwkMsgTypeMpnews,
				Mpnews: &message.WechatwkMpnews{
					Articles: []message.WechatwkMpnewsArticle{
						{
							ThumbMediaID:     "MEDIA_ID",
							Title:            "Title",
							Content:          "Content",
							Digest:           "Digest description",
							Author:           "Author",
							ContentSourceURL: "URL",
						},
					},
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
			{
				MsgType: message.WechatwkMsgTypeMiniProgramNotice,
				MiniProgramNotice: &message.WechatwkMiniProgramNotice{
					AppID:             "wx123123123123123",
					Page:              "pages/index?userid=zhangsan&orderid=123123123",
					Title:             "会议室预订成功通知",
					Description:       "4月27日 16:16",
					EmphasisFirstItem: true,
					ContentItems: []message.WechatwkMiniProgramNoticeContentItem{
						{
							Key:   "会议室",
							Value: "402",
						},
						{
							Key:   "会议地点",
							Value: "广州TIT-402会议室",
						},
						{
							Key:   "会议时间",
							Value: "2018年8月1日 09:00-09:30",
						},
						{
							Key:   "参与人员",
							Value: "周剑轩",
						},
					},
				},
				Safe:                   0,
				EnableDuplicateCheck:   1,
				DuplicateCheckInterval: 1800,
			},
		},
	}
}
