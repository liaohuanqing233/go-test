package config

import "goblog/pkg/config"

func init() {
	config.SetDefault("database.mysql.host", "127.0.0.1")
	config.SetDefault("database.mysql.port", "3306")
	config.SetDefault("database.mysql.database", "gotest")
	config.SetDefault("database.mysql.username", "")
	config.SetDefault("database.mysql.password", "")
	config.SetDefault("database.mysql.charset", "utf8mb4")
	config.SetDefault("database.mysql.max_idle_connections", 100)
	config.SetDefault("database.mysql.max_open_connections", 25)
	config.SetDefault("database.mysql.max_life_seconds", 300)
}
