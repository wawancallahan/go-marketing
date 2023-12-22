package router

import "github.com/gofiber/fiber/v2"

func MarketingEventRouter(api fiber.Router) {
	route := api.Group("/marketing-event")

	route.Get("/")
	route.Post("/create")
	route.Get("/:id")
	route.Put("/update/:id")
	route.Delete("/delete/:id")
}
