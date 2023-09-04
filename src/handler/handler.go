package handler

import s "e-wallet/src/controllers"

type Handler struct {
	userService s.UserController
	authService s.UserController
}

type HandlerConfig struct {
	UserController s.UserController
	AuthController s.AuthController
}


func NewHandler (c *HandlerConfig) *Handler {
	return &Handler{
		userService: c.UserController,
	}
}