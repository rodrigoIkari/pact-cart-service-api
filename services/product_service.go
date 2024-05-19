package services

import (
	"github.com/rodrigoikari/pact-cart-service-api/services/models"
)

type ProductService interface {
	GetProductBySKU(sku string) (models.ProductResponse, error)
}
