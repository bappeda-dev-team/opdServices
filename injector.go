//go:build wireinject
// +build wireinject

package main

import (
	"opdServices/app"

	"opdServices/controller"
	"opdServices/repository"
	"opdServices/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var opdSet = wire.NewSet(
	repository.NewOpdRepositoryImpl,
	wire.Bind(new(repository.OpdRepository), new(*repository.OpdRepositoryImpl)),
	service.NewOpdServiceImpl,
	wire.Bind(new(service.OpdService), new(*service.OpdServiceImpl)),
	controller.NewOpdControllerImpl,
	wire.Bind(new(controller.OpdController), new(*controller.OpdControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		opdSet,
		app.NewRouter,
	)
	return nil
}
