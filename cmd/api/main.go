package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ViaCEPResponse struct {
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
	Erro        bool   `json:"erro"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	city, err := getCityFromCEP("35780000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(city)

	temperature, err := getCityTemperature(city)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(temperature)
}

func getCityFromCEP(cep string) (string, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var viaCEPResp ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return "", err
	}

	if viaCEPResp.Erro {
		return "", fmt.Errorf("CEP not found")
	}

	return viaCEPResp.Localidade, nil
}

func getCityTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHERAPI_API_KEY")
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=%s", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	var weatherResponse WeatherAPIResponse
	if err := json.NewDecoder(res.Body).Decode(&weatherResponse); err != nil {
		return 0, err
	}

	return weatherResponse.Current.TempC, nil
}
