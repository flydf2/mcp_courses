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
func (c *EDTClient) CreateCourse(req models.CreateCourseRequest) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/courses", c.BaseURL)

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

	// 记录请求
	log.Printf("CreateCourse Request: %q", url)
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
	log.Printf("CreateCourse Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
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

	// 记录请求
	log.Printf("CreateChapter Request: %q", url)
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

	// 记录请求
	log.Printf("CreateLesson Request: %s", url)
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

// GetCourseByID 获取课程详情
func (c *EDTClient) GetCourseByID(courseID string) (string, error) {
	url := fmt.Sprintf("%s/api/v1/courses/%s", c.BaseURL, courseID)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("GetCourseByID Request: %q", url)

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
	log.Printf("GetCourseByID Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// CreateCategory 创建分类
func (c *EDTClient) CreateCategory(req models.CreateCategoryRequest) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories", c.BaseURL)

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

	// 记录请求
	log.Printf("CreateCategory Request: %q", url)
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
	log.Printf("CreateCategory Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// GetCategories 获取分类列表
func (c *EDTClient) GetCategories(page, pageSize, parentID int) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories?page=%d&pageSize=%d&parentID=%d", c.BaseURL, page, pageSize, parentID)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("GetCategories Request: %q", url)
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
	log.Printf("GetCategories Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// GetCategoryByID 获取分类详情
func (c *EDTClient) GetCategoryByID(categoryID int) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories/%d", c.BaseURL, categoryID)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("GetCategoryByID Request: %q", url)
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
	log.Printf("GetCategoryByID Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// UpdateCategory 更新分类
func (c *EDTClient) UpdateCategory(categoryID int, req models.UpdateCategoryRequest) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories/%d", c.BaseURL, categoryID)

	reqBody, err := json.Marshal(req)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("UpdateCategory Request: %q", url)
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
	log.Printf("UpdateCategory Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// DeleteCategory 删除分类
func (c *EDTClient) DeleteCategory(categoryID int) (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories/%d", c.BaseURL, categoryID)

	httpReq, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("DeleteCategory Request: %q", url)
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
	log.Printf("DeleteCategory Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}

// GetAllCategories 获取所有分类
func (c *EDTClient) GetAllCategories() (string, error) {
	url := fmt.Sprintf("%s/api/v1/external/categories/all", c.BaseURL)

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.AuthToken != "" {
		httpReq.Header.Set("Authorization", c.AuthToken)
	}

	// 记录请求
	log.Printf("GetAllCategories Request: %q", url)
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
	log.Printf("GetAllCategories Response: %s", string(bodyBytes))

	// 直接返回完整的响应字符串
	return string(bodyBytes), nil
}
