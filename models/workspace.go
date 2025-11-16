package models

import "time"

type Workspace struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	OwnerID   int64     `gorm:"column:owner_id"`
	Type      bool      `gorm:"column:type"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CreateWorkspaceReqs struct {
	Name   string `json:"name"`
	Type   bool   `json:"type"`
	UserID int64  `json:"user_id"`
}

type CreateWorkspaceResp struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	RedirectUrl string    `json:"redirect_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetWorkspaceByIdReqs struct {
	WorkspaceID int64 `json:"workspace_id"`
	UserID      int64 `json:"user_id"`
}

type ChannelDetailsResp struct {
	WorkspaceID   int64  `json:"workspace_id"`
	UserID        int64  `json:"user_id"`
	ChannelID     int64  `json:"channel_id"`
	ChannelName   string `json:"channel_name"`
	Type          bool   `json:"type"`
	OwnerID       int64  `json:"owner_id"`
	OwnerUsername string `json:"owner_username"`
	Role          string `json:"role"`
	DocumentID    *int64 `json:"document_id,omitempty"`
	CredentialID  *int64 `json:"credential_id,omitempty"`
}

type GetWorkspaceByIdResp struct {
	WorkspaceID int64                `json:"workspace_id"`
	Workspace   string               `json:"workspace"` // name/title of workspace
	UserID      int64                `json:"user_id"`
	Channels    []ChannelDetailsResp `json:"channels"`
}

type AddUserInWorkspaceReqs struct {
	WorkspaceID     int64  `json:"workspace_id"`
	UserEmail       string `json:"user_email"`
	Role            string `json:"role"`
	AddedByID       int64  `json:"added_by_id"`
	AddedByUserRole string `json:"added_by_user_role"`
}

type AcceptInviteReqs struct {
	UserID      int64  `json:"user_id"`
	Email       string `json:"email"`
	WorkspaceID int64  `json:"workspace_id"`
}

type GetAllWorkspaceByUserIdResp struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	OwnerID   int64     `json:"owner_id"`
	Type      bool      `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
