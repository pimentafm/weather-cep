package usecase

import (
	"github.com/pimentafm/weatherapi/internal/domain/entity"
	"github.com/pimentafm/weatherapi/internal/domain/repository"
)

type GetTemperatureUseCase struct {
	CityRepo        repository.CityRepository
	TemperatureRepo repository.TemperatureRepository
}

func NewGetTemperatureUseCase(cityRepo repository.CityRepository, tempRepo repository.TemperatureRepository) *GetTemperatureUseCase {
	return &GetTemperatureUseCase{
		CityRepo:        cityRepo,
		TemperatureRepo: tempRepo,
	}
}

func (uc *GetTemperatureUseCase) Execute(cep string) (*entity.Temperature, error) {
	city, err := uc.CityRepo.GetCityByCEP(cep)
	if err != nil {
		return nil, err
	}

	tempCelsius, err := uc.TemperatureRepo.GetTemperatureByCity(city.Localidade)
	if err != nil {
		return nil, err
	}

	return entity.NewTemperature(tempCelsius), nil
}
