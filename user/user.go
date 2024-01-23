package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/biswaone/go-blogs/auth"
	"github.com/jackc/pgx/v5"
)

type RegisterUser struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type Response struct {
	Message string       `json:"message"`
	User    RegisterUser `json:"user"`
}

var (
	ErrInvalidName     = errors.New("EmptyNameNotAllowed")
	ErrInvalidEmail    = errors.New("EmptyEmailNotAllowed")
	ErrInvalidPassword = errors.New("EmptyPasswordNotAllowed")
	ErrInvladUsername  = errors.New("EmptyUsernameNotAllowed")
)

func (u RegisterUser) ValidateNewUser() error {
	if u.Name == "" {
		return ErrInvalidName
	}
	if u.Email == "" {
		return ErrInvalidEmail
	}
	if u.Password == nil {
		return ErrInvalidPassword
	}
	if u.Name == "" {
		return ErrInvladUsername
	}
	return nil

}

func (u RegisterUser) CreateUser(db *pgx.Conn) (*RegisterUser, error) {
	hashedPassword, err := auth.HashPassword(*u.Password)
	if err != nil {
		return &RegisterUser{}, err
	}
	query := `
	INSERT INTO Users (Name, Email, Password)
		VALUES ($1, $2, $3)
		RETURNING UserID, RegistrationDate
	`
	user, err := db.Exec(context.Background(), query, u.Name, u.Email, hashedPassword)
	if err != nil {
		return &RegisterUser{}, err
	}

	fmt.Println(hashedPassword)
	fmt.Println(user)

	return &RegisterUser{Name: u.Name, Email: u.Email}, nil
}
