package domain

import "time"

type User struct {
	tableName struct{}
	FirstName string
	LastName  string
	Email     string `pg:",pk"`
	Password  string
	CreatedOn time.Time
}
