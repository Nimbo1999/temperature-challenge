package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web/dto"
	"github.com/spf13/viper"
)

type CepRepository interface {
	GetAddressByCep(cep string) (map[string]string, error)
}

type cepRepository struct{}

func NewCepRepository() *cepRepository {
	return &cepRepository{}
}

func (repo *cepRepository) GetAddressByCep(cep string) (map[string]string, error) {
	response, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	payload := make(map[string]string)
	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload, nil
}

type WeatherRepository interface {
	GetWeatherByAddress(address entity.Address) (*dto.WeatherApiResponseDTO, error)
}

type weatherRepository struct{}

func NewWeatherRepository() *weatherRepository {
	return &weatherRepository{}
}

func (repo *weatherRepository) GetWeatherByAddress(address entity.Address) (*dto.WeatherApiResponseDTO, error) {
	response, err := http.Get(
		fmt.Sprintf(
			"http://api.weatherapi.com/v1/current.json?q=%s&lang=pt&key=%s",
			url.QueryEscape(address.Localidade),
			url.QueryEscape(viper.GetString("WEATHER_API_KEY")),
		),
	)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var payload dto.WeatherApiResponseDTO
	if err = json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return &payload, nil
}
