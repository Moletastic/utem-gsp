package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New creates a new Echo Router
func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	//e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//Format: "[${method}] ${status} ${uri}\n",
	//}))
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Change after ...
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
