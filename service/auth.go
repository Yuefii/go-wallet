package service

import (
	"go-wallet/dto"
	errorhandler "go-wallet/errorHandler"
	"go-wallet/helper"
	"go-wallet/models"
	"go-wallet/repository"
)

type AuthService interface {
	Register(req *dto.Register) error
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *dto.Register) error {
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {
		return &errorhandler.BadRequestError{
			Message: "Username Already Exist",
		}
	}
	if req.Password != req.PasswordConf {
		return &errorhandler.BadRequestError{
			Message: "password not match",
		}
	}

	passwordHashed, err := helper.HashedPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := models.User{
		Name:     req.Name,
		Username: req.Username,
		Password: passwordHashed,
		Gender:   req.Gender,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}
