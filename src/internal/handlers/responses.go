package handlers

import "time"

type UserAuthenticatedResponse struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}