#!/usr/bin/env python3
import json
import sys

# 测试获取课程详情功能
test_request = {
    "jsonrpc": "2.0",
    "method": "get_course_by_id",
    "params": {
        "id": "1"
    },
    "id": 1
}

# 发送请求到服务
print(json.dumps(test_request))
sys.stdout.flush()

# 读取响应
response = sys.stdin.read()
print("Response:")
print(response)
