package repository

type TemperatureRepository interface {
	GetCityByCEP(cep string) (string, error)
	GetTemperatureByCity(city string) (float64, error)
}
