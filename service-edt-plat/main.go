package main

import (
	"fmt"
	"log"
	"net/http"
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
	servicePort := utils.GetEnvOrDefault("SERVICE_PORT", "6081")

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

	// 强制使用HTTP模式启动服务器
	log.Println("强制使用HTTP模式启动服务器")
	startHTTPServer(mcpServer, utils.AtoiToInt(servicePort))
}

// startHTTPServer 启动HTTP服务器
func startHTTPServer(mcpServer *server.MCPServer, servicePort int) {
	// 创建MCP HTTP服务器
	mcpHttpServer := server.NewStreamableHTTPServer(mcpServer)

	// 注册所有请求的处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s", r.Method, r.URL.Path)
		
		// 检查是否是获取课程详情的请求
		if r.Method == "GET" && strings.HasPrefix(r.URL.Path, "/api/v1/external/courses/") {
			// 解析课程ID
			courseID := r.URL.Path[len("/api/v1/external/courses/"):]
			if courseID == "" {
				http.Error(w, "Missing course ID", http.StatusBadRequest)
				return
			}
			
			// 加载配置
			edtApiBaseURL := utils.GetEnvOrDefault("EDT_API_BASE_URL", "http://localhost:8084")
			edtApiAuthToken := utils.GetEnvOrDefault("EDT_API_AUTH_TOKEN", "")
			
			// 初始化EDT客户端
			edtClient := client.NewEDTClient(edtApiBaseURL, edtApiAuthToken)
			
			// 获取课程详情
			response, err := edtClient.GetCourseByID(courseID)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to get course: %v", err), http.StatusInternalServerError)
				return
			}
			
			// 设置响应头
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			
			// 写入响应
			w.Write([]byte(response))
			return
		}
		
		// 检查是否是MCP请求
		if strings.HasPrefix(r.URL.Path, "/mcp") {
			mcpHttpServer.ServeHTTP(w, r)
			return
		}
		
		// 其他路径返回404
		log.Printf("404 Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
	})

	serverAddr := fmt.Sprintf(":%d", servicePort)
	log.Printf("HTTP服务器启动在 %s\n", serverAddr)
	
	// 启动HTTP服务器
	http.ListenAndServe(serverAddr, nil)
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
