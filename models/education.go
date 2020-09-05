package models

import "github.com/jinzhu/gorm"

type Department struct {
	gorm.Model
	GSPModel
	Name    string   `json:"name" mapstructure:"name" gorm:"unique"`
	Careers []Career `json:"careers" gorm:"->"`
}

func NewDepartment(d Department) *Department {
	d.InitGSP("edu:department")
	return &d
}

type Career struct {
	gorm.Model
	GSPModel
	Code         uint        `json:"code" mapstructure:"code" gorm:"unique"`
	Name         string      `json:"name" mapstructure:"name"`
	Department   *Department `json:"department,omitempty" mapstructure:"department" gorm:"foreignKey:DepartmentID"`
	DepartmentID int         `json:"department_id,omitempty" mapstructure:"department_id" gorm:"column:department_id"`
	Students     []Student   `json:"students" gorm:"->"`
}

func NewCareer(c Career) *Career {
	c.InitGSP("edu:career")
	return &c
}

type Student struct {
	gorm.Model
	GSPModel
	FirstName string
	LastName  string
	RUT       string
	CareerID  uint
	Career    *Career `json:"career"`
	EntryYear int     `json:"entry_year"`
}

func NewStudent(s Student) *Student {
	s.InitGSP("edu:student")
	return &s
}
