package main

import (
	"log"

	"github.com/nimbo1999/temperature-challenge/internal/config"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
)

func init() {
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}
}

func main() {
	web := web.WebInfra{
		Port: ":8080",
	}
	log.Fatal(web.ListenAndServe())
}
