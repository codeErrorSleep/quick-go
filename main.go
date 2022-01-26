package main

import (
	"go.uber.org/zap"
	"quick-go/bootstrap"
	"quick-go/bootstrap/conf"
)

func main() {
	// 1.读取配置相关的信息
	err := bootstrap.RegisterConfig()
	if err != nil {
		print(err.Error())
	}

	// 2.日志系统
	err = bootstrap.RegisterLog()
	if err != nil {
		print(err.Error())
	}

	conf.InfoLogger.Info("test", zap.Any("fsdfds", "fffff"))

}
