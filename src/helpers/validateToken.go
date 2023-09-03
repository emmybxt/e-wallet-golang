package helpers

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID int, email string) (string, error)
	ValidateToken(encodedToken string)(string, error)

}

type JWTConfig struct {

}


func NewJwtService(e * JWTConfig) JWTService {
	return &JWTService{}
}

var SECRET_KEY = os.Getenv("SECRET_KEY")
var JWT_TTL, _ = strconv.Atoi(os.Getenv("JWT_TTL"))
var ISSUER = os.Getenv("JWT_ISSUER")



type idTokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
	Email string `json:"email"`
}

func (n *JWTService) GenerateToken(userId int, email string) (string, error) {
	payload := idTokenClaims{}

	payload.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Minute * time.Duration(JWT_TTL))}
	payload.UserID = userId
	payload.Email = email
	payload.IssuedAt = &jwt.NumericDate{Time: time.Now()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken , err := token.SignedString(SECRET_KEY)
	
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (n *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}