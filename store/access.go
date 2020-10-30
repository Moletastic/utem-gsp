package store

import (
	"errors"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"github.com/jinzhu/gorm"
)

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

func (as *AccessStore) GetByID(id uint) (*models.Account, error) {
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

func (as *AccessStore) Create(u *models.Account) (err error) {
	return as.db.Create(u).Error
}

func (as *AccessStore) Update(u *models.Account) error {
	return as.db.Model(u).Save(u).Error
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
