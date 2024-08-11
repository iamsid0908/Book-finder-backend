package domain

import (
	"core/config"
	"core/models"
)

type PaymentDomain interface {
	Insert(param models.Payment) error
	Update(param models.Payment) error
	Get() ([]models.Payment, error)
}

type PaymentDomainCtx struct{}

func (c *PaymentDomainCtx) Insert(param models.Payment) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *PaymentDomainCtx) Get() ([]models.Payment, error) {
	db := config.DbManager()
	var payment []models.Payment
	err := db.Find(&payment).Error
	if err != nil {
		return []models.Payment{}, nil
	}
	return payment, nil
}

func (c *PaymentDomainCtx) Update(param models.Payment) error {
	db := config.DbManager().Model(&models.Payment{})
	userID := param.UserID
	update := map[string]interface{}{}
	if param.UserID != 0 {
		update["user_id"] = userID
	}
	if param.Amount != 0 {
		update["amount"] = param.Amount
	}
	if param.BillingID != 0 {
		update["billing_id"] = param.BillingID
	}
	err := db.Where("user_id = ?", userID).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
