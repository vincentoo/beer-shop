// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/biz"
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/conf"
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/data"
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/server"
	"github.com/go-kratos/beer-shop/project/app/shop-interface/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewShopService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService)
	grpcServer := server.NewGRPCServer(confServer, userService)
	app := newApp(logger, httpServer, grpcServer)
	return app, nil
}
