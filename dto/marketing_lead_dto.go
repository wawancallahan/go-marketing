package dto

import (
	"time"

	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadDTO struct {
	ProductCategory  string `json:"productCategory" validate:"required"`
	FullName         string `json:"fullName" validate:"required"`
	CompanyName      string `json:"companyName" validate:"required"`
	Address          string `json:"address" validate:"required"`
	Email            string `json:"email" validate:"required"`
	PhoneNumber      string `json:"phoneNumber" validate:"required"`
	Province         string `json:"province" validate:"required"`
	City             string `json:"city" validate:"required"`
	District         string `json:"district" validate:"required"`
	RegisteredDate   string `json:"registeredDate" validate:"omitempty"`
	SourceType       string `json:"sourceType" validate:"required"`
	Status           string `json:"status" validate:"required"`
	ActivationStatus string `json:"activationStatus" validate:"required"`
	FollowUpBy       string `json:"followUpBy" validate:"required"`
	Description      string `json:"description" validate:"required"`
	SupportName      string `json:"supportName" validate:"required"`
}

func (d *MarketingLeadDTO) ToModel() model.MarketingLead {
	registeredDate, _ := time.Parse("2006-01-02T15:04:05Z07:00", time.Now().String())
	return model.MarketingLead{
		ProductCategory:  d.ProductCategory,
		FullName:         d.FullName,
		CompanyName:      d.CompanyName,
		Address:          d.Address,
		Email:            d.Email,
		PhoneNumber:      d.PhoneNumber,
		Province:         d.Province,
		City:             d.City,
		District:         d.District,
		RegisteredDate:   registeredDate,
		SourceType:       d.SourceType,
		Status:           d.Status,
		ActivationStatus: d.ActivationStatus,
		FollowUpBy:       d.FollowUpBy,
		Description:      d.Description,
		SupportName:      d.SupportName,
	}
}
