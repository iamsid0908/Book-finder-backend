package models

import "time"

type Cart struct {
	ID        int64     `gorm:"column:id;unique"`
	UserID    int64     `gorm:"user_id"`
	BookID    int64     `gorm:"book_id"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gotm:"updated_at"`
}

type InsertCartParam struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}

type CartResponse struct {
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	WritterName string `json:"writter_name"`
}

type CartResp struct {
	BookID      int64  `json:"book_id"`
	Title       string `json:"title"`
	Thumbnail   string `json:"thumbnail"`
	WritterName string `json:"writter_name"`
	Cart        bool   `json:"cart"`
}

type RemoveFromCartReqs struct {
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}
