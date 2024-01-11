package service

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"matsukana.cloud/go-marketing/config"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	httprequest "matsukana.cloud/go-marketing/http_request"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type BlogBannerService interface {
	Index() (*[]model.BlogBanner, error)
	Update(itemDTO dto.BlogBannerUpdateDTO, id string) (*model.BlogBanner, error)
	UploadImage(itemDTO dto.BlogBannerUpdateDTO, id string) (*dto.UploadImageResult, error)
	DeleteImage(urlPath string) error
}

type BlogBannerServiceImpl struct {
	Db                   *database.Database
	Config               *config.Config
	BlogBannerRepository repository.BlogBannerRepository
}

func NewBlogBannerService(Db *database.Database, Config *config.Config, BlogBannerRepository repository.BlogBannerRepository) *BlogBannerServiceImpl {
	return &BlogBannerServiceImpl{Db: Db, Config: Config, BlogBannerRepository: BlogBannerRepository}
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

	uploadImageItem, err := s.UploadImage(itemDTO, id)

	if err != nil {
		return nil, err
	}

	if blogBannerItem.FileName.Valid && blogBannerItem.FileName.String != "" {
		err = s.DeleteImage(fmt.Sprintf("%s/%s", blogBannerItem.Path.String, blogBannerItem.FileName.String))

		log.Println(err)
	}

	item := model.BlogBanner{
		ID:   uuid.MustParse(id),
		Name: blogBannerItem.Name,
		FileName: sql.NullString{
			String: uploadImageItem.FileName,
			Valid:  true,
		},
		MimeType: sql.NullString{
			String: uploadImageItem.MimeType,
			Valid:  true,
		},
		Path: sql.NullString{
			String: uploadImageItem.FilePath,
			Valid:  true,
		},
		Url: sql.NullString{
			String: uploadImageItem.UrlPath,
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

func (s *BlogBannerServiceImpl) UploadImage(itemDTO dto.BlogBannerUpdateDTO, id string) (*dto.UploadImageResult, error) {
	openFile, err := itemDTO.File.Open()

	if err != nil {
		return nil, errors.New("File Cannot Open")
	}

	defer openFile.Close()

	byteFile, err := io.ReadAll(openFile)

	if err != nil {
		return nil, errors.New("File Cannot be Read")
	}

	mtype := mimetype.Detect(byteFile)

	filePath := "upload/blog/banners"

	extension := mtype.Extension()

	fileName := fmt.Sprintf("%s_%s%s", uuid.New(), id, extension)

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

func (s *BlogBannerServiceImpl) DeleteImage(urlPath string) error {
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
