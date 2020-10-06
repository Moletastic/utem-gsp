package store

import (
	"errors"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"gorm.io/gorm"
)

type AccessStore struct {
	db      *gorm.DB
	User    *services.CRUDHandler
	Related []*services.CRUDHandler
}

func NewAccessStore(db *gorm.DB) *AccessStore {
	teacher := services.NewCrudService(
		&models.Teacher{},
		"access:teacher",
		[]string{"User"},
		db,
	)
	user := services.NewCRUDHandler("user", services.NewCrudService(
		&models.User{},
		"access:user",
		[]string{},
		db,
	))
	admin := services.NewCrudService(
		&models.Admin{},
		"access:admin",
		[]string{"User"},
		db,
	)
	related := []*services.CRUDHandler{
		services.NewCRUDHandler("teacher", teacher),
		services.NewCRUDHandler("admin", admin),
	}
	return &AccessStore{
		db:      db,
		Related: related,
		User:    user,
	}
}

func (as *AccessStore) GetByID(id uint) (*models.User, error) {
	var m models.User
	if err := as.db.First(&m, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccessStore) GetByEmail(e string) (*models.User, error) {
	var m models.User
	if err := as.db.Where(&models.User{Email: e}).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccessStore) Create(u *models.User) (err error) {
	return as.db.Create(u).Error
}

func (as *AccessStore) Update(u *models.User) error {
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
	if err := as.db.Preload("User").Where("user_id = ?", u.ID).First(&t).Error; err != nil {
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
	if err := as.db.Preload("User").Where("user_id = ?", u.ID).First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}
