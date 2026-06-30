package config

import "goblog/pkg/config"

func init() {
	config.SetDefault("session.default", "cookie")
	config.SetDefault("session.session_name", "goblog-session")
}
