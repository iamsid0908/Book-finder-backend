package handler

import (
	"core/config"
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type WorkspaceHandler struct {
	WorkspaceService service.WorkspaceService
}

func (workspaceHandler *WorkspaceHandler) CreateWorkspace(c echo.Context) error {
	param := models.CreateWorkspaceReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	userId := c.Get("id").(int64)
	param.UserID = userId
	if param.Name == "" {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: "workspace name is required"})
	}
	if param.UserID == 0 {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: "invalid user id"})
	}
	data, err := workspaceHandler.WorkspaceService.CreateWorkspace(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (workspaceHandler *WorkspaceHandler) GetWorkspaceById(c echo.Context) error {
	param := models.GetWorkspaceByIdReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	userId := c.Get("id").(int64)
	fmt.Println(userId)
	param.UserID = userId
	data, err := workspaceHandler.WorkspaceService.GetWorkspaceById(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)

}

func (workspaceHandler *WorkspaceHandler) AddUserInWorkspace(c echo.Context) error {
	param := models.AddUserInWorkspaceReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	userId := c.Get("id").(int64)
	role := c.Get("role").(string)

	fmt.Println(userId)
	param.AddedByID = userId
	param.AddedByUserRole = role
	data, err := workspaceHandler.WorkspaceService.AddUserInWorkspace(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (workspaceHandler *WorkspaceHandler) AcceptInvite(c echo.Context) error {
	param := models.AcceptInviteReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	userID, email, err := ExtractInviteToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.BasicResp{Message: err.Error()})
	}
	param.UserID = userID
	param.Email = email
	fmt.Println("param--", param)
	data, err := workspaceHandler.WorkspaceService.AcceptInvite(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func ExtractInviteToken(c echo.Context) (int64, string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0, "", fmt.Errorf("missing authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return 0, "", fmt.Errorf("invalid authorization header")
	}

	tokenStr := parts[1]

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil || !parsedToken.Valid {
		return 0, "", fmt.Errorf("invalid or expired token")
	}

	userID := int64(claims["user_id"].(float64))
	email := claims["email"].(string)

	return userID, email, nil
}

func (workspaceHandler *WorkspaceHandler) GetAllWorkspace(c echo.Context) error {
	userId := c.Get("id").(int64)
	fmt.Println(userId)
	if userId == 0 {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: "invalid user id"})
	}
	data, err := workspaceHandler.WorkspaceService.GetAllWorkspace(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}
