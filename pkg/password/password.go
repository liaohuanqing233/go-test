package password

import (
	"goblog/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

// Hash 加密密码
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogError(err)
	return string(bytes)
}

// CheckHash 密码对比
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogError(err)
	return err == nil
}

// IsHashed 判断是否被hash过
func IsHashed(str string) bool {
	// 加密后长度为60
	return len(str) == 60
}
