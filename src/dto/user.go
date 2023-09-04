package dto

import "e-wallet/src/models"



type UserRequestParams struct {
	UserID int `uri:"id" binding:"required"`
}

type UserRequestQuery struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

type UserResponseBody struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUser(user *models.User) UserResponseBody {
	formattedUser := UserResponseBody{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	return formattedUser
}

func FormatUsers(authors []*models.User) []UserResponseBody {
	formattedUsers := []UserResponseBody{}
	for _, user := range authors {
		formattedUser := FormatUser(user)
		formattedUsers = append(formattedUsers, formattedUser)
	}
	return formattedUsers
}

type UserDetailResponse struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Wallet WalletResponse `json:"wallet"`
}

func FormatUserDetail(user *models.User, wallet *models.Wallet) UserDetailResponse {
	formattedUser := UserDetailResponse{}
	formattedUser.ID = user.ID
	formattedUser.Name = user.Name
	formattedUser.Email = user.Email
	formattedUser.Wallet = FormatWallet(wallet)
	return formattedUser
}