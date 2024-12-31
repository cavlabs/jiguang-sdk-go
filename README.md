# jiguang-sdk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/calvinit/jiguang-sdk-go.svg)](https://pkg.go.dev/github.com/calvinit/jiguang-sdk-go)
[![GitHub release](https://img.shields.io/github/v/release/calvinit/jiguang-sdk-go)](https://github.com/calvinit/jiguang-sdk-go/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/calvinit/jiguang-sdk-go)](https://golang.org/doc/devel/release.html)
[![Go Report Card](https://goreportcard.com/badge/github.com/calvinit/jiguang-sdk-go)](https://goreportcard.com/report/github.com/calvinit/jiguang-sdk-go)
[![GitHub issues](https://img.shields.io/github/issues/calvinit/jiguang-sdk-go)](https://github.com/calvinit/jiguang-sdk-go/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/calvinit/jiguang-sdk-go)](https://github.com/calvinit/jiguang-sdk-go/pulls)
[![GitHub License](https://img.shields.io/github/license/calvinit/jiguang-sdk-go)](https://github.com/calvinit/jiguang-sdk-go?tab=Apache-2.0-1-ov-file#readme)

## 简介

`jiguang-sdk-go` 是基于 Go 的极光 REST API
封装开发包，参考了极光官方提供的 [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java) 实现。
它致力于为开发者提供便捷、高效的服务端集成方式，并支持最新的 API 功能。

### 特性

- 全面支持 “极光推送（JPush）” 相关功能模块；
- 简单易用的 Go 接口；
- 支持 Go 1.16 及其以上版本。

---

## 一、极光文档

以下是 SDK 支持的极光 REST API 功能模块及官方文档链接：

### 1. 极光推送（JPush）

- [x] [应用管理 - Admin API v1](https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1)
- [x] [设备/标签/别名 - Device API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device)
- [x] [推送 - Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push)
- [x] [分组推送 - Group Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_grouppush)
- [x] [定时任务 - Schedule API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule)
- [x] [文件管理 - File API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file)
- [x] [图片管理 - Image API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image)
- [x] [推送统计 - Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)
- [x] [分组推送统计 - Group Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)

---

## 二、快速开始

1. 使用以下命令安装 SDK：
    ```bash
    go get github.com/calvinit/jiguang-sdk-go@latest
    ```

---

## 三、使用示例

1. 在项目中引入 SDK：
    ```go
    import sdk "github.com/calvinit/jiguang-sdk-go"
    ```

2. 示例代码（假设用于极光应用的「普通推送」）：
    ```go
    package main

    import (
        "context"
        "fmt"
        "os"
    
        "github.com/calvinit/jiguang-sdk-go/api/jpush/device/platform"
        "github.com/calvinit/jiguang-sdk-go/api/jpush/push"
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

3. 查看完整示例代码：https://github.com/calvinit/jiguang-sdk-go/tree/main/examples

---

## 四、支持与贡献

- 如果遇到问题，请在 [Issues 页面](https://github.com/calvinit/jiguang-sdk-go/issues/new) 提交。
- 欢迎提交 Pull Request，为项目贡献代码。

---

## 五、许可证

本项目采用 [Apache 2.0](https://github.com/calvinit/jiguang-sdk-go?tab=Apache-2.0-1-ov-file#readme) 许可证。

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

## 七、参考链接

- [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java)
- [极光文档中心](https://docs.jiguang.cn)