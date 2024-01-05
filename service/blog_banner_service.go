package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type BlogBannerService interface {
	Index() (*[]model.BlogBanner, error)
	Update(itemDTO dto.BlogBannerUpdateDTO, id string) (*model.BlogBanner, error)
}

type BlogBannerServiceImpl struct {
	Db                   *database.Database
	BlogBannerRepository repository.BlogBannerRepository
}

func NewBlogBannerService(Db *database.Database, BlogBannerRepository repository.BlogBannerRepository) *BlogBannerServiceImpl {
	return &BlogBannerServiceImpl{Db: Db, BlogBannerRepository: BlogBannerRepository}
}

func (s *BlogBannerServiceImpl) Index() (*[]model.BlogBanner, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.BlogBannerRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return items, nil
}

func (s *BlogBannerServiceImpl) Update(itemDTO dto.BlogBannerUpdateDTO, id string) (*model.BlogBanner, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	blogBannerItem, err := s.BlogBannerRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	if blogBannerItem == nil {
		return nil, errors.New("Data Not Found")
	}

	openFile, err := itemDTO.File.Open()

	defer openFile.Close()

	if err != nil {
		return nil, errors.New("File Cannot Open")
	}

	byteFile, err := io.ReadAll(openFile)

	if err != nil {
		return nil, errors.New("File Cannot be Read")
	}

	mtype := mimetype.Detect(byteFile)

	log.Printf("%s", mtype.String())

	filePath := "upload/blog/banners"

	// TODO: Upload to OSS

	urlPath := ""

	item := model.BlogBanner{
		ID:   uuid.MustParse(id),
		Name: blogBannerItem.Name,
		FileName: sql.NullString{
			String: fmt.Sprintf("%s_%s%s", uuid.New(), id, mtype.Extension()),
			Valid:  true,
		},
		MimeType: sql.NullString{
			String: mtype.String(),
			Valid:  true,
		},
		Path: sql.NullString{
			String: filePath,
			Valid:  true,
		},
		Url: sql.NullString{
			String: urlPath,
			Valid:  true,
		},
		CreatedAt: blogBannerItem.CreatedAt,
	}

	if err := s.BlogBannerRepository.Update(tx, &item); err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}
