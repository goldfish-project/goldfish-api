package domain

import (
	"github.com/google/uuid"
	"time"
)

type Workspace struct {
	WorkspaceId   uuid.UUID
	Name          string
	CreatedOn     time.Time
	Owner         User
	Collaborators []User
	Variables     []Variable
	Collections   []Collection
}

type WorkspaceToUser struct {
	WorkspaceId uuid.UUID
	UserEmail   string
}
