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
	wire.Bind(new(repository.MarketingEventRepository), new(*repository.MarketingEventRepositoryImpl)),
	service.NewMarketingEventService,
	wire.Bind(new(service.MarketingEventService), new(*service.MarketingEventServiceImpl)),
	controller.NewMarketingEventController,
	wire.Bind(new(controller.MarketingEventController), new(*controller.MarketingEventControllerImpl)),
)

func InitializedServer() *App {
	wire.Build(
		config.New,
		database.New,
		NewApp,
	)

	return nil
}

func InitializedRouter(Db *database.Database) *fiber.App {
	wire.Build(
		config.New,
		MarketingEventSet,
		router.New,
	)

	return nil
}
