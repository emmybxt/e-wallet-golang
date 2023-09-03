package middleware

import (
	n "e-wallet/src/helpers"
	"e-wallet/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleWare(JWTService n.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			response := utils.ErrorResponse("Unauthorized", "Unauthorized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authorizationHeader, " ")

		if len(arrayToken) != 2 {
			response := utils.ErrorResponse("Unauthorized", "Unauthorized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)

			return
		}

		encodedToken = arrayToken[1]

		token, err := JWTService.ValidateToken(encodedToken)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized",  err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.ErrorResponse("Unauthorized", "not a valid bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].float(64))
		Email := string(payload["email"])


		// c.Set("User", user)




	}
}
