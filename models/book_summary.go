package models

import "time"

type BookSummary struct {
	ID            int64     `gorm:"column:id;primaryKey"`
	BookID        int64     `gorm:"column:book_id;unique"`
	Thumbnail     string    `gorm:"thumbnail"`
	AuthorDetails string    `gorm:"author_details"`
	Summary       string    `gorm:"summary"`
	PublishedDate time.Time `gorm:"published_date"`
	CreatedAt     time.Time `gorm:"created_at"`
	UpdatedAt     time.Time `gotm:"updated_at"`
}

type BookSummaryReqs struct {
	BookID        int64     `json:"book_id"`
	Thumbnail     string    `json:"thumbnil"`
	AuthorDetails string    `json:"author_details"`
	Summary       string    `json:"summary"`
	PublishedDate time.Time `json:"published_date"`
}

type GetBookSummaryDetailsReqs struct {
	BookID int64 `query:"book_id"`
}

type GetBookSummaryDetailsResp struct {
	Thumbnail     string    `json:"thumbnil"`
	AuthorDetails string    `json:"author_details"`
	Summary       string    `json:"summary"`
	PublishedDate time.Time `json:"published_date"`
	Title         string    `json:"title"`
	WritterName   string    `json:"writter_name"`
	Category      string    `json:"category"`
}

type GetBookSummaryDetailsResponse struct {
	BookID        int64     `json:"book_id"`
	Thumbnail     string    `json:"thumbnil"`
	AuthorDetails string    `json:"author_details"`
	Summary       string    `json:"summary"`
	PublishedDate time.Time `json:"published_date"`
	Title         string    `json:"title"`
	WritterName   string    `json:"writter_name"`
	Category      string    `json:"category"`
}
