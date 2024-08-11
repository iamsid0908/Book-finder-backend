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
	if param.Role == 0 {
		return utils.ErrEmptyRole
	}

	return nil
}
