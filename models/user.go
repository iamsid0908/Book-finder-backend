package models

import "time"

type User struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Email     string    `gorm:"column:email;unique"`
	Password  *string   `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	Role      string    `gorm:"column:role"`
	Language  string    `gorm:"column:language"`
	IsActive  bool      `gorm:"column:is_active;default:false"`
	Otp       string    `gorm:"column:otp"`
	OtpExpiry time.Time `gorm:"column:otp_expiry"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

type GetUserParam struct {
	ID    int64
	Email string
}

type UserData struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  *string   `json:"password"`
	Role      string    `json:"roles"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ListOfUser struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  *string   `json:"password"`
	Role      string    `json:"roles"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type UpdateUserParam struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Language string `json:"language"`
}

type UserDataResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"roles"`
	Language  string    `json:"language"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
