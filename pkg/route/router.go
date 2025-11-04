package route

import (
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

// SetRouter 给router初始化
func SetRouter(r *mux.Router) {
	router = r
}

// NameToURL 路由名称转换为可访问的路由地址
func NameToURL(routeName string, pairs ...string) string {
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return url.String()
}

// GetRouteVariable 获取路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
