package session

import (
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"net/http"

	"github.com/gorilla/sessions"
)

// Store gorilla sessions的储存库
var Store = sessions.NewCookieStore([]byte(config.GetString("app.key")))

// Session 会话操作对象
var Session *sessions.Session

// Request 获取会话
var Request *http.Request

// Response 写入会话
var Response http.ResponseWriter

// StartSession 开启会话
func StartSession(w http.ResponseWriter, r *http.Request) {
	var err error

	Session, err = Store.Get(r, config.GetString("session.session_name"))
	logger.LogError(err)

	Request = r
	Response = w
}

// Put 写入会话数据
func Put(key string, value interface{}) {
	Session.Values[key] = value
	Save()
}

// Get 获取会话数据
func Get(key string) interface{} {
	return Session.Values[key]
}

// Forget 删除某个会话
func Forget(key string) {
	delete(Session.Values, key)
	Save()
}

// Flush 删除当前会话
func Flush() {
	Session.Options.MaxAge = -1
	Save()
}

// Save 保存会话
func Save() {
	err := Session.Save(Request, Response)
	logger.LogError(err)
}
