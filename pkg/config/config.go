package config

import (
	"log"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Viper viper库实例
var Viper *viper.Viper

func init() {
	Viper = viper.New()
	Viper.SetConfigName("config")
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(".")
	bindEnvs()
}

// Load 读取配置文件，需在 defaults 注册之后调用
func Load() {
	err := Viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Println(err)
		}
	}
}

func bindEnvs() {
	envs := map[string]string{
		"app.name":                            "APP_NAME",
		"app.env":                             "APP_ENV",
		"app.debug":                           "APP_DEBUG",
		"app.port":                            "APP_PORT",
		"app.key":                             "APP_KEY",
		"database.mysql.host":                 "DB_HOST",
		"database.mysql.port":                 "DB_PORT",
		"database.mysql.database":             "DB_DATABASE",
		"database.mysql.username":             "DB_USERNAME",
		"database.mysql.password":             "DB_PASSWORD",
		"database.mysql.max_idle_connections": "DB_MAX_IDLE_CONNECTIONS",
		"database.mysql.max_open_connections": "DB_MAX_OPEN_CONNECTIONS",
		"database.mysql.max_life_seconds":     "DB_MAX_LIFE_SECONDS",
		"session.default":                     "SESSION_DEFAULT",
		"session.session_name":                "SESSION_NAME",
	}
	for key, env := range envs {
		_ = Viper.BindEnv(key, env)
	}
}

// SetDefault 设置默认值
func SetDefault(key string, value interface{}) {
	Viper.SetDefault(key, value)
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
