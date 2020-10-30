package main

import (
	"fmt"
	"log"

	"github.com/Moletastic/utem-gsp/config"
	"github.com/Moletastic/utem-gsp/db"
	"github.com/Moletastic/utem-gsp/handler"
	"github.com/Moletastic/utem-gsp/pop"
	"github.com/Moletastic/utem-gsp/router"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/joho/godotenv"
)

const (
	Version = "0.5"
	Banner  = `
 _   _ _____ ___ __  __  ___ ___ ___
| | | |_   _| __|  \/  |/ __/ __| _ \
| |_| | | | | _|| |\/| | (_ \__ \  _/
 \___/  |_| |___|_|  |_|\___|___/_|
	`
)

func printInfo() {
	fmt.Println(Banner)
	fmt.Printf("Version: %s\n", Version)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	d, err := db.NewDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	if config.Pop {
		fmt.Println("Populate Database...")
		err = pop.Populate(d)
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
	printInfo()
	r.Logger.Fatal(r.Start(config.GetAddress()))
}
