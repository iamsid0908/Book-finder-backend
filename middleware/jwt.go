package middleware

import (
	"core/domain"
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func JWTVerify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := service.ExtractJWT(c)
			if err != nil {
				return err
			}
			fmt.Println(claims)
			c.Set("id", claims.ID)
			c.Set("email", claims.Email)
			c.Set("name", claims.Name)
			c.Set("role", claims.Role)
			c.Set("language", claims.Language)

			userParams := &models.User{Email: claims.Email}
			user, err := domain.UserDomain.GetLoginUser(&domain.UserDomainCtx{}, userParams)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			if user == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrUserTokenNotExist.Error())
			}
			return next(c)
		}

	}
}
