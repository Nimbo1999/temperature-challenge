package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web/dto"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
	"github.com/nimbo1999/temperature-challenge/internal/services"
)

type WebInfra struct {
	Port string
}

func handleCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	viacepService := services.NewViaCepService(repository.NewCepRepository())
	address, err := viacepService.GetData(cep)

	switch err {
	case services.ErrCepNotValid:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error() + "\n"))
		return
	case services.ErrCepNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error() + "\n"))
		return
	case nil:
		break
	default:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	weatherService := services.NewWeatherService(repository.NewWeatherRepository())
	temperature, err := weatherService.GetData(*address)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TemperatureResponseDTO{
		Celsius:    temperature.Celsius,
		Fahrenheit: temperature.Fahrenheit,
		Kelvin:     temperature.Kelvin,
	})
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
