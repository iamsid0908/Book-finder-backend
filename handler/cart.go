package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type CartHandler struct {
	CartService service.CartService
}

func (cartHandler *CartHandler) Insert(c echo.Context) error {
	param := models.InsertCartParam{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	param.UserID = c.Get("id").(int64)
	err = cartHandler.CartService.Insert(param)
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
func (cartHandler *CartHandler) GetCartByUserId(c echo.Context) error {
	userID := c.Get("id").(int64)
	data, err := cartHandler.CartService.GetCartByUserId(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (cartHandler *CartHandler) GetSizeofCart(c echo.Context) error {
	userID := c.Get("id").(int64)
	data, err := cartHandler.CartService.GetSizeofCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (cartHandler *CartHandler) RemoveFromCart(c echo.Context) error {
	param := models.RemoveFromCartReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	param.UserID = c.Get("id").(int64)
	err = cartHandler.CartService.RemoveFromCart(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
