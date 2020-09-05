package handler

import (
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
)

type RegisterForm struct {
	Username string `json:"username" validated:"required"`
	Email    string `json:"email" validated:"required,email"`
	Password string `json:"password" validated:"required"`
}

type UserRegisterReq struct {
	Form RegisterForm `json:"user"`
}

func (r *UserRegisterReq) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.Nick = r.Form.Username
	u.Email = r.Form.Email
	if r.Form.Password != u.Password {
		h, err := u.HashPassword(r.Form.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	return nil
}

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

func NewLoginResp(u *models.User) *LoginResponse {
	r := new(LoginResponse)
	r.Token = utils.GenerateJWT(*u)
	r.User = *u
	return r
}

// SignUp registers a new User
func (h *Handler) SignUp(c echo.Context) error {
	var u models.User
	req := &UserRegisterReq{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	c.Logger().Info(utils.Pretty(u))
	if err := h.accessStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) error {
	req := &UserLoginReq{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.accessStore.GetByEmail(req.Credentials.Email)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, "")
	}
	if !u.CheckPassword(req.Credentials.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, NewLoginResp(u))
}
