package db

import (
	"fmt"
	"os"

	"github.com/Moletastic/utem-gsp/config"
	"github.com/jinzhu/gorm"
)

// NewSQLite creates a new Database connection
func NewSQLite(config config.DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Name)
	if err != nil {
		return nil, err
	}
	return db, err
}

func refreshSQLite(config config.DBConfig) error {
	if err := os.Remove(config.Name); err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}
