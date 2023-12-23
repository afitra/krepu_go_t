package http

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/logger"
	"krepu_go_t/models"
	"net/http"
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
	echoGroup.API.POST("/user/register", handler.registerUser)
	echoGroup.API.POST("/user/login", handler.loginUser)
}

func (uh *UserHandler) registerUser(c echo.Context) error {
	var request models.UserPayload
	uh.response, uh.respErrors = models.NewResponse()
	if err := c.Bind(&request); err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(models.MessageUnprocessableEntity)
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}

	if err := c.Validate(request); err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(models.ErrSomethingWrong.Error())
		uh.respErrors.AddError(err.Error())
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}

	resp, err := uh.usecase.URegister(c, request)
	if err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(err.Error())
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}
	return c.JSON(http.StatusCreated, resp)
}

func (uh *UserHandler) loginUser(c echo.Context) error {
	var request models.LoginPayload
	uh.response, uh.respErrors = models.NewResponse()
	if err := c.Bind(&request); err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(models.MessageUnprocessableEntity)
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}

	if err := c.Validate(request); err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(models.ErrSomethingWrong.Error())
		uh.respErrors.AddError(err.Error())
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}

	resp, err := uh.usecase.ULogin(c, request)
	if err != nil {
		logger.Make(c, nil).Debug(err)
		uh.respErrors.SetTitle(err.Error())
		uh.response.SetResponse("", &uh.respErrors)
		return uh.response.Body(c, err)
	}
	return c.JSON(http.StatusOK, resp)
}
