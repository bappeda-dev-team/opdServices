package app

import (
	"opdServices/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(opdController controller.OpdController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/opd/create", opdController.Create)
	e.PUT("/opd/update/:id", opdController.Update)
	e.DELETE("/opd/delete/:id", opdController.Delete)
	e.GET("/opd/detail/:id", opdController.FindById)
	e.GET("/opd/findall", opdController.FindAll)
	e.GET("/opd/findbyopd/:kode_opd", opdController.FindByKodeOpd)
	e.GET("/opd/findallopd", opdController.FindAllOnlyOpd)

	return e
}
