package domain

import (
	"noob-server/pkg/models"
	"noob-server/pkg/types"
)

type IUserRepo interface {
	GetUser(username *string) (*models.UserDetail, error)
	CreateUser(user *models.UserDetail) error
}

type IAuthService interface {
	LoginUser(loginRequest *types.LoginRequest) (*types.LoginResponse, error)
	SignupUser(registerRequest *types.SignupRequest) error
}
