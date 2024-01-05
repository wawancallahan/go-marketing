package model

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type MarketingLeadAttachment struct {
	ID               uuid.UUID     `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	FileName         null.String   `gorm:"column:file_name"`
	Path             null.String   `gorm:"column:path"`
	MimeType         null.String   `gorm:"column:mime_type"`
	MarketingLeadsId uuid.UUID     `gorm:"column:marketing_leads_id"`
	CreatedAt        time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time     `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	MarketingLead    MarketingLead `gorm:"foreignKey:marketing_leads_id;references:id"`
}

func (m *MarketingLeadAttachment) TableName() string {
	return "marketing_leads_attachments"
}
