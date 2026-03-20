# AI生成课程信息规则

## 1. 概述

本文档定义了AI生成课程信息的规则，基于外部课程API接口和数据示例，确保生成的课程信息符合系统要求并具有高质量。

## 2. 数据结构要求

### 2.1 课程结构

```json
{
  "title": "课程标题",
  "description": "课程描述",
  "detail": "课程详情",
  "cover_image": "课程封面图片URL",
  "price": 课程价格,
  "original_price": 课程原价,
  "category_id": 课程分类ID,
  "level": "课程难度级别",
  "duration": 课程时长（分钟）,
  "status": "课程状态",
  "start_date": "课程开始日期",
  "end_date": "课程结束日期",
  "target_students": "目标学员",
  "requirements": "课程要求",
  "learning_objectives": "学习目标",
  "certificate": "证书信息",
  "course_highlights": ["课程亮点1", "课程亮点2"],
  "teaching_methods": ["教学方法1", "教学方法2"],
  "support_services": ["支持服务1", "支持服务2"],
  "career_prospects": ["职业前景1", "职业前景2"]
}
```

### 2.2 章节结构

```json
{
  "title": "章节标题",
  "description": "章节描述",
  "sort_order": 排序顺序,
  "lessons": [
    {
      "title": "课时标题",
      "description": "课时描述",
      "video_url": "视频URL",
      "document_url": "文档URL",
      "content": "课时内容",
      "duration": 课时时长（秒）,
      "sort_order": 排序顺序,
      "is_free": 是否免费,
      "type": "课时类型"
    }
  ]
}
```

## 3. 字段生成规则

### 3.1 课程字段规则

| 字段名 | 类型 | 必填 | 生成规则 |
|--------|------|------|----------|
| `title` | string | 是 | 简洁明了，反映课程主题，2-50个字符 |
| `description` | string | 是 | 简要描述课程内容和价值，50-200个字符 |
| `detail` | string | 否 | 详细介绍课程内容、结构和学习收获，300-1000个字符 |
| `cover_image` | string | 否 | 有效的图片URL，建议使用教育相关图片 |
| `price` | number | 是 | 合理的课程价格，范围1-9999 |
| `original_price` | number | 否 | 原价，大于等于价格，默认0表示无原价 |
| `category_id` | integer | 是 | 有效的分类ID，根据系统中存在的分类选择 |
| `level` | string | 否 | 课程难度级别：beginner, intermediate, advanced |
| `duration` | integer | 否 | 课程总时长（分钟），根据课程内容合理估计 |
| `status` | string | 否 | 课程状态，默认"published" |
| `start_date` | string | 否 | 课程开始日期，ISO格式，默认"1990-01-01T00:00:00+08:00" |
| `end_date` | string | 否 | 课程结束日期，ISO格式，默认"1990-01-01T00:00:00+08:00" |
| `target_students` | string | 否 | 描述目标学员群体，50-200个字符 |
| `requirements` | string | 否 | 课程前置要求，30-150个字符 |
| `learning_objectives` | string | 否 | 学习目标和收获，50-200个字符 |
| `certificate` | string | 否 | 证书信息，30-150个字符 |
| `course_highlights` | array<string> | 否 | 课程亮点，3-5项，每项20-50个字符 |
| `teaching_methods` | array<string> | 否 | 教学方法，2-4项，每项15-40个字符 |
| `support_services` | array<string> | 否 | 支持服务，2-4项，每项15-40个字符 |
| `career_prospects` | array<string> | 否 | 职业前景，2-4项，每项15-40个字符 |

### 3.2 章节字段规则

| 字段名 | 类型 | 必填 | 生成规则 |
|--------|------|------|----------|
| `title` | string | 是 | 章节标题，反映章节内容，5-30个字符 |
| `description` | string | 否 | 章节描述，30-100个字符 |
| `sort_order` | integer | 否 | 排序顺序，从1开始递增 |
| `lessons` | array | 否 | 课时列表，每个章节至少1个课时 |

### 3.3 课时字段规则

| 字段名 | 类型 | 必填 | 生成规则 |
|--------|------|------|----------|
| `title` | string | 是 | 课时标题，5-30个字符 |
| `description` | string | 否 | 课时描述，20-80个字符 |
| `video_url` | string | 否 | 视频URL，格式正确 |
| `document_url` | string | 否 | 文档URL，格式正确 |
| `content` | string | 否 | 课时内容，100-500个字符 |
| `duration` | integer | 否 | 课时时长（秒），合理范围300-3600 |
| `sort_order` | integer | 否 | 排序顺序，从1开始递增 |
| `is_free` | boolean | 否 | 是否免费，默认false |
| `type` | string | 否 | 课时类型：video, document, text |

