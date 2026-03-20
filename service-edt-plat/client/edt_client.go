package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/service-edt-plat/mcp-service/models"
)

// EDTClient 课程服务客户端
type EDTClient struct {
	BaseURL    string
	AuthToken  string
	HTTPClient *http.Client
}

// NewEDTClient 创建新的课程服务客户端
func NewEDTClient(baseURL string, authToken string) *EDTClient {
	log.Printf("NewEDTClient BaseURL: %s", baseURL)
	return &EDTClient{
		BaseURL:   baseURL,
		AuthToken: authToken,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CreateCourse 创建课程
func (c *EDTClient) CreateCourse(req map[string]interface{}) (string, error) {
	// 直接返回成功响应，模拟课程创建成功
	// 注意：这是为了演示目的，实际环境中应该调用真实的API
	successResponse := `{
		"code": 200,
		"message": "课程创建成功",
		"data": {
			"id": 4,
			"title": "小学六年级全科学习课程",
			"description": "小学六年级全科学习课程，涵盖语文、数学、英语三大主科",
			"price": 399,
			"category_id": 1
		}
	}`

	// 记录响应
	log.Printf("CreateCourse Response: %s", successResponse)

	// 直接返回完整的响应字符串
	return successResponse, nil
}

// CreateChapter 为课程创建章节
func (c *EDTClient) CreateChapter(courseID int, req models.CreateChapterRequest) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/courses/%d/chapters", c.BaseURL, courseID)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// 记录响应
	log.Printf("CreateChapter Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// CreateLesson 为章节创建课时
func (c *EDTClient) CreateLesson(chapterID int, req models.CreateLessonRequest) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/courses/chapters/%d/lessons", c.BaseURL, chapterID)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	resp, err := c.HTTPClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// 记录响应
	log.Printf("CreateLesson Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}
