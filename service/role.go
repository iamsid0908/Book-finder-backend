package service

import (
	"core/domain"
	"core/models"
)

type RoleService struct {
	RoleDomain domain.RoleDomain
}

func (c *RoleService) Insert(param models.RoleReqs) (models.Role, error) {
	useParam := models.Role{
		Role: param.Role,
	}
	data, err := c.RoleDomain.Insert(useParam)
	if err != nil {
		return models.Role{}, err
	}
	return data, nil
}

func (c *RoleService) FindAll() ([]models.Role, error) {
	data, err := c.RoleDomain.Get()
	if err != nil {
		return []models.Role{}, err
	}
	return data, nil
}
