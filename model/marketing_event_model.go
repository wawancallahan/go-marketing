package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketingEvent struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create" json:"id"`
	EventName        string    `gorm:"column:event_name" json:"eventName"`
	EventTime        time.Time `gorm:"column:event_time" json:"eventTime"`
	EventLocation    string    `gorm:"column:event_location" json:"eventLocation"`
	EventType        string    `gorm:"column:event_type" json:"eventType"`
	ChannelEvent     string    `gorm:"column:channel_event" json:"channelEvent"`
	MeasurementEvent string    `gorm:"column:measurement_event" json:"measurementEvent"`
	Status           string    `gorm:"column:status" json:"status"`
	Province         string    `gorm:"column:province" json:"province"`
	City             string    `gorm:"column:city" json:"city"`
	Participant      int64     `gorm:"column:participant" json:"participant"`
	PicName          string    `gorm:"column:pic_name" json:"picName"`
	SupportName      string    `gorm:"column:support_name" json:"supportName"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updatedAt"`
}

func (m *MarketingEvent) TableName() string {
	return "marketing_events"
}
