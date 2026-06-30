package config

import "goblog/pkg/config"

func init() {
	config.SetDefault("log.storage", "storage/logs")
	config.SetDefault("log.async_buffer_size", 1024)
	config.SetDefault("log.channels.app.path", "app")
	config.SetDefault("log.channels.app.level", "info")
	config.SetDefault("log.channels.app.max_age", 30)
	config.SetDefault("log.channels.app.async", false)
	config.SetDefault("log.channels.error.path", "error")
	config.SetDefault("log.channels.error.level", "error")
	config.SetDefault("log.channels.error.max_age", 30)
	config.SetDefault("log.channels.error.async", false)
	config.SetDefault("log.channels.sql.path", "sql")
	config.SetDefault("log.channels.sql.level", "error")
	config.SetDefault("log.channels.sql.max_age", 7)
	config.SetDefault("log.channels.sql.async", false)
	config.SetDefault("log.channels.curl.path", "curl")
	config.SetDefault("log.channels.curl.level", "info")
	config.SetDefault("log.channels.curl.max_age", 7)
	config.SetDefault("log.channels.curl.async", true)
}
