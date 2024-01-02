package dto

import (
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
