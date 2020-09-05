package handler

import (
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	access := v1.Group("/users")
	access.POST("", h.SignUp)
	access.POST("/login", h.Login)

	edu := v1.Group("/user")
	edu.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: utils.JWTSecret,
		Claims:     &utils.GSPClaim{},
		AuthScheme: "Bearer",
	}))
	edu.GET("/teacher", h.ListTeachers)
	edu.POST("/project", h.CreateProject)
	edu.GET("/project", h.ListProjects)
	access.POST("/careers", h.CreateCareer)
	access.GET("/careers", h.ListCareers)
}
