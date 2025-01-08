package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pimentafm/weather-cep/internal/infrastructure/api"
	handlers "github.com/pimentafm/weather-cep/internal/infrastructure/http"
	"github.com/pimentafm/weather-cep/internal/usecase"
)

type TemperatureRepository struct {
	viaCEPAPI  *api.ViaCEPAPI
	weatherAPI *api.WeatherAPI
}

func (r *TemperatureRepository) GetCityByCEP(cep string) (string, error) {
	return r.viaCEPAPI.GetCity(cep)
}

func (r *TemperatureRepository) GetTemperatureByCity(city string) (float64, error) {
	return r.weatherAPI.GetTemperature(city)
}

func main() {
	// Initialize dependencies

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	viaCEPAPI := api.NewViaCEPAPI()
	weatherAPI := api.NewWeatherAPI(os.Getenv("WEATHER_API_KEY"))

	temperatureRepo := &TemperatureRepository{
		viaCEPAPI:  viaCEPAPI,
		weatherAPI: weatherAPI,
	}

	getTemperatureUseCase := usecase.NewGetTemperatureUseCase(temperatureRepo)
	temperatureHandler := handlers.NewTemperatureHandler(getTemperatureUseCase)

	// Setup routes
	http.HandleFunc("/temperature/", temperatureHandler.GetTemperature)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
