package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherAPI struct {
	apiKey string
}

type weatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func NewWeatherAPI(apiKey string) *WeatherAPI {
	return &WeatherAPI{
		apiKey: apiKey,
	}
}

func (w *WeatherAPI) GetTemperature(city string) (float64, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?q=%s&key=%s", city, w.apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	var weatherResp weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return 0, err
	}

	return weatherResp.Current.TempC, nil
}
