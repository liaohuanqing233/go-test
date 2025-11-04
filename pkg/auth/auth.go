package auth

import (
	"errors"
	"goblog/app/models/user"
	"goblog/pkg/session"

	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

// User 获取用户登录信息
func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

// Attempt 尝试登录
func Attempt(email string, password string) error {
	_user, err := user.GetByEmail(email)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("系统内部错误，请稍后再重新尝试")
		}
	}

	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}

	session.Put("uid", _user.GetStringID())
	return nil
}

// Login 登录
func Login(_user user.User) {
	session.Put("uid", _user.GetStringID())
}

// Logout 登出
func Logout() {
	session.Forget("uid")
	session.Save()
}

// Check 检测是否以及登录
func Check() bool {
	return len(_getUID()) > 0
}
