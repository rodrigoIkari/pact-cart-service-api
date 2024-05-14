package services_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/rodrigoikari/pact-cart-service-api/models"
	"github.com/rodrigoikari/pact-cart-service-api/services"
	"github.com/stretchr/testify/assert"
)

func Test_Currency_Service_ConvertCurrencyValue_Succesfully(t *testing.T) {
	value := 15.0
	sourceCurrencyCode := "USD"
	destinationCurrencyCode := "BRL"
	exchangeRate := 5.34
	convertedValue := 80.1 // 15.0 * 5.34

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, req.URL.String(), "/exchange/convert")

			converted_amount_response, _ := json.Marshal(models.ConvertCurrencyValueResponse{
				Value:                   value,
				ConvertedValue:          convertedValue,
				SourceCurrencyCode:      sourceCurrencyCode,
				DestinationCurrencyCode: destinationCurrencyCode,
				ExchangeRate:            exchangeRate,
			})

			rw.Write([]byte(converted_amount_response))
		}))

	defer server.Close()

	u, _ := url.Parse(server.URL)

	currency_service := services.CurrencyServiceImpl{
		BaseUrl: u,
	}

	ca, err := currency_service.ConvertCurrencyValue(value, sourceCurrencyCode, destinationCurrencyCode)
	assert.NoError(t, err)
	assert.Equal(t, ca, convertedValue)

}

func Test_Currency_Service_ConvertCurrencyValue_Error_RateNotFound(t *testing.T) {
	value := 15.0
	sourceCurrencyCode := "USD"
	destinationCurrencyCode := "ILS"

	server := httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			assert.Equal(t, req.URL.String(), "/exchange/convert")

			rw.WriteHeader(http.StatusUnprocessableEntity)

		}))

	defer server.Close()

	u, _ := url.Parse(server.URL)

	currency_service := services.CurrencyServiceImpl{
		BaseUrl: u,
	}

	_, err := currency_service.ConvertCurrencyValue(value, sourceCurrencyCode, destinationCurrencyCode)
	assert.Error(t, err)

}
