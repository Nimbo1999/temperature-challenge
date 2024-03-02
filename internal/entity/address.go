package entity

import (
	"regexp"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (address Address) FormattCep() (string, error) {
	reg, err := regexp.Compile("\\D")
	if err != nil {
		return "", err
	}
	return reg.ReplaceAllString(address.Cep, ""), nil
}

func (address Address) IsCepValid() bool {
	formattedCep, err := address.FormattCep()
	if err != nil {
		return false
	}
	return len(formattedCep) == 8
}

func (address *Address) SetAddressFromMap(payload map[string]string) error {
	address.Bairro = payload["bairro"]
	address.Cep = payload["cep"]
	address.Complemento = payload["complemento"]
	address.Ddd = payload["ddd"]
	address.Gia = payload["gia"]
	address.Ibge = payload["ibge"]
	address.Logradouro = payload["logradouro"]
	address.Localidade = payload["localidade"]
	address.Siafi = payload["siafi"]
	address.Uf = payload["uf"]
	formattedCep, err := address.FormattCep()
	if err != nil {
		return err
	}
	address.Cep = formattedCep
	return nil
}
