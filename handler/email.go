package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type EmailHandler struct {
	EmailService service.EmailService
}

func (emailHandler *EmailHandler) SendOneEmail(c echo.Context) error {
	param := models.Sendsingleemail{}
	err := c.Bind(&param)
	if err != nil {
		return err
	}
	err = emailHandler.EmailService.SendOneEmail(param)
	if err != nil {
		return err
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
