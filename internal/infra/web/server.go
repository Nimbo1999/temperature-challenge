package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type WebInfra struct {
	Port string
}

func (web *WebInfra) ListenAndServe(handler func(w http.ResponseWriter, r *http.Request)) error {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{middleware.RequestIDHeader},
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Get("/{cep}", handler)
	server := http.Server{
		Addr:    web.Port,
		Handler: r,
	}
	return server.ListenAndServe()
}
