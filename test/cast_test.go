package test

import (
	"testing"

	"github.com/Moletastic/utem-gsp/handler"
	"github.com/Moletastic/utem-gsp/models"
	"github.com/parnurzeal/gorequest"
)

func TestCast(t *testing.T) {
	loginreq := gorequest.New()
	loginuri := "http://localhost:1323/api/access/login"
	login := handler.UserLoginReq{
		Credentials: handler.Credentials{
			Email:    "jacob@utem.cl",
			Password: "admin123",
		},
	}
	var data Tokened
	resp, body, errs := loginreq.Post(loginuri).Send(login).EndStruct(&data)
	if resp.StatusCode != 200 {
		if len(errs) > 0 {
			t.Error(errs)
		}
		t.Error(body)
		t.Error(resp.StatusCode)
		return
	}
	url1 := "http://localhost:1323/api/gsp/project/26"
	url2 := "http://localhost:1323/api/gsp/project/11"
	request := gorequest.New()
	p1 := new(models.Project)
	p2 := new(models.Project)
	resp, _, errs = request.Get(url1).Set("Authorization", "Bearer "+data.Token).EndStruct(p1)
	if len(errs) > 0 {
		t.Error(errs)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode")
		return
	}
	resp, _, errs = request.Get(url2).Set("Authorization", "Bearer "+data.Token).EndStruct(p2)
	if len(errs) > 0 {
		t.Error(errs)
		return
	}
	if resp.StatusCode != 200 {
		t.Error(resp.StatusCode)
		return
	}
	if len(p1.Authors) == 0 || len(p2.Authors) == 0 {
		t.Error("No authors in p1 or p2 projects")
		return
	}
	if p1.Authors[0].ID == p2.Authors[0].ID {
		t.Errorf("\np1_id: %d\np2_id: %d\n", p1.Authors[0].ID, p2.Authors[0].ID)
		t.Error("Same author")
		return
	}
}
