package domain

import (
	"core/config"
	"core/models"
)

type CustomerDomain interface {
	Insert(param models.Customer) error
	Get() ([]models.Customer, error)
	Update(param models.Customer) error
}

type CustomerDomainCtx struct{}

func (c *CustomerDomainCtx) Insert(param models.Customer) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerDomainCtx) Get() ([]models.Customer, error) {
	db := config.DbManager()
	var customers []models.Customer
	err := db.Find(&customers).Error
	if err != nil {
		return []models.Customer{}, nil
	}
	return customers, nil
}

func (c *CustomerDomainCtx) Update(param models.Customer) error {
	db := config.DbManager().Model(&models.Customer{})
	userID := param.UserID
	update := map[string]interface{}{}

	if param.Phone != "" {
		update["phone"] = param.Phone
	}
	if param.Address != "" {
		update["address"] = param.Address
	}
	if param.LastOrder != "" {
		update["last_order"] = param.LastOrder
	}

	err := db.Where("user_id = ?", userID).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
