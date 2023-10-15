package user

import (
	"github.com/labstack/echo"
	"krepu_go_t/models"
)

type Usecase interface {
	URegister(c echo.Context, payload models.UserPayload) (interface{}, error)
	//ULogin(c echo.Context) (interface{}, error)
}
