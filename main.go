package main

import (
	"log"

	"github.com/Moletastic/utem-gsp/db"
	"github.com/Moletastic/utem-gsp/handler"
	"github.com/Moletastic/utem-gsp/router"
	"github.com/Moletastic/utem-gsp/store"
)

const (
	Version = "0.1"
	Banner  = "GSP"
)

func main() {
	db.DropTestDB()
	//d := db.NewMySQL()
	d, err := db.TestDB("./utem-gsp-test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(d)
	if err != nil {
		log.Fatal(err)
	}
	err = pop(d)
	if err != nil {
		log.Fatal(err)
	}
	r := router.New()
	as := store.NewAccessStore(d)
	es := store.NewEducationStore(d)
	ps := store.NewProjectStore(d)
	h := handler.NewHandler(*as, *es, *ps)
	v1 := r.Group("/api")
	h.Register(v1)
	r.Logger.Fatal(r.Start("127.0.0.1:1323"))
}
