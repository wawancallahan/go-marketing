package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type MarketingLeadAttachmentRepository struct {
}

func NewMarketingLeadAttachmentRepository() interface{} {
	return &MarketingLeadAttachmentRepository{}
}

func (r *MarketingLeadAttachmentRepository) FindAll(tx *gorm.DB) ([]model.MarketingLeadAttachment, error) {
	var marketingLeadAttachment []model.MarketingLeadAttachment

	err := tx.Find(&marketingLeadAttachment).Error

	if err != nil {
		return marketingLeadAttachment, nil
	}

	return marketingLeadAttachment, err
}

func (r *MarketingLeadAttachmentRepository) Create(tx *gorm.DB, marketingLeadAttachment model.MarketingLeadAttachment) error {
	err := tx.Create(&marketingLeadAttachment).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadAttachmentRepository) Find(tx *gorm.DB, id string) (model.MarketingLeadAttachment, error) {
	var marketingLeadAttachment model.MarketingLeadAttachment

	err := tx.Find(&marketingLeadAttachment, id).Error

	if err != nil {
		return marketingLeadAttachment, err
	}

	return marketingLeadAttachment, nil
}

func (r *MarketingLeadAttachmentRepository) Update(tx *gorm.DB, marketingLeadAttachment model.MarketingLeadAttachment) error {
	err := tx.Save(&marketingLeadAttachment).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *MarketingLeadAttachmentRepository) Delete(tx *gorm.DB, id string) error {
	var marketingLeadAttachment model.MarketingLeadAttachment

	err := tx.Where("id = ?", id).Delete(marketingLeadAttachment).Error

	if err != nil {
		return err
	}

	return nil
}
