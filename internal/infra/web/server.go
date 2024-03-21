package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type WebInfra struct {
	Port   string
	server *http.Server
}

type handleFuncAdapter func(http.ResponseWriter, *http.Request)

func (h handleFuncAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func (web *WebInfra) ListenAndServe(handlerParameter func(w http.ResponseWriter, r *http.Request)) error {
	handler := web.getHandler(handlerParameter)
	web.server = &http.Server{
		Addr:    web.Port,
		Handler: otelhttp.NewHandler(handler, "/"),
	}
	return web.server.ListenAndServe()
}

func (web *WebInfra) ShutDown(ctx context.Context) error {
	return web.server.Shutdown(ctx)
}

func (web *WebInfra) getHandler(handlerParameter func(w http.ResponseWriter, r *http.Request)) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{middleware.RequestIDHeader},
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	handler := otelhttp.WithRouteTag("/", handleFuncAdapter(handlerParameter))
	r.Handle("/{cep}", handler)
	return r
}
