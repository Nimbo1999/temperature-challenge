package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nimbo1999/temperature-challenge/internal/models"
)

var (
	ErrCepNotValid = errors.New("invalid zipcode")
	ErrCepNotFound = errors.New("can not find zipcode")
)

type ViaCEPService struct {
}

func (service *ViaCEPService) GetData(cep string) (*models.Address, error) {
	address := models.Address{Cep: cep}
	if !address.IsCepValid() {
		return nil, ErrCepNotValid
	}

	response, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	payload := make(map[string]string)
	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if _, ok := payload["erro"]; ok {
		return nil, ErrCepNotFound
	}

	if err = address.SetAddressFromMap(payload); err != nil {
		return nil, err
	}
	return &address, nil
}
