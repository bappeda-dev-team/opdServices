package main

import (
	"fmt"
	"opdServices/app"
	_ "opdServices/docs"
	"opdServices/helper"
	"os"

	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) *echo.Echo {
	return e
}

// @title Operasional Daerah Service API
// @version 1.0
// @description API For Master Operasional Daerah Services
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8002
// @BasePath /
// @schemes http

func main() {

	app.RunFlyway()

	server := InitializedServer()
	host := os.Getenv("host")
	port := os.Getenv("port")

	addr := fmt.Sprintf("%s:%s", host, port)

	err := server.Start(addr)
	helper.PanicIfError(err)
}
