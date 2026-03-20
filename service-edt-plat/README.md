# EDT Course Service MCP

这是一个基于MCP（Model Context Protocol）的课程服务，用于与外部课程API进行交互，支持创建课程、章节和课时。

## 项目结构

```
├── client/            # 外部服务客户端实现
│   └── edt_client.go  # 与外部课程API通信的客户端
├── config/            # 配置文件目录
│   ├── config.env.example  # 配置文件示例
│   └── config.env     # 配置文件
├── models/            # 数据模型定义
│   └── models.go      # 定义数据结构
├── tools/             # 工具实现目录
│   ├── course.go      # 创建课程工具
│   ├── chapter.go     # 创建章节工具
│   └── lesson.go      # 创建课时工具
├── utils/             # 工具函数目录
│   └── utils.go       # 通用工具函数
├── main.go            # 主入口文件
├── go.mod             # Go模块文件
├── go.sum             # Go依赖校验文件
└── README.md          # 项目说明文档
```

## 功能特性

- **创建课程**：支持创建新的课程，包括标题、描述、价格等参数
- **创建章节**：支持为课程创建章节，包括标题、描述、排序等参数
- **创建课时**：支持为章节创建课时，包括标题、描述、视频URL等参数
- **多模式运行**：支持HTTP、SSE和标准输入输出三种运行模式

## 安装和使用

### 前提条件

- Go 1.25.0 或更高版本

### 安装步骤

1. 克隆项目

2. 安装依赖
   ```bash
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

配置文件 `config/config.env` 包含以下配置项：

- `SERVICE_PORT`：服务端口，默认为8080
- `SERVICE_MODE`：服务模式，支持 `http`、`sse`、`stdio`
- `EDT_API_BASE_URL`：外部课程API的基础URL

## 工具说明

### create_course

创建新课程

**参数**：
- `title`（必需）：课程标题
- `description`（必需）：课程描述
- `price`（必需）：课程价格
- `category_id`（必需）：课程分类ID
- `detail`（可选）：课程详情
- `cover_image`（可选）：课程封面图片URL
- `original_price`（可选）：课程原价
- `level`（可选）：课程难度级别
- `duration`（可选）：课程时长（分钟）
- `status`（可选）：课程状态
- `start_date`（可选）：课程开始日期
- `end_date`（可选）：课程结束日期
- `target_students`（可选）：目标学生
- `requirements`（可选）：课程要求
- `learning_objectives`（可选）：学习目标
- `certificate`（可选）：证书信息
- `course_highlights`（可选）：课程亮点（JSON数组字符串）
- `teaching_methods`（可选）：教学方法（JSON数组字符串）
- `support_services`（可选）：支持服务（JSON数组字符串）
- `career_prospects`（可选）：职业前景（JSON数组字符串）

### create_chapter

创建新章节

**参数**：
- `course_id`（必需）：课程ID
- `title`（必需）：章节标题
- `description`（可选）：章节描述
- `sort_order`（可选）：排序顺序

### create_lesson

创建新课时

**参数**：
- `chapter_id`（必需）：章节ID
- `title`（必需）：课时标题
- `description`（可选）：课时描述
- `video_url`（可选）：视频URL
- `document_url`（可选）：文档URL
- `content`（可选）：课时内容
- `duration`（可选）：课时时长（分钟）
- `sort_order`（可选）：排序顺序
- `is_free`（可选）：是否免费
- `type`（可选）：课时类型

## MCP使用方法

```json
{
  "mcpServers": {
    "edu-stdio": {
      "command": "/System/Volumes/Data/webcode/mcp-tools-v2/output/mcp-server",
      "env": {
        "SERVER_MODE": "stdio",
        "APIBaseURL": "https://edu.eluup.com/",
        "APIKey": "{}"
      }
    }
  }
}
```

## 示例调用

### HTTP模式

```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": "1",
    "method": "initialize",
    "params": {
      "clientInfo": {
        "name": "test-client",
        "version": "1.0.0"
      },
      "capabilities": {
        "tools": {
          "listChanged": true
        },
        "prompts": {
          "listChanged": true
        },
        "resources": {
          "subscribe": true,
          "listChanged": true
        }
      }
    }
  }'
```

```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": "2",
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
  }'
```

### Stdio模式

```bash
echo '{"jsonrpc":"2.0","id":"1","method":"initialize","params":{"clientInfo":{"name":"test-client","version":"1.0.0"},"capabilities":{"tools":{"listChanged":true},"prompts":{"listChanged":true},"resources":{"subscribe":true,"listChanged":true}}}}' | ./mcp-service
```

```bash
echo '{"jsonrpc":"2.0","id":"2","method":"call","params":{"thought":"创建一个新的测试课程","name":"create_course","params":{"title":"测试课程","description":"这是一个测试课程","price":"99.99","category_id":"1"}}}' | ./mcp-service
```

## 技术栈

- Go 1.25.0
- mcp-go v0.45.0
- 标准库：net/http, encoding/json, fmt, log

## 许可证

MIT

## mcp使用方法：
{
  "mcpServers": {
    "edu-stdio": {
      "command": "/System/Volumes/Data/webcode/mcp_courses/service-edt-plat/mcp-service",
      "env": {
        "SERVER_MODE": "stdio",
        "edtApiBaseURL": "https://edu.eluup.com/",
        "servicePort": "6080",
        "serviceMode": "stdio"
      }
    }
  }
}