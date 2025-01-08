package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/pimentafm/weatherapi/internal/usecase"
	"github.com/pimentafm/weatherapi/pkg/cerrors"
)

type TemperatureHandler struct {
	getTemperatureUseCase *usecase.GetTemperatureUseCase
}

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewTemperatureHandler(useCase *usecase.GetTemperatureUseCase) *TemperatureHandler {
	return &TemperatureHandler{
		getTemperatureUseCase: useCase,
	}
}

func (h *TemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cep := r.URL.Path[len("/temperature/"):]

	// Validate CEP format
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	if !match {
		h.respondWithError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	temp, err := h.getTemperatureUseCase.Execute(cep)
	if err != nil {
		switch err {
		case cerrors.ErrCEPNotFound:
			h.respondWithError(w, http.StatusNotFound, "can not find zipcode")
		default:
			h.respondWithError(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response := TemperatureResponse{
		TempC: temp.Celsius,
		TempF: temp.Fahrenheit,
		TempK: temp.Kelvin,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TemperatureHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
