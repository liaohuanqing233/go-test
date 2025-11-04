package user

import (
	"goblog/pkg/password"

	"gorm.io/gorm"
)

// BeforeCreate 创建模型前调用
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = password.Hash(u.Password)
	return
}

// BeforeUpdate 保存前调用
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return
}
