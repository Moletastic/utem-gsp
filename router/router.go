package router

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewLogFile() (*os.File, error) {
	file, err := os.OpenFile(
		"logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	return file, err
}

// New creates a new Echo Router
func New() *echo.Echo {
	e := echo.New()
	logconfig := middleware.LoggerConfig{}
	logconfig.Format = "[${time_rfc3339_nano}] [${method}] ${status} ${uri}\n"
	log, err := NewLogFile()
	if err != nil {
		fmt.Println("Cannot create log file. Ignoring")
	} else {
		logconfig.Output = log
	}
	e.Use(middleware.LoggerWithConfig(logconfig))
	e.Pre(middleware.RemoveTrailingSlash())
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Change after ...
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}
