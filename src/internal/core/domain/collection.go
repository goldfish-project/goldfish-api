package domain

import "github.com/google/uuid"

type Collection struct {
	CollectionId uuid.UUID
	Name string
}