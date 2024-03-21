package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nimbo1999/temperature-challenge/internal/config"
	"github.com/nimbo1999/temperature-challenge/internal/infra/observability"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	if err := config.LoadEnvVariables(); err != nil {
		panic(err)
	}
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("ZIPKIN_HOST", "http://zipkin:9411")
	viper.SetDefault("ZIPKIN_SERVICE_NAME", "weather-service")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	url := flag.String("zipkin", fmt.Sprintf("%s/api/v2/spans", viper.GetString("ZIPKIN_HOST")), "zipkin url")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry
	otelShutdown, err := observability.SetupOTelSDK(ctx, *url)
	if err != nil {
		return err
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	server := web.WebInfra{
		Port: fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.ListenAndServe(web.WeatherHandler)
	}()

	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return err
	case <-ctx.Done():
		// Wait for first CTRL+C
		// Stop receiving signal notifications as soon as possible
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	return server.ShutDown(context.Background())
}
