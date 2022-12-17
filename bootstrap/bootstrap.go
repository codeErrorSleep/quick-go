package bootstrap

import (
	"quick-go/global"
)

func Bootstrap(envMode string) {

	// 1.读取配置相关的信息
	err := global.InitConfig(envMode)
	if err != nil {
		print(err.Error())
	}

	// 2.日志系统
	err = global.InitLog()
	if err != nil {
		print(err.Error())
	}

	// // 3.初始化mysql
	// err = global.InitMysql()
	// if err != nil {
	// 	print(err.Error())
	// }

	// // 4.初始化redis
	// err = global.InitRedis()
	// if err != nil {
	// 	print(err.Error())
	// }

	// // 5.创建kafka连接
	// err = global.InitKafka()
	// if err != nil {
	// 	print(err.Error())
	// }

	// // 异步处理请求
	// go async.AsyncGoodsDetail()

	// rpc 服务
	// grpcServer := grpc.NewServer()
	// rpc.RegisterHelloServiceServer(grpcServer, &service.HelloService{})
	// listener, err := net.Listen("tcp", "localhost:"+global.Env.GetString("grpcPort"))
	// if err != nil {
	// 	fmt.Println("net Listen err: ", err)
	// 	return
	// }
	// grpcServer.Serve(listener)
	// defer listener.Close()

}
