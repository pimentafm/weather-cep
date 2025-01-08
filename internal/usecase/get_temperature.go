package usecase

import (
	"github.com/pimentafm/weatherapi/internal/domain/entity"
	"github.com/pimentafm/weatherapi/internal/domain/repository"
)

type GetTemperatureUseCase struct {
	temperatureRepo repository.TemperatureRepository
}

func NewGetTemperatureUseCase(repo repository.TemperatureRepository) *GetTemperatureUseCase {
	return &GetTemperatureUseCase{
		temperatureRepo: repo,
	}
}

func (uc *GetTemperatureUseCase) Execute(cep string) (*entity.Temperature, error) {
	city, err := uc.temperatureRepo.GetCityByCEP(cep)
	if err != nil {
		return nil, err
	}

	tempCelsius, err := uc.temperatureRepo.GetTemperatureByCity(city)
	if err != nil {
		return nil, err
	}

	return entity.NewTemperature(tempCelsius), nil
}
