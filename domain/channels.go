package domain

import (
	"core/config"
	"core/models"
)

type ChannelDomain interface {
	Create(params models.Channels) (models.Channels, error)
}

type ChannelDomainCtx struct {
}

func (c *ChannelDomainCtx) Create(params models.Channels) (models.Channels, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.Channels{}, err
	}
	return params, nil
}
