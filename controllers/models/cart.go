package models

type CartItem struct {
	SKU      string `json:"sku" validate:"required,min=3"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

type CartRequest struct {
	Items            []*CartItem `json:"items" validate:"required,min=1,dive"`
	Buyer_Asset_Code string      `json:"buyer-asset-code"  validate:"required,iso4217"`
}

type CartResponse struct {
	Items            []*CartItem `json:"items" validate:"required,min=1,dive"`
	Buyer_Asset_Code string      `json:"buyer-asset-code"  validate:"required,iso4217"`
	Total_Amount     float64     `json:"total-amount"`
}
