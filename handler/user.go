package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	UserService service.UserService
}

func (userHandler *UserHandler) ListUser(c echo.Context) error {
	data, err := userHandler.UserService.List()
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (userHandler *UserHandler) Update(c echo.Context) error {
	param := models.UpdateUserParam{}
	err := c.Bind(&param)
	if err != nil {
		return err
	}
	err = userHandler.UserService.Update(param)
	if err != nil {
		return err
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (userHandler *UserHandler) GetUserName(c echo.Context) error {
	userId := c.Get("id").(int64)
	data, err := userHandler.UserService.GetUserName(userId)
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}
