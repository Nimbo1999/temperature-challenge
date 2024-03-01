package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nimbo1999/temperature-challenge/internal/infra/web/dto"
	"github.com/nimbo1999/temperature-challenge/internal/models"
	"github.com/spf13/viper"
)

type WeatherService struct {
}

func (service *WeatherService) GetData(address models.Address) (*models.Temperature, error) {
	response, err := http.Get(
		fmt.Sprintf(
			"http://api.weatherapi.com/v1/current.json?q=%s&lang=pt&key=%s",
			url.QueryEscape(address.Localidade),
			viper.GetString("WEATHER_API_KEY"),
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
	return &models.Temperature{
		Celsius:    payload.Current.Celsius,
		Fahrenheit: payload.Current.Fahrenheit,
		Kelvin:     payload.Current.Celsius + 273,
	}, nil
}
