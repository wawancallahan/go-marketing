package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/model"
)

type MarketingEventRepository interface {
	FindAll(tx *gorm.DB) (*[]model.MarketingEvent, error)
	Create(tx *gorm.DB, marketingEvent *model.MarketingEvent) error
	Find(tx *gorm.DB, id string) (*model.MarketingEvent, error)
	Update(tx *gorm.DB, marketingEvent model.MarketingEvent) error
	Delete(tx *gorm.DB, id string) error
}

type MarketingEventRepositoryImpl struct {
	Db *database.Database
}

func NewMarketingEventRepository(Db *database.Database) *MarketingEventRepositoryImpl {
	return &MarketingEventRepositoryImpl{Db: Db}
}

func (r *MarketingEventRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.MarketingEvent, error) {
	var marketingEvent []model.MarketingEvent

	err := tx.Find(&marketingEvent).Error

	if err != nil {
		return nil, nil
	}

	return &marketingEvent, err
}

func (r *MarketingEventRepositoryImpl) Create(tx *gorm.DB, marketingEvent *model.MarketingEvent) error {
	err := tx.Create(&marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingEventRepositoryImpl) Find(tx *gorm.DB, id string) (*model.MarketingEvent, error) {
	var marketingEvent model.MarketingEvent

	err := tx.Take(&marketingEvent, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &marketingEvent, nil
}

func (r *MarketingEventRepositoryImpl) Update(tx *gorm.DB, marketingEvent model.MarketingEvent) error {
	err := tx.Save(&marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingEventRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var marketingEvent model.MarketingEvent

	err := tx.Where("id = ?", id).Delete(&marketingEvent).Error

	if err != nil {
		return err
	}

	return nil
}
