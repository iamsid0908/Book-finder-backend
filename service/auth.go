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

func (c *AuthService) RegisterUser(param *models.RegisterUserRequest) error {
	// err := c.validateRegisterUser(param)
	// if err != nil {
	// 	return err
	// }
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), 10)
	if err != nil {
		return err
	}
	password := string(hashedPassword)

	err = c.UserDomain.Insert(models.User{
		Email:    param.Email,
		Password: &password,
		Name:     param.Name,
		RoleId:   param.Role,
		Language: utils.UserLanguageEn,
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
	user, err := c.UserDomain.GetWithRole(models.GetUserParam{Email: param.Email})
	if err != nil {
		return models.LogInResponse{}, err
	}
	err = c.validateLogIn(param, user)
	if err != nil {
		return models.LogInResponse{}, err
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
	}
	return resp, nil
}

func (c *AuthService) validateLogIn(param models.LogInRequest, user models.UserData) error {
	if user.ID == 0 {
		return utils.ErrUserNotExist
	}

	if user.Password == nil {
		return utils.ErrPasswordNotExist
	}

	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(param.Password))
	if err != nil {
		return utils.ErrWrongPassword
	}

	return nil
}

func (c *AuthService) LoginGoogleUser(param models.GoogleUserRequest) (models.LogInResponse, error) {
	user, err := c.UserDomain.Get(models.GetUserParam{Email: param.Email})
	if err != nil {
		user, err = c.UserDomain.Create(models.User{Email: param.Email, Name: param.Name})
		if err != nil {
			return models.LogInResponse{}, err
		}
	}
	now := time.Now()
	userData := models.UserData{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Language: user.Language,
	}
	payload := ParseJWTParamFromUser(userData, now)

	token, err := GenerateJWT(payload)
	if err != nil {
		return models.LogInResponse{}, err
	}
	resp := models.LogInResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Token:     token,
	}
	return resp, nil
}

func ParseJWTParamFromUser(user models.UserData, now time.Time) models.JWTPayload {
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
	cookie, err := e.Cookie("accessToken")
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
