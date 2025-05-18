# Gnostic 项目简介

## 项目概述
Gnostic 是一个用 Go 语言开发的命令行工具，用于在 JSON/YAML 格式的 OpenAPI 描述和 Protocol Buffer 表示之间进行转换。

## 核心功能
1. OpenAPI 描述转换
   - 支持 JSON 和 YAML 格式的 OpenAPI 描述
   - 转换为 Protocol Buffer 表示
   - 支持反向转换（Protocol Buffer 到 OpenAPI）

2. 插件系统
   - 提供插件接口，支持多种语言的集成
   - 现有插件包括：
     - gnostic-analyze：分析 OpenAPI 描述
     - gnostic-vocabulary：生成 API 接口词汇摘要
     - gnostic-grpc：生成 gRPC 相关代码
     - gnostic-go-generator：生成 Go 客户端代码

3. 代码生成
   - 自动生成 Protocol Buffer 模型
   - 支持类型安全的代码生成
   - 特别适用于强类型语言（如 Go 和 Dart）

## 技术栈
- 主要语言：Go
- 核心依赖：
  - Protocol Buffer Compiler (protoc)
  - Go modules
  - Protocol Buffer 工具链

## 项目状态
- 当前版本：预发布阶段
- 开发状态：积极维护
- 许可证：Apache 2.0

## 项目目标
1. 提供稳定可靠的 OpenAPI 描述转换工具
2. 支持 API 工具链的自动化集成
3. 确保类型安全和代码质量
4. 提供可扩展的插件系统 