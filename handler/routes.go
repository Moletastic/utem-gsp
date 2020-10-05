package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	//config := middleware.JWTConfig{
	//SigningKey: utils.JWTSecret,
	//Claims:     &utils.GSPClaim{},
	//AuthScheme: "Bearer",
	//}
	users := v1.Group("/users")
	users.POST("", h.SignUp)
	users.POST("/login", h.Login)

	restricted := v1.Group("/gsp")
	//restricted.Use(middleware.JWTWithConfig(config))
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
		uri := fmt.Sprintf("/%s", handler.Name)
		uriID := fmt.Sprintf("/%s/:id", handler.Name)
		restricted.GET(uriID, handler.GetByID)
		restricted.GET(uri, handler.List)
		restricted.POST(uri, handler.Create)
		restricted.PUT(uriID, handler.Update)
		restricted.DELETE(uriID, handler.Delete)
	}
}
