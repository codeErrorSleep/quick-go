package main

import (
	"quick-go/bootstrap"
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
}
