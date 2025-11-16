package domain

import (
	"core/config"
	"core/models"
)

type ManageWorkspaceDomain interface {
	Create(params models.ManageWorkspace) (models.ManageWorkspace, error)
	GetByWorkspaceIdAndUserId(params models.ManageWorkspace) (models.ManageWorkspace, error)
	GetByWorkspaceIdAndUser(params models.ManageWorkspace) (models.ManageWorkspace, error)
	CountOfUsersInWorkspace(params models.ManageWorkspace) (int64, error)
	UpdateIsAccepted(params models.ManageWorkspace) error
}

type ManageWorkspaceDomainCtx struct {
}

func (c *ManageWorkspaceDomainCtx) Create(params models.ManageWorkspace) (models.ManageWorkspace, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.ManageWorkspace{}, err
	}
	return params, nil
}

func (c *ManageWorkspaceDomainCtx) GetByWorkspaceIdAndUserId(params models.ManageWorkspace) (models.ManageWorkspace, error) {
	db := config.DbManager()
	err := db.Where("workspace_id = ? AND joined_user_id = ?", params.WorkspaceID, params.JoinedUserID).First(&params)
	if err != nil {
		return models.ManageWorkspace{}, err.Error
	}

	return params, nil
}

func (c *ManageWorkspaceDomainCtx) GetByWorkspaceIdAndUser(params models.ManageWorkspace) (models.ManageWorkspace, error) {
	db := config.DbManager()

	var result models.ManageWorkspace
	err := db.Where("workspace_id = ? AND joined_user_id = ?", params.WorkspaceID, params.JoinedUserID).First(&result).Error
	if err != nil {
		return models.ManageWorkspace{}, err
	}

	return result, nil
}

func (c *ManageWorkspaceDomainCtx) CountOfUsersInWorkspace(params models.ManageWorkspace) (int64, error) {
	var count int64
	db := config.DbManager()
	err := db.Model(&models.ManageWorkspace{}).Where("workspace_id = ? AND is_accepted = ?", params.WorkspaceID, true).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (c *ManageWorkspaceDomainCtx) UpdateIsAccepted(params models.ManageWorkspace) error {
	db := config.DbManager()
	if params.IsAccepted {
		err := db.Model(&models.ManageWorkspace{}).Where("workspace_id = ? AND joined_user_id = ?", params.WorkspaceID, params.JoinedUserID).Update("is_accepted", params.IsAccepted).Error
		if err != nil {
			return err
		}
	}
	return nil
}
