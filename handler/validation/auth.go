package validation

import (
	"core/models"
	"core/utils"
)

func RegisterUser(param *models.RegisterUserRequest) error {
	if param.Email == "" {
		return utils.ErrEmptyEmail
	}

	if param.Name == "" {
		return utils.ErrEmptyName
	}
	if param.Password == "" {
		return utils.ErrEmptyPassword
	}

	return nil
}

func ResendOTP(param *models.ResendOTPRequest) error {
	if param.Email == "" {
		return utils.ErrEmptyEmail
	}
	if param.Id == 0 {
		return utils.ErrEmptyUserID
	}

	return nil
}

func VerifyOTP(param *models.VerifyOTPRequest) error {
	if param.Otp == "" {
		return utils.ErrEmptyOTP
	}
	if param.Id == 0 {
		return utils.ErrEmptyUserID
	}
	if param.Email == "" {
		return utils.ErrEmptyEmail
	}

	return nil
}
