package service

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/enum"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
	"matsukana.cloud/go-marketing/util"
)

type MarketingLeadService interface {
	Index() (*[]model.MarketingLead, error)
	Create(itemDTO *dto.MarketingLeadDTO) (*model.MarketingLead, error)
	Find(id string) (*model.MarketingLead, error)
	Update(itemDTO *dto.MarketingLeadDTO, id string) (*model.MarketingLead, error)
	Delete(id string) error
	SourceType() (*[]util.ResultList, error)
	ProductCategory() (*[]util.ResultList, error)
}

type MarketingLeadServiceImpl struct {
	Db                      *database.Database
	MarketingLeadRepository repository.MarketingLeadRepository
}

func NewMarketingLeadService(Db *database.Database, MarketingLeadRepository repository.MarketingLeadRepository) *MarketingLeadServiceImpl {
	return &MarketingLeadServiceImpl{Db: Db, MarketingLeadRepository: MarketingLeadRepository}
}

func (s *MarketingLeadServiceImpl) Index() (*[]model.MarketingLead, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.MarketingLeadRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return items, nil
}

func (s *MarketingLeadServiceImpl) Create(itemDTO *dto.MarketingLeadDTO) (*model.MarketingLead, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	if itemDTO.Status.Valid {
		itemDTO.Status = null.NewString("ONGOING", true)
	}

	if itemDTO.Status.Equal(null.StringFrom("REJECTED")) || itemDTO.Status.Equal(null.StringFrom("ONGOING")) {
		itemDTO.FollowUpBy = null.NewString("", false)
	}

	item := itemDTO.ToModel()

	err := s.MarketingLeadRepository.Create(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil

}

func (s *MarketingLeadServiceImpl) Find(id string) (*model.MarketingLead, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item, err := s.MarketingLeadRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return item, nil
}

func (s *MarketingLeadServiceImpl) Update(itemDTO *dto.MarketingLeadDTO, id string) (*model.MarketingLead, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	if itemDTO.Status.Valid {
		itemDTO.Status = null.NewString("ONGOING", true)
	}

	if itemDTO.Status.Equal(null.StringFrom("REJECTED")) || itemDTO.Status.Equal(null.StringFrom("ONGOING")) {
		itemDTO.FollowUpBy = null.NewString("", false)
	}

	item := itemDTO.ToModel()
	item.ID = uuid.MustParse(id)

	err := s.MarketingLeadRepository.Update(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}

func (s *MarketingLeadServiceImpl) Delete(id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	err := s.MarketingLeadRepository.Delete(tx, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (s *MarketingLeadServiceImpl) SourceType() (*[]util.ResultList, error) {
	sourceType := enum.SourceTypeEnum{}.List()

	sourceTypeResult := make([]util.ResultList, 0)

	for k, v := range sourceType {
		sourceTypeResult = append(sourceTypeResult, util.ResultList{
			ID:   k,
			Name: v,
		})
	}

	return &sourceTypeResult, nil
}

func (s *MarketingLeadServiceImpl) ProductCategory() (*[]util.ResultList, error) {
	productCategory := enum.ProductCategoryEnum{}.List()

	productCategoryResult := make([]util.ResultList, 0)

	for k, v := range productCategory {
		productCategoryResult = append(productCategoryResult, util.ResultList{
			ID:   k,
			Name: v,
		})
	}

	return &productCategoryResult, nil
}
