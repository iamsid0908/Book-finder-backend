package models

import "time"

type ManageWorkspace struct {
	ID           int64     `json:"id"`
	WorkspaceID  int64     `json:"workspace_id"`
	JoinedUserID int64     `json:"joined_user_id"`
	Role         string    `json:"role"`
	IsAccepted   bool      `json:"is_accepted"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
