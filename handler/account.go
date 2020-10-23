package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Tokened struct {
	Token string `json:"token"`
}

func (h *Handler) Me(c echo.Context) error {
	t := Tokened{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	token := strings.Split(t.Token, "Bearer ")[1]
	bytes, err := jwt.DecodeSegment(token)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, token)
	}
	profiled := &models.ProfiledUser{}
	err = json.Unmarshal(bytes, profiled)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, token)
	}
	return c.JSON(http.StatusOK, profiled)
}
