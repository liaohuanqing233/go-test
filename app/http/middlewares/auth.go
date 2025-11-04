package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

// Auth 登录操作
func Auth(next HTTPHandlerFunc) HTTPHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			flash.Warning("登录用户才能访问该页面，请登录后操作！！！")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
