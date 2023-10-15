package user

import (
	"krepu_go_t/models"
)

type Repository interface {
	RCreateUser(payload models.UserPayload) error
	RGetUserByUserName(user_name string) (models.User, error)
}
