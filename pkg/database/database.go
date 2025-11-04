package database

import (
	"database/sql"
	"goblog/pkg/logger"
	"time"

	"github.com/go-sql-driver/mysql"
)

// DB 数据库操作对象
var DB *sql.DB

// Initialize 数据库初始化
func Initialize() {
	InitDB()
	CreateTables()
}

// InitDB 初始化连接数据库
func InitDB() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "gotest",
		AllowNativePasswords: true,
	}

	// 准备数据库连接池
	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	// 设置最大连接数
	DB.SetMaxOpenConns(25)
	// 设置最大空闲数
	DB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	DB.SetConnMaxLifetime(5 * time.Minute)
	// 尝试连接，失败报错
	err = DB.Ping()
	logger.LogError(err)
}

// CreateTables 创建数据库表
func CreateTables() {
	createArticlesSQL := `
	CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci
	);
	`
	_, err := DB.Exec(createArticlesSQL)
	logger.LogError(err)
}
