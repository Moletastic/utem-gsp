package services

import (
	"reflect"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
)

type Criteria map[string]interface{}

type ListParams struct {
	Limit    int      `json:"limit"`
	Criteria Criteria `json:"criteria"`
}

type ICRUDService interface {
	GetByID(interface{}, int64) error
	GetByUID(string, interface{}) error
	Create(interface{}) error
	Update(interface{}, int64) error
	Delete(interface{}, int64) error
	GetAll() (interface{}, int64, error)
	List(l ListParams) (interface{}, int64, error)
	Where(condition string, args ...interface{}) (interface{}, int64, error)
	Joins(join string) (interface{}, int64, error)
}

type CRUDService struct {
	Model      models.Model
	Type       reflect.Type
	Preloads   []string
	Results    interface{}
	Entity     models.Entity
	EntityName string
	db         *gorm.DB
}

func NewCrudService(model models.Model, v models.Entity, entity string, preloads []string, d *gorm.DB) *CRUDService {
	cs := new(CRUDService)
	t := reflect.TypeOf(v)
	cs.Entity = v
	cs.Type = t
	cs.EntityName = entity
	cs.Preloads = preloads
	cs.Results = reflect.New(reflect.SliceOf(t)).Interface()
	cs.db = d.Model(model)
	return cs
}

func (cs *CRUDService) GetNew() models.Model {
	return cs.Entity.New()
}

func (cs *CRUDService) Create(obj interface{}) error {
	db := cs.db
	err := db.Create(obj).Error
	return err
}

func (cs *CRUDService) Update(obj interface{}, id int64) error {
	db := cs.db
	return db.Where("id = ?", id).Update(obj).Error
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

func (cs *CRUDService) GetByID(item interface{}, id int64) error {
	db := cs.preload()
	return db.Where("id = ?", id).First(item).Error
}

func (cs *CRUDService) GetByUID(uid string, item interface{}) error {
	db := cs.preload()
	return db.Where("uid = ?", uid).First(item).Error
}

func (cs *CRUDService) List(l *ListParams) (interface{}, int64, error) {
	var count int64
	var err error
	db := cs.preload().Count(&count)
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

func (cs *CRUDService) Where(condition string, args ...interface{}) (interface{}, int64, error) {
	var count int64
	db := cs.preload().Count(&count)
	err := db.Where(condition, args).Find(cs.Results).Error
	if err != nil {
		return cs.Results, count, err
	}
	return cs.Results, count, nil
}

func (cs *CRUDService) Joins(join string) (interface{}, int64, error) {
	var count int64
	db := cs.preload().Count(&count)
	err := db.Joins(join).Find(cs.Results).Error
	if err != nil {
		return cs.Results, count, err
	}
	return cs.Results, count, nil
}
