package models

import "time"

type Document struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	ChannelID int64     `gorm:"column:channel_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
