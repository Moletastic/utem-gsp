package db

import (
	"fmt"
	"os/exec"

	"github.com/Moletastic/utem-gsp/config"
	"github.com/jinzhu/gorm"
)

func GetMySQLDSN(c config.DBConfig) string {
	return fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", c.User, c.Host, c.Port, c.Name)
}

func NewMySQL(config config.DBConfig) (*gorm.DB, error) {
	dsn := GetMySQLDSN(config)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func refreshMySQL(c config.DBConfig) error {
	statement := fmt.Sprintf("drop database %s; create database %s", c.Name, c.Name)
	cmd := exec.Command("mysql", "-u", "root", "-p", "-e", statement)
	return cmd.Run()
}
