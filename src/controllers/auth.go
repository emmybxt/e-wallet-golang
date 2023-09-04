package controllers

import (
	"context"
	"e-wallet/src/dto"
	"e-wallet/src/models"
	r "e-wallet/src/repository"
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)


type AuthController interface {
	Login(input *dto.LoginRequestBody) (*models.User, error)
}


type authController struct {
	userRepository r.UserRepository
}

type ASConfig struct {
	UserRepository          r.UserRepository
	// PasswordResetRepository r.PassowrdResetRepository
}

func NewAuthController(c *ASConfig) AuthController {
	return &authController{
		userRepository:          c.UserRepository,
	}
}

func (s *authController) Login(ctx context.Context, input *dto.LoginRequestBody) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	_, err := mail.ParseAddress(input.Email)

	if err != nil {
		return nil, errors.New("Email address is not valid")
	}

	user, err := s.userRepository.FindByEmail(input.Email)

	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return nil, errors.New("UserId does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return nil, errors.New("Incorrect credentials")
	}

	return user, nil

}
