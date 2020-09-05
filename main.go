package main

import (
	"github.com/Moletastic/utem-gsp/db"
	"github.com/Moletastic/utem-gsp/handler"
	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/router"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/jinzhu/gorm"
)

func pop(d *gorm.DB) {
	es := store.NewEducationStore(d)
	department := &models.Department{
		Name: "Departamento de Informática y Computación",
	}
	career_1 := models.NewCareer(models.Career{
		Code:       21030,
		Name:       "Ingenieria en informática",
		Department: department,
	})
	career_2 := models.NewCareer(models.Career{
		Code:       21041,
		Name:       "Ingeniería civil en computación, mención informática",
		Department: department,
	})
	es.CreateCareer(career_1)
	es.CreateCareer(career_2)
	es.CreateDepartment(department)
}

func main() {
	db.DropTestDB()
	d := db.TestDB("./utem-gsp-test.db")
	db.AutoMigrate(d)
	r := router.New()
	as := store.NewAccessStore(d)
	es := store.NewEducationStore(d)
	ps := store.NewProjectStore(d)
	h := handler.NewHandler(*as, *es, *ps)
	v1 := r.Group("/api")
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:1323"))

}
