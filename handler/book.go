package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type BookHandler struct {
	BookService service.BookService
}

func (bookHandler *BookHandler) Insert(c echo.Context) error {
	param := models.BookReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = bookHandler.BookService.Insert(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (bookHandler *BookHandler) GellAllBook(c echo.Context) error {
	userID := c.Get("id").(int64)
	data, err := bookHandler.BookService.GellAllBook(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)

}