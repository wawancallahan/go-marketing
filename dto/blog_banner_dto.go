package dto

import (
	"mime/multipart"

	"matsukana.cloud/go-marketing/model"
)

type BlogBannerDTO struct {
	Name string `json:"name" validate:"required"`
}

func (d *BlogBannerDTO) ToModel() model.BlogBanner {
	return model.BlogBanner{
		Name: d.Name,
	}
}

type BlogBannerUpdateDTO struct {
	File *multipart.FileHeader `json:"-" form:"-"`
}
