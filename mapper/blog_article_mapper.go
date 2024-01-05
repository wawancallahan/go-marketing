package mapper

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type BlogArticleMapper struct {
	ID                 uuid.UUID   `json:"id"`
	Title              string      `json:"title"`
	BlogCategoryId     uuid.UUID   `json:"blogCategoryId"`
	Visibility         string      `json:"visibility"`
	PublishDate        time.Time   `json:"publish_date"`
	Content            null.String `json:"content"`
	SeoTitle           null.String `json:"seoTitle"`
	SeoSlug            null.String `json:"seoSlug"`
	SeoKeywords        interface{} `json:"seoKeywords"`
	SeoMetaDescription null.String `json:"seoMetaDescription"`
	TotalViews         int64       `json:"totalViews"`
	CreatedAt          time.Time   `json:"createdAt"`
	UpdatedAt          time.Time   `json:"updatedAt"`
}
