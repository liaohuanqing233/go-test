package main

import (
	"net/http"

	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/config"
	c "goblog/pkg/config"
	"goblog/pkg/logger"
)

func init() {
	config.Initialize()
}

func main() {
	defer logger.Sync()

	// 数据库连接及数据库表初始化
	bootstrap.SetupDB()

	// 初始化路由对象
	router := bootstrap.SetupRoute()

	// 服务
	http.ListenAndServe(":"+c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
}
