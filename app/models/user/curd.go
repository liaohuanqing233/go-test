package user

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
)

// Create 创建用户数据
func (user *User) Create() (err error) {
	if err = model.DB.Create(&user).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
