package model

import (
	"time"

	"github.com/google/uuid"
)

type MarketingLead struct {
	ID                      uuid.UUID                 `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4();<-:create" json:"id"`
	ProductCategory         string                    `gorm:"column:product_category" json:"productCategory"`
	FullName                string                    `gorm:"column:full_name" json:"fullName"`
	CompanyName             string                    `gorm:"column:company_name" json:"companyName"`
	Address                 string                    `gorm:"column:address" json:"address"`
	Email                   string                    `gorm:"column:email" json:"email"`
	PhoneNumber             string                    `gorm:"column:phone_number" json:"phoneNumber"`
	Province                string                    `gorm:"column:province" json:"province"`
	City                    string                    `gorm:"column:city" json:"city"`
	District                string                    `gorm:"column:district" json:"district"`
	RegisteredDate          time.Time                 `gorm:"column:registered_date" json:"registeredDate"`
	SourceType              string                    `gorm:"column:source_type" json:"sourceType"`
	Status                  string                    `gorm:"column:status" json:"status"`
	ActivationStatus        string                    `gorm:"column:activation_status" json:"activationStatus"`
	FollowUpBy              string                    `gorm:"column:follow_up_by" json:"followUpBy"`
	Description             string                    `gorm:"column:description" json:"description"`
	SupportName             string                    `gorm:"column:support_name" json:"supportName"`
	CreatedAt               time.Time                 `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt               time.Time                 `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
	MarketingLeadAttachment []MarketingLeadAttachment `gorm:"foreignKey:marketing_leads_id;references:id" json:"marketingLeadAttachment"`
}

func (m *MarketingLead) TableName() string {
	return "marketing_leads"
}
