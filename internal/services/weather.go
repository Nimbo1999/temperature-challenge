package services

import (
	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
)

type weatherService struct {
	repository repository.WeatherRepository
}

func NewWeatherService(repository repository.WeatherRepository) *weatherService {
	return &weatherService{repository}
}

func (service *weatherService) GetData(address entity.Address) (*entity.Temperature, error) {
	payload, err := service.repository.GetWeatherByAddress(address)
	if err != nil {
		return nil, err
	}
	return &entity.Temperature{
		Celsius:    payload.Current.Celsius,
		Fahrenheit: payload.Current.Fahrenheit,
		Kelvin:     payload.Current.Celsius + 273,
	}, nil
}
