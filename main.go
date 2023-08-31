package main

import (
	"e-wallet/src/config"
	// "fmt"
	"log"
	// "os"

	"github.com/gin-gonic/gin"
)


func main() {
	db := config.GetConnection()

	log.Print(db)

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