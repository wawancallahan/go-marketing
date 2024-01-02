package controller

import (
	"github.com/gofiber/fiber/v2"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
	"matsukana.cloud/go-marketing/validation"
)

type BlogCategoryController interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type BlogCategoryControllerImpl struct {
	BlogCategoryService service.BlogCategoryService
}

func NewBlogCategoryController(BlogCategoryService service.BlogCategoryService) *BlogCategoryControllerImpl {
	return &BlogCategoryControllerImpl{BlogCategoryService: BlogCategoryService}
}

func (c *BlogCategoryControllerImpl) Index(ctx *fiber.Ctx) error {
	items, err := c.BlogCategoryService.Index()

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

func (c *BlogCategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var itemDTO dto.BlogCategoryDTO

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

	item, err := c.BlogCategoryService.Create(&itemDTO)

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

func (c *BlogCategoryControllerImpl) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, err := c.BlogCategoryService.Find(id)

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

func (c *BlogCategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	var itemDTO dto.BlogCategoryDTO
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

	item, err := c.BlogCategoryService.Update(&itemDTO, id)

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

func (c *BlogCategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.BlogCategoryService.Delete(id)

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
