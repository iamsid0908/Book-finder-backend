package handler

import (
	"core/models"
	"core/service"

	"github.com/labstack/echo"
)

type ChannelHandler struct {
	ChannelService service.ChannelService
}

func (channelHandler *ChannelHandler) CreateChannel(c echo.Context) error {
	param := models.CreateChannelReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(400, models.BasicResp{Message: err.Error()})
	}
	userId := c.Get("id").(int64)
	param.OwnerID = userId
	data, err := channelHandler.ChannelService.CreateChannel(param)
	if err != nil {
		return c.JSON(500, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: "success",
		Data:    data,
	}
	return c.JSON(200, resp)
}

func (channelHandler *ChannelHandler) AddUserInChannel(c echo.Context) error {
	param := models.AddUserInChannelReqs{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(400, models.BasicResp{Message: err.Error()})
	}
	userId := c.Get("id").(int64)
	param.UserID = userId
	data, err := channelHandler.ChannelService.AddUserInChannel(param)
	if err != nil {
		return c.JSON(500, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: "success",
		Data:    data,
	}
	return c.JSON(200, resp)
}
