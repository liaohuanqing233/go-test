package config

import "goblog/pkg/config"

func init() {
	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{
			// 数据库连接配置
			"host":     config.Env("DB_HOST", "127.0.0.1"),
			"port":     config.Env("DB_PORT", "3306"),
			"database": config.Env("DB_DATABASE", "gotest"),
			"username": config.Env("DB_USERNAME", ""),
			"password": config.Env("DB_PASSWORD", ""),
			"charset":  "utf8mb4",
			// 连接池配置
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
			"max_open_connections": config.Env("MAX_OPEN_CONNECTIONS", 25),
			"max_life_seconds":     config.Env("MAX_LIFE_SECONDS", 300),
		},
	})
}
