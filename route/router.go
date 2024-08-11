package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitHttp() *echo.Echo {
	app := App()
	e := echo.New()
	e.Use(middleware.CORS())
	v1Routes(e.Group("/v1"), app)
	return e

}
