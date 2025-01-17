package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pimentafm/weatherapi/internal/domain/entity"
)

type CityAPI struct{}

func NewCityAPI() *CityAPI {
	return &CityAPI{}
}

func (c *CityAPI) GetCityByCEP(cep string) (*entity.City, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid zipcode")
	}

	var cityResp entity.CityResponse
	if err := json.NewDecoder(resp.Body).Decode(&cityResp); err != nil {
		return nil, err
	}

	if cityResp.Erro == "true" {
		return nil, fmt.Errorf("can not find zipcode")
	}

	return entity.NewCity(cityResp), nil
}
