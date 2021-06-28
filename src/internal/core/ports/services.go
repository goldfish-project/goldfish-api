package ports

import (
	"goldfish-api/internal/core/domain"
	"time"
)

type UserService interface {
	// Authenticate checks the user credentials and returns a valid JWT
	Authenticate(email, password string) (jwt string, expiration time.Time, err error)
	// Create creates a user and saves it to the database
	Create(user *domain.User) (jwt string, expiration time.Time, err error)
	// Update updates the user information
	Update(user *domain.User) error
}