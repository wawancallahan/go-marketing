package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func MasterRouter(api fiber.Router, controller controller.MasterController) {
	router := api.Group("/master")

	router.Get("/marketing/events/event-type", controller.MarketingEventEventType)
	router.Get("/marketing/events/channel-event", controller.MarketingEventChannelEvent)
	router.Get("/marketing/events/measurement-event", controller.MarketingEventMeasurementEvent)
	router.Get("/marketing/leads/source-type", controller.MarketingLeadSourceType)
	router.Get("/marketing/leads/product-category", controller.MarketingLeadProductCategory)
}
