package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type RoleHandler struct {
	RoleService service.RoleService
}

func (roleHandler *RoleHandler) Insert(c echo.Context) error {
	param := new(models.RoleReqs)
	err := c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	data, err := roleHandler.RoleService.Insert(*param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (roleHandler *RoleHandler) FindAll(c echo.Context) error {
	data, err := roleHandler.RoleService.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}
