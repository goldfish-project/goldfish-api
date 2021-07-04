package ports

import (
	"github.com/google/uuid"
	"goldfish-api/internal/core/domain"
)

type UserRepository interface {
	// Get retrieves an user by its user id
	Get(userId uuid.UUID) (domain.User, error)
	// GetByEmail retrieves an user by its email
	GetByEmail(email string) (domain.User, error)
	// Save saves an user to the database
	Save(user *domain.User) error
	// Update updates the user's information
	Update(user *domain.User) error
}

type WorkspaceRepository interface {
	// Get retrieves a workspace by its id
	Get(workspaceId uuid.UUID) (domain.Workspace, error)
	// Save creates a new workspace
	Save(workspace *domain.Workspace, ownerUserId string) error
	// Delete deletes a given workspace
	Delete(workspaceId uuid.UUID) error
	// GetForUser retrieves all workspaces for a given user no matter whether their the owner or an collaborator
	GetForUser(userId string) ([]domain.Workspace, error)
}

type CollectionRepository interface {
	// Get retrieves a collection by its id
	Get(collectionId uuid.UUID) (domain.Workspace, error)
	// Save creates a new collection inside a workspace
	Save(collectionId *domain.Workspace, workspaceId uuid.UUID) error
	// Delete deletes a given collection
	Delete(collectionId uuid.UUID) error
}

type RequestRepository interface {
	// Get retrieves all requests for a given collection
	Get(collectionId uuid.UUID) ([]domain.Request, error)
	// Save saves a new request inside a collection
	Save(request *domain.Request) (string, error)
	// Update updates the request's information
	Update(request *domain.Request) error
	// Delete deletes a request
	Delete(requestId uuid.UUID) error
}