package usecase

import (
	"github.com/pimentafm/weatherapi/internal/domain/entity"
	"github.com/pimentafm/weatherapi/internal/domain/repository"
)

type GetCityUseCase struct {
	CityRepo repository.CityRepository
}

func NewGetCityUseCase(repo repository.CityRepository) *GetCityUseCase {
	return &GetCityUseCase{
		CityRepo: repo,
	}
}

func (uc *GetCityUseCase) Execute(cep string) (*entity.City, error) {
	return uc.CityRepo.GetCityByCEP(cep)
}
