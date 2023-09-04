package dto

import "e-wallet/src/models"

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,email"`
}

type RegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,email"`
	Name string `json:"name" binding:"required,name"`
}

type ForgotPasswordRequestBody struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestBody struct {
	Token           string `json:"token" binding:"required"`
	Password        string `json:"password" binding:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=5"`
}

type ForgotPasswordResponseBody struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginResponseBody struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	WalletNumber string `json:"wallet"`
	Token        string `json:"token"`
}

func FormatLogin(user *models.User, wallet *models.Wallet, token string) LoginResponseBody {
	return LoginResponseBody{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		WalletNumber: wallet.Number,
		Token:        token,
	}
}

// func FormatForgotPassword(passwordReset *models.PasswordReset) ForgotPasswordResponseBody {
// 	return ForgotPasswordResponseBody{
// 		Email: passwordReset.User.Email,
// 		Token: passwordReset.Token,
// 	}
// }