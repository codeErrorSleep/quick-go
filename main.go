package main

import (
	"quick-go/global"
	"quick-go/routers"

	"go.uber.org/zap"
)

func main() {
	defer resourceCloser()

	// 1.读取配置相关的信息
	err := global.InitConfig()
	if err != nil {
		print(err.Error())
	}

	// 2.日志系统
	err = global.InitLog()
	if err != nil {
		print(err.Error())
	}

	// 3.初始化mysql
	err = global.InitMysql()
	if err != nil {
		print(err.Error())
	}
	// 4.初始化redis
	err = global.InitRedis()
	if err != nil {
		print(err.Error())
	}

	// 5.创建kafka连接
	err = global.InitKafka()
	if err != nil {
		print(err.Error())
	}

	// 6.注册路由
	r := routers.InitApiRouter(false)
	r.Run(":" + global.Env.GetString("port"))

}

// 关闭所有的连接资源
func resourceCloser() {
	for _, resourceclose := range global.ResourceCloses {
		if err := resourceclose(); err != nil {
			global.ErrorLogger.Info("资源关闭异常", zap.Error(err))
		}
	}
}
