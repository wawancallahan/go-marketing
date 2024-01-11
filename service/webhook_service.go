package service

import (
	"gorm.io/gorm"
	"matsukana.cloud/go-marketing/database"
	"matsukana.cloud/go-marketing/repository"
	"matsukana.cloud/go-marketing/util"
)

type WebhookService interface {
	MarketingEventStatus() (*util.ResultNone, error)
	MarketingLeadActivationStatus() (*util.ResultNone, error)
	MarketingLeadDuplicateStatus() (*util.ResultNone, error)
}

type WebhookServiceImpl struct {
	Db                       *database.Database
	MarketingEventRepository repository.MarketingEventRepository
	MarketingLeadRepository  repository.MarketingLeadRepository
}

func NewWebhookService(Db *database.Database, MarketingEventRepository repository.MarketingEventRepository, MarketingLeadRepository repository.MarketingLeadRepository) *WebhookServiceImpl {
	return &WebhookServiceImpl{Db: Db, MarketingEventRepository: MarketingEventRepository, MarketingLeadRepository: MarketingLeadRepository}
}

func (s *WebhookServiceImpl) MarketingEventStatus() (*util.ResultNone, error) {
	s.MarketingEventUpdateOngoingStatus()
	s.MarketingEventUpdateCompletedStatus()

	return &util.ResultNone{}, nil
}

func (s *WebhookServiceImpl) MarketingEventUpdateOngoingStatus() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingEventRepository.UpdateOngoingStatus(tx)

		return nil
	})

	return nil
}

func (s *WebhookServiceImpl) MarketingEventUpdateCompletedStatus() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingEventRepository.UpdateCompletedStatus(tx)

		return nil
	})

	return nil
}

func (s *WebhookServiceImpl) MarketingLeadActivationStatus() (*util.ResultNone, error) {
	s.MarketingLeadUpdatePartiallyRegistered()
	s.MarketingLeadUpdateRegistered()

	return &util.ResultNone{}, nil
}

func (s *WebhookServiceImpl) MarketingLeadUpdatePartiallyRegistered() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingLeadRepository.UpdatePartiallyRegistered(tx)

		return nil
	})

	return nil
}

func (s *WebhookServiceImpl) MarketingLeadUpdateRegistered() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingLeadRepository.UpdateRegistered(tx)

		return nil
	})

	return nil
}

func (s *WebhookServiceImpl) MarketingLeadDuplicateStatus() (*util.ResultNone, error) {
	s.MarketingLeadUpdateDuplicateEntry()
	s.MarketingLeadUpdateNonDuplicateEntry()

	return &util.ResultNone{}, nil
}

func (s *WebhookServiceImpl) MarketingLeadUpdateDuplicateEntry() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingLeadRepository.UpdateDuplicateEntry(tx)

		return nil
	})

	return nil
}

func (s *WebhookServiceImpl) MarketingLeadUpdateNonDuplicateEntry() error {
	_ = s.Db.Transaction(func(tx *gorm.DB) error {
		s.MarketingLeadRepository.UpdateNonDuplicateEntry(tx)

		return nil
	})

	return nil
}
