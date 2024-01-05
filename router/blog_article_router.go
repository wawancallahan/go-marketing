package router

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/controller"
)

func BlogArticleRouter(api fiber.Router, controller controller.BlogArticleController) {
	route := api.Group("/blog/article")

	route.Get("/", controller.Index)
	route.Post("/create", controller.Create)
	route.Get("/:id", controller.Find)
	route.Put("/update/:id", controller.Update)
	route.Delete("/delete/:id", controller.Delete)
}
