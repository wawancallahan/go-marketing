// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/controller"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/repository"
	"matsukana.cloud/go-marketing/router"
	"matsukana.cloud/go-marketing/service"
)

// Injectors from injector.go:

func InitializedServer() *App {
	configConfig := config.New()
	databaseDatabase := database.New(configConfig)
	app := NewApp(configConfig, databaseDatabase)
	return app
}

func InitializedRouter(Db *database.Database, Config *config.Config) *fiber.App {
	marketingEventRepositoryImpl := repository.NewMarketingEventRepository()
	marketingEventServiceImpl := service.NewMarketingEventService(Db, marketingEventRepositoryImpl)
	marketingEventControllerImpl := controller.NewMarketingEventController(marketingEventServiceImpl)
	marketingLeadRepositoryImpl := repository.NewMarketingLeadRepository()
	marketingLeadServiceImpl := service.NewMarketingLeadService(Db, marketingLeadRepositoryImpl)
	marketingLeadControllerImpl := controller.NewMarketingLeadController(marketingLeadServiceImpl)
	blogCategoryRepositoryImpl := repository.NewBlogCategoryRepository()
	blogCategoryServiceImpl := service.NewBlogCategoryService(Db, blogCategoryRepositoryImpl)
	blogCategoryControllerImpl := controller.NewBlogCategoryController(blogCategoryServiceImpl)
	blogBannerRepositoryImpl := repository.NewBlogBannerRepository()
	blogBannerServiceImpl := service.NewBlogBannerService(Db, Config, blogBannerRepositoryImpl)
	blogBannerControllerImpl := controller.NewBlogBannerController(blogBannerServiceImpl)
	blogArticleRepositoryImpl := repository.NewBlogArticleRepository()
	blogArticleServiceImpl := service.NewBlogArticleService(Db, blogArticleRepositoryImpl)
	blogArticleControllerImpl := controller.NewBlogArticleController(blogArticleServiceImpl)
	app := router.New(marketingEventControllerImpl, marketingLeadControllerImpl, blogCategoryControllerImpl, blogBannerControllerImpl, blogArticleControllerImpl)
	return app
}

// injector.go:

var MarketingEventSet = wire.NewSet(repository.NewMarketingEventRepository, service.NewMarketingEventService, controller.NewMarketingEventController, wire.Bind(new(repository.MarketingEventRepository), new(*repository.MarketingEventRepositoryImpl)), wire.Bind(new(service.MarketingEventService), new(*service.MarketingEventServiceImpl)), wire.Bind(new(controller.MarketingEventController), new(*controller.MarketingEventControllerImpl)))

var MarketingLeadSet = wire.NewSet(repository.NewMarketingLeadRepository, service.NewMarketingLeadService, controller.NewMarketingLeadController, wire.Bind(new(repository.MarketingLeadRepository), new(*repository.MarketingLeadRepositoryImpl)), wire.Bind(new(service.MarketingLeadService), new(*service.MarketingLeadServiceImpl)), wire.Bind(new(controller.MarketingLeadController), new(*controller.MarketingLeadControllerImpl)))

var BlogCategorySet = wire.NewSet(repository.NewBlogCategoryRepository, service.NewBlogCategoryService, controller.NewBlogCategoryController, wire.Bind(new(repository.BlogCategoryRepository), new(*repository.BlogCategoryRepositoryImpl)), wire.Bind(new(service.BlogCategoryService), new(*service.BlogCategoryServiceImpl)), wire.Bind(new(controller.BlogCategoryController), new(*controller.BlogCategoryControllerImpl)))

var BlogBannerSet = wire.NewSet(repository.NewBlogBannerRepository, service.NewBlogBannerService, controller.NewBlogBannerController, wire.Bind(new(repository.BlogBannerRepository), new(*repository.BlogBannerRepositoryImpl)), wire.Bind(new(service.BlogBannerService), new(*service.BlogBannerServiceImpl)), wire.Bind(new(controller.BlogBannerController), new(*controller.BlogBannerControllerImpl)))

var BlogArticleSet = wire.NewSet(repository.NewBlogArticleRepository, service.NewBlogArticleService, controller.NewBlogArticleController, wire.Bind(new(repository.BlogArticleRepository), new(*repository.BlogArticleRepositoryImpl)), wire.Bind(new(service.BlogArticleService), new(*service.BlogArticleServiceImpl)), wire.Bind(new(controller.BlogArticleController), new(*controller.BlogArticleControllerImpl)))
