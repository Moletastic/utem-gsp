package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UType int

const (
	UTeacher UType = iota + 1
	UCoordinator
	UAdmin
)

func (t UType) String() string {
	s := ""
	switch t {
	case 1:
		s = "Teacher"
	case 2:
		s = "Coordinator"
	case 3:
		s = "Admin"
	}
	return s
}

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nick      string `json:"nick"`
	RUT       string `json:"rut"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	UserType  string `json:"user_type"`
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("Contraseña vacía")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
