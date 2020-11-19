package models

// Department embeds multiple careers
type Department struct {
	GSPModel
	Name    string   `json:"name" mapstructure:"name"`
	Careers []Career `json:"careers" gorm:"->"`
}

func (d Department) Bind(v interface{}) {
	v = Department{}
}

func (d Department) New() Model {
	return &Department{}
}

// Career has many students and one Department associated
type Career struct {
	GSPModel
	Code         int64       `json:"code" mapstructure:"code" gorm:"unique"`
	Name         string      `json:"name" mapstructure:"name"`
	Department   *Department `json:"department,omitempty" mapstructure:"department" gorm:"foreignKey:DepartmentID"`
	DepartmentID int64       `json:"department_id,omitempty" mapstructure:"department_id" gorm:"column:department_id"`
	Students     []Student   `json:"students" mapstructure:"students" gorm:"->"`
}

func (c Career) Bind(v interface{}) {
	v = Career{}
}

func (c Career) New() Model {
	return &Career{}
}

// Student has one associated Career
type Student struct {
	GSPModel
	FirstName string    `mapstructure:"first_name" json:"first_name"`
	LastName  string    `mapstructure:"last_name" json:"last_name"`
	RUT       string    `mapstructure:"rut" json:"rut"`
	CareerID  int64     `mapstructure:"career_id" json:"career_id"`
	Career    *Career   `mapstructure:"career" json:"career"`
	EntryYear int       `mapstructure:"entry_year" json:"entry_year"`
	Projects  []Project `json:"projects" mapstructure:"projects" gorm:"many2many:project_authors;"`
}

func (s Student) Bind(v interface{}) {
	v = Student{}
}

func (s Student) New() Model {
	return &Student{}
}

func NewStudent(f string, l string, cid int64) Student {
	s := Student{
		FirstName: f,
		LastName:  l,
		CareerID:  cid,
		EntryYear: 2010,
	}
	s.InitGSP("edu:student")
	return s
}
