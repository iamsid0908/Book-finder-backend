package service

import (
	"core/config"
	"core/domain"
	"core/models"
	"core/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	AuthDomain domain.AuthDomain
	UserDomain domain.UserDomain
}

func (c *AuthService) RegisterUser(param *models.RegisterUserRequest) (models.ResisterResp, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), 10)
	if err != nil {
		return models.ResisterResp{}, err
	}
	password := string(hashedPassword)
	otp := utils.GenerateOTP()
	otpExpiry := time.Now().Add(10 * time.Minute) // OTP valid for 10 minutes

	user, err := c.UserDomain.Insert(models.User{
		Email:     param.Email,
		Password:  &password,
		Name:      param.Name,
		Role:      param.Role,
		Language:  utils.UserLanguageEn,
		IsActive:  param.IsActive,
		Otp:       otp,
		OtpExpiry: otpExpiry,
	})
	if err != nil {
		return models.ResisterResp{}, err
	}
	go func() {
		emailTemplate := GetEmailTemplate(utils.AuthTypeRegister)
		sendRequest := models.SendMail{
			SendTo:   param.Email,
			UserName: param.Name,
			OTP:      (otp),
		}
		subject := fmt.Sprintf("%s %s!", GetEmailSubject(utils.AuthTypeRegister), param.Name)
		utils.SendMail(emailTemplate, sendRequest, subject)
	}()
	resp := models.ResisterResp{
		UserID:   user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Redirect: fmt.Sprintf("/verifyotp/%d/%s", user.ID, user.Email),
	}
	return resp, nil
}
func GetEmailTemplate(templateType string) string {
	var emailTemplate string

	switch templateType {
	case utils.AuthTypeRegister:
		emailTemplate = "./template/register.html"
	case utils.AuthTypeResetPassword:
		emailTemplate = "./template/reset_password.html"
	}

	return emailTemplate
}

func GetEmailSubject(subjectType string) string {
	var emailSubject string

	switch subjectType {
	case utils.AuthTypeRegister:
		emailSubject = "Verify your email to complete registration"
	case utils.AuthTypeResetPassword:
		emailSubject = "Reset your password"
	}

	return emailSubject
}

func (c *AuthService) ResendOTP(param *models.ResendOTPRequest) error {
	otp := utils.GenerateOTP()
	otpExpiry := time.Now().Add(10 * time.Minute) // OTP valid for 10 minutes

	err := c.UserDomain.Update(models.User{
		ID:        param.Id,
		Email:     param.Email,
		Otp:       otp,
		OtpExpiry: otpExpiry,
	})
	if err != nil {
		return err
	}
	go func() {
		emailTemplate := GetEmailTemplate(utils.AuthTypeRegister)
		sendRequest := models.SendMail{
			SendTo:   param.Email,
			UserName: param.Email,
			OTP:      (otp),
		}
		subject := fmt.Sprintf("%s %s!", GetEmailSubject(utils.AuthTypeRegister), param.Email)
		utils.SendMail(emailTemplate, sendRequest, subject)
	}()

	return nil
}

func (c *AuthService) VerifyOTP(param models.VerifyOTPRequest) error {
	fmt.Println(param)
	user, err := c.UserDomain.Get(models.GetUserParam{ID: param.Id, Email: param.Email})
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return utils.ErrUserNotExist
	}
	if user.Otp != param.Otp {
		return utils.ErrWrongOTP
	}
	if time.Now().After(user.OtpExpiry) {
		fmt.Println("Current time is after OTP expiry time")
		return utils.ErrWrongOTP
	}

	err = c.UserDomain.Update(models.User{
		ID:       param.Id,
		Email:    param.Email,
		IsActive: true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *AuthService) validateRegisterUser(param *models.RegisterUserRequest) error {
	user, err := c.UserDomain.Get(models.GetUserParam{Email: param.Email})
	if err != nil {
		return err
	}

	if user.ID != 0 {
		return utils.ErrEmailExist
	}

	return nil
}

func (c *AuthService) LoginUser(param models.LogInRequest) (models.LogInResponse, error) {
	user, err := c.UserDomain.Get(models.GetUserParam{Email: param.Email})
	if err != nil {
		return models.LogInResponse{}, err
	}
	err = c.validateLogIn(param, user)
	if err != nil {
		return models.LogInResponse{}, utils.LogError(err, nil)
	}
	now := time.Now()
	payload := ParseJWTParamFromUser(user, now)

	token, err := GenerateJWT(payload)
	if err != nil {
		return models.LogInResponse{}, err
	}
	resp := models.LogInResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
		Redirect:  "/dashboard",
	}
	return resp, nil
}

func (c *AuthService) validateLogIn(param models.LogInRequest, user models.User) error {
	if user.ID == 0 {
		return utils.ErrUserNotExist
	}

	if user.Password == nil {
		return utils.ErrPasswordNotExist
	}
	if !user.IsActive {
		return utils.ErrUserNotActive
	}
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(param.Password))
	if err != nil {
		return utils.ErrWrongPassword
	}

	return nil
}

func ParseJWTParamFromUser(user models.User, now time.Time) models.JWTPayload {
	payload := models.JWTPayload{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Role:     user.Role,
		Language: user.Language,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Hour * 72).Unix(),
		},
	}

	return payload
}

func GenerateJWT(claims models.JWTPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractJWT(e echo.Context) (*models.JWTPayload, error) {
	cookie, err := e.Cookie("Bearer")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, utils.EmptyAuth)
	}

	tokenStr := cookie.Value
	if tokenStr == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, utils.EmptyAuth)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &models.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(utils.UnexpectedSigning, token.Header["alg"])
		}
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	if claims, ok := token.Claims.(*models.JWTPayload); token.Valid && ok {
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		var errorStr string
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errorStr = "Token has expired"
		} else {
			errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
		}
		return nil, echo.NewHTTPError(http.StatusUnauthorized, errorStr)
	} else {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unknown token error")
	}
}
