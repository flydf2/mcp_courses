package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// 全局变量存储配置文件内容
var configMap map[string]string

// loadConfigFile 加载配置文件
func LoadConfigFile() {
	configMap = make(map[string]string)

	// 配置文件路径
	configPath := "./config/config.env"

	// 检查配置文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("配置文件 %s 不存在，将只使用环境变量", configPath)
		return
	}

	// 读取配置文件
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("无法打开配置文件 %s: %v", configPath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释行
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析 KEY=VALUE 格式
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("配置文件第 %d 行格式错误: %s", lineNum, line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 移除值两边的引号（如果有）
		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		configMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Printf("读取配置文件错误: %v", err)
		return
	}

	log.Printf("成功加载配置文件 %s，共 %d 个配置项", configPath, len(configMap))
}

// getConfigValue 从配置文件获取值
func getConfigValue(key string) (string, bool) {
	if configMap == nil {
		return "", false
	}
	value, exists := configMap[key]
	return value, exists
}

// GetEnvOrDefault 获取环境变量或默认值（优先环境变量，然后配置文件，最后默认值）
func GetEnvOrDefault(key, defaultValue string) string {
	// 首先尝试从环境变量获取
	if value := os.Getenv(key); value != "" {
		return value
	}

	// 然后尝试从配置文件获取
	if value, exists := getConfigValue(key); exists {
		return value
	}

	// 最后返回默认值
	return defaultValue
}

// GetBoolEnv 获取布尔环境变量（优先环境变量，然后配置文件，最后默认值）
func GetBoolEnv(key string, defaultValue bool) bool {
	// 首先尝试从环境变量获取
	if value := os.Getenv(key); value != "" {
		return value == "true"
	}

	// 然后尝试从配置文件获取
	if value, exists := getConfigValue(key); exists {
		return value == "true"
	}

	// 最后返回默认值
	return defaultValue
}

// GetStringSliceEnv 获取字符串切片环境变量（优先环境变量，然后配置文件，最后默认值）
func GetStringSliceEnv(key string, defaultValue []string) []string {
	// 首先尝试从环境变量获取
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ",")
	}

	// 然后尝试从配置文件获取
	if value, exists := getConfigValue(key); exists {
		return strings.Split(value, ",")
	}

	// 最后返回默认值
	return defaultValue
}

// GetIntEnv 获取整数环境变量（优先环境变量，然后配置文件，最后默认值）
func GetIntEnv(key string, defaultValue int) int {
	// 首先尝试从环境变量获取
	if value := os.Getenv(key); value != "" {
		intValue := 0
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}

	// 然后尝试从配置文件获取
	if value, exists := getConfigValue(key); exists {
		intValue := 0
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}

	// 最后返回默认值
	return defaultValue
}

// GetDurationEnv 获取持续时间环境变量（优先环境变量，然后配置文件，最后默认值）
func GetDurationEnv(key string, defaultValue time.Duration) time.Duration {
	// 首先尝试从环境变量获取
	if value := os.Getenv(key); value != "" {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}

	// 然后尝试从配置文件获取
	if value, exists := getConfigValue(key); exists {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}

	// 最后返回默认值
	return defaultValue
}

// AtoiToInt 将字符串转换为整数（失败时返回默认值）
func AtoiToInt(str string) int {
	intValue := 0
	if _, err := fmt.Sscanf(str, "%d", &intValue); err == nil {
		return intValue
	}
	return 0
}
