//go:build wireinject
// +build wireinject

package main

import (
	"quick-go/app/http/controllers"
	"quick-go/app/repo"
	"quick-go/app/service"
	"quick-go/bootstrap"
	"quick-go/global"
	"quick-go/routers"

	"github.com/google/wire"
)

func InitServer() (*bootstrap.Server, func(), error) {
	wire.Build(
		global.NewData,
		repo.NewMysqlStockRepository,
		service.NewStockService,
		controllers.NewStockController,
		// repo.NewMysqlSpuRepository,
		// service.NewSpuService,
		// controllers.NewSpuController,
		routers.NewRouter,
		bootstrap.NewServer,
		bootstrap.NewGinEngine,
	)
	return &bootstrap.Server{}, nil, nil
}
