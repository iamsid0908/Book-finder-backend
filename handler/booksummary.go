package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BookSummaryHandler struct {
	BookSummaryService service.BookSummaryService
}

func (bookSummaryHandler *BookSummaryHandler) Insert(c echo.Context) error {
	param := models.BookSummaryReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = bookSummaryHandler.BookSummaryService.Insert(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (bookSummaryHandler *BookSummaryHandler) GetBookDetails(c echo.Context) error {
	param := models.GetBookSummaryDetailsReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	param.BookID, err = strconv.ParseInt(c.Param("book_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	fmt.Println(param.BookID)
	data, err := bookSummaryHandler.BookSummaryService.GetBookDetails(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Data:    data,
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}
