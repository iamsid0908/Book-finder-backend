package validation

import (
	"core/models"
	"errors"
)

func CustomerInsertValidation(param *models.GetCustomerParam) error {
	if param.Phone == "" {
		return errors.New("please enter mobile number")
	}
	if param.Address == "" {
		return errors.New("please enter Address")
	}
	if param.LastOrder == "" {
		return errors.New("please enter Last Order")
	}
	return nil
}
