package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type MarketingEventRepository interface {
	FindAll(tx *gorm.DB) (*[]model.MarketingEvent, error)
	Create(tx *gorm.DB, marketingEvent *model.MarketingEvent) error
	Find(tx *gorm.DB, id string) (*model.MarketingEvent, error)
	Update(tx *gorm.DB, marketingEvent *model.MarketingEvent) error
	Delete(tx *gorm.DB, id string) error
	UpdateOngoingStatus(tx *gorm.DB) error
	UpdateCompletedStatus(tx *gorm.DB) error
}

type MarketingEventRepositoryImpl struct {
}

func NewMarketingEventRepository() *MarketingEventRepositoryImpl {
	return &MarketingEventRepositoryImpl{}
}

func (r *MarketingEventRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.MarketingEvent, error) {
	var marketingEvent []model.MarketingEvent

	err := tx.Find(&marketingEvent).Error

	if err != nil {
		return nil, err
	}

	return &marketingEvent, nil
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

func (r *MarketingEventRepositoryImpl) Update(tx *gorm.DB, marketingEvent *model.MarketingEvent) error {
	err := tx.Updates(&marketingEvent).Error

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

func (r *MarketingEventRepositoryImpl) UpdateOngoingStatus(tx *gorm.DB) error {
	var marketingEvent model.MarketingEvent

	err := tx.Model(&marketingEvent).Where("status = ?", "ONGOING").Where("date_trunc('day', event_time) = CURRENT_DATE").Updates(&model.MarketingEvent{
		Status: "ONGOING",
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingEventRepositoryImpl) UpdateCompletedStatus(tx *gorm.DB) error {
	var marketingEvent model.MarketingEvent

	err := tx.Model(&marketingEvent).Where(tx.Where("status = ?", "ONGOING").Or("status = ?", "COMPLETED")).Where("date_trunc('day', event_time) < CURRENT_DATE").Updates(&model.MarketingEvent{
		Status: "COMPLETED",
	}).Error

	if err != nil {
		return err
	}

	return nil
}
