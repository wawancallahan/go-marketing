package repository

import (
	"errors"

	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type BlogArticleRepository interface {
	FindAll(tx *gorm.DB) (*[]model.BlogArticle, error)
	Create(tx *gorm.DB, BlogArticle *model.BlogArticle) error
	Find(tx *gorm.DB, id string) (*model.BlogArticle, error)
	Update(tx *gorm.DB, BlogArticle *model.BlogArticle) error
	Delete(tx *gorm.DB, id string) error
}

type BlogArticleRepositoryImpl struct{}

func NewBlogArticleRepository() *BlogArticleRepositoryImpl {
	return &BlogArticleRepositoryImpl{}
}

func (r *BlogArticleRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.BlogArticle, error) {
	var BlogArticle []model.BlogArticle

	err := tx.Find(&BlogArticle).Error

	if err != nil {
		return nil, err
	}

	return &BlogArticle, nil
}

func (r *BlogArticleRepositoryImpl) Create(tx *gorm.DB, BlogArticle *model.BlogArticle) error {
	err := tx.Create(&BlogArticle).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogArticleRepositoryImpl) Find(tx *gorm.DB, id string) (*model.BlogArticle, error) {
	var BlogArticle model.BlogArticle

	err := tx.Take(&BlogArticle, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &BlogArticle, nil
}

func (r *BlogArticleRepositoryImpl) Update(tx *gorm.DB, BlogArticle *model.BlogArticle) error {
	err := tx.Save(&BlogArticle).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogArticleRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var BlogArticle model.BlogArticle

	err := tx.Where("id = ?", id).Delete(&BlogArticle).Error

	if err != nil {
		return err
	}

	return nil
}
