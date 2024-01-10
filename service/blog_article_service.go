package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	httprequest "matsukana.cloud/go-marketing/http_request"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type BlogArticleService interface {
	Index() (*[]model.BlogArticle, error)
	Create(itemDTO *dto.BlogArticleDTO) (*model.BlogArticle, error)
	Find(id string) (*model.BlogArticle, error)
	Update(itemDTO *dto.BlogArticleDTO, id string) (*model.BlogArticle, error)
	Delete(id string) error
	SaveImageArticle(itemDTO *dto.BlogArticleDTO, blogArticle model.BlogArticle) error
	UploadImage(itemDTO *dto.BlogArticleDTO, blogArticle model.BlogArticle) (*dto.UploadImageResult, error)
	DeleteImage(urlPath string) error
}

type BlogArticleServiceImpl struct {
	Db                              *database.Database
	Config                          *config.Config
	BlogArticleRepository           repository.BlogArticleRepository
	BlogArticleAttachmentRepository repository.BlogArticleAttachmentRepository
}

func NewBlogArticleService(Db *database.Database, Config *config.Config, BlogArticleRepository repository.BlogArticleRepository, BlogArticleAttachmentRepository repository.BlogArticleAttachmentRepository) *BlogArticleServiceImpl {
	return &BlogArticleServiceImpl{Db: Db, Config: Config, BlogArticleRepository: BlogArticleRepository, BlogArticleAttachmentRepository: BlogArticleAttachmentRepository}
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

	if itemDTO.File != nil {
		err = s.UploadAndSaveImageAttachment(tx, itemDTO, item)
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

	if itemDTO.File != nil {
		err = s.UploadAndSaveImageAttachment(tx, itemDTO, item)
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

func (s *BlogArticleServiceImpl) UploadAndSaveImageAttachment(tx *gorm.DB, itemDTO *dto.BlogArticleDTO, blogArticle model.BlogArticle) error {
	uploadImageResult, err := s.UploadImage(itemDTO, blogArticle)

	if err != nil {
		return err
	}

	err = s.SaveAttachment(tx, uploadImageResult, blogArticle)

	if err != nil {
		return err
	}

	return nil
}

func (s *BlogArticleServiceImpl) SaveAttachment(tx *gorm.DB, itemDTO *dto.UploadImageResult, blogArticle model.BlogArticle) error {
	blogArticleAttachment, err := s.BlogArticleAttachmentRepository.Find(tx, "PRIMARY_IMAGES")

	if err != nil {
		return err
	}

	if blogArticleAttachment != nil {
		s.DeleteImage(fmt.Sprintf("%s/%s", blogArticleAttachment.Path.String, blogArticleAttachment.FileName.String))

		updateBlogArticleAttachment := model.BlogArticleAttachment{
			FileName: null.NewString(itemDTO.FileName, true),
			Path:     null.NewString(itemDTO.FilePath, true),
			MimeType: null.NewString(itemDTO.MimeType, true),
			Url:      null.NewString(itemDTO.UrlPath, true),
		}

		err = s.BlogArticleAttachmentRepository.Update(tx, &updateBlogArticleAttachment)

		if err != nil {
			return err
		}
	} else {
		newBlogArticleAttachment := model.BlogArticleAttachment{
			Name:           "PRIMARY_IMAGES",
			FileName:       null.NewString(itemDTO.FileName, true),
			Path:           null.NewString(itemDTO.FilePath, true),
			MimeType:       null.NewString(itemDTO.MimeType, true),
			Url:            null.NewString(itemDTO.UrlPath, true),
			BlogArticlesId: blogArticle.ID,
		}

		err = s.BlogArticleAttachmentRepository.Create(tx, &newBlogArticleAttachment)

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *BlogArticleServiceImpl) UploadImage(itemDTO *dto.BlogArticleDTO, blogArticle model.BlogArticle) (*dto.UploadImageResult, error) {
	openFile, err := itemDTO.File.Open()

	if err != nil {
		return nil, err
	}

	defer openFile.Close()

	byteFile, err := io.ReadAll(openFile)

	if err != nil {
		return nil, err
	}

	mtype := mimetype.Detect(byteFile)

	filePath := "upload/blog/banners"

	extension := mtype.Extension()

	fileName := fmt.Sprintf("%s_%s%s", uuid.New(), blogArticle.ID, extension)

	fieldData := map[string]string{
		"filename": fileName,
		"folder":   filePath,
		"isPublic": "true",
	}

	var requestBody bytes.Buffer

	multiPartWriter := multipart.NewWriter(&requestBody)

	// Populate File
	fileWriter, err := multiPartWriter.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, openFile)
	if err != nil {
		return nil, err
	}
	// End Populate File

	// Populate Other Field

	for k, v := range fieldData {
		fieldWriter, err := multiPartWriter.CreateFormField(k)
		if err != nil {
			return nil, err
		}

		_, err = fieldWriter.Write([]byte(v))
		if err != nil {
			return nil, err
		}
	}
	// End Populate

	multiPartWriter.Close()

	var result map[string]interface{}

	err = httprequest.RequestPostForm(string(http.MethodPost), fmt.Sprintf("%s/%s", s.Config.GetString("STORAGE_SERVICE_URL"), "upload"), multiPartWriter, requestBody, &result)

	if err != nil {
		return nil, err
	}

	urlPath, ok := result["result"].(string)

	if !ok {
		return nil, errors.New("Failed Upload Image")
	}
	return &dto.UploadImageResult{
		FileName:  fileName,
		Extension: extension,
		MimeType:  mtype.String(),
		FilePath:  filePath,
		UrlPath:   urlPath,
	}, nil
}

func (s *BlogArticleServiceImpl) DeleteImage(urlPath string) error {
	var result map[string]interface{}

	requestBody := map[string]interface{}{
		"fullPathObject": urlPath,
	}

	err := httprequest.RequestPost(string(http.MethodDelete), fmt.Sprintf("%s/%s", s.Config.GetString("STORAGE_SERVICE_URL"), "delete"), &requestBody, &result)

	if err != nil {
		return err
	}

	return nil
}
