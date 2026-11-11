# jiguang-sdk-go

[![Go Reference](https://pkg.go.dev/badge/github.com/cavlabs/jiguang-sdk-go/v2.svg)](https://pkg.go.dev/github.com/cavlabs/jiguang-sdk-go/v2)
[![Support Go 1.27+](https://img.shields.io/badge/Go-1.27+-blue.svg?style=flat-square)](https://go.dev/doc/devel/release)
[![Release](https://img.shields.io/github/v/release/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/releases)
![CI Status](https://img.shields.io/github/actions/workflow/status/cavlabs/jiguang-sdk-go/ci.yml?style=flat-square&label=CI&logo=github)
[![Go Report Card](https://goreportcard.com/badge/github.com/cavlabs/jiguang-sdk-go?style=flat-square)](https://goreportcard.com/report/github.com/cavlabs/jiguang-sdk-go)
[![Issues](https://img.shields.io/github/issues/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/issues)
[![PRs](https://img.shields.io/github/issues-pr/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go/pulls)
[![License: Apache-2.0](https://img.shields.io/github/license/cavlabs/jiguang-sdk-go.svg?style=flat-square)](https://github.com/cavlabs/jiguang-sdk-go?tab=Apache-2.0-1-ov-file#readme)

Go 语言的极光 REST API 服务端 SDK，参考官方 [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java)，面向服务端集成，并跟进最新 API 能力。

**模块路径**：`github.com/cavlabs/jiguang-sdk-go/v2` · **最低 Go**：1.27+ · **根 module**：零第三方依赖

### 特性

- 覆盖极光推送（JPush）、极光短信（JSMS v1）、极光统一消息（JUMS v1）
- 产品级入口 `jpush.New` / `jsms.New` / `jums.New`，与各 `NewAPIv*Builder` 双轨并存
- 可注入自定义 HTTP Client / Logger；示例见 [`examples/adapter`](https://github.com/cavlabs/jiguang-sdk-go/tree/v2/examples/adapter)
- 内嵌 SM2 能力（Push `SendWithSM2` 等），无需额外依赖

---

## 极光文档与能力覆盖

### [极光推送（JPush）](https://docs.jiguang.cn/jpush/server/push)

- [x] [应用管理 - Admin API v1](https://docs.jiguang.cn/jpush/server/push/rest_api_admin_api_v1)
- [x] [设备/标签/别名 - Device API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_device)
- [x] [推送 - Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push)
- [x] [分组推送 - Group Push API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push_grouppush)
- [x] [定时任务 - Schedule API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_push_schedule)
- [x] [文件管理 - File API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_file)
- [x] [图片管理 - Image API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_image)
- [x] [推送统计 - Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)
- [x] [分组推送统计 - Group Report API v3](https://docs.jiguang.cn/jpush/server/push/rest_api_v3_report)

### [极光短信（JSMS v1）](https://docs.jiguang.cn/jsms/server/restapi)

- [x] [短信签名 - Sign API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_sign)
- [x] [短信模板 - Template API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_templates)
- [x] [短信发送 - Code/Message API](https://docs.jiguang.cn/jsms/server/rest_api_jsms)
- [x] [短信定时发送 - Schedule Message API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_schedule)
- [x] [短信余量查询 - Account Dev/App Balance API](https://docs.jiguang.cn/jsms/server/rest_jsms_api_account)
- [x] [短信回执 - Inquire Report/Reply API](https://docs.jiguang.cn/jsms/server/rest_api_jsms_inquire)
- [x] [回调接口 - Callback Server](https://docs.jiguang.cn/jsms/server/callback)

### [极光统一消息（JUMS v1）](https://docs.jiguang.cn/jums/server/restapi)

- [x] [普通消息发送 - Custom Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_custom_message)
- [x] [模板消息发送 - Template Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_template_message)
- [x] [消息撤回 - Retract Message API](https://docs.jiguang.cn/jums/server/rest_api_jums_retract_message)
- [x] [用户管理 - User API](https://docs.jiguang.cn/jums/server/rest_api_jums_user)
- [x] [素材管理 - Material API](https://docs.jiguang.cn/jums/server/rest_api_jums_material)
- [x] [获取通道 Token - Token API](https://docs.jiguang.cn/jums/server/rest_api_jums_token)
- [x] [回调接口 - Callback Server](https://docs.jiguang.cn/jums/advanced/callback)

---

## 安装

```bash
go get github.com/cavlabs/jiguang-sdk-go/v2@latest
```

当前主线为 v2（见 `version.go`）。请使用带 `/v2` 的 import path。

---

## 快速开始

推荐使用产品级入口。下面以 JPush 普通推送为例（凭证请用环境变量或自有配置注入，勿硬编码）：

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/cavlabs/jiguang-sdk-go/v2/api/jpush"
    "github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/device/platform"
    "github.com/cavlabs/jiguang-sdk-go/v2/api/jpush/push"
)

func main() {
    client, err := jpush.New(
        jpush.WithApp(os.Getenv("JPUSH_APP_KEY"), os.Getenv("JPUSH_MASTER_SECRET")),
    )
    if err != nil {
        panic(err)
    }

    param := &push.SendParam{
        Platform: platform.All,
        Audience: push.BroadcastAuds,
        Notification: &push.Notification{
            Alert: "Hello, JPush!",
        },
    }

    // 默认即为 Push REST v3；也可写 client.Push.V3.Send(...)
    result, err := client.Push.Send(context.Background(), param)
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

### 其他入口

```go
// JSMS
sms, err := jsms.New(jsms.WithApp(appKey, masterSecret))

// JUMS
ums, err := jums.New(jums.WithChannel(channelKey, masterSecret))

// 进阶/兼容：仍可按能力使用 Builder
pushAPIv3, err := push.NewAPIv3Builder().
SetAppKey(appKey).
SetMasterSecret(masterSecret).
Build()
```

JPush 侧常用鉴权组合：

| Option      | 用途                                                |
|-------------|-----------------------------------------------------|
| `WithApp`   | Push / Device / File / Image / Schedule / Report 等 |
| `WithGroup` | GPush / GReport                                     |
| `WithDev`   | Admin 等                                            |

缺少某类凭证时 `New` 仍可成功；调用未配置的子 API 会返回对应的 `api.ErrNil*`。判断业务是否成功请用结果的 `IsSuccess()`，勿只看 HTTP 状态码。

完整示例（含各 API 与 Logger/HTTP 适配）：https://github.com/cavlabs/jiguang-sdk-go/tree/v2/examples

---

## 开发与贡献

- 问题反馈：[Issues](https://github.com/cavlabs/jiguang-sdk-go/issues/new)
- 欢迎提交 Pull Request
- 本地多 module 开发可复制 `go.work.example` 为 `go.work`（该文件已 gitignore）
- 可选启用 pre-push（校验 release tag 与 `version.go` 一致）：

```bash
ln -s ../../.githooks/pre-push .git/hooks/pre-push
chmod +x .githooks/pre-push
```

---

## 许可证

本项目采用 [Apache 2.0](LICENSE) 许可证。详见根目录 `LICENSE` 与 `NOTICE`。

---

## 致谢与参考

- [tjfoc/gmsm](https://github.com/tjfoc/gmsm) — 内嵌 SM2/SM3 实现的来源
- [jiguang-sdk-java](https://github.com/jpush/jiguang-sdk-java)
- [jsms-api-java-client](https://github.com/jpush/jsms-api-java-client)
- [极光文档中心](https://docs.jiguang.cn)
