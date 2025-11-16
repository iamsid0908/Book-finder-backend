package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

type ResisterResp struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Redirect string `json:"redirect"`
}
type ResendOTPRequest struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

type VerifyOTPRequest struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LogInResponse struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token"`
	Redirect  string    `json:"redirect"`
}

type JWTPayload struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Language string `json:"language"`
	jwt.StandardClaims
}

type GoogleUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
