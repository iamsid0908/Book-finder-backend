package service

import (
	"core/domain"
	"core/models"

	"github.com/lib/pq"
)

type BillingService struct {
	BillingDomain domain.BillingDomain
}

func (c *BillingService) Insert(param models.InsertBillingParam) error {
	useParam := models.Billing{
		UserID:   param.UserID,
		Number:   param.Number,
		Amount:   param.Amount,
		Status:   param.Status,
		Payments: pq.Int64Array(param.Payment),
	}
	err := c.BillingDomain.Insert(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *BillingService) Update(param models.InsertBillingParam) error {
	useParam := models.Billing{
		UserID:   param.UserID,
		Number:   param.Number,
		Amount:   param.Amount,
		Status:   param.Status,
		Payments: pq.Int64Array(param.Payment),
	}
	err := c.BillingDomain.Update(useParam)
	if err != nil {
		return err
	}
	return nil
}

func (c *BillingService) List() ([]models.ListBillingResp, error) {
	data, err := c.BillingDomain.Get()
	if err != nil {
		return []models.ListBillingResp{}, err
	}
	response := make([]models.ListBillingResp, len(data))
	for i, resp := range data {
		response[i] = models.ListBillingResp{
			ID:        resp.ID,
			UserID:    resp.UserID,
			Number:    resp.Number,
			Amount:    resp.Amount,
			Status:    resp.Status,
			Payment:   resp.Payments,
			CreatedAt: resp.CreatedAt,
			UpdatedAt: resp.UpdatedAt,
		}
	}
	return response, nil
}
