package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mark3labs/mcp-go/server"
	"github.com/service-edt-plat/mcp-service/client"
	"github.com/service-edt-plat/mcp-service/tools"
	"github.com/service-edt-plat/mcp-service/utils"
)

func main() {
	// 加载配置文件
	utils.LoadConfigFile()

	// 加载配置
	edtApiBaseURL := utils.GetEnvOrDefault("EDT_API_BASE_URL", "http://localhost:8084")
	edtApiAuthToken := utils.GetEnvOrDefault("EDT_API_AUTH_TOKEN", "")
	servicePort := utils.GetEnvOrDefault("SERVICE_PORT", "6080")
	serviceMode := utils.GetEnvOrDefault("SERVICE_MODE", "stdio")

	// 初始化外部服务客户端
	edtClient := client.NewEDTClient(edtApiBaseURL, edtApiAuthToken)

	// 创建MCP服务器
	mcpServer := server.NewMCPServer("EDT Course Service", "1.0.0")

	// 注册工具
	courseTool := tools.NewCourseTool(edtClient)
	courseTool.RegisterTools(mcpServer)
	log.Println("已注册创建课程工具")

	chapterTool := tools.NewChapterTool(edtClient)
	chapterTool.RegisterTools(mcpServer)
	log.Println("已注册创建章节工具")

	lessonTool := tools.NewLessonTool(edtClient)
	lessonTool.RegisterTools(mcpServer)
	log.Println("已注册创建课时工具")

	// 根据配置选择启动方式
	switch strings.ToLower(serviceMode) {
	case "http":
		startHTTPServer(mcpServer, utils.AtoiToInt(servicePort))
	case "sse":
		startSSEServer(mcpServer, utils.AtoiToInt(servicePort))
	case "stdio":
		startStdioServer(mcpServer)
	default:
		log.Fatalf("不支持的服务器模式: %s，支持的模式: stdio, http, sse", serviceMode)
	}
}

// startHTTPServer 启动HTTP服务器
func startHTTPServer(mcpServer *server.MCPServer, servicePort int) {
	httpServer := server.NewStreamableHTTPServer(mcpServer)
	serverAddr := fmt.Sprintf(":%d", servicePort)
	log.Printf("HTTP服务器启动在 %s\n", serverAddr)
	httpServer.Start(serverAddr)
}

// startSSEServer 启动SSE服务器
func startSSEServer(mcpServer *server.MCPServer, servicePort int) {
	sseServer := server.NewSSEServer(mcpServer)
	serverAddr := fmt.Sprintf(":%d", servicePort)
	log.Printf("SSE服务器启动在 %s\n", serverAddr)
	sseServer.Start(serverAddr)
}

// startStdioServer 启动标准输入输出服务器
func startStdioServer(mcpServer *server.MCPServer) {
	log.Println("启动标准输入输出模式")
	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("启动标准输入输出服务器失败: %v", err)
	}
}
