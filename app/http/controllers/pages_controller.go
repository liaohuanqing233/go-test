package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 静态页面处理器
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是goblog</h1>")
}

// About 关于页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "这里是关于goblog的内容12123123"+
		"<a href=\"mailto:1915443708@qq.com\">1915443708@qq.com</a>")
}

// NotFound 未找到页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "_我不知道你在说什么123")
}
