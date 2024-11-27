package models

import "time"

type Books struct {
	ID          int64     `gorm:"column:id;"`
	Title       string    `gorm:"title"`
	Thumbnail   string    `gorm:"thumbnail"`
	WritterName string    `gorm:"writter_name"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

type BookReqs struct {
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	WritterName string    `json:"writter_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BooksResp struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Thumbnail   string    `json:"thumbnail"`
	WritterName string    `json:"writter_name"`
	Cart        bool      `json:"cart"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookWithCart struct {
	ID          int64     `gorm:"column:id"`
	Title       string    `gorm:"column:title"`
	Thumbnail   string    `gorm:"column:thumbnail"`
	WritterName string    `gorm:"column:writter_name"`
	Cart        bool      `gorm:"column:cart"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
