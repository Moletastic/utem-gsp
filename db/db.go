package db

import (
	"fmt"

	"os"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// New creates a new Database connection
func New() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./utem-gsp.db")
	if err != nil {
		fmt.Println("Error de almacenamiento: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func NewMySQL() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/utem_gsp?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)
	if err != nil {
		fmt.Println("Error de almacenamiento: ", err)
	}
	return db
}

// TestDB creates a new Test Database connection
func TestDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return db, err
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db, nil
}

// DropTestDB removes test database file
func DropTestDB() error {
	if err := os.Remove("./utem-gsp-test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
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
}
