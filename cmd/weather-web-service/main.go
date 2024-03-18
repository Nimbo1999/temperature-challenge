package main

import (
	"fmt"
	"log"

	"github.com/nimbo1999/temperature-challenge/internal/config"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}
	viper.SetDefault("SERVER_PORT", "8080")
}

func main() {
	server := web.WebInfra{
		Port: fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
	}
	log.Fatal(server.ListenAndServe(web.WeatherHandler))
}
