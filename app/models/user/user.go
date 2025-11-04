package user

import (
	"goblog/pkg/model"
	"goblog/pkg/password"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// User 用户模型
type User struct {
	model.BaseModel
	Name            string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email           string `gorm:"type:varchar(255);not null;unique" valid:"email"`
	Password        string `gorm:"type:varchar(255)" valid:"password"`
	PasswordConfirm string `gorm:"-" valid:"password_confirm"`
}

// Get 根据ID获取单篇文章
func Get(idstr string) (User, error) {
	var user User
	id := types.StringToInt(idstr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail 根据邮箱获取用户信息
func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// ComparePassword 对比密码是否相同
func (u User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}

// Link 方法用来生成用户链接
func (u User) Link() string {
	return route.NameToURL("users.show", "id", u.GetStringID())
}
