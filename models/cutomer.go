package models

import (
	"time"

	"github.com/lib/pq"
)

type Customer struct {
	ID           int64         `gorm:"column:id;"`
	UserID       int64         `gorm:"column:user_id;"`
	Phone        string        `gorm:"column:phone;"`
	Address      string        `gorm:"column:address;"`
	LastOrder    string        `gorm:"column:last_order;"`
	BillingIds   pq.Int64Array `gorm:"type:integer[];column:billing_ids;"`
	LastOrderdAt time.Time     `gorm:"column:lastorder_at;"`
}

type GetCustomerParam struct {
	UserID    int64  `json:"user_id"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	LastOrder string `json:"last_order"`
}
