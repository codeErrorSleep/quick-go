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

}
