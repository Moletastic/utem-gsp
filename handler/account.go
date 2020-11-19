package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Tokened struct {
	Token string `json:"token"`
}

func (h *Handler) Me(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("Token inválido"))
	}
	claims, ok := token.Claims.(*utils.GSPClaim)
	if !ok {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("Estructura de token inválida"))
	}

	a, err := h.AccStore.GetByID(claims.User.Account.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("Cuenta no encontrada"))
	}
	var userID int64
	account, err := h.AccStore.GetByEmail(claims.User.Account.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	profiled, err := h.AccStore.GetUser(a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	profiled.ID = userID
	profiled.AccountID = account.ID
	return c.JSON(http.StatusOK, profiled)
}

type NewPasswordForm struct {
	CurrentPass string `json:"current_pass"`
	NewPass     string `json:"new_pass"`
}

type NewPasswordReq struct {
	Data NewPasswordForm `json:"data"`
}

type UserUpdateResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func (h *Handler) ChangePassword(c echo.Context) error {
	req := new(NewPasswordReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		b := errors.New("Token Inválido")
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(b))
	}
	claims, ok := token.Claims.(*utils.GSPClaim)
	if !ok {
		b := errors.New("Estructura de token inválida")
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(b))
	}
	u, err := h.AccStore.GetByID(claims.User.AccountID)
	if err != nil {
		b := errors.New("Cuenta no encontrada")
		return c.JSON(http.StatusNotFound, utils.NewError(b))
	}
	if !u.CheckPassword(req.Data.CurrentPass) {
		b := errors.New("Contraseña ingresada no coincide con la almacenada")
		return c.JSON(http.StatusBadRequest, utils.NewError(b))
	}
	hash, err := u.HashPassword(req.Data.NewPass)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u.Password = hash
	if err = h.AccStore.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	profiled, err := h.AccStore.GetUser(u)
	t, err := utils.GenerateJWT(*profiled)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	res := new(UserUpdateResponse)
	res.Token = t
	res.User = *profiled
	return c.JSON(http.StatusOK, res)
}

type UpdateAccountForm struct {
	Account models.Account `json:"account"`
}

type UpdateAccountReq struct {
	Data UpdateAccountForm `json:"data"`
}

func (h *Handler) UpdateAccount(c echo.Context) error {
	req := new(UpdateAccountReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	account := req.Data.Account
	u, err := h.AccStore.GetByID(account.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusUnprocessableEntity, account)
	}
	account.Password = u.Password
	account.ID = u.ID
	if err = h.AccStore.Update(&account); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err = h.AccStore.GetByID(account.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusUnprocessableEntity, account)
	}
	profiled, err := h.AccStore.GetUser(&account)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, account)
	}
	t, err := utils.GenerateJWT(*profiled)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	res := new(UserUpdateResponse)
	res.Token = t
	res.User = *profiled
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(models.User)
	req := &store.UserRegisterReq{}
	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	pass := fmt.Sprintf("%s.%s", u.Account.FirstName, u.Account.LastName)
	hash, err := u.Account.HashPassword(pass)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u.Account.Password = hash
	profiled, err := h.AccStore.CreateUser(req.Form, u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, profiled)
}
