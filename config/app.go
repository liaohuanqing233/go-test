package config

import "goblog/pkg/config"

func init() {
	config.SetDefault("app.name", "GoBlog")
	config.SetDefault("app.env", "production")
	config.SetDefault("app.debug", false)
	config.SetDefault("app.port", 3000)
	config.SetDefault("app.key", "liaohq1234567890")
}
