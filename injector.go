//go:build wireinject
// +build wireinject

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

var MarketingEventSet = wire.NewSet(
	repository.NewMarketingEventRepository,
	service.NewMarketingEventService,
	controller.NewMarketingEventController,
	wire.Bind(new(repository.MarketingEventRepository), new(*repository.MarketingEventRepositoryImpl)),
	wire.Bind(new(service.MarketingEventService), new(*service.MarketingEventServiceImpl)),
	wire.Bind(new(controller.MarketingEventController), new(*controller.MarketingEventControllerImpl)),
)

var MarketingLeadSet = wire.NewSet(
	repository.NewMarketingLeadRepository,
	service.NewMarketingLeadService,
	controller.NewMarketingLeadController,
	wire.Bind(new(repository.MarketingLeadRepository), new(*repository.MarketingLeadRepositoryImpl)),
	wire.Bind(new(service.MarketingLeadService), new(*service.MarketingLeadServiceImpl)),
	wire.Bind(new(controller.MarketingLeadController), new(*controller.MarketingLeadControllerImpl)),
)

var BlogCategorySet = wire.NewSet(
	repository.NewBlogCategoryRepository,
	service.NewBlogCategoryService,
	controller.NewBlogCategoryController,
	wire.Bind(new(repository.BlogCategoryRepository), new(*repository.BlogCategoryRepositoryImpl)),
	wire.Bind(new(service.BlogCategoryService), new(*service.BlogCategoryServiceImpl)),
	wire.Bind(new(controller.BlogCategoryController), new(*controller.BlogCategoryControllerImpl)),
)

var BlogBannerSet = wire.NewSet(
	repository.NewBlogBannerRepository,
	service.NewBlogBannerService,
	controller.NewBlogBannerController,
	wire.Bind(new(repository.BlogBannerRepository), new(*repository.BlogBannerRepositoryImpl)),
	wire.Bind(new(service.BlogBannerService), new(*service.BlogBannerServiceImpl)),
	wire.Bind(new(controller.BlogBannerController), new(*controller.BlogBannerControllerImpl)),
)

var BlogArticleSet = wire.NewSet(
	repository.NewBlogArticleRepository,
	service.NewBlogArticleService,
	controller.NewBlogArticleController,
	wire.Bind(new(repository.BlogArticleRepository), new(*repository.BlogArticleRepositoryImpl)),
	wire.Bind(new(service.BlogArticleService), new(*service.BlogArticleServiceImpl)),
	wire.Bind(new(controller.BlogArticleController), new(*controller.BlogArticleControllerImpl)),
)

var BlogArticleAttachmentSet = wire.NewSet(
	repository.NewBlogArticleAttachmentRepository,
	wire.Bind(new(repository.BlogArticleAttachmentRepository), new(*repository.BlogArticleAttachmentRepositoryImpl)),
)

var WebhookSet = wire.NewSet(
	service.NewWebhookService,
	controller.NewWebhookController,
	wire.Bind(new(service.WebhookService), new(*service.WebhookServiceImpl)),
	wire.Bind(new(controller.WebhookController), new(*controller.WebhookControllerImpl)),
)

func InitializedServer() *App {
	wire.Build(
		config.New,
		database.New,
		NewApp,
	)

	return nil
}

func InitializedRouter(Db *database.Database, Config *config.Config) *fiber.App {
	wire.Build(
		MarketingEventSet,
		MarketingLeadSet,
		BlogCategorySet,
		BlogBannerSet,
		BlogArticleAttachmentSet,
		BlogArticleSet,
		WebhookSet,
		router.New,
	)

	return nil
}
