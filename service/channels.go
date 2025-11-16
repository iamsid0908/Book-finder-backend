package service

import (
	"core/domain"
	"core/models"
	"core/utils"
	"fmt"
)

type ChannelService struct {
	ChannelDomain         domain.ChannelDomain
	ManageChannelsDomain  domain.ManageChannelsDomain
	UserDomain            domain.UserDomain
	ManageWorkspaceDomain domain.ManageWorkspaceDomain
}

func (c *ChannelService) CreateChannel(param models.CreateChannelReqs) (models.Channels, error) {
	manageChannelParam := models.ManageChannels{
		WorkspaceID: param.WorkspaceID,
	}
	count, err := c.ManageChannelsDomain.GetCountOfChannels(manageChannelParam)
	if err != nil {
		return models.Channels{}, err
	}
	if count >= int64(utils.CountOfChannelsInWorkspace) {
		return models.Channels{}, fmt.Errorf("you can't create more channels in this workspace")
	}
	channelParam := models.Channels{
		Name:    param.Name,
		Type:    param.Type,
		OwnerID: param.OwnerID,
	}
	channel, err := c.ChannelDomain.Create(channelParam)
	if err != nil {
		return models.Channels{}, err
	}
	manageChannelsParam := models.ManageChannels{
		WorkspaceID:  param.WorkspaceID,
		ChannelID:    channel.ID,
		JoinedUserID: param.OwnerID,
		Role:         "admin",
	}
	_, err = c.ManageChannelsDomain.Create(manageChannelsParam)
	if err != nil {
		return models.Channels{}, err
	}
	return channel, nil
}

func (c *ChannelService) AddUserInChannel(param models.AddUserInChannelReqs) (models.AddUserInChannelResp, error) {
	//TODO CHECK IF USER ALREADY IN THAT CHANNEL
	manageChannelsParam := models.ManageChannels{
		ChannelID: param.ChannelID,
	}
	count, err := c.ManageChannelsDomain.GetCountOfUsersInChannel(manageChannelsParam)
	if err != nil {
		return models.AddUserInChannelResp{}, err
	}
	if count >= int64(utils.CountOfUsersInChannel) {
		return models.AddUserInChannelResp{}, fmt.Errorf("you can't add more users in this channel")
	}
	userParam := models.User{Email: param.UserEmail}
	user, err := c.UserDomain.GetUserByEmail(userParam)
	if err != nil {
		return models.AddUserInChannelResp{}, err
	}
	workspaceParams := models.ManageWorkspace{
		WorkspaceID:  param.WorkspaceID,
		JoinedUserID: user.ID,
	}
	_, err = c.ManageWorkspaceDomain.GetByWorkspaceIdAndUser(workspaceParams)
	if err != nil {
		return models.AddUserInChannelResp{}, fmt.Errorf("user is not part of this workspace")
	}
	manageChannelsParam = models.ManageChannels{
		ChannelID:    param.ChannelID,
		JoinedUserID: user.ID,
		Role:         param.Role,
		WorkspaceID:  param.WorkspaceID,
	}
	ManageChannelData, err := c.ManageChannelsDomain.Create(manageChannelsParam)
	if err != nil {
		return models.AddUserInChannelResp{}, err
	}
	resp := models.AddUserInChannelResp{
		ID:           ManageChannelData.ID,
		ChannelID:    ManageChannelData.ChannelID,
		JoinedUserID: ManageChannelData.JoinedUserID,
		Role:         ManageChannelData.Role,
		WorkspaceID:  ManageChannelData.WorkspaceID,
		CreatedAt:    ManageChannelData.CreatedAt,
		UpdatedAt:    ManageChannelData.UpdatedAt,
	}
	return resp, nil
}
