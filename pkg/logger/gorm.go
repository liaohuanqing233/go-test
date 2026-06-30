package logger

import (
	"context"
	"errors"
	"time"

	"goblog/pkg/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// GormLogger GORM 日志适配器，SQL 错误同步写入 sql channel
type GormLogger struct {
	level gormlogger.LogLevel
}

// NewGormLogger 创建 GORM logger
func NewGormLogger() *GormLogger {
	level := gormlogger.Error
	if config.GetBool("app.debug") {
		level = gormlogger.Warn
	}
	return &GormLogger{level: level}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return &GormLogger{level: level}
}

func (l *GormLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Info {
		Channel("sql").Sugar().Infof(msg, data...)
	}
}

func (l *GormLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Warn {
		Channel("sql").Sugar().Warnf(msg, data...)
	}
}

func (l *GormLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Error {
		Channel("sql").Sugar().Errorf(msg, data...)
	}
}

func (l *GormLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.level <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	switch {
	case err != nil && l.level >= gormlogger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
		Channel("sql").Error("sql error",
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.Duration("elapsed", elapsed),
			zap.Error(err),
		)
	case elapsed > 200*time.Millisecond && l.level >= gormlogger.Warn:
		Channel("sql").Warn("slow sql",
			zap.String("sql", sql),
			zap.Int64("rows", rows),
			zap.Duration("elapsed", elapsed),
		)
	}
}
