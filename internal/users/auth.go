package users

import (
	"errors"
	"time"

	"github.com/biswaone/go-blogs/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
)

var (
	ErrUserNotFound      = errors.New("UserNotFound")
	ErrIncorrectPassword = errors.New("IncorrectPassword")
)

type LoginRequest struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token,omitempty"`
	Message     string `json:"message"`
}

func (r LoginRequest) ValidateLoginRequest(db *pgx.Conn) (*User, error) {
	user, err := GetUserByEmail(db, r.Email)
	if err == pgx.ErrNoRows {
		return user, ErrUserNotFound
	} else if err != nil {
		return user, err
	}

	return user, nil
}

func (r LoginRequest) Login(u User) (string, error) {
	if !utils.CheckPasswordHash(*r.Password, *u.Password) {
		return "", ErrIncorrectPassword
	}
	claims := utils.UserClaims{
		Email: r.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	jwtToken, err := utils.NewAccessToken(claims)
	if err != nil {
		return "", err
	}
	return jwtToken, nil

}
