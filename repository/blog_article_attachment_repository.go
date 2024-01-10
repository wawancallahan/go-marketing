package repository

import (
	"errors"

	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/model"
)

type BlogArticleAttachmentRepository interface {
	FindAll(tx *gorm.DB) (*[]model.BlogArticleAttachment, error)
	Create(tx *gorm.DB, BlogArticleAttachment *model.BlogArticleAttachment) error
	Find(tx *gorm.DB, id string) (*model.BlogArticleAttachment, error)
	FindByName(tx *gorm.DB, id string) (*model.BlogArticleAttachment, error)
	Update(tx *gorm.DB, BlogArticleAttachment *model.BlogArticleAttachment) error
	Delete(tx *gorm.DB, id string) error
}

type BlogArticleAttachmentRepositoryImpl struct{}

func NewBlogArticleAttachmentRepository() *BlogArticleAttachmentRepositoryImpl {
	return &BlogArticleAttachmentRepositoryImpl{}
}

func (r *BlogArticleAttachmentRepositoryImpl) FindAll(tx *gorm.DB) (*[]model.BlogArticleAttachment, error) {
	var BlogArticleAttachment []model.BlogArticleAttachment

	err := tx.Find(&BlogArticleAttachment).Error

	if err != nil {
		return nil, err
	}

	return &BlogArticleAttachment, nil
}

func (r *BlogArticleAttachmentRepositoryImpl) Create(tx *gorm.DB, BlogArticleAttachment *model.BlogArticleAttachment) error {
	err := tx.Create(&BlogArticleAttachment).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogArticleAttachmentRepositoryImpl) Find(tx *gorm.DB, id string) (*model.BlogArticleAttachment, error) {
	var BlogArticleAttachment model.BlogArticleAttachment

	err := tx.Take(&BlogArticleAttachment, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &BlogArticleAttachment, nil
}

func (r *BlogArticleAttachmentRepositoryImpl) FindByName(tx *gorm.DB, name string) (*model.BlogArticleAttachment, error) {
	var BlogArticleAttachment model.BlogArticleAttachment

	err := tx.Take(&BlogArticleAttachment, "name = ?", name).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &BlogArticleAttachment, nil
}

func (r *BlogArticleAttachmentRepositoryImpl) Update(tx *gorm.DB, BlogArticleAttachment *model.BlogArticleAttachment) error {
	err := tx.Updates(&BlogArticleAttachment).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BlogArticleAttachmentRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var BlogArticleAttachment model.BlogArticleAttachment

	err := tx.Where("id = ?", id).Delete(&BlogArticleAttachment).Error

	if err != nil {
		return err
	}

	return nil
}
