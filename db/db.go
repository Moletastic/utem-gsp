package db

import (
	"errors"
	"fmt"

	"github.com/Moletastic/utem-gsp/config"
	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewDB(dbconfig *config.DBConfig) (*gorm.DB, error) {
	db := new(gorm.DB)
	switch dbconfig.Engine {
	case "sqlite":
		if dbconfig.Refresh {
			err := refreshSQLite(*dbconfig)
			if err != nil {
				return nil, err
			}
		}
		sqlite, err := NewSQLite(*dbconfig)
		if err != nil {
			return nil, err
		}
		db = sqlite
	case "mysql":
		if dbconfig.Refresh {
			fmt.Println("Not implemented yet")
			//err := refreshMySQL(*dbconfig)
			//if err != nil {
			//return nil, err
			//}
		}
		mysql, err := NewMySQL(*dbconfig)
		if err != nil {
			return nil, err
		}
		db = mysql
	default:
		message := fmt.Sprintf("Unsupported engine: %s", dbconfig.Engine)
		return nil, errors.New(message)
	}
	if dbconfig.Refresh {
		err := AutoMigrate(db)
		if err != nil {
			return nil, err
		}
	}
	if dbconfig.LogMode {
		db.LogMode(true)
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Project{},
		&models.Department{},
		&models.Career{},
		&models.Account{},
		&models.Student{},
		&models.Teacher{},
		&models.Admin{},
		&models.LinkType{},
		&models.Link{},
		&models.Subject{},
		&models.Channel{},
		&models.Meet{},
		&models.Commit{},
		&models.Milestone{},
		&models.Progress{},
		&models.ProjectState{},
		&models.ProjectType{},
		&models.Rubric{},
		&models.Review{},
	).Error
	db.Model(&models.Career{}).AddForeignKey("department_id", "departments(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Student{}).AddForeignKey("career_id", "careers(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Teacher{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Admin{}).AddForeignKey("account_id", "accounts(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Link{}).AddForeignKey("link_type_id", "link_types(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Link{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Meet{}).AddForeignKey("channel_id", "channels(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Meet{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Commit{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Milestone{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Progress{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Review{}).AddForeignKey("reviewer_id", "teachers(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Review{}).AddForeignKey("rubric_id", "rubrics(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Review{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Project{}).AddForeignKey("project_state_id", "project_states(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Project{}).AddForeignKey("project_type_id", "project_types(id)", "RESTRICT", "RESTRICT")

	db.Table("project_authors").AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Table("project_authors").AddForeignKey("student_id", "students(id)", "CASCADE", "CASCADE")
	db.Table("project_guides").AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Table("project_guides").AddForeignKey("teacher_id", "teachers(id)", "CASCADE", "CASCADE")
	db.Table("project_subjects").AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	db.Table("project_subjects").AddForeignKey("subject_id", "subjects(id)", "CASCADE", "CASCADE")

	if err != nil {
		return err
	}
	return err
}
