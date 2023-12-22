package router

import "github.com/gofiber/fiber/v2"

func HealthRouter(api fiber.Router) {
	route := api.Group("/health")

	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "Service Running Well",
		})
	})
}
