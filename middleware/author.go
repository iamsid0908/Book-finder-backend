package middleware

import (
	"core/models"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

func VerifyRoles(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("role")
			for _, roles := range allowedRoles {
				if role == roles {
					return next(c)
				}
			}
			return echo.NewHTTPError(http.StatusForbidden, "Access denied")
		}
	}
}

func VerifyAccountantAuthor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role != "Accountant" {
			return c.JSON(http.StatusForbidden, models.BasicRespMesg{
				Message: utils.ErrWrongPerson,
			})
		}
		return next(c)
	}
}

func VerifyHRAuthor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role != "HR" {
			return c.JSON(http.StatusForbidden, models.BasicRespMesg{
				Message: utils.ErrWrongPerson,
			})
		}
		return next(c)
	}
}

func VerifyAdministratorAuthor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role")
		if role != "Administrator" {
			return c.JSON(http.StatusForbidden, models.BasicRespMesg{
				Message: utils.ErrWrongPerson,
			})
		}
		return next(c)
	}
}
