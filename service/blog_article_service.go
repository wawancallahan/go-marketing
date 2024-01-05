package service

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type BlogArticleService interface {
	Index() (*[]model.BlogArticle, error)
	Create(itemDTO *dto.BlogArticleDTO) (*model.BlogArticle, error)
	Find(id string) (*model.BlogArticle, error)
	Update(itemDTO *dto.BlogArticleDTO, id string) (*model.BlogArticle, error)
	Delete(id string) error
}

type BlogArticleServiceImpl struct {
	Db                    *database.Database
	BlogArticleRepository repository.BlogArticleRepository
}

func NewBlogArticleService(Db *database.Database, BlogArticleRepository repository.BlogArticleRepository) *BlogArticleServiceImpl {
	return &BlogArticleServiceImpl{Db: Db, BlogArticleRepository: BlogArticleRepository}
}

func (s *BlogArticleServiceImpl) Index() (*[]model.BlogArticle, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.BlogArticleRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return items, nil
}

func (s *BlogArticleServiceImpl) Create(itemDTO *dto.BlogArticleDTO) (*model.BlogArticle, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	seoSlug := slug.Make(itemDTO.SeoTitle)

	seoKeywords, err := json.Marshal(itemDTO.SeoKeywords)

	if err != nil {
		return nil, err
	}

	item := model.BlogArticle{
		Title:          itemDTO.Title,
		BlogCategoryId: uuid.MustParse(itemDTO.BlogCategoryId),
		Visibility:     itemDTO.Visibility,
		PublishDate:    itemDTO.PublishDate,
		Content: sql.NullString{
			String: itemDTO.Content,
			Valid:  true,
		},
		SeoTitle: sql.NullString{
			String: itemDTO.SeoTitle,
			Valid:  true,
		},
		SeoSlug: sql.NullString{
			String: seoSlug,
			Valid:  true,
		},
		SeoKeywords: seoKeywords,
		SeoMetaDescription: sql.NullString{
			String: itemDTO.SeoMetaDescription,
			Valid:  true,
		},
		TotalViews: 0,
	}

	err = s.BlogArticleRepository.Create(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil

}

func (s *BlogArticleServiceImpl) Find(id string) (*model.BlogArticle, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item, err := s.BlogArticleRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	if item == nil {
		return nil, errors.New("Data Not Found")
	}

	tx.Commit()

	return item, nil
}

func (s *BlogArticleServiceImpl) Update(itemDTO *dto.BlogArticleDTO, id string) (*model.BlogArticle, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	seoSlug := slug.Make(itemDTO.SeoTitle)

	seoKeywords, err := json.Marshal(itemDTO.SeoKeywords)

	if err != nil {
		return nil, err
	}

	item := model.BlogArticle{
		ID:             uuid.MustParse(id),
		Title:          itemDTO.Title,
		BlogCategoryId: uuid.MustParse(itemDTO.BlogCategoryId),
		Visibility:     itemDTO.Visibility,
		PublishDate:    itemDTO.PublishDate,
		Content: sql.NullString{
			String: itemDTO.Content,
			Valid:  true,
		},
		SeoTitle: sql.NullString{
			String: itemDTO.SeoTitle,
			Valid:  true,
		},
		SeoSlug: sql.NullString{
			String: seoSlug,
			Valid:  true,
		},
		SeoKeywords: seoKeywords,
		SeoMetaDescription: sql.NullString{
			String: itemDTO.SeoMetaDescription,
			Valid:  true,
		},
	}

	if err := s.BlogArticleRepository.Update(tx, &item); err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}

func (s *BlogArticleServiceImpl) Delete(id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	err := s.BlogArticleRepository.Delete(tx, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
