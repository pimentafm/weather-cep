package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pimentafm/weatherapi/internal/usecase"
)

type CityHandler struct {
	GetCityUseCase *usecase.GetCityUseCase
}

func NewCityHandler(getCityUseCase *usecase.GetCityUseCase) *CityHandler {
	return &CityHandler{
		GetCityUseCase: getCityUseCase,
	}
}

func (h *CityHandler) GetCity(w http.ResponseWriter, r *http.Request) {
	cep := strings.TrimPrefix(r.URL.Path, "/city/")
	city, err := h.GetCityUseCase.Execute(cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(city)
}
