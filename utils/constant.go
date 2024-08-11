package utils

import "errors"

var (
	Success          = "Success"
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyName     = errors.New("name cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrEmptyRole     = errors.New("role cannot be empty")
	ErrEmailExist    = errors.New("email already exist")

	ErrUserNotExist     = errors.New("user is not exist")
	ErrPasswordNotExist = errors.New("password is not exist")
	ErrWrongPassword    = errors.New("password is incorrect")
)

// user
var (
	UserLanguageEn = "en"
)

// auth
var (
	UnexpectedSigning    = "unexpected signing method: %v"
	EmptyAuth            = "authorization is empty"
	ErrUserTokenNotExist = errors.New("user token is not exist")
	ErrWrongPerson       = "You are not allowed to access this!!!"
)
