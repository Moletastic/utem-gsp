package test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Moletastic/utem-gsp/handler"
	"github.com/parnurzeal/gorequest"
)

func TestUserRegistration(t *testing.T) {

	cr := handler.Credentials{
		Email:    "boxjack.romero@gmail.com",
		Password: "pepe123",
	}
	err := LoginUser(cr)
	if err != nil {
		t.Error(err)
	}
}

func TestToken(t *testing.T) {
	cr := handler.Credentials{
		Email:    "boxjack.romero@gmail.com",
		Password: "pepe123",
	}
	request := gorequest.New()
	form := handler.UserLoginReq{
		Credentials: cr,
	}
	url := "http://localhost:1323/api/users/login"
	var data handler.LoginResponse
	resp, _, errs := request.
		Post(url).
		Send(form).
		EndStruct(&data)
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode")
	}
	url = "http://localhost:1323/api/user/teacher"
	resp, _, errs = request.Get(url).Set("Authorization", "Bearer "+data.Token).End()
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode")
	}
}

func RegisterUser(form handler.RegisterForm) error {
	request := gorequest.New()
	url := "http://localhost:1323/api/users"
	resp, _, errs := request.Post(url).Send(form).End()
	if len(errs) > 0 {
		return errs[0]
	}
	if resp.StatusCode != 201 {
		return errors.New(fmt.Sprintf("StatusCode: %d", resp.StatusCode))
	}
	return nil
}

func LoginUser(cr handler.Credentials) error {
	request := gorequest.New()
	form := handler.UserLoginReq{
		Credentials: cr,
	}
	url := "http://localhost:1323/api/users/login"
	resp, _, errs := request.Post(url).Send(form).End()
	if len(errs) > 0 {
		return errs[0]
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("StatusCode: %d", resp.StatusCode))

	}
	return nil
}
