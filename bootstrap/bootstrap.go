package bootstrap

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"quick-go/bootstrap/conf"
)

// RegisterConfig 初始化配置
func RegisterConfig() error {
	// todo 通过环境变量指定
	viperObj := viper.New()
	viperObj.SetConfigName("configs/config")
	viperObj.SetConfigType("yaml")
	viperObj.AddConfigPath(".")
	err := viperObj.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
		return err
	}

	conf.Conf = viperObj

	fmt.Println(conf.Conf.Get("db_mysql01"))
	ff := conf.Conf.Get("db_mysql01")
	fmt.Println(ff)

	return nil
}
