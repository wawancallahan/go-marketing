package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/mapper"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
	"matsukana.cloud/go-marketing/validation"
)

type BlogArticleController interface {
	Index(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type BlogArticleControllerImpl struct {
	BlogArticleService service.BlogArticleService
}

func NewBlogArticleController(BlogArticleService service.BlogArticleService) *BlogArticleControllerImpl {
	return &BlogArticleControllerImpl{BlogArticleService: BlogArticleService}
}

func (c *BlogArticleControllerImpl) Index(ctx *fiber.Ctx) error {
	items, err := c.BlogArticleService.Index()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	result := make([]mapper.BlogArticleMapper, 0)

	if items != nil {
		for _, item := range *items {
			result = append(result, mapper.BlogArticleMapper{
				ID:                 item.ID,
				Title:              item.Title,
				BlogCategoryId:     item.BlogCategoryId,
				Visibility:         item.Visibility,
				PublishDate:        item.PublishDate,
				Content:            null.NewString(item.Content.String, item.Content.Valid),
				SeoTitle:           null.NewString(item.SeoTitle.String, item.SeoTitle.Valid),
				SeoSlug:            null.NewString(item.SeoSlug.String, item.SeoSlug.Valid),
				SeoKeywords:        item.SeoKeywords,
				SeoMetaDescription: null.NewString(item.SeoMetaDescription.String, item.SeoMetaDescription.Valid),
				TotalViews:         item.TotalViews,
				CreatedAt:          item.CreatedAt,
				UpdatedAt:          item.UpdatedAt,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (c *BlogArticleControllerImpl) Create(ctx *fiber.Ctx) error {
	var itemDTO dto.BlogArticleDTO

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

	file, err := ctx.FormFile("file")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	if file == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  errors.New("File Required").Error(),
		})
	}

	itemDTO.File = file

	item, err := c.BlogArticleService.Create(&itemDTO)

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
		Data: mapper.BlogArticleMapper{
			ID:                 item.ID,
			Title:              item.Title,
			BlogCategoryId:     item.BlogCategoryId,
			Visibility:         item.Visibility,
			PublishDate:        item.PublishDate,
			Content:            null.NewString(item.Content.String, item.Content.Valid),
			SeoTitle:           null.NewString(item.SeoTitle.String, item.SeoTitle.Valid),
			SeoSlug:            null.NewString(item.SeoSlug.String, item.SeoSlug.Valid),
			SeoKeywords:        item.SeoKeywords,
			SeoMetaDescription: null.NewString(item.SeoMetaDescription.String, item.SeoMetaDescription.Valid),
			TotalViews:         item.TotalViews,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		},
	})
}

func (c *BlogArticleControllerImpl) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	item, err := c.BlogArticleService.Find(id)

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
		Data: mapper.BlogArticleMapper{
			ID:                 item.ID,
			Title:              item.Title,
			BlogCategoryId:     item.BlogCategoryId,
			Visibility:         item.Visibility,
			PublishDate:        item.PublishDate,
			Content:            null.NewString(item.Content.String, item.Content.Valid),
			SeoTitle:           null.NewString(item.SeoTitle.String, item.SeoTitle.Valid),
			SeoSlug:            null.NewString(item.SeoSlug.String, item.SeoSlug.Valid),
			SeoKeywords:        item.SeoKeywords,
			SeoMetaDescription: null.NewString(item.SeoMetaDescription.String, item.SeoMetaDescription.Valid),
			TotalViews:         item.TotalViews,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		},
	})
}

func (c *BlogArticleControllerImpl) Update(ctx *fiber.Ctx) error {
	var itemDTO dto.BlogArticleDTO
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
			Error:  errs,
		})
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	itemDTO.File = file

	item, err := c.BlogArticleService.Update(&itemDTO, id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Error:  err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data: mapper.BlogArticleMapper{
			ID:                 item.ID,
			Title:              item.Title,
			BlogCategoryId:     item.BlogCategoryId,
			Visibility:         item.Visibility,
			PublishDate:        item.PublishDate,
			Content:            null.NewString(item.Content.String, item.Content.Valid),
			SeoTitle:           null.NewString(item.SeoTitle.String, item.SeoTitle.Valid),
			SeoSlug:            null.NewString(item.SeoSlug.String, item.SeoSlug.Valid),
			SeoKeywords:        item.SeoKeywords,
			SeoMetaDescription: null.NewString(item.SeoMetaDescription.String, item.SeoMetaDescription.Valid),
			TotalViews:         item.TotalViews,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		},
	})
}

func (c *BlogArticleControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := c.BlogArticleService.Delete(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   nil,
	})
}
