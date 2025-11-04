package types

import (
	"goblog/pkg/logger"
	"strconv"
)

// Int64ToString int转换为string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToInt 字符串转数字
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		logger.LogError(err)
	}
	return i
}

// Uint64ToString uint64转字符串
func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}
