package users

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/biswaone/go-blogs/internal/utils"
	"github.com/jackc/pgx/v5"
)

type User struct {
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	Password       *string `json:"password"`
	ProfilePicture string  `json:"profile_picture"`
}

type Response struct {
	Message          string     `json:"message"`
	RegsitrationDate *time.Time `json:"registration_date,omitempty"`
	Exception        *string    `json:"exception,omitempty"`
}

var (
	ErrInvalidName     = errors.New("EmptyNameNotAllowed")
	ErrInvalidEmail    = errors.New("EmptyEmailNotAllowed")
	ErrInvalidPassword = errors.New("EmptyPasswordNotAllowed")
	ErrInvladUsername  = errors.New("EmptyUsernameNotAllowed")
)

func (u User) ValidateNewUser() error {
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

func (u User) CreateUser(db *pgx.Conn) (*Response, error) {
	hashedPassword, err := utils.HashPassword(*u.Password)
	if err != nil {
		return &Response{Message: "Cannot Create User"}, err
	}
	var userID int
	var registrationDate time.Time
	query := `
	INSERT INTO Users (Name, Email, Password)
		VALUES ($1, $2, $3)
		RETURNING UserID, RegistrationDate
	`
	err = db.QueryRow(context.Background(), query, u.Name, u.Email, hashedPassword).Scan(&userID, &registrationDate)
	if err != nil {
		return &Response{}, err
	}

	return &Response{Message: "User Created Successfully", RegsitrationDate: &registrationDate}, nil
}

func GetUserByEmail(db *pgx.Conn, email string) (*User, error) {
	var name, password, profilePicture sql.NullString
	query := `	
	SELECT Name, Password, ProfilePicture FROM Users WHERE Email = $1
	`
	err := db.QueryRow(context.Background(), query, email).Scan(&name, &password, &profilePicture)
	if err != nil {
		return &User{}, err
	}
	return &User{Name: name.String, Email: email, Password: &password.String, ProfilePicture: profilePicture.String}, nil

}
