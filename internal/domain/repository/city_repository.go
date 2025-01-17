package repository

import "github.com/pimentafm/weatherapi/internal/domain/entity"

type CityRepository interface {
	GetCityByCEP(cep string) (*entity.City, error)
}
