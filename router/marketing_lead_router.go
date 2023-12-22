package router

import (
	"github.com/gofiber/fiber/v2"
)

func MarketingLeadRouter(api fiber.Router) {
	route := api.Group("/marketing-lead")

	route.Get("/")
	route.Post("/create")
	route.Get("/:id")
	route.Put("/update/:id")
	route.Delete("/delete/:id")
}
