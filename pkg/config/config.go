package config

import (
	"goblog/pkg/logger"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Viper viper库实例
var Viper *viper.Viper

// StrMap 简写
type StrMap map[string]interface{}

func init() {
	// 初始化viper库
	Viper = viper.New()
	// 设置文件名
	Viper.SetConfigName(".env")
	// 设置文件类型
	Viper.SetConfigType("env")
	// 设置配置文件查询路径
	Viper.AddConfigPath(".")
	// 读取配置文件
	err := Viper.ReadInConfig()
	logger.LogError(err)
	// 设置环境变量前缀
	Viper.SetEnvPrefix("appenv")
	// Viper.Get() 优先读取环境变量
	Viper.AutomaticEnv()
}

// Env 读取env配置
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Get 获取配置，允许点式获取
func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// Add 新增配置
func Add(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// GetString 获取string类型数据
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取int类型数据
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取int64类型数据
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取uint类型数据
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取bool类型数据
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
