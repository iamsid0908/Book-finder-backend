package domain

import (
	"core/config"
	"core/models"
)

type ManageChannelsDomain interface {
	Create(params models.ManageChannels) (models.ManageChannels, error)
	GetByWorkspaceIdAndUser(params models.ManageChannels) ([]models.ChannelDetailsResp, error)
	GetCountOfChannels(params models.ManageChannels) (int64, error)
	GetCountOfUsersInChannel(param models.ManageChannels) (int64, error)
	GetPublicChannelByWorkspaceId(param models.ManageChannels) ([]models.ManageChannels, error)
}

type ManageChannelsDomainCtx struct {
}

func (c *ManageChannelsDomainCtx) Create(params models.ManageChannels) (models.ManageChannels, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.ManageChannels{}, err
	}
	return params, nil
}

func (c *ManageChannelsDomainCtx) GetByWorkspaceIdAndUser(params models.ManageChannels) ([]models.ChannelDetailsResp, error) {
	var results []models.ChannelDetailsResp
	db := config.DbManager()
	err := db.Raw(`
	SELECT
		mc.workspace_id,
		mc.joined_user_id AS user_id,
		c.id AS channel_id,
		c.name AS channel_name,
		c.type,
		c.owner_id,
		u.name AS owner_name,
		mc.role,
		d.id AS document_id,
		cr.id AS credential_id
	FROM manage_channels mc
	JOIN channels c ON mc.channel_id = c.id
	JOIN users u ON c.owner_id = u.id
	LEFT JOIN document d ON d.channel_id = c.id
	LEFT JOIN credentials cr ON cr.channel_id = c.id
	WHERE mc.workspace_id = ? AND mc.joined_user_id = ?`,
		params.WorkspaceID, params.JoinedUserID,
	).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c *ManageChannelsDomainCtx) GetCountOfChannels(params models.ManageChannels) (int64, error) {
	var count int64
	db := config.DbManager()
	err := db.Model(&models.ManageChannels{}).Where("workspace_id = ?", params.WorkspaceID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *ManageChannelsDomainCtx) GetCountOfUsersInChannel(param models.ManageChannels) (int64, error) {
	var count int64
	db := config.DbManager()
	err := db.Model(&models.ManageChannels{}).Where("channel_id = ?", param.ChannelID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *ManageChannelsDomainCtx) GetPublicChannelByWorkspaceId(param models.ManageChannels) ([]models.ManageChannels, error) {
	var results []models.ManageChannels
	db := config.DbManager()
	err := db.Raw(`
	SELECT
		mc.id,
		mc.channel_id,
		mc.joined_user_id,
		mc.role,
		mc.workspace_id,
		mc.created_at,
		mc.updated_at
	FROM manage_channels mc
	JOIN channels c ON mc.channel_id = c.id
	WHERE mc.workspace_id = ? AND c.type = true`, param.WorkspaceID).Scan(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
