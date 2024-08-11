package handler

import (
	"core/handler/validation"
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type CustomerHandler struct {
	CustomerService service.CustomerService
}

func (customerHandler *CustomerHandler) Insert(c echo.Context) error {
	param := models.GetCustomerParam{}
	param.UserID = c.Get("id").(int64)
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = validation.CustomerInsertValidation(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = customerHandler.CustomerService.Insert(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (customerHandler *CustomerHandler) GetAll(c echo.Context) error {
	data, err := customerHandler.CustomerService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (customerHandler *CustomerHandler) UpdateCustomer(c echo.Context) error {
	// u need to give user id, it update that specific customer
	param := models.GetCustomerParam{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	if param.UserID == 0 {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: "user does not exist"})
	}

	err = customerHandler.CustomerService.UpdateCustomer(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}

	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
