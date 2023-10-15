package usecase

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/helpers"
	"krepu_go_t/models"
	"math/rand"
)

type UserUseCase struct {
	response      models.Response
	responseError models.ErrorResponse
	userRepo      user.Repository
}

func NewUserUseCase(userRepository user.Repository) user.Usecase {

	return &UserUseCase{userRepo: userRepository}
}

func (ur *UserUseCase) URegister(c echo.Context, payload models.UserPayload) (interface{}, error) {

	var err error
	payload.Password, err = helpers.HashPassword(payload.Password)
	payload.Tenor = rand.Intn(4) + 1
	if err = ur.userRepo.RCreateUser(payload); err != nil {
		return nil, err
	}

	return ur.reverseSuccessResponse(models.CodeCreated, models.ResponseSuccess, models.MessageDataProcessing, nil, err)
}

func (ur *UserUseCase) ULogin(c echo.Context, payload models.LoginPayload) (interface{}, error) {
	var user models.User
	var token string
	var err error
	if user, err = ur.userRepo.RGetUserByUserName(payload.UserName); err != nil {
		return nil, err
	}

	if err = helpers.VerifyPassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	if token, err = helpers.GenerateToken(user.UserName); err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"token": token,
	}

	return ur.reverseSuccessResponse(models.CodeSuccess, models.ResponseSuccess, models.MessageDataProcessing, result, err)
}

func (ur *UserUseCase) reverseSuccessResponse(code string, status string, message string, data interface{}, err error) (interface{}, error) {
	var resp models.Response
	if code == "" {
		code = models.CodeSuccess
	}
	resp.Code = code
	resp.Status = status
	resp.Message = message
	resp.Data = data
	return resp, err
}
