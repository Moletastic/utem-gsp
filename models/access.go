package models

import (
	"fmt"
	"strings"
)

// Teacher is a DBModel for Teacher Entity
type Teacher struct {
	GSPModel
	UserID    int64     `json:"user_id" mapstructure:"user_id"`
	User      User      `json:"user" mapstructure:"user"`
	EntryYear int       `json:"entry_year" mapstructure:"entry_year"`
	Projects  []Project `json:"projects" mapstructure:"projects"`
}

func NewTeacher(f string, l string) Teacher {
	nick := fmt.Sprintf("%s.%s", strings.ToLower(f), strings.ToLower(l))
	u := User{
		Email:     nick + "@utem.cl",
		FirstName: f,
		LastName:  l,
		Nick:      nick,
	}
	u.InitGSP("access:user")
	t := Teacher{
		EntryYear: 2010,
		User:      u,
	}
	t.InitGSP("access:teacher")
	return t
}

// Coordinator is a DBModel for Coordinator Entity
type Coordinator struct {
	GSPModel
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	User      User  `json:"user" mapstructure:"user"`
	EntryYear int   `json:"entry_year" mapstructure:"entry_year"`
}

// Admin is a DBModel for Admin Entity
type Admin struct {
	GSPModel
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	User      User  `json:"user" mapstructure:"user"`
	EntryYear int   `json:"entry_year" mapstructure:"entry_year"`
}
