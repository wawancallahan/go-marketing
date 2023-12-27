package dto

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadDTO struct {
	ProductCategory  string      `json:"productCategory" validate:"required;oneof='Pinjam Modal Inventory' 'Pinjam Modal Usaha' 'Pinjam Modal Toko' 'Pinjam Modal Karyawan'"`
	FullName         string      `json:"fullName" validate:"required"`
	CompanyName      string      `json:"companyName" validate:"required"`
	Address          null.String `json:"address" validate:"omitempty"`
	Email            string      `json:"email" validate:"required"`
	PhoneNumber      string      `json:"phoneNumber" validate:"required"`
	Province         string      `json:"province" validate:"required"`
	City             string      `json:"city" validate:"required"`
	District         string      `json:"district" validate:"required"`
	RegisteredDate   string      `json:"registeredDate" validate:"omitempty"`
	SourceType       string      `json:"sourceType" validate:"required;oneof='EVENT_OFFLINE' 'EVENT_ONLINE' 'DASHBOARD_INTERNAL' 'OFFICIAL_WEBSITE'"`
	Status           null.String `json:"status" validate:"omitempty"`
	ActivationStatus string      `json:"activationStatus" validate:"required"`
	FollowUpBy       null.String `json:"followUpBy" validate:"omitempty"`
	Description      string      `json:"description" validate:"required"`
	SupportName      string      `json:"supportName" validate:"required"`
}

func (d *MarketingLeadDTO) ToModel() model.MarketingLead {
	tz, _ := time.LoadLocation("Asia/Jakarta")

	registeredDate := time.Now().In(tz)

	return model.MarketingLead{
		ProductCategory: d.ProductCategory,
		FullName:        d.FullName,
		CompanyName:     d.CompanyName,
		Address: sql.NullString{
			String: d.Address.String,
			Valid:  d.Address.Valid,
		},
		Email:          d.Email,
		PhoneNumber:    d.PhoneNumber,
		Province:       d.Province,
		City:           d.City,
		District:       d.District,
		RegisteredDate: registeredDate,
		SourceType:     d.SourceType,
		Status: sql.NullString{
			String: d.Status.String,
			Valid:  d.Status.Valid,
		},
		ActivationStatus: d.ActivationStatus,
		FollowUpBy: sql.NullString{
			String: d.FollowUpBy.String,
			Valid:  d.FollowUpBy.Valid,
		},
		Description: d.Description,
		SupportName: d.SupportName,
	}
}
