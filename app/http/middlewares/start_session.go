package middlewares

import (
	"goblog/pkg/session"
	"net/http"
)

// StartSession 开启会话
func StartSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 启动会话
		session.StartSession(w, r)
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
