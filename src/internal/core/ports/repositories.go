package ports

import (
	"github.com/google/uuid"
	"goldfish-api/internal/core/domain"
)

type UserRepository interface {
	// get retrieves an user by its email
	get(email string) (domain.User, error)
	// save saves an user to the database
	save(user *domain.User) error
	// update updates the user's information
	update(user *domain.User) error
}

type WorkspaceRepository interface {
	// get retrieves a workspace by its id
	get(workspaceId uuid.UUID) (domain.Workspace, error)
	// save creates a new workspace
	save(workspace *domain.Workspace, ownerEmail string) error
	// delete deletes a given workspace
	delete(workspaceId uuid.UUID) error
	// getForUser retrieves all workspaces for a given user no matter whether their the owner or an collaborator
	getForUser(email string) ([]domain.Workspace, error)
}

type CollectionRepository interface {
	// get retrieves a collection by its id
	get(collectionId uuid.UUID) (domain.Workspace, error)
	// save creates a new collection inside a workspace
	save(collectionId *domain.Workspace, workspaceId uuid.UUID) error
	// delete deletes a given collection
	delete(collectionId uuid.UUID) error
}

type RequestRepository interface {
	// get retrieves all requests for a given collection
	get(collectionId uuid.UUID) ([]domain.Request, error)
	// save saves a new request inside a collection
	save(request *domain.Request) (string, error)
	// update updates the request's information
	update(request *domain.Request) error
	// delete deletes a request
	delete(requestId uuid.UUID) error
}