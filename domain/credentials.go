package domain

import (
	"core/config"
	"core/models"
)

type CredentialsDomain interface {
	Create(params models.Credentials) (models.Credentials, error)
}

type CredentialsDomainCtx struct {
}

func (c *CredentialsDomainCtx) Create(params models.Credentials) (models.Credentials, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.Credentials{}, err
	}
	return params, nil
}
