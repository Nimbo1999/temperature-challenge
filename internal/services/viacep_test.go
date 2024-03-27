package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CepTestSuit struct {
	suite.Suite
	service *viaCEPService
}

type cepRepositoryMock struct {
	payload map[string]any
}

var mock = &cepRepositoryMock{}

func (repo *cepRepositoryMock) GetAddressByCep(cep string) (map[string]any, error) {
	return repo.payload, nil
}

func (suite *CepTestSuit) Test_ShouldReturnCepIsNotValid() {
	mock.payload = map[string]any{
		"cep": "71111",
	}
	data, err := suite.service.GetData(context.Background(), mock.payload["cep"].(string))
	suite.Nil(data)
	suite.Error(err, ErrCepNotValid)
}

func (suite *CepTestSuit) Test_ShouldReturnCepNotFoundError() {
	mock.payload = map[string]any{
		"erro": "true",
	}
	data, err := suite.service.GetData(context.Background(), "")
	fmt.Println(data)
	suite.Nil(data)
	suite.Error(err, ErrCepNotFound)
}

func (suite *CepTestSuit) Test_ShouldReturnAddressWhenValidCepIsProvided() {
	mock.payload = map[string]any{
		"cep":         "01001-000",
		"logradouro":  "Praça da Sé",
		"complemento": "lado ímpar",
		"bairro":      "Sé",
		"localidade":  "São Paulo",
		"uf":          "SP",
		"ibge":        "3550308",
		"gia":         "1004",
		"ddd":         "11",
		"siafi":       "7107",
	}
	data, err := suite.service.GetData(context.Background(), mock.payload["cep"].(string))
	suite.Nil(err)
	suite.NotNil(data)
	suite.Equal(data.Cep, "01001000")
}

func (suite *CepTestSuit) SetupTest() {
	suite.service = NewViaCepService(mock)
}

func TestStartSuit(t *testing.T) {
	suite.Run(t, new(CepTestSuit))
}
