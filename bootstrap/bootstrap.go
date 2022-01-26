package bootstrap

import (
	"flag"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"quick-go/bootstrap/conf"
)

var cmdEnv struct {
	Env  string
	Port int64
}

// RegisterLog 初始化链接
func RegisterLog() error {
	// 注册zap日志
	ErrorLogger := initLogger("error.log", "info")
	conf.ErrorLogger = ErrorLogger
	DebugLogger := initLogger("debug.log", "debug")
	conf.DebugLogger = DebugLogger
	InfoLogger := initLogger("info.log", "info")
	conf.InfoLogger = InfoLogger
	CallLogger := initLogger("call.log", "info")
	conf.CallLogger = CallLogger
	return nil
}

// loglevel 日志级别
func initLogger(logPath string, loglevel string) *zap.Logger {
	// 加上要记录的路径
	logPath = conf.Env.GetString("logPath") + logPath
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

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}

// RegisterConfig 初始化配置
func RegisterConfig() error {
	// 读取命令行参数
	registerCmdEnvConfig()

	// 读取配置的conf文件
	err := registerEnvConfig()
	if err != nil {
		return err
	}

	// 设置一下其他的变量
	conf.Env.Set("port", cmdEnv.Port)

	return nil
}

// 读取命令行参数
func registerCmdEnvConfig() {
	envFile := flag.String("e", "config", "配置文件")
	port := flag.Int64("p", 8080, "启动端口")
	flag.Parse()
	cmdEnv.Env = *envFile
	cmdEnv.Port = *port
}

// 读取config配置文件
func registerEnvConfig() error {
	viperObj := viper.New()
	viperObj.SetConfigName("configs/" + cmdEnv.Env)
	viperObj.SetConfigType("yaml")
	viperObj.AddConfigPath(".")
	err := viperObj.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
		return err
	}
	// 监听配置修改
	viperObj.WatchConfig()
	conf.Env = viperObj
	return nil
}
