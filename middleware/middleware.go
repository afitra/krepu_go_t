package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/models"
	"net/http"
	"os"
)

type customMiddleware struct {
	e        *echo.Echo
	userRepo user.Repository
}

var echGroup models.EchoGroup

func InitMiddleware(ech *echo.Echo, echoGroup models.EchoGroup, repo user.Repository) {
	cm := &customMiddleware{ech, repo}
	echGroup = echoGroup
	echGroup.AUTH.Use(cm.checkTokenMiddleware)
	echGroup.ADMIN.Use(cm.checkAdminMiddleware)
}

type CustomClaims struct {
	UserName string `json:"user_name"`
}

func (cm *customMiddleware) checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		respError := map[string]string{
			"code":    "401",
			"status":  "failed",
			"message": "you must login first",
		}
		tokenString := c.Request().Header.Get("Authorization")

		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, respError)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var user models.User
			user, _ = cm.userRepo.RGetUserByUserName(claims["user_name"].(string))
			c.Set("decode", user)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, respError)
	}
}

func (cm *customMiddleware) checkAdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		respError := map[string]string{
			"code":    "401",
			"status":  "failed",
			"message": "you dont have access this feature",
		}
		tokenString := c.Request().Header.Get("Authorization")
		var user models.User
		token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusUnauthorized, respError)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user, _ = cm.userRepo.RGetUserByUserName(claims["user_name"].(string))
		}

		if user.Role == "admin" {
			c.Set("decode", user)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, respError)
	}
}
