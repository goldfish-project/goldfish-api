package domain

import "time"

type User struct {
	tableName struct{}
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedOn time.Time
}
