package service

import (
	"core/domain"
	"core/models"
)

type PaymentService struct {
	PaymentDomain domain.PaymentDomain
}

func (c *PaymentService) Insert(param models.PaymentParam) error {
	useParam := models.Payment{
		UserID:    param.UserID,
		BillingID: param.BillingID,
		Amount:    param.Amount,
		Method:    param.Method,
	}

	err := c.PaymentDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *PaymentService) GetAll() ([]models.PaymentParam, error) {
	data, err := c.PaymentDomain.Get()
	if err != nil {
		return []models.PaymentParam{}, err
	}
	response := make([]models.PaymentParam, len(data))
	for i, resp := range data {
		response[i] = models.PaymentParam{
			UserID:    resp.UserID,
			BillingID: resp.BillingID,
			Amount:    resp.Amount,
			Method:    resp.Method,
		}
	}
	return response, nil
}

func (c *PaymentService) Update(param models.PaymentParam) error {
	useParam := models.Payment{
		UserID:    param.UserID,
		BillingID: param.BillingID,
		Amount:    param.Amount,
		Method:    param.Method,
	}

	err := c.PaymentDomain.Update(useParam)
	if err != nil {
		return err
	}
	return nil
}
