package main

import (
	"core/route"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	// config.DbInit()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e = route.InitHttp()

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
