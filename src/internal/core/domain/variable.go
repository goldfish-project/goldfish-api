package domain

import "github.com/google/uuid"

type Variable struct {
	VariableId  uuid.UUID
	Key         string
	Value       string
}