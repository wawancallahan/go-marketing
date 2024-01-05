package model

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type BlogArticleAttachment struct {
	ID             uuid.UUID   `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	Name           string      `gorm:"column:name"`
	FileName       null.String `gorm:"column:file_name"`
	Path           null.String `gorm:"column:path"`
	MimeType       null.String `gorm:"column:mime_type"`
	Url            null.String `gorm:"column:url"`
	BlogArticlesId uuid.UUID   `gorm:"column:blog_articles_id"`
	CreatedAt      time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *BlogArticleAttachment) TableName() string {
	return "blog_article_attachments"
}
