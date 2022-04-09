package main

import (
	"quick-go/conf"
	"quick-go/db"
	"quick-go/log"
	"quick-go/routers"

	"go.uber.org/zap"
)

func main() {
	defer resourceCloser()

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

	// 5.创建kafka连接
	err = db.InitKafka()
	if err != nil {
		print(err.Error())
	}

	// 6.注册路由
	r := routers.InitApiRouter(false)
	r.Run(":" + conf.Env.GetString("port"))

}

// 关闭所有的连接资源
func resourceCloser() {
	for _, resourceclose := range db.ResourceCloses {
		if err := resourceclose(); err != nil {
			log.ErrorLogger.Info("资源关闭异常", zap.Error(err))
		}
	}
}
