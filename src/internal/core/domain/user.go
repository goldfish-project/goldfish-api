package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserId uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedOn time.Time
}
