package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogBanner struct {
	gorm.Model
	ID        uuid.UUID      `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	Name      string         `gorm:"column:name"`
	FileName  sql.NullString `gorm:"column:file_name"`
	Path      sql.NullString `gorm:"column:path"`
	Url       sql.NullString `gorm:"column:url"`
	MimeType  sql.NullString `gorm:"column:mime_type"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
}

func (m *BlogBanner) TableName() string {
	return "blog_banners"
}
