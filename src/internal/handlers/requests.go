package handlers

import (
	"goldfish-api/internal/core/domain"
	"net/mail"
)

type UserRegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// IsValid checks whether a given request object is valid or not
func (req *UserRegisterRequest) IsValid() bool {
	// check if all fields are given
	if len(req.Email) == 0 || len(req.Password) < 8 || len(req.FirstName) == 0 || len(req.LastName) == 0 {
		return false
	}

	//validate email
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return false
	}

	return true
}

// ToUser converts the request object into a domain user
func (req *UserRegisterRequest) ToUser() domain.User {
	return domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Email,
	}
}

type UserAuthenticateRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// IsValid checks whether a given request object is valid or not
func (req *UserAuthenticateRequest) IsValid() bool {
	// check if all fields are given
	if len(req.Email) == 0 || len(req.Password) < 8 {
		return false
	}

	//validate email
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return false
	}

	return true
}