package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebInfra struct {
	Port string
}

func (web *WebInfra) ListenAndServe() error {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!\n"))
	})

	server := http.Server{
		Addr:    web.Port,
		Handler: r,
	}

	return server.ListenAndServe()
}
