package main

import (
	"quick-go/bootstrap"
	"quick-go/global"
	"quick-go/global/consts"

	"go.uber.org/zap"
)

func main() {
	defer resourceCloser()
	bootstrap.Bootstrap(consts.EnvProduction)

}

// 关闭所有的连接资源
func resourceCloser() {
	for _, resourceClose := range global.ResourceCloses {
		if err := resourceClose(); err != nil {
			global.ErrorLogger.Info("资源关闭异常", zap.Error(err))
		}
	}
}
