package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/rodrigoikari/pact-cart-service-api/controllers"
)

func ConfigRoutes(api fiber.Router) {
	cartApi := api.Group("/cart/")

	ctl := new(controllers.CartController)

	cartApi.Post("simulation", ctl.SimulateCart).Name("Cart Simulation")
	cartApi.Post("checkout", ctl.Checkout).Name("Cart Checkout")
}
