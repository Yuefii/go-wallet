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
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
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

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse

	user, err := s.repository.UserByUsername(req.Username)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Username or Password"}
	}

	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorhandler.NotFoundError{Message: "Wrong Email or Password"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}
	return &data, nil
}
