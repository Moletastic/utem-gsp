package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"github.com/jinzhu/gorm"
)

type EducationStore struct {
	db      *gorm.DB
	Related []*services.CRUDHandler
}

func NewEducationStore(db *gorm.DB) *EducationStore {
	career := services.NewCrudService(
		&models.Career{},
		models.Career{},
		"edu:career",
		[]string{"Department", "Students"},
		db,
	)
	department := services.NewCrudService(
		&models.Department{},
		models.Department{},
		"edu:department",
		[]string{"Careers", "Students"},
		db,
	)
	student := services.NewCrudService(
		&models.Student{},
		models.Student{},
		"edu:student",
		[]string{"Career"},
		db,
	)
	related := []*services.CRUDHandler{
		services.NewCRUDHandler("career", career),
		services.NewCRUDHandler("department", department),
		services.NewCRUDHandler("student", student),
	}
	return &EducationStore{
		db:      db,
		Related: related,
	}
}
