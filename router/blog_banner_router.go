package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func BlogBannerRouter(api fiber.Router, controller controller.BlogBannerController) {
	route := api.Group("/blog/banner")

	route.Get("/", controller.Index)
	route.Put("/update/:id", controller.Update)
}
