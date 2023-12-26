package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func New(
	marketingEventController controller.MarketingEventController,
	marketingLeadController controller.MarketingLeadController,
) *fiber.App {
	api := fiber.New()

	MarketingEventRouter(api, marketingEventController)
	MarketingLeadRouter(api, marketingLeadController)

	return api
}
