package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BlogArticle struct {
	ID                 uuid.UUID      `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	Title              string         `gorm:"column:title"`
	BlogCategoryId     uuid.UUID      `gorm:"column:blog_category_id"`
	Visibility         string         `gorm:"column:visibility"`
	PublishDate        time.Time      `gorm:"column:publish_date"`
	Content            sql.NullString `gorm:"column:content"`
	SeoTitle           sql.NullString `gorm:"column:seo_title"`
	SeoSlug            sql.NullString `gorm:"column:seo_slug"`
	SeoKeywords        interface{}    `gorm:"column:seo_keywords"`
	SeoMetaDescription sql.NullString `gorm:"column:seo_meta_description"`
	TotalViews         int64          `gorm:"column:total_views;default:0"`
	CreatedAt          time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *BlogArticle) TableName() string {
	return "blog_articles"
}
