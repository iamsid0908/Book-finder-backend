package utils

import "errors"

var (
	Success          = "Success"
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyName     = errors.New("name cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrEmptyRole     = errors.New("role cannot be empty")
	ErrEmailExist    = errors.New("email already exist")
	ErrEmptyOTP      = errors.New("otp cannot be empty")
	ErrEmptyUserID   = errors.New("user id cannot be empty")
	ErrWrongOTP      = errors.New("otp is incorrect or expired")

	ErrUserNotExist     = errors.New("user is not exist")
	ErrPasswordNotExist = errors.New("password is not exist")
	ErrWrongPassword    = errors.New("password is incorrect")
	ErrUserNotActive    = errors.New("user is not active, please verify your email")
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

var (
	AuthTypeRegister      = "register"
	AuthTypeResetPassword = "reset password"
)

// count of channels,workspace,users
var (
	CountOfChannelsInWorkspace = 10
	CountOfUsersInWorkspace    = 30
	CountOfUsersInChannel      = 30
)
