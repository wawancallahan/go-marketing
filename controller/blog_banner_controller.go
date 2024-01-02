package controller

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/mapper"
	"matsukana.cloud/go-marketing/response"
	"matsukana.cloud/go-marketing/service"
)

type BlogBannerController interface {
	Index(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type BlogBannerControllerImpl struct {
	BlogBannerService service.BlogBannerService
}

func NewBlogBannerController(BlogBannerService service.BlogBannerService) *BlogBannerControllerImpl {
	return &BlogBannerControllerImpl{BlogBannerService: BlogBannerService}
}
func (c *BlogBannerControllerImpl) Index(ctx *fiber.Ctx) error {
	items, err := c.BlogBannerService.Index()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "NOK",
			Data:   nil,
			Error:  err.Error(),
		})
	}

	result := make([]mapper.BlogBannerMapper, 0)

	if items != nil {
		for _, item := range *items {
			result = append(result, mapper.BlogBannerMapper{
				ID:        item.ID,
				Name:      item.Name,
				FileName:  null.NewString(item.FileName.String, item.FileName.Valid),
				Path:      null.NewString(item.Path.String, item.Path.Valid),
				Url:       null.NewString(item.Url.String, item.Url.Valid),
				MimeType:  null.NewString(item.Url.String, item.Url.Valid),
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(response.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (c *BlogBannerControllerImpl) Update(ctx *fiber.Ctx) error {
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
			Error:  errors.New("FIle Required").Error(),
		})
	}

	id := ctx.Params("id")

	item, err := c.BlogBannerService.Update(file, id)

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
		Data: mapper.BlogBannerMapper{
			ID:        item.ID,
			Name:      item.Name,
			FileName:  null.NewString(item.FileName.String, item.FileName.Valid),
			Path:      null.NewString(item.Path.String, item.Path.Valid),
			Url:       null.NewString(item.Url.String, item.Url.Valid),
			MimeType:  null.NewString(item.Url.String, item.Url.Valid),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		},
	})
}
