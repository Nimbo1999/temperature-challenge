package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT", "8080")
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Post("/", web.CepHandler)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}
