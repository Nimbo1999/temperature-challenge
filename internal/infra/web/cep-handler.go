package web

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web/dto"
	"github.com/nimbo1999/temperature-challenge/internal/services"
	"github.com/spf13/viper"
)

func CepHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CepDTO
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	address := entity.Address{
		Cep: payload.Cep,
	}
	if isValid := address.IsCepValid(); !isValid {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(services.ErrCepNotValid.Error() + "\n"))
		return
	}
	response, err := http.Get(fmt.Sprintf("http://%s/%s", viper.GetString("WEATHER_API_HOST"), address.Cep))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
