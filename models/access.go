package models

import (
	"github.com/jinzhu/gorm"
)

type Teacher struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	EntryYear int       `json:"entry_year"`
	Projects  []Project `gorm:"many2many:project_guides"`
}

type Coordinator struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	User      User `json:"user"`
	EntryYear int  `json:"entry_year"`
}

type Admin struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	User      User `json:"user"`
	EntryYear int  `json:"entry_year"`
}
