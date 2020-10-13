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
			err := refreshMySQL(*dbconfig)
			if err != nil {
				return nil, err
			}
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
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Department{},
		&models.Career{},
		&models.User{},
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
		&models.Project{},
	).Error
	return err
}
