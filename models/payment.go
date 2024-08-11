package models

import "time"

type Payment struct {
	ID        int64   `gorm:"column:id;"`
	UserID    int64   `gorm:"column:user_id;"`
	BillingID int64   `gorm:"column:billing_id"`
	Amount    float64 `gorm:"column:amount"`
	Method    string  `gorm:"column:method"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentParam struct {
	UserID    int64   `json:"user_id"`
	BillingID int64   `json:"billing_id"`
	Amount    float64 `json:"amount"`
	Method    string  `json:"method"`
}
