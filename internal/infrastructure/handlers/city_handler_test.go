package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pimentafm/weatherapi/internal/domain/entity"
	"github.com/pimentafm/weatherapi/internal/usecase"
	"github.com/stretchr/testify/assert"
)

type MockCityAPI struct{}

func (m *MockCityAPI) GetCityByCEP(cep string) (*entity.City, error) {
	if cep == "35780000" {
		return &entity.City{Localidade: "Cordisburgo"}, nil
	}
	return nil, errors.New("some error")
}

func TestCityHandler_GetCity_Success(t *testing.T) {
	mockAPI := &MockCityAPI{}
	useCase := &usecase.GetCityUseCase{CityRepo: mockAPI}
	handler := NewCityHandler(useCase)

	expectedCity := map[string]string{"Localidade": "Cordisburgo"}

	req, err := http.NewRequest("GET", "/city/35780000", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetCity(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, expectedCity["Localidade"], response["Localidade"])
}

func TestCityHandler_GetCity_Error(t *testing.T) {
	mockAPI := &MockCityAPI{}
	useCase := &usecase.GetCityUseCase{CityRepo: mockAPI}
	handler := NewCityHandler(useCase)

	req, err := http.NewRequest("GET", "/city/12345678", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetCity(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, "some error\n", rr.Body.String())
}
