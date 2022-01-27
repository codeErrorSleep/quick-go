package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	conf2 "quick-go/conf"
)

var (
	// 日志
	ErrorLogger *zap.Logger
	DebugLogger *zap.Logger
	InfoLogger  *zap.Logger
	CallLogger  *zap.Logger
)

// RegisterLog 初始化链接
func RegisterLog() error {
	// 注册zap日志
	ErrorLogger = initLogger("error.log", "info")
	DebugLogger = initLogger("debug.log", "debug")
	InfoLogger = initLogger("info.log", "info")
	CallLogger = initLogger("call.log", "info")
	return nil
}

// loglevel 日志级别
func initLogger(logPath string, loglevel string) *zap.Logger {
	// 加上要记录的路径
	logPath = conf2.Env.GetString("logPath") + logPath
	lumberHook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // megabytes
		MaxBackups: 300,     // 最多保留300个备份
		MaxAge:     180,     // days
		Compress:   false,   // 是否压缩 disabled by default
	}

	w := zapcore.AddSync(&lumberHook)

	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	// 时间格式
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)

	logger := zap.New(core, zap.AddCaller())
	logger.Info(logPath + "初始化成功")

	return logger
}
