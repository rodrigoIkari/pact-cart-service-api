package services

import (
	"github.com/rodrigoikari/pact-cart-service-api/models"
)

type ProductService interface {
	GetProductBySKU(sku string) (models.Product, error)
}
