package services

import (
	"reflect"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
)

type ListParams struct {
	Limit    int
	Criteria map[string]string
}

type ICRUDService interface {
	Create(interface{}) error
	Update(interface{}, uint) error
	Delete(interface{}, uint) error
	GetAll() (interface{}, uint, error)
	List(l ListParams) (interface{}, uint, error)
	GetByID(uint, interface{}) error
	GetByUID(string, interface{}) error
}

type CRUDService struct {
	Model    models.Model
	Preloads []string
	Results  interface{}
	Entity   string
	db       *gorm.DB
}

func NewCrudService(model models.Model, entity string, preloads []string, d *gorm.DB) *CRUDService {
	cs := new(CRUDService)
	cs.Model = model
	cs.Entity = entity
	cs.Preloads = preloads
	t := reflect.TypeOf(cs.Model)
	cs.Results = reflect.New(reflect.SliceOf(t)).Interface()
	cs.db = d.Model(model)
	return cs
}

func (cs *CRUDService) Create(obj interface{}) error {
	return cs.db.Create(obj).Error
}

func (cs *CRUDService) Update(obj interface{}, id uint) error {
	return cs.db.Where("id = ?", id).Update(obj).Error
}

func (cs *CRUDService) Delete(obj interface{}, id uint) error {
	return cs.db.Where("id = ?", id).Delete(obj).Error
}

func (cs *CRUDService) GetAll() (interface{}, uint, error) {
	var count uint
	db := cs.preload().Count(&count)
	err := db.Find(cs.Results).Error
	if err != nil {
		return cs.Results, count, err
	}
	return cs.Results, count, nil
}

func (cs *CRUDService) preload() *gorm.DB {
	db := cs.db
	if len(cs.Preloads) > 0 {
		for _, preload := range cs.Preloads {
			db = db.Preload(preload)
		}
	}
	return db
}

func (cs *CRUDService) GetByID(id uint, item interface{}) error {
	db := cs.preload()
	return db.Where("id = ?", id).First(item).Error
}

func (cs *CRUDService) GetByUID(uid string, item interface{}) error {
	db := cs.preload()
	return db.Where("uid = ?", uid).First(item).Error
}

func (cs *CRUDService) List(l *ListParams) (interface{}, uint, error) {
	var count uint
	db := cs.preload().Count(&count)
	var err error
	if l == nil {
		err = db.Find(cs.Results).Error
	} else {
		limit := l.Limit
		if limit == 0 {
			limit = -1
		}
		if len(l.Criteria) == 0 {
			err = db.Limit(limit).Find(cs.Results).Error
		} else {
			err = db.Limit(limit).Find(cs.Results, l.Criteria).Error
		}
	}
	if err != nil {
		return cs.Results, count, err
	}
	return cs.Results, count, nil
}
