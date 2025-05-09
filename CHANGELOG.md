# Changelog

## [v0.4.3](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.4.3) - 2025-05-09

### 新特性

- “HMOS 通道通知” 增加 “前台展示控制” 字段；
- “批量推送参数” 增加可选的 “自定义消息转厂商通知内容” 字段。

### 优化
- 添加重试机制（示例），更新 HTTP 客户端配置。

---

## [v0.4.2](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.4.2) - 2025-04-17

### 新特性

- 新增 “推送计划管理” API 的支持。

---

## [v0.4.1](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.4.1) - 2025-04-02

### 新特性

- 新增 “测试设备管理” API 的支持；
- 新增 “测试模式推送” 选项参数；
- 新增 “蔚来系统通道” 厂商推送参数和统计数据。

---

## [v0.4.0](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.4.0) - 2025-03-24

### 重构

- 利用了 Go 1 的兼容性承诺，本 SDK 使用了最新版本的 Go SDK，但不会破坏原有的 API 兼容性承诺（Go 1.16+）；
- 添加了 GitHub 工作流，以用于自动化构建和测试，达成上述目标；
- 更新代码文档注释以提高多个文件的清晰度和一致性。

---

## [v0.3.0](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.3.0) - 2025-03-17

### 新特性

- 全面支持 “极光统一消息（JUMS v1）” 相关功能模块。

---

## [v0.2.0](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.2.0) - 2025-02-10

### 新特性

- 全面支持 “极光短信（JSMS v1）” 相关功能模块。

### 优化

- 优化了日志消息输出前缀，对不同类型 API 的请求和响应日志进行了区分，使之更加清晰明了；
- 优化了示例代码的演示对象初始化方式，使之更加简洁易懂。

---

## [v0.1.3](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.1.3) - 2025-02-07

### 修复

- **接口调用修复**：修改返回的 `nil` 接口为带有动态类型的 `nil`，避免调用接口时发生方法调用失败。

### 优化

- **安全性改进**：对 `Authorization` 请求头中的敏感令牌信息进行日志输出屏蔽，以增强安全性。

---

## [v0.1.2](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.1.2) - 2025-01-15

### 新特性

- 增加了对普通推送 API 的 SM2 加密推送功能。

---

## [v0.1.1](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.1.1) - 2024-12-31

### 修复

- 修复了 README 中的一些示例代码错误问题；
- 优化了 HTTP 日志输出时源文件位置的显示，使之更加清晰明了；
- 其他一些细节优化和文档完善。

---

## [v0.1.0](https://github.com/calvinit/jiguang-sdk-go/releases/tag/v0.1.0) - 2024-12-31

🎉🎉🎉 **历经长时间细心打磨，首个正式版本终于发布啦！** 🎉🎉🎉

### 新特性

- **SDK 初始版本**：基于 Go 语言封装了极光服务端 REST API，全面支持 “极光推送（JPush）” 相关功能模块。