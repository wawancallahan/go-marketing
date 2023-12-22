package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketingLeadAttachment struct {
	gorm.Model
	ID               uuid.UUID     `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	FileName         string        `gorm:"column:file_name"`
	Path             time.Time     `gorm:"column:path"`
	MimeType         string        `gorm:"column:mime_type"`
	MarketingLeadsId string        `gorm:"column:marketing_leads_id"`
	CreatedAt        time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time     `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	MarketingLead    MarketingLead `gorm:"foreignKey:marketing_leads_id;references:id"`
}

func (m *MarketingLeadAttachment) TableName() string {
	return "marketing_leads_attachments"
}
