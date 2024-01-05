package dto

import (
	"mime/multipart"
	"time"
)

type BlogArticleDTO struct {
	Title              string                `json:"title" validate:"required"`
	BlogCategoryId     string                `json:"blogCategoryId" validate:"required"`
	Visibility         string                `json:"visibility" validate:"required"`
	PublishDate        time.Time             `json:"description" validate:"required"`
	Content            string                `json:"content" validate:"required"`
	SeoTitle           string                `json:"seoTitle" validate:"required"`
	SeoKeywords        string                `json:"seoKeywords" validate:"required"`
	SeoMetaDescription string                `json:"seoMetaDescription" validate:"required"`
	AppNotification    uint8                 `json:"appNotification" validate:"required,int"`
	File               *multipart.FileHeader `json:"-"`
}
