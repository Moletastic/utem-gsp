package db

import (
	"fmt"

	"os"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/jinzhu/gorm"
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

// TestDB creates a new Test Database connection
func TestDB(path string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		fmt.Println("Error de almacenamiento: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

// DropTestDB removes test database file
func DropTestDB() error {
	if err := os.Remove("./utem-gsp-test.db"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Department{},
		&models.Career{},
		&models.User{},
		&models.Student{},
		&models.Teacher{},
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
		&models.Project{},
	)
	db.Model(&models.Milestone{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
}

func GenerateDB(db *gorm.DB) {

}
