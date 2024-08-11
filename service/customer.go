package service

import (
	"core/domain"
	"core/models"
)

type CustomerService struct {
	CustomerDomain domain.CustomerDomain
}

func (c *CustomerService) Insert(param models.GetCustomerParam) error {
	useParam := models.Customer{
		UserID:    param.UserID,
		Phone:     param.Phone,
		Address:   param.Address,
		LastOrder: param.LastOrder,
	}
	err := c.CustomerDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerService) GetAll() ([]models.Customer, error) {
	data, err := c.CustomerDomain.Get()
	if err != nil {
		return []models.Customer{}, err
	}
	return data, nil
}

func (c *CustomerService) UpdateCustomer(param models.GetCustomerParam) error {
	cutomerParam := models.Customer{
		UserID:    param.UserID,
		Phone:     param.Phone,
		Address:   param.Address,
		LastOrder: param.LastOrder,
	}
	err := c.CustomerDomain.Update(cutomerParam)
	if err != nil {
		return err
	}
	return nil
}
