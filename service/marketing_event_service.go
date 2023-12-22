package service

import (
	"github.com/google/uuid"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/dto"
	"matsukana.cloud/go-marketing/model"
	"matsukana.cloud/go-marketing/repository"
)

type MarketingEventService struct {
	Db                       *database.Database
	MarketingEventRepository *repository.MarketingEventRepository
}

func NewMarketingEventService(Db *database.Database, MarketingEventRepository *repository.MarketingEventRepository) interface{} {
	return &MarketingEventService{Db: Db, MarketingEventRepository: MarketingEventRepository}
}

func (s *MarketingEventService) Index() (*[]model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	items, err := s.MarketingEventRepository.FindAll(tx)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &items, nil
}

func (s *MarketingEventService) Create(itemDTO *dto.MarketingEventDTO) (model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	participant, _ := itemDTO.Participant.Int64()

	item := model.MarketingEvent{
		EventName:        itemDTO.EventName,
		EventTime:        itemDTO.EventTime,
		EventLocation:    itemDTO.EventLocation,
		EventType:        itemDTO.EventType,
		ChannelEvent:     itemDTO.ChannelEvent,
		MeasurementEvent: itemDTO.MeasurementEvent,
		Status:           itemDTO.Status,
		Province:         itemDTO.Province,
		City:             itemDTO.City,
		Participant:      participant,
		PicName:          itemDTO.PicName,
		SupportName:      itemDTO.SupportName,
	}

	err := s.MarketingEventRepository.Create(tx, &item)

	if err != nil {
		return item, err
	}

	tx.Commit()

	return item, nil

}

func (s *MarketingEventService) Find(id string) (*model.MarketingEvent, error) {
	tx := s.Db.Begin()

	defer tx.Rollback()

	item, err := s.MarketingEventRepository.Find(tx, id)

	if err != nil {
		return nil, err
	}

	tx.Commit()

	return &item, nil
}

func (s *MarketingEventService) Update(dto *dto.MarketingEventDTO, id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	participant, _ := dto.Participant.Int64()

	item := model.MarketingEvent{
		ID:               uuid.MustParse(id),
		EventName:        dto.EventName,
		EventTime:        dto.EventTime,
		EventLocation:    dto.EventLocation,
		EventType:        dto.EventType,
		ChannelEvent:     dto.ChannelEvent,
		MeasurementEvent: dto.MeasurementEvent,
		Status:           dto.Status,
		Province:         dto.Province,
		City:             dto.City,
		Participant:      participant,
		PicName:          dto.PicName,
		SupportName:      dto.SupportName,
	}

	err := s.MarketingEventRepository.Update(tx, item)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (s *MarketingEventService) Delete(id string) error {
	tx := s.Db.Begin()

	defer tx.Rollback()

	err := s.MarketingEventRepository.Delete(tx, id)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
