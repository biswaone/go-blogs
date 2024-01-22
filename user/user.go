package user

import "errors"

type RegisterUser struct {
	Username string  `json:"username"`
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
	if u.Username == "" {
		return ErrInvladUsername
	}
	return nil

}

func (u RegisterUser) CreateUser() (*RegisterUser, error) {
	err := u.ValidateNewUser()
	if err != nil {
		return &RegisterUser{}, err
	}
	return &RegisterUser{Name: u.Name, Email: u.Email, Username: u.Username}, nil
}
