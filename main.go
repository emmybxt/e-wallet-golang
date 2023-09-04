package main

import (
	"e-wallet/src/config"
	"e-wallet/src/handler"
	jwt "e-wallet/src/helpers"
	"e-wallet/src/repository"

	// "fmt"
	"log"
	// "os"


	"github.com/gin-gonic/gin"

	service "e-wallet/src/controllers"
)

func main() {
	db := config.GetConnection()

	userRepository := repository.NewUserRepository(&repository.URConfig{DB: db})
	
	
	authService := service.NewAuthController(&service.ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})

	userController := service.NewUserService(&service.USConfig{UserRepository: userRepository})


	jwtService := jwt.NewJWTService(&jwt.JWTSConfig{})


	handle := handler.NewHandler(&handler.HandlerConfig{
		UserController: userController,
		AuthController: authService,
	})

	routes := route.NewRouter(&route.R)
	router := gin.Default()
	router.Static("/docs", "./pkg/swaggerui")
	// router.NoRoute(h.NoRoute)

	// version := os.Getenv("API_VERSION")
	//api := router.Group(fmt.Sprintf("/api/%s", version))

	// routes.Auth(api, h)
	// routes.User(api, h)
	// routes.Transaction(api, h)

	router.Run(":8000")

}
