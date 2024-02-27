package main

import (
	"log"

	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
)

func main() {
	web := web.WebInfra{
		Port: ":8080",
	}
	log.Fatal(web.ListenAndServe())
}
