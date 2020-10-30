package models

import (
	"fmt"
	"strings"
)

// Teacher is a DBModel for Teacher Entity
type Teacher struct {
	GSPModel
	AccountID int64     `json:"account_id" mapstructure:"account_id"`
	Account   Account   `json:"account" mapstructure:"account"`
	EntryYear int       `json:"entry_year" mapstructure:"entry_year"`
	Projects  []Project `gorm:"many2many:project_guides" json:"projects" mapstructure:"projects"`
}

func (t Teacher) Bind(v interface{}) {
	v = Teacher{}
}

func (t Teacher) New() Model {
	return &Teacher{}
}

func NewTeacher(f string, l string) Teacher {
	nick := fmt.Sprintf("%s.%s", strings.ToLower(f), strings.ToLower(l))
	u := Account{
		Email:     nick + "@utem.cl",
		FirstName: f,
		LastName:  l,
		Nick:      nick,
	}
	u.InitGSP("access:user")
	t := Teacher{
		EntryYear: 2010,
		Account:   u,
	}
	t.InitGSP("access:teacher")
	return t
}

// Coordinator is a DBModel for Coordinator Entity
type Coordinator struct {
	GSPModel
	AccountID int64   `json:"account_id" mapstructure:"account_id"`
	Account   Account `json:"account" mapstructure:"account"`
	EntryYear int     `json:"entry_year" mapstructure:"entry_year"`
}

// Admin is a DBModel for Admin Entity
type Admin struct {
	GSPModel
	AccountID int64   `json:"account_id" mapstructure:"account_id"`
	Account   Account `json:"account" mapstructure:"account"`
	EntryYear int     `json:"entry_year" mapstructure:"entry_year"`
}

func (a Admin) Bind(v interface{}) {
	v = Admin{}
}

func (a Admin) New() Model {
	return &Admin{}
}
