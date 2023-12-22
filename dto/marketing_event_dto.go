package dto

import (
	"encoding/json"
	"time"
)

type MarketingEventDTO struct {
	EventName        string      `json:"name" validate:"required"`
	EventTime        time.Time   `json:"eventTime" validate:"required,datetime"`
	EventLocation    string      `json:"eventLocation" validate:"required"`
	EventType        string      `json:"eventType" validate:"required"`
	ChannelEvent     string      `json:"channelEvent" validate:"required"`
	MeasurementEvent string      `json:"measurementEvent" validate:"required"`
	Status           string      `json:"status" validate:"required"`
	Province         string      `json:"province" validate:"required"`
	City             string      `json:"city" validate:"required"`
	Participant      json.Number `json:"participant" validate:"required, numeric"`
	PicName          string      `json:"picName" validate:"required"`
	SupportName      string      `json:"supportName" validate:"required"`
}
