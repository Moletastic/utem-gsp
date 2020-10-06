package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"gorm.io/gorm"
)

type EducationStore struct {
	db      *gorm.DB
	Related []*services.CRUDHandler
}

func NewEducationStore(db *gorm.DB) *EducationStore {
	career := services.NewCrudService(
		&models.Career{},
		"edu:career",
		[]string{"Department", "Students"},
		db,
	)
	department := services.NewCrudService(
		&models.Department{},
		"edu:department",
		[]string{"Careers", "Students"},
		db,
	)
	student := services.NewCrudService(
		&models.Student{},
		"edu:student",
		[]string{},
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
