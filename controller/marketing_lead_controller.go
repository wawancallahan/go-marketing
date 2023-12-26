package controller

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
	"matsukana.cloud/go-marketing/validation"
)

type MarketingLeadController interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type MarketingLeadControllerImpl struct {
	MarketingLeadService service.MarketingLeadService
}

func NewMarketingLeadController(MarketingLeadService service.MarketingLeadService) *MarketingLeadControllerImpl {
	return &MarketingLeadControllerImpl{MarketingLeadService: MarketingLeadService}
}

func (c *MarketingLeadControllerImpl) Index(ctx *fiber.Ctx) error {
	items, err := c.MarketingLeadService.Index()

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

func (c *MarketingLeadControllerImpl) Create(ctx *fiber.Ctx) error {
	var itemDTO dto.MarketingLeadDTO

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

	item, err := c.MarketingLeadService.Create(&itemDTO)

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

func (c *MarketingLeadControllerImpl) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, err := c.MarketingLeadService.Find(id)

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

func (c *MarketingLeadControllerImpl) Update(ctx *fiber.Ctx) error {
	var itemDTO dto.MarketingLeadDTO
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

	err := c.MarketingLeadService.Update(&itemDTO, id)

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

func (c *MarketingLeadControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.MarketingLeadService.Delete(id)

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
