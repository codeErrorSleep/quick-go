package main

import (
	"quick-go/global"
	"quick-go/global/consts"
	"quick-go/routers"

	"go.uber.org/zap"
)

func main() {
	defer resourceCloser()
	global.Bootstrap(consts.EnvProduction)

	r := routers.InitApiRouter(false)
	r.Run(":" + global.Env.GetString("httpPort"))
}

// 关闭所有的连接资源
func resourceCloser() {
	for _, resourceClose := range global.ResourceCloses {
		if err := resourceClose(); err != nil {
			global.ErrorLogger.Info("资源关闭异常", zap.Error(err))
		}
	}
}
