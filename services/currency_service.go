package services

type CurrencyService interface {
	ConvertCurrencyValue(value float64, sourceCurrencyCode string, destinationCurrencyCode string) (float64, error)
}
