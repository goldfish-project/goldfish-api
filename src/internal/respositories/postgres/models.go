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

// from takes a domain model and converts it into a postgres model
func (u *User) from(user *domain.User) {
	u.UserId = user.UserId
	u.Password = user.Password
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.CreatedOn = user.CreatedOn
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

// from takes a domain model and converts it into a postgres model
func (w *Workspace) from(workspace *domain.Workspace, ownerId uuid.UUID) Workspace {
	// convert collaborators
	var collaborators []*User
	for _, value := range workspace.Collaborators {
		user := User{}
		user.from(&value)
		collaborators = append(collaborators, &user)
	}

	// convert variables
	var variables []*Variable
	for _, value := range workspace.Variables {
		variable := Variable{}
		variable.from(&value, workspace.WorkspaceId)
		variables = append(variables, &variable)
	}

	// convert collections
	var collections []*Collection
	for _, value := range workspace.Collections {
		collection := Collection{}
		collection.from(&value, workspace.WorkspaceId)
		collections = append(collections, &collection)
	}

	// convert owner
	owner := &User{}
	owner.from(&workspace.Owner)

	return Workspace{
		WorkspaceId:   workspace.WorkspaceId,
		Name:          workspace.Name,
		CreatedOn:     workspace.CreatedOn,
		OwnerUserId:   ownerId,
		Owner:         owner,
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

// from takes a domain model and converts it into a postgres model
func (v *Variable) from(variable *domain.Variable, workspaceId uuid.UUID) {
	v.VariableId = variable.VariableId
	v.Key = variable.Key
	v.Value = variable.Value
	v.WorkspaceId = workspaceId
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

// from takes a domain model and converts it into a postgres model
func (c *Collection) from(collection *domain.Collection, workspaceId uuid.UUID) {
	c.CollectionId = collection.CollectionId
	c.Name = collection.Name
	c.WorkspaceId = workspaceId
}