package config

import "goblog/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"name":  config.Env("APP_NAME", "GoBlog"),
		"env":   config.Env("APP_ENV", "production"),
		"debug": config.Env("APP_DEBUG", false),
		"port":  config.Env("APP_PORT", 3000),
		"key":   config.Env("APP_KEY", "liaohq1234567890"),
	})
}
