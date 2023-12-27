package controller

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
	"matsukana.cloud/go-marketing/validation"
)

type MarketingEventController interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type MarketingEventControllerImpl struct {
	MarketingEventService service.MarketingEventService
}

func NewMarketingEventController(MarketingEventService service.MarketingEventService) *MarketingEventControllerImpl {
	return &MarketingEventControllerImpl{MarketingEventService: MarketingEventService}
}

func (c *MarketingEventControllerImpl) Index(ctx *fiber.Ctx) error {
	items, err := c.MarketingEventService.Index()

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
		Data:   &items,
	})
}

func (c *MarketingEventControllerImpl) Create(ctx *fiber.Ctx) error {
	var itemDTO dto.MarketingEventDTO

	if err := ctx.BodyParser(&itemDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}
	errs := validation.SetupValidation(itemDTO)

	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  errs,
		})
	}

	item, err := c.MarketingEventService.Create(&itemDTO)

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
		Data:   item,
	})
}

func (c *MarketingEventControllerImpl) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, err := c.MarketingEventService.Find(id)

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
		Data:   item,
	})
}

func (c *MarketingEventControllerImpl) Update(ctx *fiber.Ctx) error {
	var itemDTO dto.MarketingEventDTO
	id := ctx.Params("id")

	if err := ctx.BodyParser(&itemDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
		})
	}

	errs := validation.SetupValidation(itemDTO)

	if len(errs) > 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
		})
	}

	item, err := c.MarketingEventService.Update(&itemDTO, id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   item,
	})
}

func (c *MarketingEventControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.MarketingEventService.Delete(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
