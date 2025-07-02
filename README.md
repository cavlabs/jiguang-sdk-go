# jiguang-sdk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/cavlabs/jiguang-sdk-go.svg)](https://pkg.go.dev/github.com/cavlabs/jiguang-sdk-go)
[![Support Go 1.16+](https://img.shields.io/badge/Go-1.16+-blue.svg?style=flat-square)](https://go.dev/doc/devel/release)
[![Release](https://img.shields.io/github/v/release/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/releases)
![CI Status](https://img.shields.io/github/actions/workflow/status/cavlabs/jiguang-sdk-go/ci.yml?label=CI&logo=github)
[![Go Report Card](https://goreportcard.com/badge/github.com/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://goreportcard.com/report/github.com/cavlabs/jiguang-sdk-go)
[![Issues](https://img.shields.io/github/issues/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/issues)
[![PRs](https://img.shields.io/github/issues-pr/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/pulls)
[![License: Apache-2.0](https://img.shields.io/github/license/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go?tab=Apache-2.0-1-ov-file#readme)

## 简介

`jiguang-sdk-go` 是基于 Go 的极光 REST API
封装开发包，参考了极光官方提供的 [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java) 实现。
它致力于为开发者提供便捷、轻量和高效的服务端集成方式，并支持最新的 API 功能。

### 特性

- 全面支持 “极光推送（JPush）”、“极光短信（JSMS v1）” 和 “极光统一消息（JUMS v1）” 相关功能模块；
- 简单易用的 Go 接口；
- 支持 Go 1.16 及其以上版本。

---

## 一、极光文档

以下是本 SDK 支持的极光 REST API 功能模块及官方文档链接：

### [1. 极光推送（JPush）](https://docs.jiguang.cn/jpush/server/push)

- [x] [应用管理 - Admin API v1](https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1)
- [x] [设备/标签/别名 - Device API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device)
- [x] [推送 - Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push)
- [x] [分组推送 - Group Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_grouppush)
- [x] [定时任务 - Schedule API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule)
- [x] [文件管理 - File API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file)
- [x] [图片管理 - Image API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image)
- [x] [推送统计 - Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)
- [x] [分组推送统计 - Group Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)

### [2. 极光短信（JSMS v1）](https://docs.jiguang.cn/jsms/server/restapi)

- [x] [短信签名 - Sign API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign)
- [x] [短信模板 - Template API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates)
- [x] [短信发送 - Code/Message API](https://docs.jiguang.cn/jsms/server/rest_api_jsms)
- [x] [短信定时发送 - Schedule Message API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule)
- [x] [短信余量查询 - Account Dev/App Balance API](https://docs.jiguang.cn/jsms/server/rest_jsms_api_account)
- [x] [短信回执 - Inquire Report/Reply API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire)
- [x] [回调接口 - Callback Server (SMS_SIGN, SMS_TEMPLATE, SMS_REPORT, SMS_REPLY)](https://docs.jiguang.cn/jsms/server/callback)

### [3. 极光统一消息（JUMS v1）](https://docs.jiguang.cn/jums/server/restapi)

- [x] [普通消息发送 - Custom Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message)
- [x] [模板消息发送 - Template Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_template_message)
- [x] [消息撤回 - Retract Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_retract_message)
- [x] [用户管理 - User API](https://docs.jiguang.cn/jums/server/rest_api_jums_user)
- [x] [素材管理 - Material API](https://docs.jiguang.cn/jums/server/rest_api_jums_material)
- [x] [获取通道 Token - Token API](https://docs.jiguang.cn/jums/server/rest_api_jums_token)
- [x] [回调接口 - Callback Server (目标有效/无效, 提交成功/失败, 送达成功/失败, 点击, 撤回成功/失败)](https://docs.jiguang.cn/jums/advanced/callback)

---

## 二、快速开始

1. 使用以下命令安装 SDK：
    ```bash
    go get github.com/cavlabs/jiguang-sdk-go@latest
    ```

---

## 三、使用示例

1. 在项目中引入 SDK：
    ```go
    import sdk "github.com/cavlabs/jiguang-sdk-go"
    ```

2. 示例代码（假设用于极光应用的「普通推送」）：
    ```go
    package main

    import (
        "context"
        "fmt"
        "os"
    
        "github.com/cavlabs/jiguang-sdk-go/api/jpush/device/platform"
        "github.com/cavlabs/jiguang-sdk-go/api/jpush/push"
    )
    
    func main() {
        pushAPIv3, _ := push.NewAPIv3Builder().
    		SetAppKey(os.Getenv("JPUSH_APP_KEY")).
    		SetMasterSecret(os.Getenv("JPUSH_MASTER_SECRET")).
    		Build()
    
        param := &push.SendParam{
            Platform: platform.All,
            Audience: push.BroadcastAuds,
            Notification: &push.Notification{
                Alert: "Hello, JPush!",
            },
        }
    
        result, err := pushAPIv3.Send(context.Background(), param)
        if err != nil {
            panic(err)
        }
    
        if result.IsSuccess() {
            fmt.Printf("Send success, MsgID: %s, SendNo: %s\n", result.MsgID, result.SendNo)
        } else {
            fmt.Printf("Send failed: %s\n", result.Error)
        }
    }
    ```

3. 查看完整示例代码：https://github.com/cavlabs/jiguang-sdk-go/tree/main/examples

---

## 四、支持与贡献

- 如果遇到问题，请在 [Issues 页面](https://github.com/cavlabs/jiguang-sdk-go/issues/new) 提交。
- 欢迎提交 Pull Request，为项目贡献代码。

---

## 五、许可证

本项目采用 [Apache 2.0](https://github.com/cavlabs/jiguang-sdk-go?tab=Apache-2.0-1-ov-file#readme) 许可证。

---

## 六、Git Hooks 设置

为了确保代码质量并执行必要的预推送检查，项目包含了一些 Git Hooks 文件。如果您想使用这些 Git Hooks，请按照以下步骤操作：

1. 克隆项目并进入项目目录后，创建符号链接来启用 Git Hooks：
    ```bash
    ln -s ../../.githooks/pre-push .git/hooks/pre-push
    ```
2. 确保您为 Git Hooks 脚本添加了执行权限：
    ```bash
    chmod +x .githooks/pre-push
    ```
3. 确保您在推送之前已经正确地配置了所需的 Git Hooks。如果您没有设置，它们将不会被自动运行。

---

## 七、致谢

- 感谢 [tjfoc/gmsm](https://github.com/tjfoc/gmsm) 库提供的支持，帮助我们实现了 SM2 加解密算法的核心功能。

---

## 八、参考链接

- [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java)
- [jsms-api-java-client](https://github.com/jpush/jsms-api-java-client)
- [极光文档中心](https://docs.jiguang.cn)
