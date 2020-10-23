package handler

import (
	"errors"
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
)

type RegisterForm struct {
	FirstName string `json:"first_name" validated:"required"`
	LastName  string `json:"last_name" validated:"required"`
	Nick      string `json:"nick" validated:"required"`
	RUT       string `json:"rut" validated:"required"`
	Email     string `json:"email" gorm:"unique" validated:"required, email"`
	Password  string `json:"password" validated:"required"`
	UserType  string `json:"user_type" validated:"required"`
}

type UserRegisterReq struct {
	Form RegisterForm `json:"user"`
}

func (r *UserRegisterReq) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.FirstName = r.Form.FirstName
	u.LastName = r.Form.LastName
	u.RUT = r.Form.RUT
	u.Nick = r.Form.Nick
	u.Email = r.Form.Email
	u.UserType = r.Form.UserType
	if r.Form.Password != u.Password {
		h, err := u.HashPassword(r.Form.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	u.InitGSP("access:user")
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
	User  models.ProfiledUser `json:"user"`
	Token string              `json:"token"`
}

func NewLoginResp(u *models.ProfiledUser) (*LoginResponse, error) {
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
	req := &UserRegisterReq{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	c.Logger().Info(utils.Pretty(u))
	if err := h.AccStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if u.UserType == "Teacher" {
		t := models.Teacher{
			User:   u,
			UserID: u.ID,
		}
		if err := h.AccStore.CreateTeacher(&t); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
		}
	} else if u.UserType == "Admin" {
		a := models.Admin{
			User:   u,
			UserID: u.ID,
		}
		if err := h.AccStore.CreateAdmin(&a); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
		}
	}
	return c.JSON(http.StatusCreated, u)
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
		return c.JSON(http.StatusForbidden, utils.NewError(errors.New("User not found")))
	}
	if !u.CheckPassword(req.Credentials.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}

	var profiled models.ProfiledUser

	if u.UserType == "Teacher" {
		t, err := h.AccStore.GetTeacherByEmail(u.Email)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
		}
		profile := models.NewTeacherProfile(t)
		profiled = models.NewProfiledUser(u, &profile)
	} else if u.UserType == "Admin" {
		a, err := h.AccStore.GetAdminByEmail(u.Email)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
		}
		profile := models.NewAdminProfile(a)
		profiled = models.NewProfiledUser(u, &profile)
	}
	resp, err := NewLoginResp(&profiled)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, resp)
}
