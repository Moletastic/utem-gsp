package models

// Department embeds multiple careers
type Department struct {
	GSPModel
	Name    string   `json:"name" mapstructure:"name" gorm:"unique"`
	Careers []Career `json:"careers" gorm:"->"`
}

// Career has many students and one Department associated
type Career struct {
	GSPModel
	Code         uint        `json:"code" mapstructure:"code" gorm:"unique"`
	Name         string      `json:"name" mapstructure:"name"`
	Department   *Department `json:"department,omitempty" mapstructure:"department" gorm:"foreignKey:DepartmentID"`
	DepartmentID uint        `json:"department_id,omitempty" mapstructure:"department_id" gorm:"column:department_id"`
	Students     []Student   `json:"students" mapstructure:"students" gorm:"->"`
}

// Student has one associated Career
type Student struct {
	GSPModel
	FirstName string    `mapstructure:"first_name" json:"first_name"`
	LastName  string    `mapstructure:"last_name" json:"last_name"`
	RUT       string    `mapstructure:"rut" json:"rut"`
	CareerID  uint      `mapstructure:"career_id" json:"career_id"`
	Career    *Career   `mapstructure:"career" json:"career" gorm:"->"`
	EntryYear int       `mapstructure:"entry_year" json:"entry_year"`
	Projects  []Project `mapstructure:"projects" json:"projects"`
}

func NewStudent(f string, l string, cid uint) Student {
	s := Student{
		FirstName: f,
		LastName:  l,
		CareerID:  cid,
		EntryYear: 2010,
	}
	s.InitGSP("edu:student")
	return s
}
