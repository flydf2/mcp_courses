# MCP Courses Project

这是一个基于 MCP（Model Context Protocol）的课程服务项目，包含课程管理和 AI 课程生成等功能。

## 项目结构

```
├── docs/                  # 项目文档目录
│   ├── Demo.md                      # 演示文档
│   ├── global_course_development_summary.md  # 全局课程开发总结
│   ├── global_小学六年级全科学习课程.md       # 小学六年级全科学习课程
│   └── service_ai_course_generation_rules.md  # AI 课程生成规则
├── service-edt-plat/      # 课程服务 MCP 实现
│   ├── client/            # 外部服务客户端
│   ├── config/            # 配置文件
│   ├── models/            # 数据模型
│   ├── utils/             # 工具函数
│   ├── main.go            # 主入口
│   ├── go.mod             # Go 模块文件
│   ├── go.sum             # 依赖校验文件
│   ├── mcp-service        # 编译后的可执行文件
│   └── README.md          # 服务说明文档
```

## 项目功能

### 1. 课程服务 MCP（service-edt-plat）

- **创建课程**：支持创建新的课程，包括标题、描述、价格等参数
- **创建章节**：支持为课程创建章节，包括标题、描述、排序等参数
- **创建课时**：支持为章节创建课时，包括标题、描述、视频URL等参数
- **多模式运行**：支持 HTTP、SSE 和标准输入输出三种运行模式

### 2. 文档管理（docs/）

- **课程开发文档**：包含课程开发的总结和规范
- **AI 课程生成规则**：定义 AI 生成课程的规则和标准
- **示例课程**：提供小学六年级全科学习课程的示例

## 安装和使用

### 前提条件

- Go 1.25.0 或更高版本

### 安装步骤

1. 克隆项目

2. 进入服务目录并安装依赖
   ```bash
   cd service-edt-plat
   go mod tidy
   ```

3. 配置环境变量
   ```bash
   cp config/config.env.example config/config.env
   # 编辑 config/config.env 文件，设置相应的配置
   ```

4. 构建项目
   ```bash
   go build -o mcp-service .
   ```

5. 运行服务
   ```bash
   ./mcp-service
   ```

## 配置说明

### 服务配置（service-edt-plat/config/config.env）

- `SERVICE_PORT`：服务端口，默认为 8080
- `SERVICE_MODE`：服务模式，支持 `http`、`sse`、`stdio`
- `EDT_API_BASE_URL`：外部课程 API 的基础 URL

## MCP 使用方法

### 配置 MCP 服务器

在 Cursor 或其他 MCP 客户端中配置：

```json
{
  "mcpServers": {
    "edu-stdio": {
      "command": "/path/to/mcp_courses/service-edt-plat/mcp-service",
      "env": {
        "SERVER_MODE": "stdio",
        "edtApiBaseURL": "https://edu.eluup.com/",
        "servicePort": "6080",
        "serviceMode": "stdio"
      }
    }
  }
}
```

### 调用示例

#### 创建课程

```json
{
  "jsonrpc": "2.0",
  "id": "1",
  "method": "call",
  "params": {
    "thought": "创建一个新的测试课程",
    "name": "create_course",
    "params": {
      "title": "测试课程",
      "description": "这是一个测试课程",
      "price": "99.99",
      "category_id": "1"
    }
  }
}
```

#### 创建章节

```json
{
  "jsonrpc": "2.0",
  "id": "2",
  "method": "call",
  "params": {
    "thought": "为课程创建章节",
    "name": "create_chapter",
    "params": {
      "course_id": "1",
      "title": "第一章 课程介绍",
      "description": "介绍课程的基本内容和学习目标"
    }
  }
}
```

#### 创建课时

```json
{
  "jsonrpc": "2.0",
  "id": "3",
  "method": "call",
  "params": {
    "thought": "为章节创建课时",
    "name": "create_lesson",
    "params": {
      "chapter_id": "1",
      "title": "1.1 课程概述",
      "description": "课程的整体介绍和学习路径",
      "video_url": "https://example.com/video1.mp4"
    }
  }
}
```

## 文档说明

- **Demo.md**：项目演示文档
- **global_course_development_summary.md**：全局课程开发总结
- **global_小学六年级全科学习课程.md**：小学六年级全科学习课程示例
- **service_ai_course_generation_rules.md**：AI 课程生成规则

## 技术栈

- **Go 1.25.0**：主要开发语言
- **MCP**：Model Context Protocol，用于工具调用
- **标准库**：net/http, encoding/json, fmt, log

## 许可证

MIT

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目。

## 联系方式

- GitHub: https://github.com/flydf2/mcp_courses
