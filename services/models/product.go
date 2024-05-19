package models

type ProductResponse struct {
	SKU       string  `json:"sku"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	AssetCode string  `json:"asset_code"`
}