## 4. 数据验证规则

1. **必填字段检查**：确保所有必填字段都有值
2. **数据类型检查**：确保字段类型正确
3. **长度检查**：确保字符串字段长度符合要求
4. **范围检查**：确保数值字段在合理范围内
5. **格式检查**：确保日期、URL等格式正确
6. **分类ID检查**：确保category_id在系统中存在
7. **逻辑检查**：确保original_price >= price

## 5. 内容质量要求

1. **准确性**：课程内容描述准确，符合实际
2. **一致性**：课程、章节、课时内容相互一致
3. **完整性**：信息完整，无缺失重要内容
4. **专业性**：语言专业，符合教育领域规范
5. **吸引力**：内容有吸引力，能够激发学习兴趣

## 6. 示例数据

### 6.1 课程示例

```json
{
  "title": "Vue 3 实战开发",
  "description": "从入门到精通，全面掌握Vue 3核心概念和实战技巧",
  "detail": "本课程将带你从Vue 3的基础概念开始，逐步深入到组件化开发、状态管理、路由等高级特性，通过多个实战项目巩固所学知识，让你能够独立开发Vue 3应用。",
  "cover_image": "https://img0.baidu.com/it/u=3852706705,376344197&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=652",
  "price": 499,
  "original_price": 699,
  "category_id": 1,
  "level": "intermediate",
  "duration": 480,
  "status": "published",
  "target_students": "有一定JavaScript基础的前端开发者，想要学习Vue 3的初学者",
  "requirements": "基本的HTML、CSS、JavaScript知识，了解ES6语法",
  "learning_objectives": "掌握Vue 3核心概念，能够独立开发Vue 3应用，理解组件化开发思想",
  "certificate": "完成课程并通过考核后，可获得Vue 3开发工程师证书",
  "course_highlights": ["Vue 3 核心概念讲解清晰", "实战项目覆盖全面", "课程内容及时更新"],
  "teaching_methods": ["理论讲解 + 实践操作", "小组讨论 + 案例分析", "实时答疑"],
  "support_services": ["课程资料永久下载", "在线答疑", "社群交流"],
  "career_prospects": ["Vue 3 开发工程师", "前端架构师", "技术专家"]
}
```

### 6.2 章节和课时示例

```json
[
  {
    "title": "Vue 3 基础",
    "description": "Vue 3的基本概念和核心特性",
    "sort_order": 1,
    "lessons": [
      {
        "title": "1.1 Vue 3 简介",
        "description": "Vue 3的新特性和优势",
        "video_url": "https://example.com/videos/vue3-intro.mp4",
        "duration": 600,
        "sort_order": 1,
        "is_free": true,
        "type": "video"
      },
      {
        "title": "1.2 响应式系统",
        "description": "Vue 3的响应式原理和使用",
        "video_url": "https://example.com/videos/vue3-reactive.mp4",
        "duration": 900,
        "sort_order": 2,
        "is_free": false,
        "type": "video"
      }
    ]
  },
  {
    "title": "组件化开发",
    "description": "Vue 3的组件化开发最佳实践",
    "sort_order": 2,
    "lessons": [
      {
        "title": "2.1 组件基础",
        "description": "组件的创建和使用",
        "video_url": "https://example.com/videos/vue3-components.mp4",
        "duration": 720,
        "sort_order": 1,
        "is_free": false,
        "type": "video"
      }
    ]
  }
]
```

## 7. 生成流程

1. **分析输入**：理解用户需求和课程主题
2. **生成课程信息**：根据规则生成课程基本信息
3. **生成章节结构**：设计合理的章节划分
4. **生成课时内容**：为每个章节生成具体课时
5. **数据验证**：检查生成的数据是否符合规则
6. **格式化输出**：按照API要求的格式输出数据

## 8. 注意事项

1. 所有创建的课程默认教师ID为1，无需在生成数据中指定
2. 课程状态默认为"published"，无需手动设置
3. 数组类型的字段（如course_highlights）如果为空，会自动转换为空数组
4. 请确保提供的分类ID在系统中存在
5. 生成的内容应符合教育领域的专业标准
6. 避免生成重复或无关的内容
7. 确保生成的数据结构与API接口要求一致
