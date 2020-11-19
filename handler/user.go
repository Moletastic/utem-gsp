package handler

import (
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginReq struct {
	Credentials Credentials `json:"user"`
}

func (l *UserLoginReq) bind(c echo.Context) error {
	if err := c.Bind(l); err != nil {
		return err
	}
	return nil
}

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func NewLoginResp(u *models.User) (*LoginResponse, error) {
	r := new(LoginResponse)
	token, err := utils.GenerateJWT(*u)
	if err != nil {
		return nil, err
	}
	r.Token = token
	r.User = *u
	return r, nil
}

// SignUp registers a new User
func (h *Handler) SignUp(c echo.Context) error {
	var u models.User
	req := &store.UserRegisterReq{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if _, err := h.AccStore.CreateUser(req.Form, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, utils.NewSimpleReq("Usuario creado con éxito"))
}

func (h *Handler) Login(c echo.Context) error {
	req := &UserLoginReq{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.AccStore.GetByEmail(req.Credentials.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.NewSimpleReq("User not found"))
	}
	if !u.CheckPassword(req.Credentials.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	profiled, err := h.AccStore.GetUser(u)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	resp, err := NewLoginResp(profiled)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
