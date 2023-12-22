package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/model"
)

type MarketingEventRepository struct {
	Db *database.Database
}

func NewMarketingEventRepository(Db *database.Database) interface{} {
	return &MarketingEventRepository{Db: Db}
}

func (r *MarketingEventRepository) FindAll(tx *gorm.DB) ([]model.MarketingEvent, error) {
	var marketingEvent []model.MarketingEvent

	err := tx.Find(&marketingEvent).Error

	if err != nil {
		return marketingEvent, nil
	}

	return marketingEvent, err
}

func (r *MarketingEventRepository) Create(tx *gorm.DB, marketingEvent *model.MarketingEvent) error {
	err := tx.Create(marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingEventRepository) Find(tx *gorm.DB, id string) (model.MarketingEvent, error) {
	var marketingEvent model.MarketingEvent

	err := tx.Find(&marketingEvent, id).Error

	if err != nil {
		return marketingEvent, err
	}

	return marketingEvent, nil
}

func (r *MarketingEventRepository) Update(tx *gorm.DB, marketingEvent model.MarketingEvent) error {
	err := tx.Save(&marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingEventRepository) Delete(tx *gorm.DB, id string) error {
	var marketingEvent model.MarketingEvent

	err := tx.Where("id = ?", id).Delete(marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}
