package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadRepository struct {
	Db *database.Database
}

func NewMarketingLeadRepository(Db *database.Database) interface{} {
	return &MarketingLeadRepository{Db: Db}
}

func (r *MarketingLeadRepository) FindAll(tx *gorm.DB) ([]model.MarketingLead, error) {
	var marketingLead []model.MarketingLead

	err := tx.Find(&marketingLead).Error

	if err != nil {
		return marketingLead, nil
	}

	return marketingLead, err
}

func (r *MarketingLeadRepository) Create(tx *gorm.DB, marketingLead model.MarketingLead) error {
	err := tx.Create(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepository) Find(tx *gorm.DB, id string) (model.MarketingLead, error) {
	var marketingLead model.MarketingLead

	err := tx.Find(&marketingLead, id).Error

	if err != nil {
		return marketingLead, err
	}

	return marketingLead, nil
}

func (r *MarketingLeadRepository) Update(tx *gorm.DB, marketingLead model.MarketingLead) error {
	err := tx.Save(&marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadRepository) Delete(tx *gorm.DB, id string) error {
	var marketingLead model.MarketingLead

	err := tx.Where("id = ?", id).Delete(marketingLead).Error

	if err != nil {
		return err
	}

	return nil
}
