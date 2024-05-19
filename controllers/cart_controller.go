package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rodrigoikari/pact-cart-service-api/controllers/models"
	"github.com/rodrigoikari/pact-cart-service-api/services"
)

type CartController struct {
	currencyService services.CurrencyService
	productService  services.ProductService
}

func (ctl *CartController) SimulateCart(c *fiber.Ctx) error {

	cart := new(models.CartRequest)

	if err := c.BodyParser(cart); err != nil {
		fmt.Println("error parsing cart: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Checkout Simulation: Bad Request")
	}

	errors := services.Validate(cart)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	fmt.Println("cart request accepted")

	fmt.Println("Calculating Cart Value ...")
	totalAmount, err := ctl.CalculateCartValue(cart)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	fmt.Println("Cart Value: ", totalAmount, cart.Buyer_Asset_Code)

	cartResponse := new(models.CartResponse)
	cartResponse.Items = cart.Items
	cartResponse.Buyer_Asset_Code = cart.Buyer_Asset_Code
	cartResponse.Total_Amount = totalAmount

	return c.JSON(cartResponse)
}

func (ctl *CartController) Checkout(c *fiber.Ctx) error {
	cart := new(models.CartRequest)

	if err := c.BodyParser(cart); err != nil {
		fmt.Println("error parsing cart: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Checkout: Bad Request")
	}

	errors := services.Validate(cart)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	fmt.Println("cart request accepted")

	fmt.Println("Calculating Cart Value ...")
	totalAmount, err := ctl.CalculateCartValue(cart)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}
	fmt.Println("Cart Value: ", totalAmount, cart.Buyer_Asset_Code)

	return nil

}

func (ctl *CartController) CalculateCartValue(cart *models.CartRequest) (float64, error) {

	totalAmount := 0.0

	for _, it := range cart.Items {

		product, err := ctl.productService.GetProductBySKU(it.SKU)
		if err != nil {

			convertedValue, err := ctl.currencyService.ConvertCurrencyValue(product.Price*float64(it.Quantity), product.AssetCode, cart.Buyer_Asset_Code)
			if err == nil {
				return 0, err
			}
			totalAmount += convertedValue
		}
	}

	return totalAmount, nil
}
