package web

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/temperature-challenge/internal/services"
)

type WebInfra struct {
	Port string
}

func handleCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	viacepService := services.ViaCEPService{}
	address, err := viacepService.GetData(cep)

	switch err {
	case services.ErrCepNotValid:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error() + "\n"))
	case services.ErrCepNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error() + "\n"))
	case nil:
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(address)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
	}
}

func (web *WebInfra) ListenAndServe() error {
	r := chi.NewRouter()

	r.Get("/{cep}", handleCep)

	server := http.Server{
		Addr:    web.Port,
		Handler: r,
	}

	return server.ListenAndServe()
}
