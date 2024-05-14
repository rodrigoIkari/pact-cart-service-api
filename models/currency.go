package models

type ConvertCurrencyValueRequest struct {
	Value                   float64 `json:"value"`
	SourceCurrencyCode      string  `json:"source_currency_code"`
	DestinationCurrencyCode string  `json:"destination_currency_code"`
}

type ConvertCurrencyValueResponse struct {
	Value                   float64 `json:"value"`
	ConvertedValue          float64 `json:"converted_value"`
	SourceCurrencyCode      string  `json:"source_currency_code"`
	DestinationCurrencyCode string  `json:"destination_currency_code"`
	ExchangeRate            float64 `json:"exchange_rate"`
}
