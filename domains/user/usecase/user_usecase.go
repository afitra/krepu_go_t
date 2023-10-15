package usecase

import (
	"github.com/labstack/echo"
	"krepu_go_t/domains/user"
	"krepu_go_t/models"
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
	err = ur.userRepo.RCreateUser(c, payload)
	if err != nil {
		return nil, err
	}

	return ur.reverseSuccessResponse(models.ResponseSuccess, models.MessageDataProcessing, nil, err)
}

func (ur *UserUseCase) reverseSuccessResponse(status string, message string, data interface{}, err error) (interface{}, error) {
	var resp models.Response
	resp.Code = models.CodeSuccess
	resp.Status = status
	resp.Message = message
	resp.Data = data
	return resp, err
}
