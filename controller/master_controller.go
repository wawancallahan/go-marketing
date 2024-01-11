package controller

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
)

type MasterController interface {
	MarketingEventEventType(ctx *fiber.Ctx) error
	MarketingEventChannelEvent(ctx *fiber.Ctx) error
	MarketingEventMeasurementEvent(ctx *fiber.Ctx) error
	MarketingLeadSourceType(ctx *fiber.Ctx) error
	MarketingLeadProductCategory(ctx *fiber.Ctx) error
}

type MasterControllerImpl struct {
	MarketingEventService service.MarketingEventService
	MarketingLeadService  service.MarketingLeadService
}

func NewMasterController(MarketingEventService service.MarketingEventService, MarketingLeadService service.MarketingLeadService) *MasterControllerImpl {
	return &MasterControllerImpl{MarketingEventService: MarketingEventService, MarketingLeadService: MarketingLeadService}
}

func (c *MasterControllerImpl) MarketingEventEventType(ctx *fiber.Ctx) error {
	items, _ := c.MarketingEventService.EventType()

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   items,
	})
}

func (c *MasterControllerImpl) MarketingEventChannelEvent(ctx *fiber.Ctx) error {
	items, _ := c.MarketingEventService.ChannelEvent()

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   items,
	})
}

func (c *MasterControllerImpl) MarketingEventMeasurementEvent(ctx *fiber.Ctx) error {
	items, _ := c.MarketingEventService.MeasurementEvent()

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   items,
	})
}

func (c *MasterControllerImpl) MarketingLeadSourceType(ctx *fiber.Ctx) error {
	items, _ := c.MarketingLeadService.SourceType()

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   items,
	})
}

func (c *MasterControllerImpl) MarketingLeadProductCategory(ctx *fiber.Ctx) error {
	items, _ := c.MarketingLeadService.ProductCategory()

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   items,
	})
}
