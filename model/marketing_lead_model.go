package model

import (
	"time"

	"github.com/google/uuid"
)

type MarketingLead struct {
	ID                      uuid.UUID                 `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4();<-:create"`
	ProductCategory         string                    `gorm:"column:product_category"`
	FullName                time.Time                 `gorm:"column:full_name"`
	CompanyName             string                    `gorm:"column:company_name"`
	Address                 string                    `gorm:"column:address"`
	Email                   string                    `gorm:"column:email"`
	PhoneNumber             string                    `gorm:"column:phone_number"`
	Province                string                    `gorm:"column:province"`
	City                    string                    `gorm:"column:city"`
	District                string                    `gorm:"column:district"`
	RegisteredDate          int32                     `gorm:"column:registered_date"`
	SourceType              string                    `gorm:"column:source_type"`
	Status                  string                    `gorm:"column:status"`
	ActivationStatus        string                    `gorm:"column:activation_status"`
	FollowUpBy              string                    `gorm:"column:follow_up_by"`
	Description             string                    `gorm:"column:description"`
	SupportName             string                    `gorm:"column:support_name"`
	CreatedAt               time.Time                 `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt               time.Time                 `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	MarketingLeadAttachment []MarketingLeadAttachment `gorm:"foreignKey:marketing_leads_id;references:id"`
}

func (m *MarketingLead) TableName() string {
	return "marketing_leads"
}
