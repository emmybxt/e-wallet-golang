package controllers

import (
	"e-wallet/src/dto"
	"e-wallet/src/models"
	r "e-wallet/src/repository"
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	GetUser(input *dto.UserRequestParams)(*models.User, error)
	CreateUser(input *dto.RegisterRequestBody) (*models.User, error)

}

type userController struct {
	userRepository r.UserRepository
	// walletRepositoty r.
}

type USConfig struct {
	UserRepository   r.UserRepository
	// WalletRepository r.WalletRepository
}

func NewUserService(c *USConfig) UserController {
	return &userController{
		userRepository:   c.UserRepository,
		// walletRepository: c.WalletRepository,
	}
}

func (s *userController) GetUser(input *dto.UserRequestParams) (*models.User, error) {

	user, err := s.userRepository.FindById(input.UserID)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (s *userController) CreateUser(input *dto.RegisterRequestBody)(*models.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return nil, errors.New("Invalid email address")
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.ID != 0 {
		return nil, errors.New("user already exists")

	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}