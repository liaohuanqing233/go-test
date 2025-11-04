package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

// Flashes flash消息
type Flashes map[string]interface{}

var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据需
	// 要提前注册 gob，方便后续 gob 序列化编码、解码
	gob.Register(Flashes{})
}

// addFlash 添加消息
func addFlash(key string, message string) {
	flashes := Flashes{}
	flashes[key] = message
	session.Put(flashKey, flashes)
	session.Save()
}

// Info 提示消息
func Info(message string) {
	addFlash("info", message)
}

// Warning 告警消息
func Warning(message string) {
	addFlash("warning", message)
}

// Success 成功消息
func Success(message string) {
	addFlash("success", message)
}

// Danger 危险操作消息
func Danger(message string) {
	addFlash("danger", message)
}

// All 读取所有消息
func All() Flashes {
	flashes := session.Get(flashKey)
	// 类型转换
	flashMessages, ok := flashes.(Flashes)
	if !ok {
		return nil
	}
	// 读取后删除
	session.Forget(flashKey)
	session.Save()
	return flashMessages
}
