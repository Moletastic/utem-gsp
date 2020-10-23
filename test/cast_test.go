package test

import (
	"testing"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/parnurzeal/gorequest"
)

func TestCast(t *testing.T) {
	url1 := "http://localhost:1323/api/gsp/project/26"
	url2 := "http://localhost:1323/api/gsp/project/11"
	request := gorequest.New()
	p1 := new(models.Project)
	p2 := new(models.Project)
	resp, _, errs := request.Get(url1).EndStruct(p1)
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode")
	}
	resp, _, errs = request.Get(url2).EndStruct(p2)
	if len(errs) > 0 {
		t.Error(errs)
	}
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode")
	}
	if len(p1.Authors) == 0 || len(p2.Authors) == 0 {
		t.Error("No authors in p1 or p2 projects")
	}
	if p1.Authors[0].ID == p2.Authors[0].ID {
		t.Errorf("\np1_id: %d\np2_id: %d\n", p1.Authors[0].ID, p2.Authors[0].ID)
		t.Error("Same author")
	}
}
