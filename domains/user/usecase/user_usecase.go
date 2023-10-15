package usecase

import (
	"krepu_go_t/domains/user"
	"krepu_go_t/models"
)

type UserUseCase struct {
	response models.Response
	userRepo user.Repository
}

func NewUserUseCase(userRepository user.Repository) user.Usecase {

	return &UserUseCase{userRepo: userRepository}
}
