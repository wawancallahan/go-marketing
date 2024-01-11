package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func WebhookRouter(api fiber.Router, controller controller.WebhookController) {
	route := api.Group("/webhook")

	route.Post("/marketing/events/status", controller.MarketingEventStatus)
	route.Post("/marketing/leads/activation-status", controller.MarketingLeadActivationStatus)
	route.Post("/marketing/leads/duplicate-status", controller.MarketingLeadDuplicateStatus)
}
