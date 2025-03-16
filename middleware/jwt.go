package middleware

import (
	"context"
	"core/domain"
	"core/firebase"
	"core/models"
	"core/service"
	"core/utils"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func JWTVerify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims, err := service.ExtractJWT(c)
			if err != nil {
				return err
			}

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

func VerifyGoogleToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			app, err := firebase.InitializeFirebase()
			if err != nil {
				log.Fatalf("Failed to initialize Firebase: %v", err)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			}

			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token format"})
			}
			idToken := tokenParts[1]

			authClient, err := app.Auth(context.Background())
			if err != nil {
				log.Println("Failed to get Firebase Auth client:", err)
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Server error"})
			}

			decodedToken, err := authClient.VerifyIDToken(context.Background(), idToken)
			if err != nil {
				log.Println("Invalid Firebase ID Token:", err)
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			}

			// Extract user data directly from the token
			claims := decodedToken.Claims
			email, _ := claims["email"].(string)
			name, _ := claims["name"].(string)
			picture, _ := claims["picture"].(string)

			// Store user data in the context
			c.Set("user_email", email)
			c.Set("user_name", name)
			c.Set("user_picture", picture)

			return next(c)
		}
	}
}
