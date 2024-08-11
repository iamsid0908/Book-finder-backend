package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type BillingHandler struct {
	BillingService service.BillingService
}

func (billingHandler *BillingHandler) Insert(c echo.Context) error {
	param := models.InsertBillingParam{}
	param.UserID = c.Get("id").(int64)
	err := c.Bind(&param)
	if err != nil {
		return err
	}
	err = billingHandler.BillingService.Insert(param)
	if err != nil {
		return err
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (billingHandler *BillingHandler) Update(c echo.Context) error {
	param := models.InsertBillingParam{}
	param.UserID = c.Get("id").(int64)

	err := c.Bind(&param)
	if err != nil {
		return err
	}
	err = billingHandler.BillingService.Update(param)
	if err != nil {
		return err
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (billingHandler *BillingHandler) List(c echo.Context) error {
	data, err := billingHandler.BillingService.List()
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}
