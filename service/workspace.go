package service

import (
	"core/config"
	"core/domain"
	"core/models"
	"core/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type WorkspaceService struct {
	WorkspaceDomain       domain.WorkspaceDomain
	ManageWorkspaceDomain domain.ManageWorkspaceDomain
	ChannelDomain         domain.ChannelDomain
	DocumentDomain        domain.DocumentDomain
	CredentialsDomain     domain.CredentialsDomain
	ManageChannelsDomain  domain.ManageChannelsDomain
	UserDomain            domain.UserDomain
}

func (c *WorkspaceService) CreateWorkspace(param models.CreateWorkspaceReqs) (models.CreateWorkspaceResp, error) {
	workspaceParam := models.Workspace{
		Name:    param.Name,
		OwnerID: param.UserID,
	}
	workspace, err := c.WorkspaceDomain.Create(workspaceParam)
	if err != nil {
		return models.CreateWorkspaceResp{}, err
	}
	manageWorkspaceParam := models.ManageWorkspace{
		WorkspaceID:  workspace.ID,
		JoinedUserID: param.UserID,
		Role:         "admin",
		IsAccepted:   true,
	}
	_, err = c.ManageWorkspaceDomain.Create(manageWorkspaceParam)
	if err != nil {
		return models.CreateWorkspaceResp{}, err
	}

	channelParam := models.Channels{
		Name:    "general",
		Type:    true,
		OwnerID: param.UserID,
	}
	channels, err := c.ChannelDomain.Create(channelParam)
	if err != nil {
		return models.CreateWorkspaceResp{}, err
	}
	// docParam := models.Document{
	// 	ChannelID: channels.ID,
	// 	Name:      "This is general channel",
	// }
	// _, err = c.DocumentDomain.Create(docParam)
	// if err != nil {
	// 	return models.CreateWorkspaceResp{}, err
	// }
	// _, err = c.CredentialsDomain.Create(models.Credentials{
	// 	Text:      "This is general channel",
	// 	ChannelID: channels.ID,
	// })
	// if err != nil {
	// 	return models.CreateWorkspaceResp{}, err
	// }
	manageChannelParam := models.ManageChannels{
		ChannelID:    channels.ID,
		JoinedUserID: param.UserID,
		Role:         "admin",
		WorkspaceID:  workspace.ID,
	}
	_, err = c.ManageChannelsDomain.Create(manageChannelParam)
	if err != nil {
		return models.CreateWorkspaceResp{}, err
	}
	resp := models.CreateWorkspaceResp{
		ID:          workspace.ID,
		Name:        workspace.Name,
		RedirectUrl: "/workspace/" + strconv.FormatInt(workspace.ID, 10),
		CreatedAt:   workspace.CreatedAt,
	}
	return resp, nil
}

func (c *WorkspaceService) GetWorkspaceById(param models.GetWorkspaceByIdReqs) (models.GetWorkspaceByIdResp, error) {
	// do you belongs to this workspace
	ManageWorkspace := models.ManageWorkspace{
		WorkspaceID:  param.WorkspaceID,
		JoinedUserID: param.UserID,
	}
	_, err := c.ManageWorkspaceDomain.GetByWorkspaceIdAndUserId(ManageWorkspace)
	if err != nil {
		return models.GetWorkspaceByIdResp{}, fmt.Errorf("you don't belong to this workspace: %w", err)
	}
	workspaceParam := models.Workspace{
		ID: param.WorkspaceID,
	}
	workspace, err := c.WorkspaceDomain.GetById(workspaceParam)
	if err != nil {
		return models.GetWorkspaceByIdResp{}, err
	}
	fmt.Println("workspace-0090991211", workspace)
	manageChannelParam := models.ManageChannels{
		WorkspaceID:  param.WorkspaceID,
		JoinedUserID: param.UserID,
	}
	managechannelData, err := c.ManageChannelsDomain.GetByWorkspaceIdAndUser(manageChannelParam)
	if err != nil {
		return models.GetWorkspaceByIdResp{}, fmt.Errorf("you don't belong to this workspace: %w", err)
	}
	resp := models.GetWorkspaceByIdResp{
		WorkspaceID: workspace.ID,
		Workspace:   workspace.Name, // assuming your model has Name
		UserID:      param.UserID,
		Channels:    managechannelData,
	}

	return resp, nil

}

