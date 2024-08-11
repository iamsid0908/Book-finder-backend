package handler

import (
	"core/handler/validation"
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (authHandler *AuthHandler) RegisterUser(c echo.Context) error {
	var err error
	param := new(models.RegisterUserRequest)

	err = c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = validation.RegisterUser(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}

	err = authHandler.AuthService.RegisterUser(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})

	}
	resp := models.BasicResp{
		Message: utils.Success,
	}

	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) LoginUser(c echo.Context) error {
	var err error
	param := new(models.LogInRequest)
	err = c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusOK, models.BasicResp{Message: err.Error()})
	}
	data, err := authHandler.AuthService.LoginUser(*param)
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) Test(c echo.Context) error {
	fmt.Println(c)
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
