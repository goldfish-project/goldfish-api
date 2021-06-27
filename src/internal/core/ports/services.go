package ports

import (
	"goldfish-api/internal/core/domain"
)

type UserService interface {
	// authenticate checks the user credentials and returns a valid JWT
	Authenticate(email, password string) (jwt string, err error)
	// createUser creates a user and saves it to the database
	Create(user *domain.User) (jwt string, err error)
	// update updates the user information
	Update(user *domain.User) error
}