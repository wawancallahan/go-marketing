package controller

import "github.com/gofiber/fiber/v2"

type MasterController interface {
	MarketingEventEventType(ctx *fiber.Ctx) error
	MarketingEventChannelEvent(ctx *fiber.Ctx) error
	MarketingEventMeasurementEvent(ctx *fiber.Ctx) error
	MarketingLeadSourceType(ctx *fiber.Ctx) error
	MarketingLeadProductCategory(ctx *fiber.Ctx) error
}

type MasterControllerImpl struct {
}

func NewMasterController() *MasterControllerImpl {
	return &MasterControllerImpl{}
}

func (c *MasterControllerImpl) MarketingEventEventType(ctx *fiber.Ctx) error {
	return nil
}

func (c *MasterControllerImpl) MarketingEventChannelEvent(ctx *fiber.Ctx) error {
	return nil
}

func (c *MasterControllerImpl) MarketingEventMeasurementEvent(ctx *fiber.Ctx) error {
	return nil
}

func (c *MasterControllerImpl) MarketingLeadSourceType(ctx *fiber.Ctx) error {
	return nil
}

func (c *MasterControllerImpl) MarketingLeadProductCategory(ctx *fiber.Ctx) error {
	return nil
}
