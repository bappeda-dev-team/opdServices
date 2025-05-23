// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"opdServices/app"
	"opdServices/controller"
	"opdServices/repository"
	"opdServices/service"
)

import (
	_ "opdServices/docs"
)

// Injectors from injector.go:

func InitializedServer() *echo.Echo {
	opdRepositoryImpl := repository.NewOpdRepositoryImpl()
	db := app.GetConnection()
	v := _wireValue
	validate := validator.New(v...)
	opdServiceImpl := service.NewOpdServiceImpl(opdRepositoryImpl, db, validate)
	opdControllerImpl := controller.NewOpdControllerImpl(opdServiceImpl)
	echoEcho := app.NewRouter(opdControllerImpl)
	return echoEcho
}

var (
	_wireValue = []validator.Option{}
)

// injector.go:

var opdSet = wire.NewSet(repository.NewOpdRepositoryImpl, wire.Bind(new(repository.OpdRepository), new(*repository.OpdRepositoryImpl)), service.NewOpdServiceImpl, wire.Bind(new(service.OpdService), new(*service.OpdServiceImpl)), controller.NewOpdControllerImpl, wire.Bind(new(controller.OpdController), new(*controller.OpdControllerImpl)))
