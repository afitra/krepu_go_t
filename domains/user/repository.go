package user

import (
	"github.com/labstack/echo"
	"krepu_go_t/models"
)

type Repository interface {
	RCreateUser(c echo.Context, payload models.UserPayload) error
	RGetUserByUserName(c echo.Context, user_name string) (models.User, error)
}
