package models

import (
	"errors"

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

type Account struct {
	GSPModel
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Nick        string `json:"nick"`
	RUT         string `json:"rut"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccountType string `json:"account_type"`
}

func (ac Account) Bind(v interface{}) {
	v = Account{}
}

func (ac Account) New() Model {
	return &Account{}
}

func (u *Account) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("Contraseña vacía")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *Account) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}
