package dto

import (
	"time"

	"matsukana.cloud/go-marketing/model"
)

type MarketingEventDTO struct {
	EventName        string `json:"eventName" validate:"required"`
	EventTime        string `json:"eventTime" validate:"required"`
	EventLocation    string `json:"eventLocation" validate:"required"`
	EventType        string `json:"eventType" validate:"required"`
	ChannelEvent     string `json:"channelEvent" validate:"required"`
	MeasurementEvent string `json:"measurementEvent" validate:"required"`
	Status           string `json:"status" validate:"required"`
	Province         string `json:"province" validate:"required"`
	City             string `json:"city" validate:"required"`
	Participant      int    `json:"participant" validate:"required,numeric"`
	PicName          string `json:"picName" validate:"required"`
	SupportName      string `json:"supportName" validate:"required"`
}

func (s *MarketingEventDTO) ToModel() model.MarketingEvent {
	eventTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", s.EventTime)
	return model.MarketingEvent{
		EventName:        s.EventName,
		EventTime:        eventTime,
		EventLocation:    s.EventLocation,
		EventType:        s.EventType,
		ChannelEvent:     s.ChannelEvent,
		MeasurementEvent: s.MeasurementEvent,
		Status:           s.Status,
		Province:         s.Province,
		City:             s.City,
		Participant:      int64(s.Participant),
		PicName:          s.PicName,
		SupportName:      s.SupportName,
	}
}
