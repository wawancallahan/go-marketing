package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func New(
	marketingEventController controller.MarketingEventController,
	marketingLeadController controller.MarketingLeadController,
	blogCategoryController controller.BlogCategoryController,
	blogBannerController controller.BlogBannerController,
) *fiber.App {
	api := fiber.New()

	MarketingEventRouter(api, marketingEventController)
	MarketingLeadRouter(api, marketingLeadController)
	BlogCategoryRouter(api, blogCategoryController)
	BlogBannerRouter(api, blogBannerController)

	return api
}
