package models

import "time"

type ManageChannels struct {
	ID           int64     `gorm:"column:id;primary_key"`
	ChannelID    int64     `gorm:"column:channel_id"`
	JoinedUserID int64     `json:"joined_user_id"`
	Role         string    `gorm:"column:role"`
	WorkspaceID  int64     `gorm:"column:workspace_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
