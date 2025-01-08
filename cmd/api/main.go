package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

const baseURL = "http://viacep.com.br/ws/"

func main() {
	city, err := getCityFromCEP("35780000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(city)
}

func getCityFromCEP(cep string) (string, error) {
	url := fmt.Sprintf("%s%s/json/", baseURL, cep)
	req, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", req.StatusCode)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var data Address
	err = json.Unmarshal(res, &data)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %v", err)
	}

	if data.Localidade == "" {
		return "", fmt.Errorf("localidade not found for CEP: %s", cep)
	}

	return data.Localidade, nil
}
