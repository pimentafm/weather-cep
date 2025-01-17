package repository

type TemperatureRepository interface {
	GetTemperatureByCity(city string) (float64, error)
}
