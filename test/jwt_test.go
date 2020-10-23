package test

import (
	"testing"

	"github.com/Moletastic/utem-gsp/handler"
	"github.com/parnurzeal/gorequest"
)

type GSPClient struct {
	BaseURL string `json:"base_url"`
	Client  *gorequest.SuperAgent
}

type Tokened struct {
	Token string `json:token`
}

func NewGSPClient(b string) *GSPClient {
	gsp := new(GSPClient)
	gsp.BaseURL = b
	gsp.Client = gorequest.New()
	return gsp
}

func TestJWT(t *testing.T) {
	c := NewGSPClient("http://localhost:1323/api")
	login := handler.UserLoginReq{
		Credentials: handler.Credentials{
			Email:    "jacob@utem.cl",
			Password: "admin123",
		},
	}
	var data Tokened
	resp, _, errs := c.Client.Post(c.BaseURL + "/users/login").Send(login).EndStruct(&data)
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode == 401 || resp.StatusCode == 400 || resp.StatusCode == 422 {
		t.Error(resp.StatusCode)
	}
	resp, _, errs = c.Client.Get(c.BaseURL+"/gsp/project").Set("Authorization", ""+data.Token).End()
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode == 401 || resp.StatusCode == 400 || resp.StatusCode == 422 {
		t.Error(resp.StatusCode)
	}
}
