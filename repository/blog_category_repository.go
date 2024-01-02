package repository

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type BlogCategoryRepository interface {
	FindAll(tx *gorm.DB) (*[]model.BlogCategory, error)
	Create(tx *gorm.DB, BlogCategory *model.BlogCategory) error
	Find(tx *gorm.DB, id string) (*model.BlogCategory, error)
	Update(tx *gorm.DB, BlogCategory *model.BlogCategory) error
	Delete(tx *gorm.DB, id string) error
}

type BlogCategoryRepositoryImpl struct{}

func NewBlogCategoryRepository() *BlogCategoryRepositoryImpl {
	return &BlogCategoryRepositoryImpl{}
}

func (r *BlogCategoryRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.BlogCategory, error) {
	var BlogCategory []model.BlogCategory

	err := tx.Find(&BlogCategory).Error

	if err != nil {
		return nil, nil
	}

	return &BlogCategory, err
}

func (r *BlogCategoryRepositoryImpl) Create(tx *gorm.DB, BlogCategory *model.BlogCategory) error {
	err := tx.Create(&BlogCategory).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogCategoryRepositoryImpl) Find(tx *gorm.DB, id string) (*model.BlogCategory, error) {
	var BlogCategory model.BlogCategory

	err := tx.Take(&BlogCategory, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &BlogCategory, nil
}

func (r *BlogCategoryRepositoryImpl) Update(tx *gorm.DB, BlogCategory *model.BlogCategory) error {
	err := tx.Save(&BlogCategory).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogCategoryRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var BlogCategory model.BlogCategory

	err := tx.Where("id = ?", id).Delete(&BlogCategory).Error

	if err != nil {
		return err
	}

	return nil
}
