package models

import "time"

type Credentials struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Text      string    `gorm:"column:text"`
	Keys      string    `gorm:"column:keys"`
	Values    string    `gorm:"column:values"`
	ChannelID int64     `gorm:"column:channel_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
