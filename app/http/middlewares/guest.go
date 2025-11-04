package middlewares

import (
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"net/http"
)

// Guest 游客页面
func Guest(next HTTPHandlerFunc) HTTPHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {
			flash.Warning("你已登录！！！")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
