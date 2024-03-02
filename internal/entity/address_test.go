package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FormattCep(t *testing.T) {
	Cep := "01001-000"
	address := Address{Cep: Cep}
	assert.Equal(t, address.Cep, Cep)

	formattedCep, err := address.FormattCep()
	assert.NoError(t, err)
	assert.NotEqual(t, formattedCep, Cep)
}

func Test_IsCepValid(t *testing.T) {
	validCeps := []string{"01001-000", "01001000", "72917-160", "72917160"}
	invalidCeps := []string{"01001-00", "010add00", "aaaaaaaa", "729"}

	for _, cep := range validCeps {
		assert.True(t, Address{Cep: cep}.IsCepValid())
	}

	for _, cep := range invalidCeps {
		assert.False(t, Address{Cep: cep}.IsCepValid())
	}
}

func Test_SetAddressFromMap(t *testing.T) {
	payload := map[string]string{
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
	address := Address{}
	assert.Empty(t, address.Cep)
	assert.Empty(t, address.Bairro)
	assert.Empty(t, address.Logradouro)
	assert.False(t, address.IsCepValid())

	address.SetAddressFromMap(payload)
	assert.NotEmpty(t, address.Cep)
	assert.Len(t, address.Cep, 8)
	assert.True(t, address.IsCepValid())
	assert.Equal(t, address.Cep, "01001000")

	assert.NotEmpty(t, address.Bairro)
	assert.Equal(t, address.Bairro, payload["bairro"])
	assert.NotEmpty(t, address.Logradouro)
	assert.Equal(t, address.Logradouro, payload["logradouro"])
}
