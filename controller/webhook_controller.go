package controller

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
)

type WebhookController interface {
	MarketingEventStatus(ctx *fiber.Ctx) error
	MarketingLeadActivationStatus(ctx *fiber.Ctx) error
	MarketingLeadDuplicateStatus(ctx *fiber.Ctx) error
}

type WebhookControllerImpl struct {
	WebhookService service.WebhookService
}

func NewWebhookController(WebhookService service.WebhookService) *WebhookControllerImpl {
	return &WebhookControllerImpl{WebhookService: WebhookService}
}

func (c *WebhookControllerImpl) MarketingEventStatus(ctx *fiber.Ctx) error {
	_, err := c.WebhookService.MarketingEventStatus()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (c *WebhookControllerImpl) MarketingLeadActivationStatus(ctx *fiber.Ctx) error {
	_, err := c.WebhookService.MarketingLeadActivationStatus()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}

func (c *WebhookControllerImpl) MarketingLeadDuplicateStatus(ctx *fiber.Ctx) error {
	_, err := c.WebhookService.MarketingLeadDuplicateStatus()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
