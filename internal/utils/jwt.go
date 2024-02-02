package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", nil
	}
	return jwtToken, nil
}
