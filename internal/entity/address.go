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

func (address *Address) SetAddressFromMap(payload map[string]any) error {
	address.Bairro = payload["bairro"].(string)
	address.Cep = payload["cep"].(string)
	address.Complemento = payload["complemento"].(string)
	address.Ddd = payload["ddd"].(string)
	address.Gia = payload["gia"].(string)
	address.Ibge = payload["ibge"].(string)
	address.Logradouro = payload["logradouro"].(string)
	address.Localidade = payload["localidade"].(string)
	address.Siafi = payload["siafi"].(string)
	address.Uf = payload["uf"].(string)
	formattedCep, err := address.FormattCep()
	if err != nil {
		return err
	}
	address.Cep = formattedCep
	return nil
}
