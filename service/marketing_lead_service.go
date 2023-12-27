package service

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type MarketingLeadService interface {
	Index() (*[]model.MarketingLead, error)
	Create(itemDTO *dto.MarketingLeadDTO) (*model.MarketingLead, error)
	Find(id string) (*model.MarketingLead, error)
	Update(itemDTO *dto.MarketingLeadDTO, id string) error
	Delete(id string) error
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

func (s *MarketingLeadServiceImpl) Update(itemDTO *dto.MarketingLeadDTO, id string) error {
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

	err := s.MarketingLeadRepository.Update(tx, item)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
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
