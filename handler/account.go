package handler

import (
	"net/http"

	"github.com/Moletastic/utem-gsp/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Tokened struct {
	Token string `json:"token"`
}

func (h *Handler) Me(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*utils.GSPClaim)
	return c.JSON(http.StatusOK, claims)
}
