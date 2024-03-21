package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("ZIPKIN_HOST", "http://zipkin:9411")
	viper.SetDefault("ZIPKIN_SERVICE_NAME", "cep-service")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	url := flag.String("zipkin", fmt.Sprintf("%s/api/v2/spans", viper.GetString("ZIPKIN_HOST")), "zipkin url")
	flag.Parse()

	// Handle SIGINT (CTRL + C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry
	otelShutdown, err := setupOTelSDK(ctx, *url)
	if err != nil {
		return err
	}

	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server
	server := http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("SERVER_PORT")),
		Handler:      newHTTPHandler(),
		BaseContext:  func(l net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- server.ListenAndServe()
	}()

	// Wait for interruption.
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
	return server.Shutdown(context.Background())
}

func newHTTPHandler() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	handler := otelhttp.WithRouteTag("/", http.HandlerFunc(web.CepHandler))
	router.Handle("/", handler)
	return otelhttp.NewHandler(router, "/")
}
