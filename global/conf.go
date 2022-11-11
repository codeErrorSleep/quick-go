package global

import (
	"flag"
	"log"
	"quick-go/global/consts"

	"github.com/spf13/viper"
)

var (
	Env *viper.Viper
)

var cmdEnv struct {
	Env     string
	BaseEnv string
}

// func Bootstrap(envMode string) {

// 	// 1.读取配置相关的信息
// 	err := InitConfig(envMode)
// 	if err != nil {
// 		print(err.Error())
// 	}

// 	// 2.日志系统
// 	err = InitLog()
// 	if err != nil {
// 		print(err.Error())
// 	}

// 	// 3.初始化mysql
// 	err = InitMysql()
// 	if err != nil {
// 		print(err.Error())
// 	}

// 	// 4.初始化redis
// 	err = InitRedis()
// 	if err != nil {
// 		print(err.Error())
// 	}

// 	// // 5.创建kafka连接
// 	// err = InitKafka()
// 	// if err != nil {
// 	// 	print(err.Error())
// 	// }

// 	// // 异步处理请求
// 	// go async.AsyncGoodsDetail()

// 	// rpc 服务
// 	// grpcServer := grpc.NewServer()
// 	// rpc.RegisterHelloServiceServer(grpcServer, &service.HelloService{})
// 	// listener, err := net.Listen("tcp", "localhost:"+Env.GetString("grpcPort"))
// 	// if err != nil {
// 	// 	fmt.Println("net Listen err: ", err)
// 	// 	return
// 	// }
// 	// grpcServer.Serve(listener)
// 	// defer listener.Close()

// }

// InitConfig 初始化配置
func InitConfig(envMode string) error {
	// 读取命令行参数
	registerCmdEnvConfig(envMode)

	// 读取配置的conf文件
	err := registerEnvConfig()
	if err != nil {
		return err
	}
	return nil
}

// 读取命令行参数
func registerCmdEnvConfig(envMode string) {
	envFile := flag.String("e", "config", "配置文件")
	bashEnv := flag.String("b", "./", "根目录相对路径")
	flag.Parse()

	cmdEnv.Env = *envFile
	cmdEnv.BaseEnv = *bashEnv

	// 单测的固定启动参数
	if envMode == consts.EnvUnitTest {
		cmdEnv.BaseEnv = "../"
	}
}

// 读取config配置文件
func registerEnvConfig() error {
	viperObj := viper.New()
	viperObj.SetConfigName(cmdEnv.BaseEnv + "conf/" + cmdEnv.Env)
	viperObj.SetConfigType("yaml")
	viperObj.AddConfigPath(".")
	err := viperObj.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
		return err
	}
	// 监听配置修改
	viperObj.WatchConfig()
	Env = viperObj
	return nil
}
