package postgres

import (
	"github.com/google/uuid"
	"goldfish-api/internal/core/domain"
	"time"
)

type User struct {
	tableName struct{}
	UserId    uuid.UUID `pg:"alias:id,default:uuid_generate_v4(),pk,type:uuid"`
	FirstName string
	LastName  string
	Email     string `pg:",unique"`
	Password  string
	CreatedOn time.Time `pg:"default:now()"`
}

// domain maps the given postgres entity to a domain entity used inside the business layer
func (u *User) domain() domain.User {
	return domain.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		CreatedOn: u.CreatedOn,
		UserId:    u.UserId,
	}
}

type Workspace struct {
	WorkspaceId   uuid.UUID `pg:"alias:id,default:uuid_generate_v4(),pk,type:uuid"`
	Name          string
	CreatedOn     time.Time
	OwnerUserId   uuid.UUID     `pg:",type:uuid"`
	Owner         *User         `pg:"rel:has-one"`
	Collaborators []*User       `pg:"many2many:workspace_to_users"`
	Variables     []*Variable   `pg:"rel:has-many"`
	Collections   []*Collection `pg:"rel:has-many"`
}

// domain maps the given postgres entity to a domain entity used inside the business layer
func (w *Workspace) domain() domain.Workspace {
	// parse collaborators for domain
	var collaborators []domain.User
	for _, val := range w.Collaborators {
		collaborators = append(collaborators, val.domain())
	}

	// parse variables for domain
	var variables []domain.Variable
	for _, val := range w.Variables {
		variables = append(variables, val.domain())
	}

	// parse collections for domain
	var collections []domain.Collection
	for _, val := range w.Collections {
		collections = append(collections, val.domain())
	}

	return domain.Workspace{
		WorkspaceId:   w.WorkspaceId,
		Name:          w.Name,
		CreatedOn:     w.CreatedOn,
		Owner:         w.Owner.domain(),
		Collaborators: collaborators,
		Variables:     variables,
		Collections:   collections,
	}
}

type WorkspaceToUser struct {
	WorkspaceId uuid.UUID
	UserId      uuid.UUID
}

type Variable struct {
	VariableId  uuid.UUID `pg:"alias:id,default:uuid_generate_v4(),pk,type:uuid"`
	Key         string
	Value       string
	WorkspaceId uuid.UUID
}

// domain maps the given postgres entity to a domain entity used inside the business layer
func (v *Variable) domain() domain.Variable {
	return domain.Variable{
		VariableId: v.VariableId,
		Key:        v.Key,
		Value:      v.Value,
	}
}

type Collection struct {
	CollectionId uuid.UUID `pg:"alias:id,default:uuid_generate_v4(),pk,type:uuid"`
	Name         string
	WorkspaceId  uuid.UUID
}

// domain maps the given postgres entity to a domain entity used inside the business layer
func (c *Collection) domain() domain.Collection {
	return domain.Collection{
		CollectionId: c.CollectionId,
		Name:         c.Name,
	}
}
