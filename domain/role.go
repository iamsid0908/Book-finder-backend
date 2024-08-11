package domain

import (
	"core/config"
	"core/models"
)

type RoleDomain interface {
	Insert(param models.Role) (models.Role, error)
	Get() ([]models.Role, error)
}

type RoleDomainCtx struct{}

func (c *RoleDomainCtx) Insert(param models.Role) (models.Role, error) {
	db := config.DbManager()

	err := db.Create(&param).Error
	if err != nil {
		return models.Role{}, err
	}
	return param, nil
}

func (c *RoleDomainCtx) Get() ([]models.Role, error) {
	db := config.DbManager()
	var roles []models.Role
	err := db.Find(&roles).Error
	if err != nil {
		return []models.Role{}, nil
	}
	return roles, nil
}
