package services

import (
	"context"
	"errors"

	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	ErrCepNotValid = errors.New("invalid zipcode")
	ErrCepNotFound = errors.New("can not find zipcode")
)

type viaCEPService struct {
	repository repository.CepRepository
}

func NewViaCepService(repository repository.CepRepository) *viaCEPService {
	return &viaCEPService{
		repository: repository,
	}
}

func (service *viaCEPService) GetData(ctx context.Context, cep string) (*entity.Address, error) {
	tr := otel.GetTracerProvider().Tracer("weather-handler-component")
	_, span := tr.Start(ctx, "get-city-name")
	defer span.End()
	span.SetAttributes(attribute.String("CEP", cep))

	address := entity.Address{Cep: cep}
	if !address.IsCepValid() {
		return nil, ErrCepNotValid
	}

	payload, err := service.repository.GetAddressByCep(address.Cep)
	if err != nil {
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
