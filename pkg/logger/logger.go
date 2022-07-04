package logger

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goapi/pkg/app"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"
)

var Logger *zap.Logger

func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(Logger)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	if logType == "daily" {
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logname)
	}
	fmt.Printf("file:%s", filename)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	if app.IsLocal() {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		return zapcore.AddSync(lumberJackLogger)
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	if app.IsLocal() {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

func LogIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

func Debug(modleName string, fields ...zap.Field) {
	Logger.Debug(modleName, fields...)
}

func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
//         logger.DebugJSON("Auth", "读取登录用户", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(jsonData interface{}) string {
	b, err := json.Marshal(jsonData)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
