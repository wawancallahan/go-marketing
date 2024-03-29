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
	blogArticleController controller.BlogArticleController,
	webhookController controller.WebhookController,
	masterController controller.MasterController,
) *fiber.App {
	api := fiber.New()

	MarketingEventRouter(api, marketingEventController)
	MarketingLeadRouter(api, marketingLeadController)
	BlogCategoryRouter(api, blogCategoryController)
	BlogBannerRouter(api, blogBannerController)
	BlogArticleRouter(api, blogArticleController)
	WebhookRouter(api, webhookController)
	MasterRouter(api, masterController)

	return api
}
