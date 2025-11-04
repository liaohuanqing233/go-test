package controllers

import (
	"fmt"
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/view"
	"net/http"
)

// AuthController controller
type AuthController struct {
}

// Register 注册方法
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 注册执行方法
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 初始化用户数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		_user.Create()

		if _user.ID > 0 {
			flash.Success("注册成功！！！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}
}

// Login 跳转登录页面方法
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 执行登录操作
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if err := auth.Attempt(email, password); err == nil {
		flash.Success("欢迎回来！！！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		view.RenderSimple(w, view.D{
			"Errors":   err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}

// Logout 登出
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	flash.Success("登出成功！！！")
	auth.Logout()
	http.Redirect(w, r, "/", http.StatusFound)
}
