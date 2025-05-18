# Gnostic 产品上下文

## 应用场景
1. API 开发工具链
   - OpenAPI 描述文件的处理和分析
   - API 文档的自动生成
   - 客户端代码的自动生成

2. 微服务架构
   - 服务间接口定义
   - API 网关集成
   - 服务发现和注册

3. 开发工作流
   - CI/CD 流程中的 API 验证
   - 自动化测试
   - 文档生成

## 使用方式
1. 命令行工具
   ```bash
   # 基本转换
   gnostic --pb-out=. api.json
   
   # 使用插件
   gnostic api.json --analyze_out=.
   gnostic api.json --vocabulary_out=.
   ```

2. 插件开发
   - 基于 Protocol Buffer 的插件接口
   - 支持多种编程语言
   - 可扩展的插件系统

3. 集成方式
   - 作为独立工具使用
   - 作为库集成到其他项目
   - 通过插件系统扩展功能

## 目标用户
1. API 开发者
   - 需要处理 OpenAPI 描述文件
   - 需要生成客户端代码
   - 需要自动化 API 工具链

2. 工具开发者
   - 开发 API 相关工具
   - 需要处理 OpenAPI 描述
   - 需要扩展 API 工具链

3. 运维团队
   - 管理 API 网关
   - 监控 API 使用情况
   - 自动化部署流程

## 产品优势
1. 标准化
   - 基于 OpenAPI 规范
   - 使用 Protocol Buffer 标准
   - 支持多种格式转换

2. 可扩展性
   - 插件系统
   - 多语言支持
   - 自定义扩展

3. 开发效率
   - 自动化代码生成
   - 类型安全
   - 工具链集成 