package domain

import (
	"core/config"
	"core/models"
)

type DocumentDomain interface {
	Create(params models.Document) (models.Document, error)
}

type DocumentDomainCtx struct {
}

func (c *DocumentDomainCtx) Create(params models.Document) (models.Document, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.Document{}, err
	}
	return params, nil
}
