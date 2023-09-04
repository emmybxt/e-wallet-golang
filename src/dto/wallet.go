package dto

import "e-wallet/src/models"

type WalletRequestBody struct {
	UserID int `json:"name" binding:"required,alphanum"`
}

type WalletResponse struct {
	ID      int    `json:"id"`
	Number  string `json:"number"`
	Balance int    `json:"balance"`
}

func FormatWallet(wallet *models.Wallet) WalletResponse {
	return WalletResponse{
		ID:      int(wallet.ID),
		Number:  wallet.Number,
		Balance: wallet.Balance,
	}
}