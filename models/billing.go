package models

import (
	"time"

	"github.com/lib/pq"
)

type Billing struct {
	ID        int64         `gorm:"column:id;"`
	UserID    int64         `gorm:"column:user_id;"`
	Number    string        `gorm:"column:number;size:20;uniqueIndex;not null"`
	Amount    float64       `gorm:"column:amount"`
	Status    string        `gorm:"column:status;size:20;not null"`
	Payments  pq.Int64Array `gorm:"type:integer[];column:payment_id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InsertBillingParam struct {
	UserID  int64   `json:"user_id"`
	Number  string  `json:"number"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	Payment []int64 `json:"payment"`
}

type ListBillingResp struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Number    string    `json:"number"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Payment   []int64   `json:"payment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
