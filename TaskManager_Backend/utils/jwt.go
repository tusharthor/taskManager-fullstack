package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your secret key")

// generate jwt token
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
	}

	//with hashing256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// validate token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
