package logger

import (
	"context"
	"errors"
	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewGormLogger() GormLogger {
	return GormLogger{
		ZapLogger:     Logger,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Debugf(s, i...)
}

func (l GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Warnf(s, i...)
}

func (l GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	l.logger().Sugar().Errorf(s, i...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rowsAffected := fc()
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", elapsed.String()),
		zap.Int64("rows", rowsAffected),
	}
	if err != nil {
		if errors.Is(err, gormlogger.ErrRecordNotFound) {
			l.logger().Warn("Database ErrRecordNotFound", logFields...)
		} else {
			logFields = append(logFields, zap.Error(err))
			l.logger().Error("Database Error", logFields...)
		}
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.logger().Warn("Database Slow Log", logFields...)
	}

	l.logger().Debug("Database Query", logFields...)
}

func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

func (l GormLogger) logger() *zap.Logger {

	// 跳过 gorm 内置的调用
	var (
		gormPackage    = filepath.Join("gorm.io", "gorm")
		zapgormPackage = filepath.Join("moul.io", "zapgorm2")
	)

	// 减去一次封装，以及一次在 logger 初始化里添加 zap.AddCallerSkip(1)
	clone := l.ZapLogger.WithOptions(zap.AddCallerSkip(-2))

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			// 返回一个附带跳过行号的新的 zap logger
			return clone.WithOptions(zap.AddCallerSkip(i))
		}
	}
	return l.ZapLogger
}
