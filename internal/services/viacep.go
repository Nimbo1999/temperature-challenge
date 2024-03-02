package services

import (
	"errors"

	"github.com/nimbo1999/temperature-challenge/internal/entity"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
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

func (service *viaCEPService) GetData(cep string) (*entity.Address, error) {
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
