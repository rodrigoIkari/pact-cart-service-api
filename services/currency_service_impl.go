package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rodrigoikari/pact-cart-service-api/services/models"
)

type CurrencyServiceImpl struct {
	BaseUrl *url.URL
}

const exchange_convert_url string = "/exchange/convert"

func (c *CurrencyServiceImpl) ConvertCurrencyValue(value float64, sourceCurrencyCode string, destinationCurrencyCode string) (float64, error) {

	// Monta o Request
	req, _ := json.Marshal(models.ConvertCurrencyValueRequest{
		Value:                   value,
		SourceCurrencyCode:      sourceCurrencyCode,
		DestinationCurrencyCode: destinationCurrencyCode,
	})
	reqBuffer := bytes.NewBuffer(req)

	//Efetua o POST na API
	rel := &url.URL{Path: exchange_convert_url}
	u := c.BaseUrl.ResolveReference(rel)
	resp, err := http.DefaultClient.Post(u.String(), "application/json", reqBuffer)
	if err != nil {
		return 0.0, err
	}

	// Interpreta o Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0.0, fmt.Errorf("conversion error: HttpStatusCode %v", resp.StatusCode)
	}

	var currencyResponse models.ConvertCurrencyValueResponse
	err = json.Unmarshal(body, &currencyResponse)
	if err != nil {
		return 0.0, err
	}

	return currencyResponse.ConvertedValue, nil

}
