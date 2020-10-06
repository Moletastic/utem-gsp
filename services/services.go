package services

import (
	"fmt"
	"reflect"

	"github.com/Moletastic/utem-gsp/models"
	"gorm.io/gorm"
)

type ListParams struct {
	Limit    int
	Criteria map[string]string
}

type ICRUDService interface {
	Create(interface{}) error
	Update(interface{}, int64) error
	Delete(interface{}, int64) error
	GetAll() (interface{}, int64, error)
	List(l ListParams) (interface{}, int64, error)
	GetByID(int64, interface{}) error
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
	db := cs.db
	err := db.Create(obj).Error
	return err
}

func (cs *CRUDService) Update(obj interface{}, id int64) error {
	db := cs.db
	return db.Save(obj).Where("id = ?", id).Error
}

func (cs *CRUDService) Delete(obj interface{}, id int64) error {
	db := cs.db
	return db.Where("id = ?", id).Delete(obj).Error
}

func (cs *CRUDService) GetAll() (interface{}, int64, error) {
	var count int64
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

func (cs *CRUDService) GetByID(id int64, item interface{}) error {
	db := cs.preload()
	return db.Where("id = ?", id).First(item).Error
}

func (cs *CRUDService) GetByUID(uid string, item interface{}) error {
	db := cs.preload()
	return db.Where("uid = ?", uid).First(item).Error
}

func (cs *CRUDService) List(l *ListParams) (interface{}, int64, error) {
	var count int64
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
