package service

import (
	"time"

	"github.com/google/uuid"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type MarketingEventService interface {
	Index() (*[]model.MarketingEvent, error)
	Create(itemDTO *dto.MarketingEventDTO) (*model.MarketingEvent, error)
	Find(id string) (*model.MarketingEvent, error)
	Update(itemDTO *dto.MarketingEventDTO, id string) (*model.MarketingEvent, error)
	Delete(id string) error
}

type MarketingEventServiceImpl struct {
	Db                       *database.Database
	MarketingEventRepository repository.MarketingEventRepository
}

func NewMarketingEventService(Db *database.Database, MarketingEventRepository repository.MarketingEventRepository) *MarketingEventServiceImpl {
	return &MarketingEventServiceImpl{Db: Db, MarketingEventRepository: MarketingEventRepository}
}

func (s *MarketingEventServiceImpl) Index() (*[]model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.MarketingEventRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return items, nil
}

func (s *MarketingEventServiceImpl) Create(itemDTO *dto.MarketingEventDTO) (*model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item := itemDTO.ToModel()

	tz, _ := time.LoadLocation("Asia/Jakarta")

	now := time.Now().In(tz)

	diffDays := item.EventTime.Sub(now).Hours() / 24

	if diffDays > 0 {
		item.Status = "UPCOMING"
	} else if diffDays < 0 {
		item.Status = "COMPLETED"
	} else {
		item.Status = "ONGOING"
	}

	err := s.MarketingEventRepository.Create(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil

}

func (s *MarketingEventServiceImpl) Find(id string) (*model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item, err := s.MarketingEventRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return item, nil
}

func (s *MarketingEventServiceImpl) Update(itemDTO *dto.MarketingEventDTO, id string) (*model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item := itemDTO.ToModel()
	item.ID = uuid.MustParse(id)

	tz, _ := time.LoadLocation("Asia/Jakarta")

	now := time.Now().In(tz)

	diffDays := item.EventTime.Sub(now).Hours() / 24

	if diffDays > 0 {
		item.Status = "UPCOMING"
	} else if diffDays < 0 {
		item.Status = "COMPLETED"
	} else {
		item.Status = "ONGOING"
	}

	err := s.MarketingEventRepository.Update(tx, &item)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}

func (s *MarketingEventServiceImpl) Delete(id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	err := s.MarketingEventRepository.Delete(tx, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
