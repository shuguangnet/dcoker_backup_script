package config

import (
	"fmt"
	"os"
	"strings"
)

// Config 结构体用于存放所有配置项
type Config struct {
	Port           string
	ScriptPath     string
	CallbackSecret string
	CallbackURL    string // 新增字段用于存储回调URL
}

// LoadConfig 从指定路径读取配置文件并返回配置对象
func LoadConfig(path string) (*Config, error) {
	// 默认配置值
	config := &Config{
		Port:       "47731",
		ScriptPath: "docker-backup.sh",
	}

	// 读取配置文件
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置文件
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		// 跳过注释和空行
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}

		// 解析键值对
		parts := strings.SplitN(trimmedLine, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 根据键更新配置
		switch key {
		case "port":
			config.Port = value
		case "scriptpath":
			config.ScriptPath = value
		case "callback_secret":
			config.CallbackSecret = value
		}
	}

	// 验证必要配置项
	if config.CallbackSecret == "" {
		return nil, fmt.Errorf("callback_secret not found in config file")
	}

	return config, nil
}
