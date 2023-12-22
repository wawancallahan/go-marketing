package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MarketingEvent struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();<-:create"`
	EventName        string    `gorm:"column:event_name"`
	EventTime        time.Time `gorm:"column:event_time"`
	EventLocation    string    `gorm:"column:event_location"`
	EventType        string    `gorm:"column:event_type"`
	ChannelEvent     string    `gorm:"column:channel_type"`
	MeasurementEvent string    `gorm:"column:measurement_event"`
	Status           string    `gorm:"column:status"`
	Province         string    `gorm:"column:province"`
	City             string    `gorm:"column:city"`
	Participant      int64     `gorm:"column:participant"`
	PicName          string    `gorm:"column:pic_name"`
	SupportName      string    `gorm:"column:support_name"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (m *MarketingEvent) TableName() string {
	return "marketing_events"
}
