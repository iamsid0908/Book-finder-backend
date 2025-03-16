package handler

import (
	"core/handler/validation"
	"core/models"
	"core/service"
	"core/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (authHandler *AuthHandler) RegisterUser(c echo.Context) error {
	var err error
	param := new(models.RegisterUserRequest)

	err = c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}

	err = validation.RegisterUser(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}

	err = authHandler.AuthService.RegisterUser(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BasicResp{Message: err.Error()})

	}
	resp := models.BasicResp{
		Message: utils.Success,
	}

	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) LoginUser(c echo.Context) error {
	var err error
	param := new(models.LogInRequest)
	err = c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusOK, models.BasicResp{Message: err.Error()})
	}
	data, err := authHandler.AuthService.LoginUser(*param)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "accessToken",
		Value:    data.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   86400,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) GoogleLogin(c echo.Context) error {
	var err error
	param := new(models.GoogleUserRequest)
	email := c.Get("user_email").(string)
	name := c.Get("user_name").(string)

	if email == "" || name == "" {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: "All fields are required"})
	}

	param.Email = email
	param.Name = name

	err = c.Bind(param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.BasicResp{Message: err.Error()})
	}
	data, err := authHandler.AuthService.LoginGoogleUser(*param)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     "accessToken",
		Value:    data.Token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		MaxAge:   86400,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) Test(c echo.Context) error {
	fmt.Println(c)
	resp := models.BasicRespMesg{
		Message: utils.Success,
	}
	return c.JSON(http.StatusOK, resp)
}

func (authHandler *AuthHandler) UserLogOut(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true, // Keep it HttpOnly, but setting MaxAge -1 removes it
		// Secure:   true, // Keep it secure if using HTTPS
	}

	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
