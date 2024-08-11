package models

type Role struct {
	ID   int64  `gorm:"column:id;primary_key"`
	Role string `gorm:"column:role;"`
}

type RoleReqs struct {
	Role string `json:"role"`
}

type RoleResp struct {
	RoleId int64  `json:"role_id"`
	Role   string `json:"role"`
}
