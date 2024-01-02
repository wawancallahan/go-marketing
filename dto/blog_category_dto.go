package dto

import (
	"database/sql"

	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/model"
)

type BlogCategoryDTO struct {
	Name        string      `json:"name" validate:"required"`
	Slug        string      `json:"slug" validate:"required"`
	IsActive    bool        `json:"is_active" validate:"required"`
	Description null.String `json:"description,omitempty" validate:"omitempty"`
}

func (d *BlogCategoryDTO) ToModel() model.BlogCategory {
	return model.BlogCategory{
		Name:     d.Name,
		Slug:     d.Slug,
		IsActive: d.IsActive,
		Description: sql.NullString{
			Valid:  d.Description.Valid,
			String: d.Description.String,
		},
	}
}
