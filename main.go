package main

import (
	"quick-go/bootstrap"
)

func main() {

	err := bootstrap.RegisterConfig()

	if err != nil {
		print(err.Error())
	}

}
