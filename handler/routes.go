package handler

import (
	"fmt"

	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	config := middleware.JWTConfig{
		SigningKey: utils.JWTSecret,
		Claims:     &utils.GSPClaim{},
		AuthScheme: "Bearer",
	}
	users := v1.Group("/access")
	users.POST("/signup", h.SignUp)
	users.POST("/login", h.Login)

	restricted := v1.Group("/gsp")
	restricted.Use(middleware.JWTWithConfig(config))
	restricted.GET("/account/me", h.Me)
	restricted.PUT("/account/me", h.UpdateAccount)
	restricted.PUT("/account/password", h.ChangePassword)
	restricted.POST("/account/user", h.CreateUser)
	for _, handler := range h.AccStore.Related {
		uri := fmt.Sprintf("/%s", handler.Name)
		uriID := fmt.Sprintf("/%s/:id", handler.Name)
		restricted.GET(uriID, handler.GetByID)
		restricted.GET(uri, handler.List)
		restricted.POST(uri, handler.Create)
		restricted.PUT(uriID, handler.Update)
		restricted.DELETE(uriID, handler.Delete)
	}

	for _, handler := range h.EduStore.Related {
		uri := fmt.Sprintf("/%s", handler.Name)
		uriID := fmt.Sprintf("/%s/:id", handler.Name)
		restricted.GET(uriID, handler.GetByID)
		restricted.GET(uri, handler.List)
		restricted.POST(uri, handler.Create)
		restricted.PUT(uriID, handler.Update)
		restricted.DELETE(uriID, handler.Delete)
	}

	for _, handler := range h.ProStore.Related {
		if handler.Name != "project" {
			uri := fmt.Sprintf("/%s", handler.Name)
			uriID := fmt.Sprintf("/%s/:id", handler.Name)
			restricted.GET(uriID, handler.GetByID)
			restricted.GET(uri, handler.List)
			restricted.POST(uri, handler.Create)
			restricted.PUT(uriID, handler.Update)
			restricted.DELETE(uriID, handler.Delete)
		}
	}
	uri := fmt.Sprintf("/%s", "project")
	uriID := fmt.Sprintf("/%s/:id", "project")
	restricted.GET(uri, h.ProStore.Project.List)
	restricted.GET(uriID, h.ProStore.Related[0].GetByID)
	restricted.POST(uri, h.ProStore.Project.Create)
	restricted.PUT(uriID, h.ProStore.Project.Update)
	restricted.DELETE(uriID, h.ProStore.Related[0].Delete)
	restricted.PATCH(uriID, h.ProStore.Project.Patch)
}
