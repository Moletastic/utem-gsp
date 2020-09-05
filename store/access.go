package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
)

type AccessStore struct {
	db *gorm.DB
}

func NewAccessStore(db *gorm.DB) *AccessStore {
	return &AccessStore{
		db: db,
	}
}

func (as *AccessStore) GetByID(id uint) (*models.User, error) {
	var m models.User
	if err := as.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (as *AccessStore) GetByEmail(e string) (*models.User, error) {
	var m models.User
	if err := as.db.Where(&models.User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
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
	return as.db.Model(u).Update(u).Error
}

func (as *AccessStore) CreateTeacher(t *models.Teacher) (err error) {
	return as.db.Create(t).Error
}

func (as *AccessStore) GetTeacherByEmail(e string) (*models.Teacher, error) {
	var t models.Teacher
	u, err := as.GetByEmail(e)
	if err != nil {
		return nil, err
	}
	if err := as.db.Preload("User").Where("user_id = ?", u.ID).First(&t).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &t, nil
}

func (as *AccessStore) UpdateTeacher(t *models.Teacher) error {
	return as.db.Model(t).Update(t).Error
}
