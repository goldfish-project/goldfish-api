package domain

import (
	"github.com/google/uuid"
	"time"
)

type Workspace struct {
	WorkspaceId uuid.UUID `pg:"alias:id,default:uuid_generate_v4(),pk,type:uuid"`
	Name        string
	CreatedOn   time.Time

	OwnerEmail string
	Owner       *User `pg:"rel:has-one"`
	Collaborators []*User `pg:"many2many:workspace_to_users"`
}

type WorkspaceToUser struct {
	WorkspaceId uuid.UUID
	UserEmail       string
}
