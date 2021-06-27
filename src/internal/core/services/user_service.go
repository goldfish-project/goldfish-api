package services

import (
	"goldfish-api/internal/core/domain"
	"goldfish-api/internal/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

// NewService creates a new instance of the user service
func NewService(userRepository ports.UserRepository) ports.UserService {
	return &UserService{userRepository:userRepository}
}

// authenticate checks the user credentials and returns a valid JWT
func (service *UserService) Authenticate(email, password string) (jwt string, err error) {
	return "", nil
}

// createUser creates a user and saves it to the database
func (service *UserService) Create(user *domain.User) (jwt string, err error) {
	if err := service.userRepository.Save(user); err != nil {
		return "", err
	}

	//TODO: create jwt
	return "", nil
}

// update updates the user information
func (service *UserService) Update(user *domain.User) error {
	return nil
}