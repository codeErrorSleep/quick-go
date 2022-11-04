package main

import (
	"fmt"
	"net"
	"quick-go/app/rpc"
	"quick-go/app/service"
	"quick-go/global"
	"quick-go/routers"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	// // 5.创建kafka连接
	// err = global.InitKafka()
	// if err != nil {
	// 	print(err.Error())
	// }

	// // 异步处理请求
	// go async.AsyncGoodsDetail()

	// rpc 服务
	grpcServer := grpc.NewServer()
	rpc.RegisterHelloServiceServer(grpcServer, &service.HelloService{})
	listener, err := net.Listen("tcp", "localhost:"+global.Env.GetString("grpcPort"))
	if err != nil {
		fmt.Println("net Listen err: ", err)
		return
	}
	grpcServer.Serve(listener)
	defer listener.Close()

	// 6.注册路由
	r := routers.InitApiRouter(false)
	r.Run(":" + global.Env.GetString("httpPort"))

}

// 关闭所有的连接资源
func resourceCloser() {
	for _, resourceclose := range global.ResourceCloses {
		if err := resourceclose(); err != nil {
			global.ErrorLogger.Info("资源关闭异常", zap.Error(err))
		}
	}
}
