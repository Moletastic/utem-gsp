package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
)

type EducationStore struct {
	db *gorm.DB
}

func NewEducationStore(db *gorm.DB) *EducationStore {
	return &EducationStore{
		db: db,
	}
}

func (es *EducationStore) CreateEntity(gsp models.Model) error {
	return es.db.Create(gsp).Error
}

func (es *EducationStore) ListTeachers() ([]models.Teacher, int, error) {
	var (
		teachers []models.Teacher
		count    int
	)
	es.db.Model(&teachers).Count(&count)
	es.db.Preload("User").Find(&teachers)
	return teachers, count, nil
}

func (es *EducationStore) CreateDepartment(d *models.Department) error {
	return es.db.Create(d).Error
}

func (es *EducationStore) UpdateDepartment(d *models.Department) (*models.Department, error) {
	err := es.db.Save(d).Error
	if err != nil {
		return nil, err
	}
	return es.GetDepartmentByID(d.ID), nil
}

func (es *EducationStore) GetDepartmentByID(id uint) *models.Department {
	var d models.Department
	es.db.Preload("Careers").Where("id = ?", id).First(&d)
	return &d
}

func (es *EducationStore) ListDepartments() ([]models.Department, int, error) {
	var (
		departments []models.Department
		count       int
	)
	es.db.Model(&departments).Count(&count)
	es.db.Preload("Careers").Find(&departments)
	return departments, count, nil
}

func (es *EducationStore) CreateCareer(c *models.Career) (err error) {
	return es.db.Create(c).Error
}

func (es *EducationStore) UpdateCareer(c *models.Career) (*models.Career, error) {
	err := es.db.Save(c).Error
	if err != nil {
		return nil, err
	}
	return es.GetCareerByID(c.ID), nil
}

func (es *EducationStore) DeleteGSPModel(gsp *models.GSPModel) (err error) {
	return es.db.Delete(gsp).Error
}

func (es *EducationStore) ListCareers() ([]models.Career, int, error) {
	var (
		careers []models.Career
		count   int
	)
	es.db.Model(&careers).Count(&count)
	es.db.Preload("Department").Find(&careers)
	return careers, count, nil
}

func (es *EducationStore) GetCareerByID(id uint) *models.Career {
	var c models.Career
	es.db.Preload("Department").Where("id = ?", id).First(&c)
	return &c
}

func (es *EducationStore) GetCareerByCode(code uint) *models.Career {
	var c models.Career
	es.db.Preload("Department").Where("code = ?", code).First(&c)
	return &c
}
