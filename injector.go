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
		MarketingEventSet,
		MarketingLeadSet,
		router.New,
	)

	return nil
}
