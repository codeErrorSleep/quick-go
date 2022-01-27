package conf

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var (
	Env *viper.Viper
)

var cmdEnv struct {
	Env  string
	Port int64
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
	Env.Set("port", cmdEnv.Port)

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
	Env = viperObj
	return nil
}
