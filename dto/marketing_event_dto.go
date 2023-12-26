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

func (d *MarketingEventDTO) ToModel() model.MarketingEvent {
	eventTime, _ := time.Parse("2006-01-02T15:04:05Z07:00", d.EventTime)
	return model.MarketingEvent{
		EventName:        d.EventName,
		EventTime:        eventTime,
		EventLocation:    d.EventLocation,
		EventType:        d.EventType,
		ChannelEvent:     d.ChannelEvent,
		MeasurementEvent: d.MeasurementEvent,
		Status:           d.Status,
		Province:         d.Province,
		City:             d.City,
		Participant:      int64(d.Participant),
		PicName:          d.PicName,
		SupportName:      d.SupportName,
	}
}
