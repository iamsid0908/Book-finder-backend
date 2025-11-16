package models

import "time"

type Channels struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	OwnerID   int64     `gorm:"column:owner_id"`
	Type      bool      `gorm:"column:type"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CreateChannelReqs struct {
	Name        string `json:"name"`
	Type        bool   `json:"type"`
	OwnerID     int64  `json:"owner_id"`
	WorkspaceID int64  `json:"workspace_id"`
}

type AddUserInChannelReqs struct {
	ChannelID   int64  `json:"channel_id"`
	WorkspaceID int64  `json:"workspace_id"`
	UserID      int64  `json:"user_id"`
	UserEmail   string `json:"user_email"`
	Role        string `json:"role"`
}

type AddUserInChannelResp struct {
	ID           int64     `gorm:"column:id;primary_key"`
	ChannelID    int64     `gorm:"column:channel_id"`
	JoinedUserID int64     `json:"joined_user_id"`
	Role         string    `gorm:"column:role"`
	WorkspaceID  int64     `gorm:"column:workspace_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}
