package services

import (
	"context"

	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
	"go.opentelemetry.io/otel"
)

type weatherService struct {
	repository repository.WeatherRepository
}

func NewWeatherService(repository repository.WeatherRepository) *weatherService {
	return &weatherService{repository}
}

func (service *weatherService) GetData(ctx context.Context, address entity.Address) (*entity.Temperature, error) {
	tr := otel.GetTracerProvider().Tracer("weather-handler-component")
	_, span := tr.Start(ctx, "get-city-weather")
	defer span.End()

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
