package bootstrap

import (
	"goblog/pkg/route"
	"goblog/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	// 注册路由
	routes.RegisterWebRoutes(router)
	// 给路由包中的路由操作对象初始化赋值
	route.SetRouter(router)

	return router
}
