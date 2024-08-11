package handler

import (
	"core/models"
	"core/service"
	"net/http"

	"github.com/labstack/echo"
)

type HealthHandler struct {
	HealthService service.HealthService
}

func (healthHandler *HealthHandler) Check(c echo.Context) error {
	health, err := healthHandler.HealthService.Check()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})
	}

	resp := models.BasicResp{
		Message: "Success",
		Data:    health,
	}
	return c.JSON(http.StatusOK, resp)
}
