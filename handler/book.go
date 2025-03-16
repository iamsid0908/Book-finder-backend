package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"
	"strconv"

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
	author := c.QueryParam("writter_name")
	title := c.QueryParam("title")
	page := c.QueryParam("page")
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		pageInt = 1
	}

	param := models.SearchByInputParam{
		WritterName: author,
		Title:       title,
		UserID:      userID,
		Page:        (pageInt),
		Limit:       int64(8),
	}
	data, err := bookHandler.BookService.GellAllBook(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicRespWithMeta{
		Message: utils.Success,
		Data:    data.Data,
		Meta:    data.Meta,
	}
	return c.JSON(http.StatusOK, resp)
}

func (bookHandler *BookHandler) BulkInsert(c echo.Context) error {
	param := models.BulkInsertBookReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	param.File, err = c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	err = bookHandler.BookService.BulkInsert(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}

	resp := models.BasicResp{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (bookHandler *BookHandler) Recommend(c echo.Context) error {
	param := models.RecommendReqs{}
	err := c.Bind(&param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	data, err := bookHandler.BookService.Recommend(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)

}
