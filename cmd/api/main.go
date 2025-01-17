package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pimentafm/weatherapi/configs"
	"github.com/pimentafm/weatherapi/internal/infrastructure/api"
	"github.com/pimentafm/weatherapi/internal/infrastructure/handlers"
	"github.com/pimentafm/weatherapi/internal/usecase"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	viaCEPAPI := api.NewCityAPI()
	weatherAPI := api.NewWeatherAPI(cfg.WeatherAPIKey)

	cityRepo := viaCEPAPI
	temperatureRepo := weatherAPI

	getCityUseCase := usecase.NewGetCityUseCase(cityRepo)
	getTemperatureUseCase := usecase.NewGetTemperatureUseCase(cityRepo, temperatureRepo)

	cityHandler := handlers.NewCityHandler(getCityUseCase)
	temperatureHandler := handlers.NewTemperatureHandler(getTemperatureUseCase)

	// Setup routes
	http.HandleFunc("/city/", cityHandler.GetCity)
	http.HandleFunc("/temperature/", temperatureHandler.GetTemperature)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
