package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadRepository interface {
	FindAll(tx *gorm.DB) (*[]model.MarketingLead, error)
	Create(tx *gorm.DB, marketingLead *model.MarketingLead) error
	Find(tx *gorm.DB, id string) (*model.MarketingLead, error)
	Update(tx *gorm.DB, marketingLead model.MarketingLead) error
	Delete(tx *gorm.DB, id string) error
}

type MarketingLeadRepositoryImpl struct {
	Db *database.Database
}

func NewMarketingLeadRepository(Db *database.Database) *MarketingLeadRepositoryImpl {
	return &MarketingLeadRepositoryImpl{Db: Db}
}

func (r *MarketingLeadRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.MarketingLead, error) {
	var marketingLead []model.MarketingLead

	err := tx.Find(&marketingLead).Error

	if err != nil {
		return nil, nil
	}

	return &marketingLead, err
}

func (r *MarketingLeadRepositoryImpl) Create(tx *gorm.DB, marketingLead *model.MarketingLead) error {
	err := tx.Create(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) Find(tx *gorm.DB, id string) (*model.MarketingLead, error) {
	var marketingLead model.MarketingLead

	err := tx.Take(&marketingLead, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &marketingLead, nil
}

func (r *MarketingLeadRepositoryImpl) Update(tx *gorm.DB, marketingLead model.MarketingLead) error {
	err := tx.Save(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var marketingLead model.MarketingLead

	err := tx.Where("id = ?", id).Delete(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}
