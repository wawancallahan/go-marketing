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

var MarketingLeadSet = wire.NewSet(
	repository.NewMarketingLeadRepository,
	wire.Bind(new(repository.MarketingLeadRepository), new(*repository.MarketingLeadRepositoryImpl)),
	service.NewMarketingLeadService,
	wire.Bind(new(service.MarketingLeadService), new(*service.MarketingLeadServiceImpl)),
	controller.NewMarketingLeadController,
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

func InitializedRouter() *fiber.App {
	wire.Build(
		config.New,
		database.New,
		MarketingEventSet,
		MarketingLeadSet,
		router.New,
	)

	return nil
}
