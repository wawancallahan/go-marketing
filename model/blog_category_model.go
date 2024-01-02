package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type BlogCategory struct {
	ID          uuid.UUID      `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	Name        string         `gorm:"column:name"`
	Slug        string         `gorm:"column:slug"`
	IsActive    bool           `gorm:"column:is_active"`
	Description sql.NullString `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *BlogCategory) TableName() string {
	return "blog_categories"
}
