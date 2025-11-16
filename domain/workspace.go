package domain

import (
	"core/config"
	"core/models"
)

type WorkspaceDomain interface {
	Create(params models.Workspace) (models.Workspace, error)
	GetById(params models.Workspace) (models.Workspace, error)
	GetAllWorkspaceByUserId(userId int64) []models.Workspace
}

type WorkspaceDomainCtx struct {
}

func (c *WorkspaceDomainCtx) Create(params models.Workspace) (models.Workspace, error) {
	db := config.DbManager()
	err := db.Create(&params).Error
	if err != nil {
		return models.Workspace{}, err
	}
	return params, nil
}
func (c *WorkspaceDomainCtx) GetById(params models.Workspace) (models.Workspace, error) {
	db := config.DbManager()
	var workspace models.Workspace
	result := db.Where("id = ?", params.ID).First(&workspace)
	if result.Error != nil {
		return models.Workspace{}, result.Error
	}

	return workspace, nil
}

func (c *WorkspaceDomainCtx) GetAllWorkspaceByUserId(userId int64) []models.Workspace {
	db := config.DbManager()
	var workspaces []models.Workspace
	err := db.Joins("JOIN manage_workspace ON manage_workspace.workspace_id = workspace.id").
		Where("manage_workspace.joined_user_id = ? AND manage_workspace.is_accepted = ?", userId, true).
		Find(&workspaces).Error
	if err != nil {
		return nil
	}
	return workspaces
}