func (c *WorkspaceService) AddUserInWorkspace(param models.AddUserInWorkspaceReqs) (models.BasicResp, error) {
	manageWorkSpaceParam := models.ManageWorkspace{
		WorkspaceID:  param.WorkspaceID,
		JoinedUserID: param.AddedByID,
	}
	userWhoIsAddingRole, err := c.ManageWorkspaceDomain.GetByWorkspaceIdAndUser(manageWorkSpaceParam)
	if err != nil {
		return models.BasicResp{}, fmt.Errorf("failed to get owner id: %w", err)
	}
	fmt.Println("userWhoIsAddingRole", userWhoIsAddingRole)
	if userWhoIsAddingRole.Role != "admin" {
		return models.BasicResp{}, fmt.Errorf("only admin can add user in workspace")
	}
	count, err := c.ManageWorkspaceDomain.CountOfUsersInWorkspace(models.ManageWorkspace{WorkspaceID: param.WorkspaceID})
	if err != nil {
		return models.BasicResp{}, err
	}
	if count >= int64(utils.CountOfUsersInWorkspace) {
		return models.BasicResp{}, fmt.Errorf("you can't add more users in this workspace")
	}
	userParam := models.User{Email: param.UserEmail}
	user, err := c.UserDomain.GetUserByEmail(userParam)
	if err != nil {
		return models.BasicResp{}, err
	}
	manageWorkspaceParam := models.ManageWorkspace{
		WorkspaceID:  param.WorkspaceID,
		JoinedUserID: user.ID,
		Role:         param.Role,
		IsAccepted:   false,
	}
	_, err = c.ManageWorkspaceDomain.Create(manageWorkspaceParam)
	if err != nil {
		return models.BasicResp{}, err
	}
	// ✅ generate JWT token
	token, err := generateUserToken(int64(user.ID), user.Email)
	if err != nil {
		return models.BasicResp{}, err
	}

	// ✅ invite link
	inviteLink := fmt.Sprintf("%s/invite?jwt=%s&workspaceId=%d", config.GetConfig().FrontendUrl, token, param.WorkspaceID)

	// ✅ prepare email data
	mailData := models.SendMail{
		SendTo: user.Email,
		Data: map[string]interface{}{
			"InviteLink": inviteLink,
			"Workspace":  param.WorkspaceID,
		},
	}

	// ✅ send email with template
	go utils.SendMailForInvite("./template/invite.html", mailData, "You’ve been invited to join a workspace")

	// ✅ return only success message
	return models.BasicResp{
		Message: "Invitation sent successfully",
	}, nil
}

func generateUserToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // 1 week expiry
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().JWTSecret))
}

func (c *WorkspaceService) AcceptInvite(param models.AcceptInviteReqs) (models.BasicResp, error) {
	manageWorkspaceParam := models.ManageWorkspace{
		WorkspaceID:  param.UserID,
		JoinedUserID: param.UserID,
	}
	manageWorkspace, err := c.ManageWorkspaceDomain.GetByWorkspaceIdAndUserId(manageWorkspaceParam)
	if err != nil {
		return models.BasicResp{}, fmt.Errorf("no invitation found for this user in this workspace: %w", err)
	}
	if manageWorkspace.IsAccepted {
		return models.BasicResp{}, fmt.Errorf("you have already accepted the invitation")
	}
	manageWorkspaceParam.IsAccepted = true
	err = c.ManageWorkspaceDomain.UpdateIsAccepted(manageWorkspaceParam)
	if err != nil {
		return models.BasicResp{}, err
	}
	getPublicChannelParam := models.ManageChannels{
		WorkspaceID: param.WorkspaceID,
	}
	var publicChannels []models.ManageChannels

	publicChannels, err = c.ManageChannelsDomain.GetPublicChannelByWorkspaceId(getPublicChannelParam)
	if err != nil {
		return models.BasicResp{}, err
	}
	fmt.Println("publicChannels", publicChannels)
	for _, channel := range publicChannels {
		manageChannelParam := models.ManageChannels{
			ChannelID:    channel.ChannelID,
			JoinedUserID: param.UserID,
			Role:         "member",
			WorkspaceID:  param.WorkspaceID,
		}
		_, err = c.ManageChannelsDomain.Create(manageChannelParam)
		if err != nil {
			return models.BasicResp{}, err
		}
	}
	return models.BasicResp{
		Message: "Invitation accepted successfully",
	}, nil
}

func (c *WorkspaceService) GetAllWorkspace(userId int64) ([]models.GetAllWorkspaceByUserIdResp, error) {
	data := c.WorkspaceDomain.GetAllWorkspaceByUserId(userId)
	if data == nil {
		return nil, fmt.Errorf("no workspaces found for user: %d", userId)
	}
	response := make([]models.GetAllWorkspaceByUserIdResp, len(data))
	for i, workspace := range data {
		response[i] = models.GetAllWorkspaceByUserIdResp(workspace)
	}
	return response, nil
}
