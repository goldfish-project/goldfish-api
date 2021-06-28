package services

import (
	"errors"
	"goldfish-api/internal/core/domain"
	"goldfish-api/internal/core/ports"
	"goldfish-api/internal/handlers/middleware"
	"goldfish-api/internal/utils/hashing"
	"time"
)

type UserService struct {
	userRepository ports.UserRepository
	jwt            *middleware.JWTMiddleware
}

// NewService creates a new instance of the user service
func NewService(userRepository ports.UserRepository, jwt *middleware.JWTMiddleware) ports.UserService {
	return &UserService{userRepository: userRepository, jwt: jwt}
}

// Authenticate checks the user credentials and returns a valid JWT
func (service *UserService) Authenticate(email, password string) (jwt string, expiration time.Time, err error) {
	// get user from database
	user, err := service.userRepository.GetByEmail(email)

	if err != nil {
		return "", time.Time{}, err
	}

	// validate password
	if !hashing.VerifyPassword(user.Password, password) {
		return "", time.Time{}, errors.New("Invalid credentials.")
	}

	// create token
	token, expiration, err := service.jwt.GetToken(user.UserId.String())

	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiration, nil
}

// Create creates a user and saves it to the database
func (service *UserService) Create(user *domain.User) (jwt string, expiration time.Time, err error) {
	// hash user password
	user.Password, err = hashing.HashAndSaltPassword(user.Password)

	if err != nil {
		return "", time.Time{}, err
	}

	// save user to database
	if err := service.userRepository.Save(user); err != nil {
		return "", time.Time{}, err
	}

	// create token
	token, expiration, err := service.jwt.GetToken(user.UserId.String())

	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiration, nil
}

// update updates the user information
func (service *UserService) Update(user *domain.User) error {
	return nil
}
