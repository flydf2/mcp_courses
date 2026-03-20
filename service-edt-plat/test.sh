#!/bin/bash

echo '{"tool":"create_course","args":{"title":"Go语言从入门到精通","description":"本课程将带你从Go语言基础开始，逐步深入到高级特性，包括并发编程、网络编程、数据库操作等，帮助你成为Go语言专家。","price":"999","category_id":"1"}}' | ./mcp-service
