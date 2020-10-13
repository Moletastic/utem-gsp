package main

import (
	"fmt"
	"log"

	"github.com/Moletastic/utem-gsp/config"
	"github.com/Moletastic/utem-gsp/db"
	"github.com/Moletastic/utem-gsp/handler"
	"github.com/Moletastic/utem-gsp/router"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/joho/godotenv"
)

const (
	Version = "0.1"
	Banner  = "GSP"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(utils.Pretty(config))
	d, err := db.NewDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	if config.DBConfig.Refresh && config.Pop {
		fmt.Println("Populate Database...")
		err = pop(d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Populate process has finished")
	}
	r := router.New()
	as := store.NewAccessStore(d)
	es := store.NewEducationStore(d)
	ps := store.NewProjectStore(d)
	h := handler.NewHandler(*as, *es, *ps)
	v1 := r.Group("/api")
	h.Register(v1)
	address := utils.GetLocalAddress(config.Port)
	r.Logger.Fatal(r.Start(address))
}
