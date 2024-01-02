package service

import (
	"github.com/google/uuid"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type BlogCategoryService interface {
	Index() (*[]model.BlogCategory, error)
	Create(itemDTO *dto.BlogCategoryDTO) (*model.BlogCategory, error)
	Find(id string) (*model.BlogCategory, error)
	Update(itemDTO *dto.BlogCategoryDTO, id string) (*model.BlogCategory, error)
	Delete(id string) error
}

type BlogCategoryServiceImpl struct {
	Db                     *database.Database
	BlogCategoryRepository repository.BlogCategoryRepository
}

func NewBlogCategoryService(Db *database.Database, BlogCategoryRepository repository.BlogCategoryRepository) *BlogCategoryServiceImpl {
	return &BlogCategoryServiceImpl{Db: Db, BlogCategoryRepository: BlogCategoryRepository}
}

func (s *BlogCategoryServiceImpl) Index() (*[]model.BlogCategory, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.BlogCategoryRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return items, nil
}

func (s *BlogCategoryServiceImpl) Create(itemDTO *dto.BlogCategoryDTO) (*model.BlogCategory, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item := itemDTO.ToModel()

	err := s.BlogCategoryRepository.Create(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil

}

func (s *BlogCategoryServiceImpl) Find(id string) (*model.BlogCategory, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item, err := s.BlogCategoryRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return item, nil
}

func (s *BlogCategoryServiceImpl) Update(itemDTO *dto.BlogCategoryDTO, id string) (*model.BlogCategory, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item := itemDTO.ToModel()
	item.ID = uuid.MustParse(id)

	err := s.BlogCategoryRepository.Update(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}

func (s *BlogCategoryServiceImpl) Delete(id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	err := s.BlogCategoryRepository.Delete(tx, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
