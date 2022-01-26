package main

import (
	"quick-go/bootstrap/conf"
	"quick-go/bootstrap/log"
)

func main() {
	// 1.读取配置相关的信息
	err := conf.RegisterConfig()
	if err != nil {
		print(err.Error())
	}

	// 2.日志系统
	err = log.RegisterLog()
	if err != nil {
		print(err.Error())
	}
}
