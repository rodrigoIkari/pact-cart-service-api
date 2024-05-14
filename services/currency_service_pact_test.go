package services_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/rodrigoikari/pact-cart-service-api/models"
	"github.com/rodrigoikari/pact-cart-service-api/services"
)

func TestMain(m *testing.M) {
	var exitCode int

	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	exitCode = m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

var commonHeaders = dsl.MapMatcher{
	"Content-Type": term("application/json; charset=utf-8", `application\/json`),
}

// Common test data
var pact dsl.Pact

// Aliases
var term = dsl.Term

type request = dsl.Request

var u *url.URL
var currency_service *services.CurrencyServiceImpl

func setup() {
	pact = createPact()

	// Proactively start service to get access to the port
	pact.Setup(true)

	u, _ = url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))

	currency_service = &services.CurrencyServiceImpl{
		BaseUrl: u,
	}

}

func TestCurrencyServicePact_ConvertCurrencyValue_Successfull(t *testing.T) {
	t.Run("convert a value from original to destination currency succesfully", func(t *testing.T) {

		sourceCurrencyCode := "USD"
		destinationCurrencyCode := "BRL"
		value := 15.0
		convertedValue := 80.1
		exchangeRate := 5.34

		pact.
			AddInteraction().
			Given("USD to BRL exchange rate is updated to 5.34").
			UponReceiving("A Request to convert USD to BRL").
			WithRequest(request{
				Method: "POST",
				Path:   term("/exchange/convert", "/exchange/convert"),
				Body: models.ConvertCurrencyValueRequest{
					SourceCurrencyCode:      sourceCurrencyCode,
					DestinationCurrencyCode: destinationCurrencyCode,
					Value:                   value,
				},
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body: dsl.Like(models.ConvertCurrencyValueResponse{
					SourceCurrencyCode:      sourceCurrencyCode,
					DestinationCurrencyCode: destinationCurrencyCode,
					Value:                   value,
					ConvertedValue:          convertedValue,
					ExchangeRate:            exchangeRate,
				}),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			cv, err := currency_service.ConvertCurrencyValue(value, sourceCurrencyCode, destinationCurrencyCode)

			if err != nil {
				return err
			}
			// Assert basic fact
			if cv != convertedValue {
				return fmt.Errorf("wanted converted value %v but got %v", convertedValue, cv)
			}
			return nil
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}

func TestCurrencyServicePact_ConvertCurrencyValue_Error_TaxRateNotFound(t *testing.T) {

	var response = struct {
		Message string
	}{
		Message: "Currency Tax Rate not found for conversion",
	}

	t.Run("convert a value from original to destination currency when tax rate is not found", func(t *testing.T) {

		sourceCurrencyCode := "USD"
		destinationCurrencyCode := "XXX"
		value := 15.0

		pact.
			AddInteraction().
			Given("USD to XXX exchange rate is not found").
			UponReceiving("A Request to convert USD to XXX").
			WithRequest(request{
				Method: "POST",
				Path:   term("/exchange/convert", "/exchange/convert"),
				Body: models.ConvertCurrencyValueRequest{
					SourceCurrencyCode:      sourceCurrencyCode,
					DestinationCurrencyCode: destinationCurrencyCode,
					Value:                   value,
				},
			}).
			WillRespondWith(dsl.Response{
				Status:  422,
				Body:    dsl.Like(response),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			cv, err := currency_service.ConvertCurrencyValue(value, sourceCurrencyCode, destinationCurrencyCode)

			if err == nil {
				return fmt.Errorf("expected error != nil")
			}
			// Assert basic fact
			if cv != 0.0 {
				return fmt.Errorf("wanted converted value 0.0 but got %v", cv)
			}
			return nil
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer: "pact-cart-service-api",
		Provider: "pact-exchange-service-api",
		//LogDir:   os.Getenv("LOG_DIR"),
		//PactDir:  os.Getenv("PACT_DIR"),
		LogLevel: "INFO",
	}
}
