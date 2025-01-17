package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pimentafm/weatherapi/internal/domain/entity"
	"github.com/pimentafm/weatherapi/internal/usecase"
	"github.com/pimentafm/weatherapi/pkg/cerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWeatherAPI struct {
	mock.Mock
}

func (m *MockWeatherAPI) GetTemperatureByCity(city string) (float64, error) {
	args := m.Called(city)
	return args.Get(0).(float64), args.Error(1)
}

type MockCityRepository struct {
	mock.Mock
}

func (m *MockCityRepository) GetCityByCEP(cep string) (*entity.City, error) {
	args := m.Called(cep)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.City), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetTemperature(t *testing.T) {
	mockAPI := &MockWeatherAPI{}
	mockCityRepo := &MockCityRepository{}
	useCase := &usecase.GetTemperatureUseCase{CityRepo: mockCityRepo, TemperatureRepo: mockAPI}
	handler := NewTemperatureHandler(useCase)

	tests := []struct {
		name           string
		method         string
		cep            string
		mockCityReturn *entity.City
		mockTempReturn float64
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:           "Successful request",
			method:         http.MethodGet,
			cep:            "35780000",
			mockCityReturn: &entity.City{Localidade: "Cordisburgo"},
			mockTempReturn: 25.0,
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"temp_C": 25.0,
				"temp_F": 77.0,
				"temp_K": 298.15,
			},
		},
		{
			name:           "CEP not found",
			method:         http.MethodGet,
			cep:            "35780300",
			mockError:      cerrors.ErrCEPNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   map[string]interface{}{"message": "can not find zipcode"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockCityReturn != nil || tt.mockError != nil {
				mockCityRepo.On("GetCityByCEP", tt.cep).Return(tt.mockCityReturn, tt.mockError)
			}
			if tt.mockCityReturn != nil && (tt.mockTempReturn != 0 || tt.mockError != nil) {
				mockAPI.On("GetTemperatureByCity", tt.mockCityReturn.Localidade).Return(tt.mockTempReturn, tt.mockError)
			}

			req := httptest.NewRequest(tt.method, "/temperature/"+tt.cep, nil)
			rr := httptest.NewRecorder()

			handler.GetTemperature(rr, req)

			log.Printf("Test: %s, CEP: %s, Status: %d, Body: %s", tt.name, tt.cep, rr.Code, rr.Body.String())

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var responseBody map[string]interface{}
			if err := json.NewDecoder(rr.Body).Decode(&responseBody); err == nil {
				assert.Equal(t, tt.expectedBody, responseBody)
			} else {
				assert.Nil(t, responseBody)
			}

			mockCityRepo.AssertExpectations(t)
			mockAPI.AssertExpectations(t)
		})
	}
}
