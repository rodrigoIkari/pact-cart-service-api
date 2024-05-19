package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID           uint
	Buyer_Asset_Code string
	Total_Amount     float64
	Items            []CartItem
}

type CartItem struct {
	gorm.Model
	SKU      string
	Quantity int
}
