package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type BlogBannerRepository interface {
	FindAll(tx *gorm.DB) (*[]model.BlogBanner, error)
	Create(tx *gorm.DB, BlogBanner *model.BlogBanner) error
	Find(tx *gorm.DB, id string) (*model.BlogBanner, error)
	Update(tx *gorm.DB, BlogBanner *model.BlogBanner) error
	Delete(tx *gorm.DB, id string) error
}

type BlogBannerRepositoryImpl struct{}

func NewBlogBannerRepository() *BlogBannerRepositoryImpl {
	return &BlogBannerRepositoryImpl{}
}

func (r *BlogBannerRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.BlogBanner, error) {
	var BlogBanner []model.BlogBanner

	err := tx.Find(&BlogBanner).Error

	if err != nil {
		return nil, nil
	}

	return &BlogBanner, err
}

func (r *BlogBannerRepositoryImpl) Create(tx *gorm.DB, BlogBanner *model.BlogBanner) error {
	err := tx.Create(&BlogBanner).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogBannerRepositoryImpl) Find(tx *gorm.DB, id string) (*model.BlogBanner, error) {
	var BlogBanner model.BlogBanner

	err := tx.Take(&BlogBanner, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &BlogBanner, nil
}

func (r *BlogBannerRepositoryImpl) Update(tx *gorm.DB, BlogBanner *model.BlogBanner) error {
	err := tx.Updates(&BlogBanner).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogBannerRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var BlogBanner model.BlogBanner

	err := tx.Where("id = ?", id).Delete(&BlogBanner).Error

	if err != nil {
		return err
	}

	return nil
}
