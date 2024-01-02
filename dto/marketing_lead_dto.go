package dto

import (
	"database/sql"
	"mime/multipart"
	"time"

	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadDTO struct {
	ProductCategory  string                `form:"productCategory" validate:"required;oneof='Pinjam Modal Inventory' 'Pinjam Modal Usaha' 'Pinjam Modal Toko' 'Pinjam Modal Karyawan'"`
	FullName         string                `form:"fullName" validate:"required"`
	CompanyName      string                `form:"companyName" validate:"required"`
	Address          null.String           `form:"address,omitempty" validate:"omitempty"`
	Email            string                `form:"email" validate:"required"`
	PhoneNumber      string                `form:"phoneNumber" validate:"required"`
	Province         string                `form:"province" validate:"required"`
	City             string                `form:"city" validate:"required"`
	District         string                `form:"district" validate:"required"`
	RegisteredDate   string                `form:"registeredDate,omitempty" validate:"omitempty"`
	SourceType       string                `form:"sourceType" validate:"required;oneof='EVENT_OFFLINE' 'EVENT_ONLINE' 'DASHBOARD_INTERNAL' 'OFFICIAL_WEBSITE'"`
	Status           null.String           `form:"status,omitempty" validate:"omitempty"`
	ActivationStatus string                `form:"activationStatus" validate:"required"`
	FollowUpBy       null.String           `form:"followUpBy,omitempty" validate:"omitempty"`
	Description      string                `form:"description" validate:"required"`
	SupportName      string                `form:"supportName" validate:"required"`
	File             *multipart.FileHeader `form:"-"`
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
