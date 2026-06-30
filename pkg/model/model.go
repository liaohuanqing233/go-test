package model

import (
	"fmt"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"goblog/pkg/types"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// BaseModel 模型基类
type BaseModel struct {
	ID        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}

// GetStringID 获取string类型的ID
func (model BaseModel) GetStringID() string {
	return types.Uint64ToString(model.ID)
}

// DB 数据库操作对象
var DB *gorm.DB

// ConnectDB 连接数据库
func ConnectDB() *gorm.DB {
	var err error

	// 初始化 MySQL 连接信息
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: logger.NewGormLogger(),
	})

	logger.LogError(err)

	return DB
}
