package main

import (
	"quick-go/conf"
	"quick-go/db"
	"quick-go/log"
)

func main() {
	// 1.读取配置相关的信息
	err := conf.InitConfig()
	if err != nil {
		print(err.Error())
	}

	// 2.日志系统
	err = log.InitLog()
	if err != nil {
		print(err.Error())
	}

	// 3.初始化mysql
	err = db.InitMysql()
	if err != nil {
		print(err.Error())
	}

	// 4.初始化redis
	err = db.InitRedis()
	if err != nil {
		print(err.Error())
	}
}
