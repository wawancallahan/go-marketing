package router

import "github.com/gofiber/fiber/v2"

func RegisterRoute(api fiber.Router) {
	apiV1 := fiber.New()

	api.Mount("/v1", apiV1)
	HealthRouter(api)
}
