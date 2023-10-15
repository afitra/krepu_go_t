package http

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/models"
)

type UserHandler struct {
	response   models.Response
	respErrors models.ErrorResponse
	usecase    user.Usecase
}

func NewUserHandler(echoGroup models.EchoGroup, auc user.Usecase) {
	handler := &UserHandler{
		usecase: auc,
	}
	echoGroup.API.GET("/GetCurrentTime", handler.registerUser)
}

func (uh *UserHandler) registerUser(c echo.Context) error {
	var err error
	//var resp interface{}
	//
	//fmt.Printf(resp)
	return err
}
