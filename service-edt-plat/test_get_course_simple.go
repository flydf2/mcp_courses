package main

import (
	"fmt"
	"log"

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

	// 初始化外部服务客户端
	edtClient := client.NewEDTClient(edtApiBaseURL, edtApiAuthToken)

	// 创建MCP服务器
	mcpServer := server.NewMCPServer("EDT Course Service", "1.0.0")

	// 注册课程工具
	courseTool := tools.NewCourseTool(edtClient)
	courseTool.RegisterTools(mcpServer)
	log.Println("已注册课程工具")

	// 打印所有注册的工具
	tools := mcpServer.GetTools()
	fmt.Println("注册的工具:")
	for _, tool := range tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
	}
}
