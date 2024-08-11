package domain

import (
	"core/config"
	"core/models"
	"fmt"
)

type BillingDomain interface {
	Insert(param models.Billing) error
	Update(param models.Billing) error
	Get() ([]models.Billing, error)
}

type BillingDomainCtx struct{}

func (c *BillingDomainCtx) Insert(param models.Billing) error {
	db := config.DbManager()
	err := db.Create(&param).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *BillingDomainCtx) Update(param models.Billing) error {
	fmt.Println(param)
	db := config.DbManager().Model(&models.Billing{})
	update := map[string]interface{}{}

	if param.Number != "" {
		update["number"] = param.Number
	}
	if param.Amount != 0 {
		update["amount"] = param.Amount
	}
	if param.Status != "" {
		update["status"] = param.Status
	}
	if param.Payments != nil {
		update["payment_id"] = param.Payments
	}
	err := db.Where("user_id = ?", param.UserID).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *BillingDomainCtx) Get() ([]models.Billing, error) {
	db := config.DbManager()
	var customers []models.Billing
	err := db.Find(&customers).Error
	if err != nil {
		return []models.Billing{}, nil
	}
	return customers, nil
}
