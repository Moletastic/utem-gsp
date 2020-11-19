package store

import (
	"errors"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type UserRegisterReq struct {
	Form RegisterForm `json:"user"`
}

func (r *UserRegisterReq) Bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	u.Account = models.Account{}
	u.Account.InitGSP("access:account")
	u.Account.FirstName = r.Form.FirstName
	u.Account.LastName = r.Form.LastName
	u.Account.RUT = r.Form.RUT
	u.Account.Nick = r.Form.Nick
	u.Account.Email = r.Form.Email
	u.Account.AccountType = r.Form.AccountType
	if r.Form.Password != u.Account.Password {
		h, err := u.Account.HashPassword(r.Form.Password)
		if err != nil {
			return err
		}
		u.Account.Password = h
	}
	return nil
}

type RegisterForm struct {
	FirstName   string `json:"first_name" validated:"required"`
	LastName    string `json:"last_name" validated:"required"`
	Nick        string `json:"nick" validated:"required"`
	RUT         string `json:"rut" validated:"required"`
	Email       string `json:"email" gorm:"unique" validated:"required, email"`
	Password    string `json:"password" validated:"required"`
	AccountType string `json:"account_type" validated:"required"`
}

type AccessStore struct {
	db      *gorm.DB
	Account *services.CRUDHandler
	Related []*services.CRUDHandler
}

func NewAccessStore(db *gorm.DB) *AccessStore {
	teacher := services.NewCrudService(
		&models.Teacher{},
		models.Teacher{},
		"access:teacher",
		[]string{"Account"},
		db,
	)
	user := services.NewCRUDHandler("account", services.NewCrudService(
		&models.Account{},
		models.Account{},
		"access:user",
		[]string{},
		db,
	))
	admin := services.NewCrudService(
		&models.Admin{},
		models.Admin{},
		"access:admin",
		[]string{"Account"},
		db,
	)
	related := []*services.CRUDHandler{
		services.NewCRUDHandler("teacher", teacher),
		services.NewCRUDHandler("admin", admin),
	}
	return &AccessStore{
		db:      db,
		Related: related,
		Account: user,
	}
}

func (as *AccessStore) GetByID(id int64) (*models.Account, error) {
	var m models.Account
	if err := as.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccessStore) GetByEmail(e string) (*models.Account, error) {
	var m models.Account
	if err := as.db.Where(&models.Account{Email: e}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccessStore) CreateUser(form RegisterForm, u *models.User) (*models.User, error) {
	ac := u.Account
	if err := as.Create(&ac); err != nil {
		return nil, err
	}
	if ac.AccountType == "Teacher" {
		t := models.Teacher{
			Account:   ac,
			AccountID: ac.ID,
		}
		if err := as.CreateTeacher(&t); err != nil {
			return nil, err
		}
	} else if ac.AccountType == "Admin" {
		a := models.Admin{
			Account:   ac,
			AccountID: ac.ID,
		}
		if err := as.CreateAdmin(&a); err != nil {
			return nil, err
		}
	}
	profiled, err := as.GetUser(&ac)
	if err != nil {
		return nil, err
	}
	return profiled, nil
}

func (as *AccessStore) GetUser(a *models.Account) (*models.User, error) {
	var profiled models.User
	if a.AccountType == "Teacher" {
		t, err := as.GetTeacherByEmail(a.Email)
		if err != nil {
			return nil, err
		}
		profiled = models.NewUser(&t.Account)
		profiled.AccountID = t.Account.ID
	} else if a.AccountType == "Admin" {
		a, err := as.GetAdminByEmail(a.Email)
		if err != nil {
			return nil, err
		}
		profiled = models.NewUser(&a.Account)
		profiled.AccountID = a.Account.ID
	}
	profiled.ID = a.ID
	return &profiled, nil
}

func (as *AccessStore) Create(u *models.Account) (err error) {
	return as.db.Create(u).Error
}

func (as *AccessStore) Update(u *models.Account) error {
	return as.db.Model(u).Update(u).Error
}

func (as *AccessStore) CreateTeacher(t *models.Teacher) (err error) {
	return as.db.Create(t).Error
}

func (as *AccessStore) CreateAdmin(a *models.Admin) (err error) {
	return as.db.Create(a).Error
}

func (as *AccessStore) GetTeacherByEmail(e string) (*models.Teacher, error) {
	var t models.Teacher
	u, err := as.GetByEmail(e)
	if err != nil {
		return nil, err
	}
	if err := as.db.Preload("Account").Where("account_id = ?", u.ID).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (as *AccessStore) UpdateTeacher(t *models.Teacher) error {
	return as.db.Model(t).Save(t).Error
}

func (as *AccessStore) GetAdminByEmail(e string) (*models.Admin, error) {
	var a models.Admin
	u, err := as.GetByEmail(e)
	if err != nil {
		return nil, err
	}
	if err := as.db.Preload("Account").Where("account_id = ?", u.ID).First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}
