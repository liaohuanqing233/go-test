package config

import (
	"goblog/pkg/config"
	"goblog/pkg/logger"
)

// Initialize 配置信息初始化
func Initialize() {
	config.Load()
	logger.Initialize()
}
