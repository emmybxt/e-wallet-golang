package middleware

import (
	n "e-wallet/src/controllers"
	"e-wallet/src/dto"
	s "e-wallet/src/helpers"
	"e-wallet/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtService s.JWTService, userService n.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.ErrorResponse("Unauthorized", "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) != 2 {
			response := utils.ErrorResponse("Unauthorized", "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		encodedToken := arrayToken[1]
		token, err := jwtService.ValidateToken(encodedToken)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.ErrorResponse("Unauthorized", "not a valid bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))

		params := &dto.UserRequestParams{}
		params.UserID = userID
		user, err := userService.GetUser(params)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
