package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
)

type ProjectStore struct {
	db *gorm.DB
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

type ListConf struct {
	Criteria map[string]interface{}
	Limit    int8
}

func (ps *ProjectStore) ListProjects(l *ListConf) ([]models.Project, int, error) {
	var (
		projects []models.Project
		count    int
	)
	ps.db.Model(&projects).Count(&count)

	db := ps.db.
		Preload("Authors").
		Preload("Authors.Career").
		Preload("Milestones").
		Preload("Meets")

	if l == nil {
		db.Find(&projects)
	} else {
		limit := l.Limit
		if limit == 0 {
			limit = -1
		}
		if len(l.Criteria) == 0 {
			db.Limit(limit).Find(&projects)
		} else {
			db.Limit(limit).Find(&projects, l.Criteria)
		}
	}

	return projects, count, nil
}

func (ps *ProjectStore) GetMilestonesByProject(p models.Project) ([]models.Milestone, error) {
	var (
		milestones []models.Milestone
	)
	ps.db.Model(&milestones).Related(&p)
	ps.db.Find(&milestones)
	return milestones, nil
}

func (ps *ProjectStore) GetByID(id uint) (*models.Project, error) {
	var project models.Project
	err := ps.db.Where("id = ?", id).First(&project).Error
	return &project, err
}

func (ps *ProjectStore) CreateProject(p *models.Project) error {
	subjects := p.Subjects
	tx := ps.db.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	for _, s := range p.Subjects {
		if err := tx.Where("name = ?", s.Name).First(&s).Error; gorm.IsRecordNotFoundError(err) {
			if err := tx.Create(&s).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	p.Subjects = subjects
	return tx.Commit().Error
}

func (ps *ProjectStore) UpdateProject(p *models.Project) error {
	return ps.db.Save(p).Error
}
